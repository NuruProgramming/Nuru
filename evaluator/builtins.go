package evaluator

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/AvicennaJr/Nuru/object"
)

var builtins = map[string]*object.Builtin{
	"jaza": {
		Fn: func(args ...object.Object) object.Object {

			if len(args) > 1 {
				return newError("Samahani, kiendesha hiki kinapokea hoja 0 au 1, wewe umeweka %d", len(args))
			}

			if len(args) > 0 && args[0].Type() != object.STRING_OBJ {
				return newError(fmt.Sprintf(`Tafadhali tumia alama ya nukuu: "%s"`, args[0].Inspect()))
			}
			if len(args) == 1 {
				prompt := args[0].(*object.String).Value
				fmt.Fprint(os.Stdout, prompt)
			}

			buffer := bufio.NewReader(os.Stdin)

			line, _, err := buffer.ReadLine()
			if err != nil && err != io.EOF {
				return newError("Nimeshindwa kusoma uliyo yajaza")
			}

			return &object.String{Value: string(line)}
		},
	},
	"andika": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) == 0 {
				fmt.Println("")
			} else {
				var arr []string
				for _, arg := range args {
					if arg == nil {
						return newError("Hauwezi kufanya operesheni hii")
					}
					arr = append(arr, arg.Inspect())
				}
				str := strings.Join(arr, " ")
				fmt.Println(str)
			}
			return nil
		},
	},
	"_andika": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) == 0 {
				return &object.String{Value: "\n"}
			} else {
				var arr []string
				for _, arg := range args {
					if arg == nil {
						return newError("Hauwezi kufanya operesheni hii")
					}
					arr = append(arr, arg.Inspect())
				}
				str := strings.Join(arr, " ")
				return &object.String{Value: str}
			}
		},
	},
	"aina": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("Samahani, tunahitaji hoja 1, wewe umeweka %d", len(args))
			}

			return &object.String{Value: string(args[0].Type())}
		},
	},
	"fungua": {
		Fn: func(args ...object.Object) object.Object {

			if len(args) != 1 {
				return newError("Samahani, tunahitaji hoja 1, wewe umeweka %d", len(args))
			}
			filename := args[0].(*object.String).Value

			file, err := os.ReadFile(filename)
			if err != nil {
				return &object.Error{Message: "Tumeshindwa kusoma faili"}
			}
			return &object.File{Filename: filename, Content: string(file)}
		},
	},

	// "jumla": {
	// 	Fn: func(args ...object.Object) object.Object {
	// 		if len(args) != 1 {
	// 			return newError("Hoja hazilingani, tunahitaji=1, tumepewa=%d", len(args))
	// 		}

	// 		switch arg := args[0].(type) {
	// 		case *object.Array:

	// 			var sums float64
	// 			for _, num := range arg.Elements {

	// 				if num.Type() != object.INTEGER_OBJ && num.Type() != object.FLOAT_OBJ {
	// 					return newError("Samahani namba tu zinahitajika")
	// 				} else {
	// 					if num.Type() == object.INTEGER_OBJ {
	// 						no, _ := strconv.Atoi(num.Inspect())
	// 						floatnum := float64(no)
	// 						sums += floatnum
	// 					} else if num.Type() == object.FLOAT_OBJ {
	// 						no, _ := strconv.ParseFloat(num.Inspect(), 64)
	// 						sums += no
	// 					}

	// 				}
	// 			}

	// 			if math.Mod(sums, 1) == 0 {
	// 				return &object.Integer{Value: int64(sums)}
	// 			}

	// 			return &object.Float{Value: float64(sums)}

	// 		default:
	// 			return newError("Samahani, hii function haitumiki na %s", args[0].Type())
	// 		}
	// 	},
	// },
}
