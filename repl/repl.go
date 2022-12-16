package repl

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/AvicennaJr/Nuru/evaluator"
	"github.com/AvicennaJr/Nuru/lexer"
	"github.com/AvicennaJr/Nuru/object"
	"github.com/AvicennaJr/Nuru/parser"
)

const PROMPT = ">>> "
const ERROR_FACE = `
	███████████████████████████
	███████▀▀▀░░░░░░░▀▀▀███████
	████▀░░░░░░░░░░░░░░░░░▀████
	███│░░░░░░░░░░░░░░░░░░░│███
	██▌│░░░░░░░░░░░░░░░░░░░│▐██
	██░└┐░░░░░░░░░░░░░░░░░┌┘░██
	██░░└┐░░░░░░░░░░░░░░░┌┘░░██
	██░░┌┘▄▄▄▄▄░░░░░▄▄▄▄▄└┐░░██
	██▌░│██████▌░░░▐██████│░▐██
	███░│▐███▀▀░░▄░░▀▀███▌│░███
	██▀─┘░░░░░░░▐█▌░░░░░░░└─▀██
	██▄░░░▄▄▄▓░░▀█▀░░▓▄▄▄░░░▄██
	████▄─┘██▌░░░░░░░▐██└─▄████
	█████░░▐█─┬┬┬┬┬┬┬─█▌░░█████
	████▌░░░▀┬┼┼┼┼┼┼┼┬▀░░░▐████
	█████▄░░░└┴┴┴┴┴┴┴┘░░░▄█████
	███████▄░░░░░░░░░░░▄███████
	██████████▄▄▄▄▄▄▄██████████
	███████████████████████████

  █▄▀ █░█ █▄░█ ▄▀█   █▀ █░█ █ █▀▄ ▄▀█
  █░█ █▄█ █░▀█ █▀█   ▄█ █▀█ █ █▄▀ █▀█

`

func Read(contents string) {
	env := object.NewEnvironment()

	l := lexer.New(contents)
	p := parser.New(l)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		// fmt.Println(colorfy(ERROR_FACE, 31))
		fmt.Println("Kuna Errors Zifuatazo:")

		for _, msg := range p.Errors() {
			fmt.Println("\t" + colorfy(msg, 31))
		}

	}
	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		fmt.Println(colorfy(evaluated.Inspect(), 32))
	}

}

func Start(in io.Reader, out io.Writer) {

	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		if strings.TrimSpace(line) == "exit()" || strings.TrimSpace(line) == "toka()" {
			fmt.Println("✨🅺🅰🆁🅸🅱🆄 🆃🅴🅽🅰✨")
			os.Exit(0)
		}
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParseErrors(out, p.Errors())
			continue
		}
		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			if evaluated.Type() != object.NULL_OBJ {
				io.WriteString(out, colorfy(evaluated.Inspect(), 32))
				io.WriteString(out, "\n")
			}
		}
	}
}

func printParseErrors(out io.Writer, errors []string) {
	//io.WriteString(out, colorfy(ERROR_FACE, 31))
	io.WriteString(out, "Kuna Errors Zifuatazo:\n")

	for _, msg := range errors {
		io.WriteString(out, "\t"+colorfy(msg, 31)+"\n")
	}
}

func colorfy(str string, colorCode int) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", colorCode, str)
}
