package evaluator

import (
	"github.com/NuruProgramming/Nuru/ast"
	"github.com/NuruProgramming/Nuru/object"
)

func evalAssign(node *ast.Assign, env *object.Environment) object.Object {
	// Preserve the name of the Function for use in later cases
	switch node.Value.(type) {
	case *ast.FunctionLiteral:
		node.Value.(*ast.FunctionLiteral).Name = node.Name.Value
	}

	val := Eval(node.Value, env)
	if isError(val) {
		return val
	}

	obj := env.Set(node.Name.Value, val)
	return obj
}
