//go:build !wasm || !js

package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/NuruProgramming/Nuru/repl"
	"github.com/NuruProgramming/Nuru/styles"
	"github.com/charmbracelet/lipgloss"
)

var (
	Title = styles.TitleStyle.
		Render(`
█░░ █░█ █▀▀ █░█ ▄▀█   █▄█ ▄▀█   █▄░█ █░█ █▀█ █░█
█▄▄ █▄█ █▄█ █▀█ █▀█   ░█░ █▀█   █░▀█ █▄█ █▀▄ █▄█`)
	Version = styles.VersionStyle.Render("v0.5.18")
	Author  = styles.AuthorStyle.Render("by Nuru Org")
	NewLogo = lipgloss.JoinVertical(lipgloss.Center, Title, lipgloss.JoinHorizontal(lipgloss.Center, Author, " | ", Version))
	Help    = styles.HelpStyle.Italic(false).Render(fmt.Sprintf(`💡 Namna ya kutumia Nuru:
	%s: Kuanza programu ya Nuru
	%s: Kuendesha faili la Nuru
	%s: Kusoma nyaraka za Nuru
	%s: Kufahamu toleo la Nuru
`,
		styles.HelpStyle.Bold(true).Render("nuru"),
		styles.HelpStyle.Bold(true).Render("nuru jinaLaFile.nr"),
		styles.HelpStyle.Bold(true).Render("nuru --nyaraka"),
		styles.HelpStyle.Bold(true).Render("nuru --toleo")))
)

func hasValidExtension(file string) bool {
	return strings.HasSuffix(file, ".nr") || strings.HasSuffix(file, ".sw")
}

func readFile(file string) (string, error) {
	if !hasValidExtension(file) {
		return "", fmt.Errorf("'%s' sii faili sahihi. Tumia faili la '.nr' au '.sw'", file)
	}

	contents, err := os.ReadFile(file)
	if err != nil {
		return "", fmt.Errorf("Error: Nuru imeshindwa kusoma faili: %s", file)
	}

	return string(contents), nil
}

func printError(err error) {
	fmt.Println(styles.ErrorStyle.Render(err.Error()))
	os.Exit(1)
}

func main() {

	args := os.Args
	if len(args) < 2 {

		help := styles.HelpStyle.Render("💡 Tumia exit() au toka() kuondoka")
		fmt.Println(lipgloss.JoinVertical(lipgloss.Left, NewLogo, "\n", help))
		repl.Start()
		return
	}

	if len(args) == 2 {
		switch args[1] {
		case "msaada", "-msaada", "--msaada", "help", "-help", "--help", "-h":
			fmt.Println(Help)
		case "version", "-version", "--version", "-v", "v", "--toleo", "-toleo":
			fmt.Println(NewLogo)
		case "-docs", "--docs", "-nyaraka", "--nyaraka":
			repl.Docs()
		default:
			file := args[1]
			if content, err := readFile(file); err == nil {
				repl.Read(content)
			} else {
				printError(err)
			}
		}

		return
	}

	if len(args) > 2 {
		switch args[1] {
		case "pima":
			file := args[2]
			if content, err := readFile(file); err == nil {
				repl.Test(content)
			} else {
				printError(err)
			}
		default:
			printError(fmt.Errorf("Error: Operesheni uliyochagua haijaumdwa"))
		}

		return
	}

	fmt.Println(styles.ErrorStyle.Render("Error: Operesheni imeshindikana boss."))
	fmt.Println(Help)
	os.Exit(1)

}
