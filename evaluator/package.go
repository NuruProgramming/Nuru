package evaluator

import (
	"github.com/NuruProgramming/Nuru/ast"
	"github.com/NuruProgramming/Nuru/object"
)

func evalPackage(node *ast.Package, env *object.Environment) object.Object {
	pakeji := &object.Package{
		Name:  node.Name,
		Env:   env,
		Scope: object.NewEnclosedEnvironment(env),
	}

	Eval(node.Block, pakeji.Scope)
	env.Set(node.Name.Value, pakeji)
	return pakeji
}
