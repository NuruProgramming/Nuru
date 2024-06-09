package evaluator

import (
	"github.com/NuruProgramming/Nuru/ast"
	"github.com/NuruProgramming/Nuru/object"
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
	return applyMethod(obj, node.Method, args, defs, node.Token.Line)
}

func applyMethod(obj object.Object, method ast.Expression, args []object.Object, defs map[string]object.Object, l int) object.Object {
	switch obj := obj.(type) {
	case *object.String:
		return obj.Method(method.(*ast.Identifier).Value, args)
	case *object.File:
		return obj.Method(method.(*ast.Identifier).Value, args)
	case *object.Time:
		return obj.Method(method.(*ast.Identifier).Value, args, defs)
	case *object.Array:
		switch method.(*ast.Identifier).Value {
		case "map":
			return maap(obj, args)
		case "chuja":
			return filter(obj, args)
		default:
			return obj.Method(method.(*ast.Identifier).Value, args)
		}
	case *object.Module:
		if fn, ok := obj.Functions[method.(*ast.Identifier).Value]; ok {
			return fn(args, defs)
		}
	case *object.Instance:
		if fn, ok := obj.Package.Scope.Get(method.(*ast.Identifier).Value); ok {
			fn.(*object.Function).Env.Set("@", obj)
			ret := applyFunction(fn, args, l)
			fn.(*object.Function).Env.Del("@")
			return ret
		}
	case *object.Package:
		if fn, ok := obj.Scope.Get(method.(*ast.Identifier).Value); ok {
			fn.(*object.Function).Env.Set("@", obj)
			ret := applyFunction(fn, args, l)
			fn.(*object.Function).Env.Del("@")
			return ret
		}
	}
	return newError("Samahani, %s haina function '%s()'", obj.Inspect(), method.(*ast.Identifier).Value)
}

// ///////////////////////////////////////////////////////////////
// //////// Some methods here because of loop dependency ////////
// /////////////////////////////////////////////////////////////
func maap(a *object.Array, args []object.Object) object.Object {
	if len(args) != 1 && args[0].Type() != object.FUNCTION_OBJ {
		return newError("Samahani, hoja sii sahihi")
	}

	fn, ok := args[0].(*object.Function)
	if !ok {
		return newError("Samahani, hoja sii sahihi")
	}
	env := object.NewEnvironment()
	newArr := object.Array{Elements: []object.Object{}}
	for _, obj := range a.Elements {
		env.Set(fn.Parameters[0].Value, obj)
		r := Eval(fn.Body, env)
		if o, ok := r.(*object.ReturnValue); ok {
			r = o.Value
		}
		newArr.Elements = append(newArr.Elements, r)
	}
	return &newArr
}

func filter(a *object.Array, args []object.Object) object.Object {
	if len(args) != 1 && args[0].Type() != object.FUNCTION_OBJ {
		return newError("Samahani, hoja sii sahihi")
	}

	fn, ok := args[0].(*object.Function)
	if !ok {
		return newError("Samahani, hoja sii sahihi")
	}
	env := object.NewEnvironment()
	newArr := object.Array{Elements: []object.Object{}}
	for _, obj := range a.Elements {
		env.Set(fn.Parameters[0].Value, obj)
		cond := Eval(fn.Body, env)
		if cond.Inspect() == "kweli" {
			newArr.Elements = append(newArr.Elements, obj)
		}
	}
	return &newArr
}
