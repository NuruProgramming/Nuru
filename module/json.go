package module

import (
	"encoding/json"

	"github.com/NuruProgramming/Nuru/object"
)

var JsonFunctions = map[string]object.ModuleFunction{}

func init() {
	JsonFunctions["dikodi"] = decode
	JsonFunctions["enkodi"] = encode
}

func decode(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "Hoja hii hairuhusiwi"}
	}
	if len(args) != 1 {
		return &object.Error{Message: "Tunahitaji hoja moja tu"}
	}

	if args[0].Type() != object.STRING_OBJ {
		return &object.Error{Message: "Hoja lazima iwe neno"}
	}

	var i interface{}

	input := args[0].(*object.String).Value
	err := json.Unmarshal([]byte(input), &i)
	if err != nil {
		return &object.Error{Message: "Hii data sio jsoni"}
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
		return &object.Float{Value: v}
	case bool:
		if v {
			return &object.Boolean{Value: true}
		} else {
			return &object.Boolean{Value: false}
		}
	}
	return &object.Null{}
}

func encode(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "Hoja hii hairuhusiwi"}
	}

	input := args[0]
	i := convertObjectToWhatever(input)
	data, err := json.Marshal(i)

	if err != nil {
		return &object.Error{Message: "Siwezi kubadilisha data hii kuwa jsoni"}
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
