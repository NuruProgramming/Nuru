package repl

import (
	"fmt"
	"strings"
	"time"

	"github.com/NuruProgramming/Nuru/ast"
	"github.com/NuruProgramming/Nuru/evaluator"
	"github.com/NuruProgramming/Nuru/lexer"
	"github.com/NuruProgramming/Nuru/module"
	"github.com/NuruProgramming/Nuru/object"
	"github.com/NuruProgramming/Nuru/parser"
	"github.com/NuruProgramming/Nuru/token"
	"github.com/charmbracelet/lipgloss"
)

const (
	MsgNoTests      = "Hakuna majaribio yaliyopatikana"
	MsgSyntaxErrors = "Kuna Makosa Yafuatayo:"

	TxtPass = " IMEPITA"
	TxtFail = " IMEFELI"

	HeaderStart    = "============================= JARIBIO LIMEANZA ============================="
	HeaderFailures = "========================= MAJARIBIO YALIYOSHINDWA =========================="

	FmtCollected    = "imekusanya vipengele %d\n\n"
	FmtFailDivider  = "_________________________ %s _________________________"
	FmtFailDetail   = ">   %s\n\n"
	FmtSummaryFail  = "%d imeshindwa"
	FmtSummaryPass  = "%d imefaulu"
	FmtSummaryTime  = "kwa muda wa %.2fs"
	FmtSummaryTotal = "======================= %s ========================"
	FmtTestVerbose  = "%-30s %s"
)

var (
	stylePass     = lipgloss.NewStyle().Foreground(lipgloss.Color("42"))
	styleFail     = lipgloss.NewStyle().Foreground(lipgloss.Color("196"))
	styleFailBold = styleFail.Copy().Bold(true)
	stylePassBold = stylePass.Copy().Bold(true)
	styleHeader   = lipgloss.NewStyle().Bold(true)
)

type testFailure struct {
	name string
	msg  string
}

type testEvent struct {
	pass bool
	msg  string
}

func Test(source string) bool {
	env := object.NewEnvironment()
	var events []testEvent

	teardown := attachReporter(&events)
	defer teardown()

	program, errs := parse(source)
	if len(errs) > 0 {
		reportSyntaxErrors(errs)
		return false
	}

	hoistFunctions(program, env)
	evaluator.Eval(program, env)

	tests := collectTests(program)
	if len(tests) == 0 {
		fmt.Println(MsgNoTests)
		return true
	}

	printSessionHeader(len(tests))

	startTotal := time.Now()
	failures := runTestLoop(tests, env, &events)
	totalDuration := time.Since(startTotal).Seconds()

	fmt.Println("")

	if len(failures) > 0 {
		printFailuresSection(failures)
	}

	return printSessionSummary(len(tests), len(failures), totalDuration)
}

func attachReporter(events *[]testEvent) func() {
	module.TestReporter = func(pass bool, message string) {
		*events = append(*events, testEvent{pass, message})
	}
	return func() {
		module.TestReporter = nil
	}
}

func parse(source string) (*ast.Program, []string) {
	l := lexer.New(source)
	p := parser.New(l)
	return p.ParseProgram(), p.Errors()
}

func reportSyntaxErrors(errors []string) {
	fmt.Println(styleFail.Render(MsgSyntaxErrors))
	for _, msg := range errors {
		fmt.Println("\t" + styleFail.Render(msg))
	}
}

func hoistFunctions(program *ast.Program, env *object.Environment) {
	for _, stmt := range program.Statements {
		exprStmt, ok := stmt.(*ast.ExpressionStatement)
		if !ok {
			continue
		}
		fn, ok := exprStmt.Expression.(*ast.FunctionLiteral)
		if !ok {
			continue
		}
		if fn.Name == "" {
			continue
		}
		val := evaluator.Eval(fn, env)
		env.Set(fn.Name, val)
	}
}

func collectTests(program *ast.Program) []string {
	var tests []string
	for _, stmt := range program.Statements {
		name := identifyTest(stmt)
		if name != "" {
			tests = append(tests, name)
		}
	}
	return tests
}

func identifyTest(stmt ast.Statement) string {
	var name string
	switch s := stmt.(type) {
	case *ast.ExpressionStatement:
		if fn, ok := s.Expression.(*ast.FunctionLiteral); ok {
			name = fn.Name
		}
	case *ast.LetStatement:
		name = s.Name.Value
	}
	if strings.HasPrefix(name, "pima_") {
		return name
	}
	return ""
}

func runTestLoop(tests []string, env *object.Environment, events *[]testEvent) []testFailure {
	var failures []testFailure

	for _, name := range tests {
		*events = []testEvent{}

		call := &ast.CallExpression{
			Token:     token.Token{Type: token.IDENT, Literal: name},
			Function:  &ast.Identifier{Token: token.Token{Type: token.IDENT, Literal: name}, Value: name},
			Arguments: []ast.Expression{},
		}
		evaluator.Eval(call, env)

		testFailed := false
		var failureMsg string

		for _, e := range *events {
			if !e.pass {
				testFailed = true
				failureMsg = e.msg
			}
		}

		if testFailed {
			status := styleFail.Render(TxtFail)
			fmt.Printf(FmtTestVerbose+"\n", name, status)
			failures = append(failures, testFailure{name: name, msg: failureMsg})
		} else {
			status := stylePass.Render(TxtPass)
			fmt.Printf(FmtTestVerbose+"\n", name, status)
		}
	}

	return failures
}

func printSessionHeader(count int) {
	fmt.Println(styleHeader.Render(HeaderStart))
	fmt.Printf(FmtCollected, count)
}

func printFailuresSection(failures []testFailure) {
	fmt.Println("")
	fmt.Println(styleFailBold.Render(HeaderFailures))

	for _, f := range failures {
		header := fmt.Sprintf(FmtFailDivider, f.name)
		fmt.Println(styleFail.Render(header))
		fmt.Printf(FmtFailDetail, f.msg)
	}
}

func printSessionSummary(total, failed int, duration float64) bool {
	passed := total - failed

	parts := []string{}
	if failed > 0 {
		parts = append(parts, fmt.Sprintf(FmtSummaryFail, failed))
	}
	parts = append(parts, fmt.Sprintf(FmtSummaryPass, passed))
	parts = append(parts, fmt.Sprintf(FmtSummaryTime, duration))

	summaryText := strings.Join(parts, ", ")
	fullBar := fmt.Sprintf(FmtSummaryTotal, summaryText)

	if failed > 0 {
		fmt.Println(styleFailBold.Render(fullBar))
		return false
	}

	fmt.Println(stylePassBold.Render(fullBar))
	return true
}
