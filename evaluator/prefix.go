package evaluator

import "github.com/NuruProgramming/Nuru/object"

func evalMinusPrefixOperatorExpression(right object.Object, line int) object.Object {
	switch obj := right.(type) {

	case *object.Integer:
		return &object.Integer{Value: -obj.Value}

	case *object.Float:
		return &object.Float{Value: -obj.Value}

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
	case "*":
		if right == nil {
			return newError("Line %d: cannot dereference nil", line)
		}
		if p, ok := right.(*object.Pointer); ok {
			if p.Ref == nil {
				return newError("Line %d: pointer is nil", line)
			}
			return p.Ref
		}
		return newError("Line %d: cannot dereference non-pointer", line)
	case "&":
		if right == nil {
			return newError("Line %d: cannot take address of nil", line)
		}
		return &object.Pointer{Ref: right}
	default:
		return newError("Mstari %d: Operesheni Haieleweki: %s%s", line, operator, right.Type())
	}
}
