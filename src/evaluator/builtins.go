package evaluator

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/AvicennaJr/Nuru/object"
)

var builtins = map[string]*object.Builtin{
	"idadi": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("Hoja hazilingani, tunahitaji=1, tumepewa=%d", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("Samahani, hii function haitumiki na %s", args[0].Type())
			}
		},
	},
	"jumla": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("Hoja hazilingani, tunahitaji=1, tumepewa=%d", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Array:
				
				var sums float64
				for _,num := range arg.Elements {
				   
                   if num.Type() != object.INTEGER_OBJ && num.Type() != object.FLOAT_OBJ{
					  return newError("Samahani namba tu zinahitajika")
				   }else{
					if num.Type() == object.INTEGER_OBJ{
						no , _ := strconv.Atoi(num.Inspect())
						floatnum := float64(no)
						sums += floatnum
					}else if num.Type() == object.FLOAT_OBJ {
                        no , _ := strconv.ParseFloat(num.Inspect(), 64)
                        sums += no
					}
					   
					  
				   } 
				}

				if math.Mod(sums,1) == 0 {
					return &object.Integer{Value : int64(sums)}
				}

				return &object.Float {Value: float64(sums)}
			
			default:
				return newError("Samahani, hii function haitumiki na %s", args[0].Type())
			}
		},
	},
	"yamwisho": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("Samahani, tunahitaji Hoja moja tu, wewe umeweka %d", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("Samahani, hii function haitumiki na %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				return arr.Elements[length-1]
			}

			return NULL
		},
	},
	"sukuma": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("Samahani, tunahitaji Hoja 2, wewe umeweka %d", len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("Samahani, hii function haitumiki na %s", args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			newElements := make([]object.Object, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]

			return &object.Array{Elements: newElements}
		},
	},
	"jaza": {
		Fn: func(args ...object.Object) object.Object {

			if len(args) > 1 {
				return newError("Samahani, hii function inapokea hoja 0 au 1, wewe umeweka %d", len(args))
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
				print(str + "\n")
			}
			return nil
		},
	},
	"aina": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("Samahani, tunahitaji Hoja 1, wewe umeweka %d", len(args))
			}

			return &object.String{Value: string(args[0].Type())}
		},
	},
}
