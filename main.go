package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/AvicennaJr/Nuru/repl"
)

func main() {

	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Habari %s!, Karibu utumie lugha ya Nuru!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
