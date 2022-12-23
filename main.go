package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/AvicennaJr/Nuru/repl"
)

const (
	LOGO = `

█░░ █░█ █▀▀ █░█ ▄▀█   █▄█ ▄▀█   █▄░█ █░█ █▀█ █░█
█▄▄ █▄█ █▄█ █▀█ █▀█   ░█░ █▀█   █░▀█ █▄█ █▀▄ █▄█                                        

        | Authored by Avicenna | v0.2.0 |
`
)

func main() {

	args := os.Args
	coloredLogo := fmt.Sprintf("\x1b[%dm%s\x1b[0m", 34, LOGO)

	if len(args) < 2 {

		fmt.Println(coloredLogo)
		fmt.Println("𝑯𝒂𝒃𝒂𝒓𝒊, 𝒌𝒂𝒓𝒊𝒃𝒖 𝒖𝒕𝒖𝒎𝒊𝒆 𝒍𝒖𝒈𝒉𝒂 𝒚𝒂 𝑵𝒖𝒓𝒖 ✨")
		fmt.Println("\nTumia exit() au toka() kuondoka")

		repl.Start(os.Stdin, os.Stdout)
	}

	if len(args) == 2 {

		switch args[1] {
		case "msaada", "--msaada", "help", "--help", "-h":
			fmt.Printf("\x1b[%dm%s\x1b[0m\n", 32, "\nTumia 'nuru' kuanza program\n\nAU\n\nTumia 'nuru' ikifuatiwa na jina la file.\n\n\tMfano:\tnuru fileYangu.nr")
			os.Exit(0)
		case "version", "--version", "-v", "v":
			fmt.Println(coloredLogo)
			os.Exit(0)
		}

		file := args[1]

		if strings.HasSuffix(file, "nr") || strings.HasSuffix(file, ".sw") {
			contents, err := ioutil.ReadFile(file)
			if err != nil {
				fmt.Printf("\x1b[%dm%s%s\x1b[0m\n", 31, "Error: Nimeshindwa kusoma file: ", args[0])
				os.Exit(0)
			}

			repl.Read(string(contents))
		} else {
			fmt.Printf("\x1b[%dm%s%s\x1b[0m", 31, file, " sii file sahihi. Tumia file la '.nr' au '.sw'\n")
			os.Exit(0)
		}

	} else {
		fmt.Printf("\x1b[%dm%s\x1b[0m\n", 31, "Error: Operesheni imeshindikana boss.")
		fmt.Printf("\x1b[%dm%s\x1b[0m\n", 32, "\nTumia 'nuru' kuprogram\n\nAU\n\nTumia 'nuru' ikifuatiwa na jina la file.\n\n\tMfano:\tnuru fileYangu.nr")
		os.Exit(0)
	}
}
