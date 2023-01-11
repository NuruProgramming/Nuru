package evaluator

import (
	"github.com/AvicennaJr/Nuru/ast"
	"github.com/AvicennaJr/Nuru/module"
	"github.com/AvicennaJr/Nuru/object"
)

func evalImport(node *ast.Import, env *object.Environment) object.Object {
	for k, v := range node.Identifiers {
		if mod, ok := module.Mapper[v.Value]; ok {
			env.Set(k, mod)
		} else {
			return newError("Mstari %d: Module '%s' haipo", node.Token.Line, v.Value)
		}
	}

	return NULL
}
