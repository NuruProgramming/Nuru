package evaluator

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/NuruProgramming/Nuru/object"
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
				return nil
			}

			if args[0].Type() != "NENO" {
				// Try and give a meaningful message instead of this generic error
				return newError(fmt.Sprintf("chaguo la kwanza linafaa kuwa neno, umepena %s", args[0].Type()))
			}

			str, err := formatStr(args[0], args[1:])
			if err != nil {
				return newError(err.Error())
			}

			fmt.Printf(str)

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
				return &object.Error{Message: "Tumeshindwa kusoma faili au faili halipo"}
			}
			return &object.File{Filename: filename, Content: string(file)}
		},
	},
	"mfululizo": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 || len(args) > 3 {
				return newError("Samahani, mfululizo inahitaji hoja 1 hadi 3, wewe umeweka %d", len(args))
			}

			var start, end, step int64
			var err error

			switch len(args) {
			case 1:
				end, err = getIntValue(args[0])
				if err != nil {
					return newError("Hoja lazima iwe nambari nzima")
				}
				start, step = 0, 1
			case 2:
				start, err = getIntValue(args[0])
				if err != nil {
					return newError("Hoja ya kwanza lazima iwe nambari nzima")
				}
				end, err = getIntValue(args[1])
				if err != nil {
					return newError("Hoja ya pili lazima iwe nambari nzima")
				}
				step = 1
			case 3:
				start, err = getIntValue(args[0])
				if err != nil {
					return newError("Hoja ya kwanza lazima iwe nambari nzima")
				}
				end, err = getIntValue(args[1])
				if err != nil {
					return newError("Hoja ya pili lazima iwe nambari nzima")
				}
				step, err = getIntValue(args[2])
				if err != nil {
					return newError("Hoja ya tatu lazima iwe nambari nzima")
				}
				if step == 0 {
					return newError("Hatua haiwezi kuwa sifuri")
				}
			}

			elements := []object.Object{}
			for i := start; (step > 0 && i < end) || (step < 0 && i > end); i += step {
				elements = append(elements, &object.Integer{Value: i})
			}

			return &object.Array{Elements: elements}
		},
	},

	"badilisha": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("Samahani, badili inahitaji hoja 2, wewe umeweka %d", len(args))
			}

			value := args[0]
			targetType := args[1]

			if targetType.Type() != object.STRING_OBJ {
				return newError("Aina ya lengo lazima iwe neno")
			}

			targetTypeStr := targetType.(*object.String).Value

			switch targetTypeStr {
			case "NAMBA":
				return convertToInteger(value)
			case "DESIMALI":
				return convertToFloat(value)
			case "NENO":
				return convertToString(value)
			case "BOOLEAN":
				return convertToBoolean(value)
			default:
				return newError("Aina isiyojulikana: %s", targetTypeStr)
			}
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

func getIntValue(obj object.Object) (int64, error) {
	switch obj := obj.(type) {
	case *object.Integer:
		return obj.Value, nil
	default:
		return 0, fmt.Errorf("expected integer, got %T", obj)
	}
}

// Format the string like "Hello {0}", jina ...
func formatStr(format object.Object, options []object.Object) (string, error) {
	var str strings.Builder
	var val strings.Builder
	var check_val bool
	var opts_len int = len(options)

	// This is to enable escaping the {} braces if you want them included
	var escapeChar bool

	type optM struct {
		val bool
		obj object.Object
	}

	var optionsMap = make(map[int]optM, opts_len)

	// convert the options into a map (may not be the most efficient but we are not going fast)
	for i, optm := range options {
		optionsMap[i] = optM{val: false, obj: optm}
	}

	// Now go through the format string and do the replacement(s)
	// this has approx time complexity of O(n) (bestest case)
	for _, opt := range format.Inspect() {

		if !escapeChar && opt == '\\' {
			escapeChar = true
			continue
		}

		if opt == '{' && !escapeChar {
			check_val = true
			continue
		}

		if escapeChar {
			if opt != '{' && opt != '}' {
				str.WriteRune('\\')
			}
			escapeChar = false
		}

		if check_val && opt == '}' {
			vstr := strings.TrimSpace(val.String()) // remove accidental spaces

			// check the value and try to convert to int
			arrv, err := strconv.Atoi(vstr)
			if err != nil {
				// return non-cryptic go errors
				return "", fmt.Errorf(fmt.Sprintf("Ulichopeana si NAMBA, jaribu tena: `%s'", vstr))
			}

			oVal, exists := optionsMap[arrv]

			if !exists {
				return "", fmt.Errorf(fmt.Sprintf("Nambari ya chaguo unalolitaka %d ni kubwa kuliko ulizopeana (%d)", arrv, opts_len))
			}

			str.WriteString(oVal.obj.Inspect())
			// may be innefficient
			optionsMap[arrv] = optM{val: true, obj: oVal.obj}

			check_val = false
			val.Reset()
			continue
		}

		if check_val {
			val.WriteRune(opt)
			continue
		}

		str.WriteRune(opt)
	}

	// A check if they never closed the formatting braces e.g '{0'
	if check_val {
		return "", fmt.Errorf(fmt.Sprintf("Haukufunga '{', tuliokota kabla ya kufika mwisho `%s'", val.String()))
	}

	// Another innefficient loop
	for _, v := range optionsMap {
		if !v.val {
			return "", fmt.Errorf(fmt.Sprintf("Ulipeana hili chaguo (%s) {%s} lakini haukutumia", v.obj.Inspect(), v.obj.Type()))
		}
	}

	// !start
	// Here we can talk about wtf just happened
	// 3 loops to do just formatting, formatting for codes sake.
	// it can be done in 2 loops, the last one is just to confirm that you didn't forget anything.
	// this is not required but a nice syntatic sugar, we are not here for speed, so ergonomics matter instead of speed
	// finally we can say we are slower than python!, what an achievement
	// done!

	return str.String(), nil
}
