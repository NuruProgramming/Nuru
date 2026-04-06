//go:build wasm && js

// Modified version with of the builtins.go file with browser friendly functions.
package evaluator

import (
	"fmt"
	"strings"
	"github.com/NuruProgramming/Nuru/object"
	"syscall/js"
)

var builtins = map[string]*object.Builtin{
	"jaza": {
		Fn: func(args ...object.Object) object.Object {

			if len(args) > 1 {
				return newError("Samahani, kiendesha hiki kinapokea hoja 0 au 1, wewe umeweka %d", len(args))
			}

			if len(args) == 1 && args[0].Type() != object.STRING_OBJ {
				return newError(fmt.Sprintf(`Tafadhali tumia alama ya nukuu: "%s"`, args[0].Inspect()))
			}

			// Get the window.prompt function
			jsPromptFunction := js.Global().Get("prompt")
			if jsPromptFunction.Type() != js.TypeFunction {
				return newError("prompt function not found")
			}

			// invoke it!!
			var result js.Value
			if len(args) == 0 {
				result = jsPromptFunction.Invoke()
			} else {
				result = jsPromptFunction.Invoke(args[0].Inspect())
			}

			// fmt.Println("the arguments", args[0].Inspect())
			// fmt.Println("the result of window.prompt", result.String())

			// buffer := bufio.NewReader(os.Stdin)

			// line, _, err := buffer.ReadLine()

			if result.String() == ""|| result.String() == "null" {
				return newError("Nimeshindwa kusoma uliyo yajaza")
			}

			return &object.String{Value: string(result.String())}
		},
	},
	"andika": {
		Fn: func(args ...object.Object) object.Object {
			jsOutputReceiverFunction := js.Global().Get("nuruOutputReceiver")
			if len(args) == 0 {
				jsOutputReceiverFunction.Invoke("")
			} else {
				var arr []string
				for _, arg := range args {
					if arg == nil {
						return newError("Hauwezi kufanya operesheni hii")
					}
					arr = append(arr, arg.Inspect())
				}
				str := strings.Join(arr, " ")
				// fmt.Println(str)  // removed Print to console
				// return &object.String{Value: str} // return the output
				jsOutputReceiverFunction.Invoke(str)	
			}
			return nil
		},
	},
}

func init(){
	for name, builtin := range commonBuiltins{
		builtins[name]=builtin
	}
}