package main

import (
	"fmt"
	"os"
	"regexp"
	"unicode"

	"github.com/NuruProgramming/Nuru/repl"
)

type flagsPassed struct {
	help    bool
	version bool
	msaada  bool
}

type passedArgs struct {
	flags flagsPassed
	args  []string
}

var nuruVersion = "v0.5.1"

func main() {
	pa, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n\n", err)
		msaada(1)
	}

	flags := pa.flags
	handleFlags(flags)

	args := pa.args
	if len(args) <= 0 {
		pre_text := fmt.Sprintf(`nuru version %s
tumia toka() ama exit() kundoka.
`, nuruVersion)
		fmt.Println(pre_text)
		repl.Start()
		return
	}

	content, err := os.ReadFile(args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Nuru imeshindwa kusoma faili %s\n", args[0])
		os.Exit(1)
	}
	repl.Read(string(content))

}

// call the relevant functions or do the neccesary action
func handleFlags(flags flagsPassed) {
	if flags.help {
		help()
	}
	if flags.version {
		version()
	}
	if flags.msaada {
		msaada()
	}
}

// parse the arguments passed in from stdin (or where they are received from)
// It follows closely with the GNU Argument Syntax:
// https://www.gnu.org/software/libc/manual/html_node/Argument-Syntax.html
// Arguments and flags are each populated differently.
// If any argument is found that is not defined, then parsing stops and an error returned.
func parseArgs(args []string) (passedArgs, error) {
	var pa passedArgs
	for _, arg := range args {
		var multiple = regexp.MustCompile(`^-[\w]{2,}$`)
		if multiple.MatchString(arg) {
			for _, ml := range arg[1:] {
				if unicode.IsSpace(ml) {
					continue
				}
				switch ml {
				case 'h':
					pa.flags.help = true
				case 'm':
					pa.flags.msaada = true
				case 'v':
					pa.flags.version = true
				default:
					return pa, fmt.Errorf("-%s haijafanuliwa kama bendera lakini imepatikana", string(ml))
				}
			}

			continue
		}
		switch arg {
		case "--help", "-h":
			pa.flags.help = true
		case "--msaada", "-m":
			pa.flags.msaada = true
		case "--version", "-v", "--toleo", "-t":
			pa.flags.version = true
		default:
			var re = regexp.MustCompile(`^--?`)
			if re.MatchString(arg) {
				return pa, fmt.Errorf("%s haijafanuliwa kama bendera lakini imepatikana", arg)
			}
			pa.args = append(pa.args, arg)
		}
	}

	return pa, nil
}

func printHelpMessage(exitCode int, message string) {
	issue := "If there is an error please file a bug at: https://github.com/NuruProgramming/Nuru/issues"
	message = fmt.Sprintf("%s\n%s\n", message, issue)
	if exitCode != 0 {
		fmt.Fprintf(os.Stderr, message)
	} else {
		fmt.Fprintf(os.Stdout, message)
	}
	os.Exit(exitCode)
}

func help(args ...int) {
	exitCode := 0
	if len(args) > 0 {
		exitCode = args[0]
	}
	str := `nuru programming language

usage: nuru [option] [file]

Flags:
--help: Display this help message
--version: Display the current version
`
	printHelpMessage(exitCode, str)
}

func msaada(args ...int) {
	exitCode := 0
	if len(args) > 0 {
		exitCode = args[0]
	}
	ms := `Lugha ya programu ya Nuru

Matumizi: nuru [chaguo] [faili]

Bendera:
--msaada: Onyesha ujumbe huu wa msaada
--toleo: Onyesha toleo la sasa
`
	printHelpMessage(exitCode, ms)
}

func version() {
	ver := fmt.Sprintf(`nuru programming %s
Copyright (C) 2024 Nuru Authors.
This is free software; see the source for copying conditions.
There is NO warranty; not even for MERCHANDABILITY or FITNESS FOR A PARTICULAR PURPOSE.
`, nuruVersion)

	fmt.Println(ver)
	os.Exit(0)
}
