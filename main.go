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
		Render(`  _   _      🔥            
 | \ | |_   _ _ __ _   _ 
 |  \| | | | | '__| | | |
 | |\  | |_| | |  | |_| |
 |_| \_|\__,_|_|   \__,_|`)
	Version = styles.VersionStyle.Render("v0.5.1")
	Author  = styles.AuthorStyle.Render("copyleft 🄯 Nuru Organization")
	Logo    = lipgloss.JoinVertical(lipgloss.Center, Title, lipgloss.JoinHorizontal(lipgloss.Center, Author, " | ", Version))
	Help    = lipgloss.JoinVertical(lipgloss.Center, Title, styles.HelpStyle.Italic(false).Render(fmt.Sprintf(`💡 Namna ya kutumia Nuru:
	%s: Kuanza programu ya Nuru
	%s: Kuendesha faili la Nuru
	%s: Kusoma nyaraka za Nuru
	%s: Kufahamu toleo la Nuru
	%s: Kupata msaada
`,
		styles.HelpStyle.Bold(true).Render("nuru"),
		styles.HelpStyle.Bold(true).Render("nuru jina_la_file.nr"),
		styles.HelpStyle.Bold(true).Render("nuru -n | --nyaraka"),
		styles.HelpStyle.Bold(true).Render("nuru -t | --toleo"),
		styles.HelpStyle.Bold(true).Render("nuru -m | --msaada"))))
	EnglishHelp = lipgloss.JoinVertical(lipgloss.Center, Title, styles.HelpStyle.Italic(false).Render(fmt.Sprintf(`💡 How to use Nuru:
	%s: to start the Nuru repl
	%s: to run a Nuru script
	%s: to read the Nuru documentation
	%s: to know the Nuru version
	%s: to get this help message
`,
		styles.HelpStyle.Bold(true).Render("nuru"),
		styles.HelpStyle.Bold(true).Render("nuru file_name.nr"),
		styles.HelpStyle.Bold(true).Render("nuru -d | --docs"),
		styles.HelpStyle.Bold(true).Render("nuru -v | --version"),
		styles.HelpStyle.Bold(true).Render("nuru -h | --help"))))
)

func main() {

	args := os.Args
	if len(args) < 2 {
		help := styles.HelpStyle.Render("💡 Tumia toka() kuondoka || Use exit() to exit")
		fmt.Println(lipgloss.JoinVertical(lipgloss.Left, Logo, "\n", help))
		repl.Start()
		return
	}

	if len(args) == 2 {
		switch args[1] {
		case "-m", "--msaada":
			fmt.Println(Help)
		case "-h", "--help":
			fmt.Println(EnglishHelp)
		case "-t", "--toleo", "-v", "--version":
			fmt.Println(Logo)
		case "-n", "--nyaraka", "-d", "--docs":
			repl.Docs()
		default:
			file := args[1]

			if strings.HasSuffix(file, "nr") || strings.HasSuffix(file, ".sw") {
				contents, err := os.ReadFile(file)
				if err != nil {
					fmt.Println(styles.ErrorStyle.Render("Error: Nuru imeshindwa kusoma faili: ", args[1]))
					os.Exit(1)
				}

				repl.Read(string(contents))
			} else {
				fmt.Println(styles.ErrorStyle.Render("'"+file+"'", "sii faili sahihi. Tumia faili la '.nr' au '.sw'"))
				os.Exit(1)
			}
		}
	} else {
		fmt.Println(styles.ErrorStyle.Render("Error: idadi ya bendera sii sahihi"))
		fmt.Println(Help)
		os.Exit(1)
	}
}
