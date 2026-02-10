package evaluator

import (
	"bufio"
	"fmt"
	"io"
	"os"
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
