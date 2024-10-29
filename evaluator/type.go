package evaluator

import (
	"strconv"

	"github.com/NuruProgramming/Nuru/object"
)

func convertToInteger(obj object.Object) object.Object {
	switch obj := obj.(type) {
	case *object.Integer:
		return obj
	case *object.Float:
		return &object.Integer{Value: int64(obj.Value)}
	case *object.String:
		i, err := strconv.ParseInt(obj.Value, 10, 64)
		if err != nil {
			return newError("Haiwezi kubadilisha '%s' kuwa NAMBA", obj.Value)
		}
		return &object.Integer{Value: i}
	case *object.Boolean:
		if obj.Value {
			return &object.Integer{Value: 1}
		}
		return &object.Integer{Value: 0}
	default:
		return newError("Haiwezi kubadilisha %s kuwa NAMBA", obj.Type())
	}
}

func convertToFloat(obj object.Object) object.Object {
	switch obj := obj.(type) {
	case *object.Float:
		return obj
	case *object.Integer:
		return &object.Float{Value: float64(obj.Value)}
	case *object.String:
		f, err := strconv.ParseFloat(obj.Value, 64)
		if err != nil {
			return newError("Haiwezi kubadilisha '%s' kuwa DESIMALI", obj.Value)
		}
		return &object.Float{Value: f}
	case *object.Boolean:
		if obj.Value {
			return &object.Float{Value: 1.0}
		}
		return &object.Float{Value: 0.0}
	default:
		return newError("Haiwezi kubadilisha %s kuwa DESIMALI", obj.Type())
	}
}

func convertToString(obj object.Object) object.Object {
	return &object.String{Value: obj.Inspect()}
}

func convertToBoolean(obj object.Object) object.Object {
	switch obj := obj.(type) {
	case *object.Boolean:
		return obj
	case *object.Integer:
		return &object.Boolean{Value: obj.Value != 0}
	case *object.Float:
		return &object.Boolean{Value: obj.Value != 0}
	case *object.String:
		return &object.Boolean{Value: len(obj.Value) > 0}
	case *object.Null:
		return &object.Boolean{Value: false}
	default:
		return &object.Boolean{Value: true}
	}
}
