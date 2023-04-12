package module

import (
	"encoding/json"

	"github.com/AvicennaJr/Nuru/object"
)

var JsonFunctions = map[string]object.ModuleFunction{}

func init() {
	JsonFunctions["decode"] = decode
	JsonFunctions["encode"] = encode
}

func decode(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "Hoja hii hairuhusiwi"}
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
	input := args[0].Inspect()

	jsonBody, err := json.Marshal(input)

	if err != nil {
		return &object.Error{Message: "Hii data haiwezi kua jsoni"}
	}

	return &object.Byte{String: string(jsonBody), Value: jsonBody}
}
