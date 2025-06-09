package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/NuruProgramming/Nuru/analysis"
	"github.com/NuruProgramming/Nuru/object"
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
	%s: Kufanya uchambuzi wa msimbo
	%s: Kuonyesha uhusiano wa moduli
	%s: Onyesha taarifa zaidi wakati wa uendeshaji
`,
		styles.HelpStyle.Bold(true).Render("nuru"),
		styles.HelpStyle.Bold(true).Render("nuru jinaLaFile.nr"),
		styles.HelpStyle.Bold(true).Render("nuru --nyaraka"),
		styles.HelpStyle.Bold(true).Render("nuru --toleo"),
		styles.HelpStyle.Bold(true).Render("nuru --chambua <file_or_dir>"),
		styles.HelpStyle.Bold(true).Render("nuru --tengeneza <file_or_dir>"),
		styles.HelpStyle.Bold(true).Render("nuru --verbose jinaLaFile.nr")))
)

func main() {
	// Get a copy of the original arguments
	originalArgs := os.Args

	// First, process the --verbose flag
	object.VerboseGC = false
	args := []string{}
	for _, arg := range originalArgs {
		if arg == "--verbose" {
			object.VerboseGC = true
		} else {
			args = append(args, arg)
		}
	}

	// Now process the main command
	if len(args) < 2 {
		help := styles.HelpStyle.Render("💡 Tumia exit() au toka() kuondoka")
		fmt.Println(lipgloss.JoinVertical(lipgloss.Left, NewLogo, "\n", help))
		repl.Start()
		return
	}

	// Handle analysis commands first
	if len(args) > 1 {
		switch args[1] {
		case "--chambua", "-chambua":
			if len(args) > 2 {
				exitCode := analysis.AnalyzeCommand(args[2:])
				os.Exit(exitCode)
			} else {
				fmt.Println(styles.ErrorStyle.Render("Error: No file or directory specified for analysis."))
				os.Exit(1)
			}
			return
		case "--tengeneza", "-tengeneza":
			if len(args) > 2 {
				exitCode := analysis.HandleVisualizeDepsCommand(args[2:])
				os.Exit(exitCode)
			} else {
				fmt.Println(styles.ErrorStyle.Render("Error: No file or directory specified for visualization."))
				os.Exit(1)
			}
			return
		}
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

			if strings.HasSuffix(file, ".nr") || strings.HasSuffix(file, ".sw") {
				contents, err := os.ReadFile(file)
				if err != nil {
					fmt.Println(styles.ErrorStyle.Render("Error: Nuru imeshindwa kusoma faili: ", args[1]))
					os.Exit(1)
				}

				absPath, err := filepath.Abs(file)
				if err != nil {
					absPath = file // fallback
				}
				dirPath := absPath
				if idx := strings.LastIndex(absPath, string(os.PathSeparator)); idx != -1 {
					dirPath = absPath[:idx]
				} else {
					dirPath = "."
				}

				env := object.NewEnvironment()
				env.Set("__FILE__", &object.String{Value: absPath})
				env.Set("__DIR__", &object.String{Value: dirPath})

				repl.ReadWithEnv(string(contents), env)
			} else {
				fmt.Println(styles.ErrorStyle.Render("'"+file+"'", "sii faili sahihi. Tumia faili la '.nr' au '.sw'"))
				os.Exit(1)
			}
		}
	} else {
		fmt.Println(styles.ErrorStyle.Render("Error: Operesheni imeshindikana boss."))
		fmt.Println(Help)
		os.Exit(1)
	}
}
