package evaluator

import (
	"strings"

	"github.com/NuruProgramming/Nuru/ast"
	"github.com/NuruProgramming/Nuru/object"
)

func evalAssignEqual(node *ast.AssignEqual, env *object.Environment) object.Object {
	left := Eval(node.Left, env)
	if isError(left) {
		return left
	}

	value := Eval(node.Value, env)
	if isError(value) {
		return value
	}

	switch node.Token.Literal {
	case "+=":
		switch arg := left.(type) {
		case *object.Integer:
			switch val := value.(type) {
			case *object.Integer:
				v := arg.Value + val.Value
				return env.Set(node.Left.Token.Literal, &object.Integer{Value: v})
			case *object.Float:
				v := float64(arg.Value) + val.Value
				return env.Set(node.Left.Token.Literal, &object.Float{Value: v})
			default:
				return newError("Mstari %d: Huwezi kutumia '+=' kujumlisha %v na %v", node.Token.Line, arg.Type(), val.Type())
			}
		case *object.Float:
			switch val := value.(type) {
			case *object.Integer:
				v := arg.Value + float64(val.Value)
				return env.Set(node.Left.Token.Literal, &object.Float{Value: v})
			case *object.Float:
				v := arg.Value + val.Value
				return env.Set(node.Left.Token.Literal, &object.Float{Value: v})
			default:
				return newError("Mstari %d: Huwezi kutumia '+=' kujumlisha %v na %v", node.Token.Line, arg.Type(), val.Type())
			}
		case *object.String:
			switch val := value.(type) {
			case *object.String:
				v := arg.Value + val.Value
				return env.Set(node.Left.Token.Literal, &object.String{Value: v})
			default:
				return newError("Mstari %d: Huwezi kutumia '+=' kwa %v na %v", node.Token.Line, arg.Type(), val.Type())
			}
		default:
			return newError("Mstari %d: Huwezi kutumia '+=' na %v", node.Token.Line, arg.Type())
		}
	case "-=":
		switch arg := left.(type) {
		case *object.Integer:
			switch val := value.(type) {
			case *object.Integer:
				v := arg.Value - val.Value
				return env.Set(node.Left.Token.Literal, &object.Integer{Value: v})
			case *object.Float:
				v := float64(arg.Value) - val.Value
				return env.Set(node.Left.Token.Literal, &object.Float{Value: v})
			default:
				return newError("Mstari %d: Huwezi kutumia '-=' kujumlisha %v na %v", node.Token.Line, arg.Type(), val.Type())
			}
		case *object.Float:
			switch val := value.(type) {
			case *object.Integer:
				v := arg.Value - float64(val.Value)
				return env.Set(node.Left.Token.Literal, &object.Float{Value: v})
			case *object.Float:
				v := arg.Value - val.Value
				return env.Set(node.Left.Token.Literal, &object.Float{Value: v})
			default:
				return newError("Mstari %d: Huwezi kutumia '-=' kujumlisha %v na %v", node.Token.Line, arg.Type(), val.Type())
			}
		default:
			return newError("Mstari %d: Huwezi kutumia '-=' na %v", node.Token.Line, arg.Type())
		}
	case "*=":
		switch arg := left.(type) {
		case *object.Integer:
			switch val := value.(type) {
			case *object.Integer:
				v := arg.Value * val.Value
				return env.Set(node.Left.Token.Literal, &object.Integer{Value: v})
			case *object.Float:
				v := float64(arg.Value) * val.Value
				return env.Set(node.Left.Token.Literal, &object.Float{Value: v})
			case *object.String:
				v := strings.Repeat(val.Value, int(arg.Value))
				return env.Set(node.Left.Token.Literal, &object.String{Value: v})
			default:
				return newError("Mstari %d: Huwezi kutumia '*=' kujumlisha %v na %v", node.Token.Line, arg.Type(), val.Type())
			}
		case *object.Float:
			switch val := value.(type) {
			case *object.Integer:
				v := arg.Value * float64(val.Value)
				return env.Set(node.Left.Token.Literal, &object.Float{Value: v})
			case *object.Float:
				v := arg.Value * val.Value
				return env.Set(node.Left.Token.Literal, &object.Float{Value: v})
			default:
				return newError("Mstari %d: Huwezi kutumia '*=' kujumlisha %v na %v", node.Token.Line, arg.Type(), val.Type())
			}
		case *object.String:
			switch val := value.(type) {
			case *object.Integer:
				v := strings.Repeat(arg.Value, int(val.Value))
				return env.Set(node.Left.Token.Literal, &object.String{Value: v})
			default:
				return newError("Mstari %d: Huwezi kutumia '+=' kwa %v na %v", node.Token.Line, arg.Type(), val.Type())
			}
		default:
			return newError("Mstari %d: Huwezi kutumia '*=' na %v", node.Token.Line, arg.Type())
		}
	case "/=":
		switch arg := left.(type) {
		case *object.Integer:
			switch val := value.(type) {
			case *object.Integer:
				v := arg.Value / val.Value
				return env.Set(node.Left.Token.Literal, &object.Integer{Value: v})
			case *object.Float:
				v := float64(arg.Value) / val.Value
				return env.Set(node.Left.Token.Literal, &object.Float{Value: v})
			default:
				return newError("Mstari %d: Huwezi kutumia '/=' kujumlisha %v na %v", node.Token.Line, arg.Type(), val.Type())
			}
		case *object.Float:
			switch val := value.(type) {
			case *object.Integer:
				v := arg.Value / float64(val.Value)
				return env.Set(node.Left.Token.Literal, &object.Float{Value: v})
			case *object.Float:
				v := arg.Value / val.Value
				return env.Set(node.Left.Token.Literal, &object.Float{Value: v})
			default:
				return newError("Mstari %d: Huwezi kutumia '/=' kujumlisha %v na %v", node.Token.Line, arg.Type(), val.Type())
			}
		default:
			return newError("Mstari %d: Huwezi kutumia '/=' na %v", node.Token.Line, arg.Type())
		}
	default:
		return newError("Mstari %d: Operesheni Haifahamiki  %s", node.Token.Line, node.Token.Literal)
	}
}
