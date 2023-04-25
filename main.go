package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/AvicennaJr/Nuru/repl"
	"github.com/AvicennaJr/Nuru/styles"
	"github.com/charmbracelet/lipgloss"
)

var (
	Title = styles.TitleStyle.
		Render(`
█░░ █░█ █▀▀ █░█ ▄▀█   █▄█ ▄▀█   █▄░█ █░█ █▀█ █░█
█▄▄ █▄█ █▄█ █▀█ █▀█   ░█░ █▀█   █░▀█ █▄█ █▀▄ █▄█`)
	Version = styles.VersionStyle.Render("v0.5.1")
	Author  = styles.AuthorStyle.Render("by Avicenna")
	NewLogo = lipgloss.JoinVertical(lipgloss.Center, Title, lipgloss.JoinHorizontal(lipgloss.Center, Author, " | ", Version))
)

func main() {

	args := os.Args
	if len(args) < 2 {

		help := styles.HelpStyle.Render("💡 Tumia exit() au toka() kuondoka")
		fmt.Println(lipgloss.JoinVertical(lipgloss.Left, NewLogo, "\n", help))
		repl.Start()
		os.Exit(0)
	}

	if len(args) == 2 {

		switch args[1] {
		case "msaada", "-msaada", "--msaada", "help", "-help", "--help", "-h":
			fmt.Println(styles.HelpStyle.Italic(false).Render(fmt.Sprintf(`💡 Namna ya kutumia Nuru:
	%s: Kuanza programu ya Nuru
	%s: Kurun file la Nuru
	%s: Kusoma nyaraka za Nuru
	%s: Kufahamu toleo la Nuru
`,
				styles.HelpStyle.Bold(true).Render("nuru"),
				styles.HelpStyle.Bold(true).Render("nuru jinaLaFile.nr"),
				styles.HelpStyle.Bold(true).Render("nuru --nyaraka"),
				styles.HelpStyle.Bold(true).Render("nuru --toleo"))))
			os.Exit(0)
		case "version", "-version", "--version", "-v", "v", "--toleo", "-toleo":
			fmt.Println(NewLogo)
			os.Exit(0)
		case "-docs", "--docs", "-nyaraka", "--nyaraka":
			repl.Docs()
			os.Exit(0)
		}

		file := args[1]

		if strings.HasSuffix(file, "nr") || strings.HasSuffix(file, ".sw") {
			contents, err := ioutil.ReadFile(file)
			if err != nil {
				fmt.Println(styles.ErrorStyle.Render("Error: Nimeshindwa kusoma file: ", args[1]))
				os.Exit(0)
			}

			repl.Read(string(contents))
		} else {
			fmt.Println(styles.ErrorStyle.Render("'"+file+"'", "sii file sahihi. Tumia file la '.nr' au '.sw'"))
			os.Exit(0)
		}

	} else {
		fmt.Println(styles.ErrorStyle.Render("Error: Operesheni imeshindikana boss."))
		fmt.Println(styles.HelpStyle.Italic(false).Render(fmt.Sprintf(`💡 Namna ya kutumia Nuru:
	%s: Kuanza programu ya Nuru
	%s: Kurun file la Nuru
	%s: Kusoma nyaraka za Nuru
	%s: Kufahamu toleo la Nuru
`,
			styles.HelpStyle.Bold(true).Render("nuru"),
			styles.HelpStyle.Bold(true).Render("nuru jinaLaFile.nr"),
			styles.HelpStyle.Bold(true).Render("nuru --nyaraka"),
			styles.HelpStyle.Bold(true).Render("nuru --toleo"))))
		os.Exit(0)
	}
}
