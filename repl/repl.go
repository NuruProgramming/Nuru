package repl

import (
	"bufio"
	"embed"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/AvicennaJr/Nuru/evaluator"
	"github.com/AvicennaJr/Nuru/lexer"
	"github.com/AvicennaJr/Nuru/object"
	"github.com/AvicennaJr/Nuru/parser"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	zone "github.com/lrstanley/bubblezone"
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

//go:embed docs
var res embed.FS

func Read(contents string) {
	env := object.NewEnvironment()

	l := lexer.New(contents)
	p := parser.New(l)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		fmt.Println(colorfy(ERROR_FACE, 31))
		fmt.Println("Kuna Errors Zifuatazo:")

		for _, msg := range p.Errors() {
			fmt.Println("\t" + colorfy(msg, 31))
		}

	}
	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		if evaluated.Type() != object.NULL_OBJ {
			fmt.Println(colorfy(evaluated.Inspect(), 32))
		}
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

func Tutor() {
	zone.NewGlobal()

	languageChoice := []list.Item{
		languages{title: "Kiswahili", desc: "Soma nyaraka kwa Kiswahili", dir: "sw"},
		languages{title: "English", desc: "Read documentation in English", dir: "en"},
	}

	var p playground

	p.languageCursor = list.New(languageChoice, list.NewDefaultDelegate(), 50, 8)
	p.languageCursor.Title = "Chagua Lugha"
	p.languageCursor.SetFilteringEnabled(false)
	p.languageCursor.SetShowStatusBar(false)
	p.languageCursor.SetShowPagination(false)
	p.languageCursor.SetShowHelp(false)
	p.toc = list.New(englishItems, list.NewDefaultDelegate(), 0, 0)
	p.toc.Title = "Table of Contents"
	p.id = zone.NewPrefix()

	if _, err := tea.NewProgram(p, tea.WithMouseAllMotion()).Run(); err != nil {
		log.Fatal(err)
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

var (
	englishItems = []list.Item{
		item{title: "Arrays", desc: "🚀 Unleash the power of arrays in Nuru", filename: "arrays.md"},
		item{title: "Booleans", desc: "👍👎 Master the world of 'if' and 'else' with bools", filename: "bool.md"},
		item{title: "Builtins", desc: "💡 Reveal the secrets of builtin functions in Nuru", filename: "builtins.md"},
		item{title: "Comments", desc: "💬 Speak your mind with comments in Nuru", filename: "comments.md"},
		item{title: "Dictionaries", desc: "📚 Unlock the knowledge of dictionaries in Nuru", filename: "dictionaries.md"},
		item{title: "For", desc: "🔄 Loop like a pro with 'for' in Nuru", filename: "for.md"},
		item{title: "Function", desc: "🔧 Create powerful functions in Nuru", filename: "function.md"},
		item{title: "Identifiers", desc: "🔖 Give your variables their own identity in Nuru", filename: "identifiers.md"},
		item{title: "If Statements", desc: "🔮 Control the flow with 'if' statements in Nuru", filename: "ifStatements.md"},
		item{title: "Keywords", desc: "🔑 Learn the secret language of Nuru's keywords", filename: "keywords.md"},
		item{title: "Null", desc: "🌌 Embrace the void with Null in Nuru", filename: "null.md"},
		item{title: "Numbers", desc: "🔢 Discover the magic of numbers in Nuru", filename: "numbers.md"},
		item{title: "Operators", desc: "🧙 Perform spells with Nuru's operators", filename: "operators.md"},
		item{title: "Strings", desc: "🎼 Compose stories with strings in Nuru", filename: "strings.md"},
		item{title: "Switch", desc: "🧭 Navigate complex scenarios with 'switch' in Nuru", filename: "switch.md"},
		item{title: "While", desc: "⌛ Learn the art of patience with 'while' loops in Nuru", filename: "while.md"},
	}

	kiswahiliItems = []list.Item{
		item{title: "Maoni Katika Nuru", desc: "💬 Toa mawazo yako na maoni (comments) katika Nuru", filename: "comments.md"},
	}
)
