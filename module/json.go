package module

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/NuruProgramming/Nuru/object"
)

var JsonFunctions = map[string]object.ModuleFunction{}

func init() {
	JsonFunctions["dikodi"] = decode
	JsonFunctions["enkodi"] = encode
	JsonFunctions["soma"] = load
	JsonFunctions["hifadhi"] = dump
	JsonFunctions["pendeza"] = prettyPrint
	JsonFunctions["msailiaji"] = jsonEncoder
	JsonFunctions["msailiaji_bora"] = jsonEncoderBetter
}

func decode(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 1 {
		return &object.Error{Message: "dikodi inahitaji hoja ya kwanza (JSON string)"}
	}

	if args[0].Type() != object.STRING_OBJ {
		return &object.Error{Message: "Hoja lazima iwe neno"}
	}

	input := args[0].(*object.String).Value

	if parseFloatObj, ok := defs["parse_float"]; ok {
		if _, ok := parseFloatObj.(*object.Function); ok {
			// In a real implementation, we would use this function
			// But for now, we keep the default behavior
		}
	}

	if parseIntObj, ok := defs["parse_int"]; ok {
		if _, ok := parseIntObj.(*object.Function); ok {
			// In a real implementation, we would use this function
			// But for now, we keep the default behavior
		}
	}

	if objectHookObj, ok := defs["object_hook"]; ok {
		if _, ok := objectHookObj.(*object.Function); ok {
			// In a real implementation, we would use this function
			// But for now, we keep the default behavior
		}
	}

	var i interface{}
	err := json.Unmarshal([]byte(input), &i)
	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Hitilafu kuchanganua JSON: %s", err)}
	}

	return convertWhateverToObject(i)
}

func convertWhateverToObject(i interface{}) object.Object {
	switch v := i.(type) {
	case map[string]interface{}:
		dict := &object.Dict{}
		dict.Pairs = make(map[object.HashKey]object.DictPair)

		for k, v := range v {
			pair := object.DictPair{
				Key:   &object.String{Value: k},
				Value: convertWhateverToObject(v),
			}
			dict.Pairs[pair.Key.(object.Hashable).HashKey()] = pair
		}

		return dict
	case []interface{}:
		list := &object.Array{}
		for _, e := range v {
			list.Elements = append(list.Elements, convertWhateverToObject(e))
		}

		return list
	case string:
		return &object.String{Value: v}
	case int64:
		return &object.Integer{Value: v}
	case float64:
		if v == float64(int64(v)) {
			return &object.Integer{Value: int64(v)}
		}
		return &object.Float{Value: v}
	case bool:
		return &object.Boolean{Value: v}
	case nil:
		return &object.Null{}
	}
	return &object.Null{}
}

func encode(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 1 {
		return &object.Error{Message: "enkodi inahitaji hoja ya kwanza (object)"}
	}

	input := args[0]
	i := convertObjectToWhatever(input)
	data, err := json.Marshal(i)

	if err != nil {
		return &object.Error{Message: fmt.Sprintf("Hitilafu kutuma JSON: %s", err)}
	}

	return &object.String{Value: string(data)}
}

func convertObjectToWhatever(obj object.Object) interface{} {
	switch v := obj.(type) {
	case *object.Dict:
		m := make(map[string]interface{})
		for _, pair := range v.Pairs {
			key := pair.Key.(*object.String).Value
			m[key] = convertObjectToWhatever(pair.Value)
		}
		return m
	case *object.Array:
		list := make([]interface{}, len(v.Elements))
		for i, e := range v.Elements {
			list[i] = convertObjectToWhatever(e)
		}
		return list
	case *object.String:
		return v.Value
	case *object.Integer:
		return v.Value
	case *object.Float:
		return v.Value
	case *object.Boolean:
		return v.Value
	case *object.Null:
		return nil
	}
	return nil
}

func load(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 1 {
		return &object.Error{Message: "soma inahitaji hoja ya kwanza (filename)"}
	}

	filename, ok := args[0].(*object.String)
	if !ok {
		return &object.Error{Message: "Hoja ya kwanza lazima iwe jina la faili (neno)"}
	}

	if FSFunctions == nil {
		return &object.Error{Message: "Moduli ya 'faili' haipatikani"}
	}

	readArgs := []object.Object{filename}
	fileContent := readFile(readArgs, nil)

	if fileContent.Type() == object.ERROR_OBJ {
		return fileContent
	}

	decodeArgs := []object.Object{fileContent}
	return decode(decodeArgs, defs)
}

func dump(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 2 {
		return &object.Error{Message: "hifadhi inahitaji hoja mbili (object, filename)"}
	}

	obj := args[0]
	filename, ok := args[1].(*object.String)
	if !ok {
		return &object.Error{Message: "Hoja ya pili lazima iwe jina la faili (neno)"}
	}

	encodeArgs := []object.Object{obj}
	encodedJSON := encode(encodeArgs, defs)

	if encodedJSON.Type() == object.ERROR_OBJ {
		return encodedJSON
	}

	if FSFunctions == nil {
		return &object.Error{Message: "Moduli ya 'faili' haipatikani"}
	}

	writeArgs := []object.Object{filename, encodedJSON}
	return writeFile(writeArgs, nil)
}

