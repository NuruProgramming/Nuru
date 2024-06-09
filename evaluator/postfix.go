package evaluator

import (
	"github.com/NuruProgramming/Nuru/ast"
	"github.com/NuruProgramming/Nuru/object"
)

func evalPostfixExpression(env *object.Environment, operator string, node *ast.PostfixExpression) object.Object {
	val, ok := env.Get(node.Token.Literal)
	if !ok {
		return newError("Tumia KITAMBULISHI CHA NAMBA AU DESIMALI, sio %s", node.Token.Type)
	}
	switch operator {
	case "++":
		switch arg := val.(type) {
		case *object.Integer:
			v := arg.Value + 1
			return env.Set(node.Token.Literal, &object.Integer{Value: v})
		case *object.Float:
			v := arg.Value + 1
			return env.Set(node.Token.Literal, &object.Float{Value: v})
		default:
			return newError("Mstari %d: %s sio kitambulishi cha namba. Tumia '++' na kitambulishi cha namba au desimali.\nMfano:\tfanya i = 2; i++", node.Token.Line, node.Token.Literal)

		}
	case "--":
		switch arg := val.(type) {
		case *object.Integer:
			v := arg.Value - 1
			return env.Set(node.Token.Literal, &object.Integer{Value: v})
		case *object.Float:
			v := arg.Value - 1
			return env.Set(node.Token.Literal, &object.Float{Value: v})
		default:
			return newError("Mstari %d: %s sio kitambulishi cha namba. Tumia '--' na kitambulishi cha namba au desimali.\nMfano:\tfanya i = 2; i++", node.Token.Line, node.Token.Literal)
		}
	default:
		return newError("Haifahamiki: %s", operator)
	}
}
