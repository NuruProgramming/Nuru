package evaluator

import (
	"github.com/AvicennaJr/Nuru/object"
)

func evalDotExpression(left, right object.Object, line int) object.Object {
	// if right.Type() != object.BUILTIN_OBJ {
	// 	return newError("Mstari %d: '%s' sio function sahihi", line, right.Inspect())
	// }
	switch left.(type) {
	case *object.String:
		return evalDotStringExpression(left, right, line)
	case *object.Array:
		return evalDotArrayExpression(left, right)
	case *object.Dict:
		return evalDotDictExpression(left, right, line)
	default:
		return FALSE
	}
}

func evalDotStringExpression(left, right object.Object, line int) object.Object {
	leftVal := left.(*object.String)
	rightVal := right.(*object.String)

	switch rightVal.Inspect() {
	case "idadi":
		return &object.Integer{Value: int64(len(leftVal.Value))}
	}
	return nil
}

func evalDotDictExpression(left, right object.Object, line int) object.Object {
	// to do

	return nil
}

func evalDotArrayExpression(left, right object.Object) object.Object {
	// to do

	return nil
}
