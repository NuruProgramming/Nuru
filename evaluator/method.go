package evaluator

import (
	"github.com/AvicennaJr/Nuru/ast"
	"github.com/AvicennaJr/Nuru/object"
)

func evalMethodExpression(node *ast.MethodExpression, env *object.Environment) object.Object {
	obj := Eval(node.Object, env)
	if isError(obj) {
		return obj
	}
	args := evalExpressions(node.Arguments, env)
	if len(args) == 1 && isError(args[0]) {
		return args[0]
	}

	defs := make(map[string]object.Object)

	for k, v := range node.Defaults {
		defs[k] = Eval(v, env)
	}
	return applyMethod(obj, node.Method, args, defs)
}

func applyMethod(obj object.Object, method ast.Expression, args []object.Object, defs map[string]object.Object) object.Object {
	switch obj := obj.(type) {
	case *object.String:
		return obj.Method(method.(*ast.Identifier).Value, args)
	case *object.File:
		return obj.Method(method.(*ast.Identifier).Value, args)
	case *object.Time:
		return obj.Method(method.(*ast.Identifier).Value, args, defs)
	case *object.Array:
		return obj.Method(method.(*ast.Identifier).Value, args)
	case *object.Module:
		if fn, ok := obj.Functions[method.(*ast.Identifier).Value]; ok {
			return fn(args)
		}
	}
	return newError("Samahani, %s haina function '%s()'", obj.Inspect(), method.(*ast.Identifier).Value)
}
