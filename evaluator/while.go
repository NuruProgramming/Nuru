package evaluator

import (
	"github.com/NuruProgramming/Nuru/ast"
	"github.com/NuruProgramming/Nuru/object"
)

const MAX_ITERATIONS = 1_000_000

func evalWhileExpression(we *ast.WhileExpression, env *object.Environment) object.Object {
	var evaluated object.Object
	iterations := 0

	for {
		iterations++
		if iterations > MAX_ITERATIONS {
			return newError("mzunguko usio na mwisho umegunduliwa")
		}

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
			return evaluated
		}
	}

	return evaluated
}
