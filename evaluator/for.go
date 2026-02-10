package evaluator

import (
	"github.com/NuruProgramming/Nuru/ast"
	"github.com/NuruProgramming/Nuru/object"
)

func evalForExpression(f *ast.For, env *object.Environment) object.Object {
	// Init: evaluate starter value and set loop variable
	initVal := Eval(f.StarterValue, env)
	if isError(initVal) {
		return initVal
	}
	env.Set(f.Identifier, initVal)

	var last object.Object = NULL
	for {
		// Condition
		cond := Eval(f.Condition, env)
		if isError(cond) {
			return cond
		}
		if !isTruthy(cond) {
			break
		}

		// Block
		last = Eval(f.Block, env)
		if isError(last) {
			return last
		}
		if last != nil {
			switch last.Type() {
			case object.BREAK_OBJ:
				return NULL
			case object.CONTINUE_OBJ:
				// fall through to update
				last = NULL
			case object.RETURN_VALUE_OBJ:
				return last
			}
		}

		// Update (closer)
		_ = Eval(f.Closer, env)
		// Re-read loop variable in case closer was e.g. i = i + 1 (assignment is in env)
	}

	return last
}
