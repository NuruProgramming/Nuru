package evaluator

import (
	"github.com/NuruProgramming/Nuru/ast"
	"github.com/NuruProgramming/Nuru/object"
)

func evalAt(node *ast.At, env *object.Environment) object.Object {
	if at, ok := env.Get("@"); ok {
		return at
	}
	return newError("Iko nje ya scope")
}
