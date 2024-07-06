package main

import (
	"cmp"
	"fmt"
	"os"
	"regexp"
	"strings"

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

	flags, args := parseArgs(os.Args[1:])

	handleFlags(flags)

	env_val := cmp.Or(os.Getenv("RANGI"), "1")
	os.Setenv("RANGI", env_val)

	if len(args) <= 0 {
		pre_text := fmt.Sprintf(`nuru version %s
tumia toka() ama exit() kundoka.
`, nuruVersion)
		fmt.Println(pre_text)
		repl.Start()
		return
	}

	var filename = args[0]
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Nuru imeshindwa kusoma faili %s\n", args[0])
		os.Exit(1)
	}
	repl.Read(filename, string(content))

}

type msaada_st struct {
	fupi, refu string
	msaada     string
	hatua      func()
}

func pataMaandishiMsaada() []msaada_st {
	return []msaada_st{
		{fupi: "-m", refu: "--msaada", msaada: "Onyesha huu ujumbe", hatua: onyeshaMsaada},
		{fupi: "-h", refu: "--help", msaada: "Show help message (in swahili)", hatua: onyeshaMsaada},
		{fupi: "-v", refu: "--version", msaada: "Show version", hatua: version},
		{fupi: "-t", refu: "--toleo", msaada: "Onyesha toleo", hatua: version},
		{fupi: "-r", refu: "--hakuna-rangi", msaada: "Usionyeshe rangi", hatua: onyeshaRangi},
	}
}

// call the relevant functions or do the neccesary action
func handleFlags(flags []string) {
	var found bool
	for _, flag := range flags {
		for _, msaada := range pataMaandishiMsaada() {
			if flag == msaada.fupi || flag == msaada.refu {
				found = true

				if msaada.hatua != nil {
					msaada.hatua()
				} else {
					fmt.Printf("bendera '%s' imefafanuliwa lakini hakuna hatua yoyote ilitolewa kuitekeleza\n", flag)
					os.Exit(1)
				}
			}

		}
		if !found {
			fmt.Printf("'%s' haijafanuliwa kama bendera lakini imepatikana\n", flag)
			os.Exit(2)
		}
	}

}

func onyeshaRangi() {
	os.Setenv("RANGI", "0")
}

func onyeshaMsaada() {
	var pre_text string = fmt.Sprintf("Nuru %s\nUtumiaji [faili] | [chaguzi]\n", nuruVersion)
	var post_text string = "\nIkiwa kuna hitilafu tafadhali tuma ripoti ya hitilafu (bug report) kwa: https://github.com/NuruProgramming/Nuru/issues"
	var msaada_v strings.Builder
	msaada_v.WriteString(pre_text)
	for _, m := range pataMaandishiMsaada() {
		msaada_v.WriteString(fmt.Sprintf("%s\t%s\t%s\n", m.fupi, m.refu, m.msaada))
	}

	msaada_v.WriteString(post_text)

	fmt.Println(msaada_v.String())
	os.Exit(0)
}

// Do a simplistic job of going over the args passed and seperating the args from the flags
func parseArgs(passedArgs []string) ([]string, []string) {
	var flags []string
	var args []string
	for _, arg := range passedArgs {
		var multiple = regexp.MustCompile(`^-[\w]{2,}$`)
		if multiple.MatchString(arg) {
			for _, ml := range arg[1:] {
				flags = append(flags, fmt.Sprintf("-%s", string(ml)))
			}
			continue
		}

		var arg_r = regexp.MustCompile(`^--?`)
		if arg_r.MatchString(arg) {
			flags = append(flags, arg)
			continue
		}

		args = append(args, arg)
	}

	return flags, args
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
