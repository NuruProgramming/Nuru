package evaluator

import (
	"strings"

	"github.com/NuruProgramming/Nuru/object"
)

func evalInExpression(left, right object.Object, line int) object.Object {
	switch right.(type) {
	case *object.String:
		return evalInStringExpression(left, right)
	case *object.Array:
		return evalInArrayExpression(left, right)
	case *object.Dict:
		return evalInDictExpression(left, right, line)
	default:
		return FALSE
	}
}

func evalInStringExpression(left, right object.Object) object.Object {
	if left.Type() != object.STRING_OBJ {
		return FALSE
	}
	leftVal := left.(*object.String)
	rightVal := right.(*object.String)
	found := strings.Contains(rightVal.Value, leftVal.Value)
	return nativeBoolToBooleanObject(found)
}

func evalInDictExpression(left, right object.Object, line int) object.Object {
	leftVal, ok := left.(object.Hashable)
	if !ok {
		return newError("Mstari %d: Huwezi kutumia kama 'key': %s", line, left.Type())
	}
	key := leftVal.HashKey()
	rightVal := right.(*object.Dict).Pairs
	_, ok = rightVal[key]
	return nativeBoolToBooleanObject(ok)
}

func evalInArrayExpression(left, right object.Object) object.Object {
	rightVal := right.(*object.Array)
	switch leftVal := left.(type) {
	case *object.Null:
		for _, v := range rightVal.Elements {
			if v.Type() == object.NULL_OBJ {
				return TRUE
			}
		}
	case *object.String:
		for _, v := range rightVal.Elements {
			if v.Type() == object.STRING_OBJ {
				elem := v.(*object.String)
				if elem.Value == leftVal.Value {
					return TRUE
				}
			}
		}
	case *object.Integer:
		for _, v := range rightVal.Elements {
			if v.Type() == object.INTEGER_OBJ {
				elem := v.(*object.Integer)
				if elem.Value == leftVal.Value {
					return TRUE
				}
			}
		}
	case *object.Float:
		for _, v := range rightVal.Elements {
			if v.Type() == object.FLOAT_OBJ {
				elem := v.(*object.Float)
				if elem.Value == leftVal.Value {
					return TRUE
				}
			}
		}
	}
	return FALSE
}
