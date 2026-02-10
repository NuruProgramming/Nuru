package evaluator

import (
	"github.com/NuruProgramming/Nuru/ast"
	"github.com/NuruProgramming/Nuru/object"
)

func evalWhileExpression(we *ast.WhileExpression, env *object.Environment) object.Object {
	var evaluated object.Object

	for {
		condition := Eval(we.Condition, env)
		if isError(condition) {
			return condition
		}

		if !isTruthy(condition) {
			break
		}

		evaluated = Eval(we.Consequence, env)
		if isError(evaluated) {
			return evaluated
		}

		if evaluated != nil && evaluated.Type() == object.BREAK_OBJ {
			return NULL
		}
	}

	// Don't expose Break/Continue as the loop value (align with for-in)
	if evaluated != nil && (evaluated.Type() == object.BREAK_OBJ || evaluated.Type() == object.CONTINUE_OBJ) {
		return NULL
	}
	return evaluated
}
