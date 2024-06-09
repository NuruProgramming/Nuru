package evaluator

import (
	"github.com/NuruProgramming/Nuru/ast"
	"github.com/NuruProgramming/Nuru/object"
)

func evalWhileExpression(we *ast.WhileExpression, env *object.Environment) object.Object {
	condition := Eval(we.Condition, env)
	var evaluated object.Object
	if isError(condition) {
		return condition
	}
	if isTruthy(condition) {
		evaluated = Eval(we.Consequence, env)
		if isError(evaluated) {
			return evaluated
		}
		if evaluated != nil && evaluated.Type() == object.BREAK_OBJ {
			return evaluated
		}
		evaluated = evalWhileExpression(we, env)
	}
	return evaluated
}
