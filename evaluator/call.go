package evaluator

import (
	"github.com/NuruProgramming/Nuru/ast"
	"github.com/NuruProgramming/Nuru/object"
)

func evalCall(node *ast.CallExpression, env *object.Environment) object.Object {
	function := Eval(node.Function, env)

	if isError(function) {
		return function
	}

	var args []object.Object

	switch fn := function.(type) {
	case *object.Function:
		args = evalArgsExpressions(node, fn, env)
	case *object.Package:
		obj, ok := fn.Scope.Get("andaa")
		if !ok {
			return newError("Pakeji haina 'andaa'")
		}
		args = evalArgsExpressions(node, obj.(*object.Function), env)
	default:
		args = evalExpressions(node.Arguments, env)
	}

	if len(args) == 1 && isError(args[0]) {
		return args[0]
	}

	return applyFunction(function, args, node.Token.Line)
}

func evalArgsExpressions(node *ast.CallExpression, fn *object.Function, env *object.Environment) []object.Object {
	argsList := &object.Array{}
	argsHash := &object.Dict{}
	argsHash.Pairs = make(map[object.HashKey]object.DictPair)
	for _, exprr := range node.Arguments {
		switch exp := exprr.(type) {
		case *ast.Assign:
			val := Eval(exp.Value, env)
			if isError(val) {
				return []object.Object{val}
			}
			var keyHash object.HashKey
			key := &object.String{Value: exp.Name.Value}
			keyHash = key.HashKey()
			pair := object.DictPair{Key: key, Value: val}
			argsHash.Pairs[keyHash] = pair
		default:
			evaluated := Eval(exp, env)
			if isError(evaluated) {
				return []object.Object{evaluated}
			}
			argsList.Elements = append(argsList.Elements, evaluated)
		}
	}

	var result []object.Object
	var params = map[string]bool{}
	for _, exp := range fn.Parameters {
		params[exp.Value] = true
		if len(argsList.Elements) > 0 {
			result = append(result, argsList.Elements[0])
			argsList.Elements = argsList.Elements[1:]
		} else {
			keyParam := &object.String{Value: exp.Value}
			keyParamHash := keyParam.HashKey()
			if valParam, ok := argsHash.Pairs[keyParamHash]; ok {
				result = append(result, valParam.Value)
				delete(argsHash.Pairs, keyParamHash)
			} else {
				if _e, _ok := fn.Defaults[exp.Value]; _ok {
					evaluated := Eval(_e, env)
					if isError(evaluated) {
						return []object.Object{evaluated}
					}
					result = append(result, evaluated)
				} else {
					return []object.Object{&object.Error{Message: "Tumekosa Hoja"}}
				}
			}
		}
	}

	for _, pair := range argsHash.Pairs {
		if _, ok := params[pair.Key.(*object.String).Value]; ok {
			return []object.Object{&object.Error{Message: "Tumepewa hoja nyingi kwa parameter moja"}}
		}
	}

	return result

}
