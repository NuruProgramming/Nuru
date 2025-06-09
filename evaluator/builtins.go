package evaluator

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"encoding/base64"

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
	"namba": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("Samahani, namba inahitaji hoja 1, wewe umeweka %d", len(args))
			}
			value := args[0]
			return convertToInteger(value)
		},
	},
	"tungo": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("Samahani, tungo inahitaji hoja 1, wewe umeweka %d", len(args))
			}
			value := args[0]
			return convertToString(value)
		},
	},

	// Garbage collection functions
	"safishaMemori": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 0 {
				return newError("Samahani, safishaMemori haihitaji hoja, wewe umeweka %d", len(args))
			}
			object.RunGC()
			return &object.Null{}
		},
	},
	"takwimuMemori": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 0 {
				return newError("Samahani, takwimuMemori haihitaji hoja, wewe umeweka %d", len(args))
			}
			object.PrintObjectStats()
			return &object.Integer{Value: int64(object.GetObjectCount())}
		},
	},
	"kumbukumbaDhaifu": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("Samahani, kumbukumbaDhaifu inahitaji hoja 1, wewe umeweka %d", len(args))
			}
			return object.NewWeakReference(args[0])
		},
	},
	"takwimuMemoriKwa": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("Samahani, takwimuMemoriKwa inahitaji hoja 1, wewe umeweka %d", len(args))
			}

			if args[0].Type() != object.DICT_OBJ {
				return newError("Samahani, takwimuMemoriKwa inahitaji kamusi")
			}

			// Create a new dict to hold object stats
			result := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}

			// Get total count
			totalCount := object.GetObjectCount()
			totalKey := &object.String{Value: "jumla"}
			totalValue := &object.Integer{Value: int64(totalCount)}
			totalHash, _ := totalKey.HashKey(), true
			result.Pairs[totalHash] = object.DictPair{Key: totalKey, Value: totalValue}

			// Get type counts - use the provided function instead of accessing internal fields
			stats := getTypeCountStats()

			// Add type counts to result dict
			for objType, count := range stats {
				typeKey := &object.String{Value: string(objType)}
				typeValue := &object.Integer{Value: int64(count)}
				typeHash, _ := typeKey.HashKey(), true
				result.Pairs[typeHash] = object.DictPair{Key: typeKey, Value: typeValue}
			}

			return result
		},
	},

	// Base64 encoding/decoding functions
	"kodeBase64": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 || len(args) > 2 {
				return newError("Samahani, kodeBase64 inahitaji hoja 1 au 2, wewe umeweka %d", len(args))
			}

			var data []byte
			var urlSafe bool

			// Check for urlsafe option
			if len(args) == 2 {
				if s, ok := args[1].(*object.String); ok && s.Value == "urlsafe" {
					urlSafe = true
				} else {
					return newError("Chaguo la pili lazima liwe 'urlsafe' kama unataka URL-safe encoding")
				}
			}

			switch arg := args[0].(type) {
			case *object.String:
				data = []byte(arg.Value)
			case *object.Byte:
				data = arg.Value
			case *object.Array:
				// Support encoding array of bytes/integers
				arr := arg.Elements
				data = make([]byte, len(arr))
				for i, v := range arr {
					switch vv := v.(type) {
					case *object.Integer:
						if vv.Value < 0 || vv.Value > 255 {
							return newError("Thamani ya orodha lazima iwe kati ya 0 na 255 kwa byte encoding")
						}
						data[i] = byte(vv.Value)
					case *object.Byte:
						if len(vv.Value) != 1 {
							return newError("Kila Byte kwenye orodha lazima iwe na urefu wa 1")
						}
						data[i] = vv.Value[0]
					default:
						return newError("Orodha lazima iwe na namba au Byte pekee, imepata %s", v.Type())
					}
				}
			default:
				return newError("Samahani, kodeBase64 inatumika na Neno, Byte, au Orodha ya namba/Byte tu, umeweka %s", args[0].Type())
			}

			var encoded string
			if urlSafe {
				encoded = base64.URLEncoding.EncodeToString(data)
			} else {
				encoded = base64.StdEncoding.EncodeToString(data)
			}
			return &object.String{Value: encoded}
		},
	},

	"katuaBase64": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 || len(args) > 3 {
				return newError("Samahani, katuaBase64 inahitaji hoja 1 hadi 3, wewe umeweka %d", len(args))
			}

			if args[0].Type() != object.STRING_OBJ {
				return newError("Samahani, katuaBase64 inatumika na Neno pekee, umeweka %s", args[0].Type())
			}

			encoded := args[0].(*object.String).Value
			var urlSafe bool
			var asByte bool
			if len(args) > 1 {
				for i := 1; i < len(args); i++ {
					if s, ok := args[i].(*object.String); ok {
						switch s.Value {
						case "urlsafe":
							urlSafe = true
						case "byte":
							asByte = true
						case "array":
							// handled below
						default:
							return newError("Chaguo lisilojulikana: %s. Tumia 'urlsafe', 'byte', au 'array' tu.", s.Value)
						}
					} else {
						return newError("Chaguo la ziada lazima liwe Neno ('urlsafe', 'byte', 'array')")
					}
				}
			}

			var data []byte
			var err error
			if urlSafe {
				data, err = base64.URLEncoding.DecodeString(encoded)
			} else {
				data, err = base64.StdEncoding.DecodeString(encoded)
			}
			if err != nil {
				return &object.Error{Message: fmt.Sprintf("Shida wakati wa kuondoa usimbaji Base64: %s", err.Error())}
			}

			if asByte {
				return &object.Byte{Value: data, String: string(data)}
			}
			// If 'array' is specified, return as array of integers
			if len(args) > 1 {
				for i := 1; i < len(args); i++ {
					if s, ok := args[i].(*object.String); ok && s.Value == "array" {
						elements := make([]object.Object, len(data))
						for j, b := range data {
							elements[j] = &object.Integer{Value: int64(b)}
						}
						return &object.Array{Elements: elements}
					}
				}
			}
			return &object.String{Value: string(data)}
		},
	},
}

func getIntValue(obj object.Object) (int64, error) {
	switch obj := obj.(type) {
	case *object.Integer:
		return obj.Value, nil
	default:
		return 0, fmt.Errorf("expected integer, got %T", obj)
	}
}

// getTypeCountStats returns a map of object types to their counts
func getTypeCountStats() map[object.ObjectType]int {
	return object.GetTypeStats()
}
