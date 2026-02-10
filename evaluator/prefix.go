package evaluator

import (
	"math/big"

	"github.com/NuruProgramming/Nuru/object"
)

func evalMinusPrefixOperatorExpression(right object.Object, line int) object.Object {
	switch obj := right.(type) {
	case *object.Integer:
		return &object.Integer{Value: -obj.Value}
	case *object.Float:
		return &object.Float{Value: -obj.Value}
	case *object.BigInteger:
		if obj.Value == nil {
			return &object.BigInteger{Value: big.NewInt(0)}
		}
		return &object.BigInteger{Value: new(big.Int).Neg(obj.Value)}
	default:
		return newError("Mstari %d: Operesheni Haieleweki: -%s", line, right.Type())
	}
}
func evalPlusPrefixOperatorExpression(right object.Object, line int) object.Object {
	switch obj := right.(type) {
	case *object.Integer:
		return &object.Integer{Value: obj.Value}
	case *object.Float:
		return &object.Float{Value: obj.Value}
	case *object.BigInteger:
		return obj
	default:
		return newError("Mstari %d: Operesheni Haieleweki: +%s", line, right.Type())
	}
}

func evalPrefixExpression(operator string, right object.Object, line int) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right, line)
	case "+":
		return evalPlusPrefixOperatorExpression(right, line)
	default:
		return newError("Mstari %d: Operesheni Haieleweki: %s%s", line, operator, right.Type())
	}
}
