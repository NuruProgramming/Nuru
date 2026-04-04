//go:build wasm && js

package main

import (
	"fmt"

	"github.com/NuruProgramming/Nuru/evaluator"
	"github.com/NuruProgramming/Nuru/lexer"
	"github.com/NuruProgramming/Nuru/object"
	"github.com/NuruProgramming/Nuru/parser"

	"syscall/js"
)

func Read(contents string) {
	jsOutputReceiverFunction := js.Global().Get("nuruOutputReceiver")

	env := object.NewEnvironment()

	l := lexer.New(contents)
	p := parser.New(l)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		fmt.Println("Kuna makosa yafuatayo:")
		jsOutputReceiverFunction.Invoke("Kuna makosa yafuatayo:", true)

		for _, msg := range p.Errors() {
			// fmt.Println("\t" + msg)
			jsOutputReceiverFunction.Invoke("\t" + msg, true)
		}

	}
	evaluated := evaluator.Eval(program, env)
	if evaluated != nil {
		if evaluated.Type() != object.NULL_OBJ {
			jsOutputReceiverFunction.Invoke(evaluated.Inspect(), true)
		}
	}

}

func runCode(this js.Value, args []js.Value) interface{} {
	code := args[0].String()
	Read(code)
	return nil
}

func main() {
	fmt.Println("Go WASM initialized")
	js.Global().Set("runCode", js.FuncOf(runCode))
	<-make(chan bool)
}
