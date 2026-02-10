package evaluator

import (
	"math"
	"strconv"

	"github.com/NuruProgramming/Nuru/object"
)

func convertToInteger(obj object.Object) object.Object {
	switch obj := obj.(type) {
	case *object.Integer:
		return obj
	case *object.BigInteger:
		if v, ok := obj.Int64(); ok {
			return &object.Integer{Value: v}
		}
		return newError("Namba kubwa haiwezi kubadilishwa kuwa NAMBA (inazidi kikomo)")
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
	case *object.BigInteger:
		if obj.Value == nil {
			return &object.Float{Value: 0}
		}
		f, _ := obj.Value.Float64()
		if math.IsInf(f, 0) || math.IsNaN(f) {
			return newError("Namba kubwa haiwezi kubadilishwa kuwa DESIMALI")
		}
		return &object.Float{Value: f}
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
	case *object.BigInteger:
		return &object.Boolean{Value: obj.Value != nil && obj.Value.Sign() != 0}
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

func convertToBigInteger(obj object.Object) object.Object {
	switch obj := obj.(type) {
	case *object.BigInteger:
		return obj
	case *object.Integer:
		return object.NewBigIntegerFromInt64(obj.Value)
	case *object.String:
		bi, ok := object.NewBigIntegerFromString(obj.Value)
		if !ok {
			return newError("Haiwezi kubadilisha '%s' kuwa NAMBA_KUBWA", obj.Value)
		}
		return bi
	default:
		return newError("Haiwezi kubadilisha %s kuwa NAMBA_KUBWA", obj.Type())
	}
}
