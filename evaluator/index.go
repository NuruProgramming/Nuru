package evaluator

import "github.com/NuruProgramming/Nuru/object"

func evalIndexExpression(left, index object.Object, line int) object.Object {
	switch {
	case left.Type() == object.ARRAY_OBJ && index.Type() == object.INTEGER_OBJ:
		return evalArrayIndexExpression(left, index)
	case left.Type() == object.ARRAY_OBJ && index.Type() != object.INTEGER_OBJ:
		return newError("Mstari %d: Tafadhali tumia number, sio: %s", line, index.Type())
	case left.Type() == object.DICT_OBJ:
		return evalDictIndexExpression(left, index, line)
	default:
		return newError("Mstari %d: Operesheni hii haiwezekani kwa: %s", line, left.Type())
	}
}

func evalArrayIndexExpression(array, index object.Object) object.Object {
	arrayObject := array.(*object.Array)
	idx := index.(*object.Integer).Value
	max := int64(len(arrayObject.Elements) - 1)

	if idx < 0 || idx > max {
		return NULL
	}

	return arrayObject.Elements[idx]
}

func evalDictIndexExpression(dict, index object.Object, line int) object.Object {
	dictObject := dict.(*object.Dict)

	key, ok := index.(object.Hashable)
	if !ok {
		return newError("Mstari %d: Samahani, %s haitumiki kama ufunguo", line, index.Type())
	}

	pair, ok := dictObject.Pairs[key.HashKey()]
	if !ok {
		return NULL
	}

	return pair.Value
}
