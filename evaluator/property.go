package evaluator

import (
	"github.com/NuruProgramming/Nuru/ast"
	"github.com/NuruProgramming/Nuru/object"
)

func evalPropertyExpression(node *ast.PropertyExpression, env *object.Environment) object.Object {
	left := Eval(node.Object, env)
	if isError(left) {
		return left
	}
	switch left.(type) {
	case *object.Instance:
		obj := left.(*object.Instance)
		prop := node.Property.(*ast.Identifier).Value
		if val, ok := obj.Env.Get(prop); ok {
			return val
		}
	case *object.Package:
		obj := left.(*object.Package)
		prop := node.Property.(*ast.Identifier).Value
		if val, ok := obj.Env.Get(prop); ok {
			return val
			}
	case *object.Dict:
		dict := left.(*object.Dict)
		prop := node.Property.(*ast.Identifier).Value
		key := &object.String{Value: prop}
		if val, ok := dict.Get(key); ok {
			return val
		}
	}
	return newError("Value %s sii sahihi kwenye %s", node.Property.(*ast.Identifier).Value, left.Inspect())
}

func evalPropertyAssignment(name *ast.PropertyExpression, val object.Object, env *object.Environment) object.Object {
	left := Eval(name.Object, env)
	if isError(left) {
		return left
	}
	switch left.(type) {
	case *object.Instance:
		obj := left.(*object.Instance)
		prop := name.Property.(*ast.Identifier).Value
		if _, ok := obj.Env.Get(prop); ok {
			obj.Env.Set(prop, val)
			return NULL
		}
		obj.Env.Set(prop, val)
		return NULL
	case *object.Package:
		obj := left.(*object.Package)
		prop := name.Property.(*ast.Identifier).Value
		if _, ok := obj.Env.Get(prop); ok {
			obj.Env.Set(prop, val)
			return NULL
		}
		obj.Env.Set(prop, val)
		return NULL
	default:
		return newError("Imeshindikana kuweka kwenye pakiti %s", left.Type())
	}
}