func prettyPrint(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 1 {
		return &object.Error{Message: "pendeza inahitaji hoja ya kwanza (JSON string or object)"}
	}

	indent := 4
	sortKeys := false

	if indentObj, ok := defs["indent"]; ok {
		if intObj, ok := indentObj.(*object.Integer); ok {
			indent = int(intObj.Value)
		}
	}

	if sortKeysObj, ok := defs["sort_keys"]; ok {
		if boolObj, ok := sortKeysObj.(*object.Boolean); ok {
			sortKeys = boolObj.Value
		}
	}

	encodeDefs := map[string]object.Object{
		"indent":    &object.Integer{Value: int64(indent)},
		"sort_keys": &object.Boolean{Value: sortKeys},
	}

	var jsonObj object.Object
	switch arg := args[0].(type) {
	case *object.String:
		decodeArgs := []object.Object{arg}
		jsonObj = decode(decodeArgs, nil)
		if jsonObj.Type() == object.ERROR_OBJ {
			return jsonObj
		}
	default:
		jsonObj = arg
	}

	encodeArgs := []object.Object{jsonObj}
	return encode(encodeArgs, encodeDefs)
}

func jsonEncoder(args []object.Object, defs map[string]object.Object) object.Object {
	encoder := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}

	skipKeys := false
	ensureAscii := true
	indent := ""
	sortKeys := false
	separators := []string{", ", ": "}

	if len(args) > 0 {
		if options, ok := args[0].(*object.Dict); ok {
			if skipKeysObj, ok := options.Get(&object.String{Value: "skipkeys"}); ok {
				if boolObj, ok := skipKeysObj.(*object.Boolean); ok {
					skipKeys = boolObj.Value
				}
			}

			if ensureAsciiObj, ok := options.Get(&object.String{Value: "ensure_ascii"}); ok {
				if boolObj, ok := ensureAsciiObj.(*object.Boolean); ok {
					ensureAscii = boolObj.Value
				}
			}

			if indentObj, ok := options.Get(&object.String{Value: "indent"}); ok {
				switch i := indentObj.(type) {
				case *object.Integer:
					spaces := int(i.Value)
					if spaces > 0 {
						indent = strings.Repeat(" ", spaces)
					}
				case *object.String:
					indent = i.Value
				}
			}

			if sortKeysObj, ok := options.Get(&object.String{Value: "sort_keys"}); ok {
				if boolObj, ok := sortKeysObj.(*object.Boolean); ok {
					sortKeys = boolObj.Value
				}
			}

			if separatorsObj, ok := options.Get(&object.String{Value: "separators"}); ok {
				if sepArray, ok := separatorsObj.(*object.Array); ok && len(sepArray.Elements) == 2 {
					if itemSep, ok := sepArray.Elements[0].(*object.String); ok {
						if keySep, ok := sepArray.Elements[1].(*object.String); ok {
							separators = []string{itemSep.Value, keySep.Value}
						}
					}
				}
			}
		}
	}

	encoder.Set(&object.String{Value: "skipkeys"}, &object.Boolean{Value: skipKeys})
	encoder.Set(&object.String{Value: "ensure_ascii"}, &object.Boolean{Value: ensureAscii})
	encoder.Set(&object.String{Value: "indent"}, &object.String{Value: indent})
	encoder.Set(&object.String{Value: "sort_keys"}, &object.Boolean{Value: sortKeys})

	sepArray := &object.Array{Elements: []object.Object{
		&object.String{Value: separators[0]},
		&object.String{Value: separators[1]},
	}}
	encoder.Set(&object.String{Value: "separators"}, sepArray)

	encoder.Set(&object.String{Value: "enkodi"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 {
				return &object.Error{Message: "enkodi inahitaji hoja ya kwanza (object)"}
			}

			defs := map[string]object.Object{}

			if skipKeys, ok := encoder.Get(&object.String{Value: "skipkeys"}); ok {
				defs["skipkeys"] = skipKeys
			}

			if ensureAscii, ok := encoder.Get(&object.String{Value: "ensure_ascii"}); ok {
				defs["ensure_ascii"] = ensureAscii
			}

			if indent, ok := encoder.Get(&object.String{Value: "indent"}); ok {
				defs["indent"] = indent
			}

			if sortKeys, ok := encoder.Get(&object.String{Value: "sort_keys"}); ok {
				defs["sort_keys"] = sortKeys
			}

			if separators, ok := encoder.Get(&object.String{Value: "separators"}); ok {
				defs["separators"] = separators
			}

			return encode([]object.Object{args[0]}, defs)
		},
	})

	encoder.Set(&object.String{Value: "iterenkodi"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 {
				return &object.Error{Message: "iterenkodi inahitaji hoja ya kwanza (object)"}
			}

			encoded := encoder.Pairs[(&object.String{Value: "enkodi"}).HashKey()].Value.(*object.Builtin).Fn(args[0])
			if encoded.Type() == object.ERROR_OBJ {
				return encoded
			}

			return &object.Array{Elements: []object.Object{encoded}}
		},
	})

	encoder.Set(&object.String{Value: "default"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 {
				return &object.Error{Message: "default inahitaji hoja ya kwanza (object)"}
			}

			return &object.Error{Message: fmt.Sprintf("Object %s haiwezi kuwa JSON", args[0].Type())}
		},
	})

	return encoder
}

func jsonEncoderBetter(args []object.Object, defs map[string]object.Object) object.Object {
	encoder := jsonEncoder(args, defs).(*object.Dict)

	encoder.Set(&object.String{Value: "default"}, &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) < 1 {
				return &object.Error{Message: "default inahitaji hoja ya kwanza (object)"}
			}

			return &object.Error{Message: fmt.Sprintf("Object %s haiwezi kuwa JSON", args[0].Type())}
		},
	})

	return encoder
}
