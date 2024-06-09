package evaluator

import (
	"fmt"

	"github.com/NuruProgramming/Nuru/ast"
	"github.com/NuruProgramming/Nuru/object"
)

var (
	NULL     = &object.Null{}
	TRUE     = &object.Boolean{Value: true}
	FALSE    = &object.Boolean{Value: false}
	BREAK    = &object.Break{}
	CONTINUE = &object.Continue{}
)

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	case *ast.Program:
		return evalProgram(node, env)

	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)

	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}

	case *ast.FloatLiteral:
		return &object.Float{Value: node.Value}

	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)

	case *ast.PrefixExpression:
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right, node.Token.Line)

	case *ast.InfixExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		right := Eval(node.Right, env)
		if isError(right) && right != nil {
			return right
		}
		return evalInfixExpression(node.Operator, left, right, node.Token.Line)
	case *ast.PostfixExpression:
		return evalPostfixExpression(env, node.Operator, node)

	case *ast.BlockStatement:
		return evalBlockStatement(node, env)

	case *ast.IfExpression:
		return evalIfExpression(node, env)

	case *ast.ReturnStatement:
		val := Eval(node.ReturnValue, env)
		if isError(val) {
			return val
		}
		return &object.ReturnValue{Value: val}

	case *ast.LetStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}

		env.Set(node.Name.Value, val)

	case *ast.Identifier:
		return evalIdentifier(node, env)

	case *ast.FunctionLiteral:
		return evalFunction(node, env)

	case *ast.MethodExpression:
		return evalMethodExpression(node, env)

	case *ast.Import:
		return evalImport(node, env)

	case *ast.CallExpression:
		return evalCall(node, env)

	case *ast.StringLiteral:
		return &object.String{Value: node.Value}
	case *ast.At:
		return evalAt(node, env)
	case *ast.ArrayLiteral:
		elements := evalExpressions(node.Elements, env)
		if len(elements) == 1 && isError(elements[0]) {
			return elements[0]
		}
		return &object.Array{Elements: elements}
	case *ast.IndexExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}
		index := Eval(node.Index, env)
		if isError(index) {
			return index
		}
		return evalIndexExpression(left, index, node.Token.Line)
	case *ast.DictLiteral:
		return evalDictLiteral(node, env)
	case *ast.WhileExpression:
		return evalWhileExpression(node, env)
	case *ast.Break:
		return evalBreak(node)
	case *ast.Continue:
		return evalContinue(node)
	case *ast.SwitchExpression:
		return evalSwitchStatement(node, env)
	case *ast.Null:
		return NULL
	// case *ast.For:
	// 	return evalForExpression(node, env)
	case *ast.ForIn:
		return evalForInExpression(node, env, node.Token.Line)
	case *ast.Package:
		return evalPackage(node, env)
	case *ast.PropertyExpression:
		return evalPropertyExpression(node, env)
	case *ast.PropertyAssignment:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}
		return evalPropertyAssignment(node.Name, val, env)
	case *ast.Assign:
		return evalAssign(node, env)
	case *ast.AssignEqual:
		return evalAssignEqual(node, env)

	case *ast.AssignmentExpression:
		left := Eval(node.Left, env)
		if isError(left) {
			return left
		}

		value := Eval(node.Value, env)
		if isError(value) {
			return value
		}

		// This is an easy way to assign operators like +=, -= etc
		// for index expressions (arrays and dicts) where applicable
		op := node.Token.Literal
		if len(op) >= 2 {
			op = op[:len(op)-1]
			value = evalInfixExpression(op, left, value, node.Token.Line)
			if isError(value) {
				return value
			}
		}

		if ident, ok := node.Left.(*ast.Identifier); ok {
			env.Set(ident.Value, value)
		} else if ie, ok := node.Left.(*ast.IndexExpression); ok {
			obj := Eval(ie.Left, env)
			if isError(obj) {
				return obj
			}

			if array, ok := obj.(*object.Array); ok {
				index := Eval(ie.Index, env)
				if isError(index) {
					return index
				}
				if idx, ok := index.(*object.Integer); ok {
					if int(idx.Value) >= len(array.Elements) {
						return newError("Index imezidi idadi ya elements")
					}
					array.Elements[idx.Value] = value
				} else {
					return newError("Hauwezi kufanya operesheni hii na %#v", index)
				}
			} else if hash, ok := obj.(*object.Dict); ok {
				key := Eval(ie.Index, env)
				if isError(key) {
					return key
				}
				if hashKey, ok := key.(object.Hashable); ok {
					hashed := hashKey.HashKey()
					hash.Pairs[hashed] = object.DictPair{Key: key, Value: value}
				} else {
					return newError("Hauwezi kufanya operesheni hii na %T", key)
				}
			} else {
				return newError("%T haifanyi operesheni hii", obj)
			}
		} else {
			return newError("Tumia neno kama kibadala, sio %T", left)
		}

	}

	return nil
}

func evalProgram(program *ast.Program, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range program.Statements {
		result = Eval(statement, env)

		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}

	return result
}

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

func isTruthy(obj object.Object) bool {
	switch obj {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}

	return false
}

func evalExpressions(exps []ast.Expression, env *object.Environment) []object.Object {
	var result []object.Object

	for _, e := range exps {
		evaluated := Eval(e, env)
		if isError(evaluated) {
			return []object.Object{evaluated}
		}

		result = append(result, evaluated)
	}

	return result
}

func applyFunction(fn object.Object, args []object.Object, line int) object.Object {
	switch fn := fn.(type) {
	case *object.Function:
		extendedEnv := extendedFunctionEnv(fn, args)
		evaluated := Eval(fn.Body, extendedEnv)
		return unwrapReturnValue(evaluated)
	case *object.Builtin:
		if result := fn.Fn(args...); result != nil {
			return result
		}
		return NULL
	case *object.Package:
		obj := &object.Instance{
			Package: fn,
			Env:     object.NewEnclosedEnvironment(fn.Env),
		}
		obj.Env.Set("@", obj)
		node, ok := fn.Scope.Get("andaa")
		if !ok {
			return newError("Hamna andaa kiendesha")
		}
		node.(*object.Function).Env.Set("@", obj)
		applyFunction(node, args, fn.Name.Token.Line)
		node.(*object.Function).Env.Del("@")
		return obj
	default:
		if fn != nil {
			return newError("Mstari %d: Hiki sio kiendesha: %s", line, fn.Type())
		} else {
			return newError("Bro how did you even get here??? Contact language maker asap!")
		}
	}

}

func extendedFunctionEnv(fn *object.Function, args []object.Object) *object.Environment {
	env := object.NewEnclosedEnvironment(fn.Env)

	for paramIdx, param := range fn.Parameters {
		if paramIdx < len(args) {
			env.Set(param.Value, args[paramIdx])
		}
	}
	return env
}

func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}

	return obj
}

func evalBreak(node *ast.Break) object.Object {
	return BREAK
}

func evalContinue(node *ast.Continue) object.Object {
	return CONTINUE
}

// func evalForExpression(fe *ast.For, env *object.Environment) object.Object {
// 	obj, ok := env.Get(fe.Identifier)
// 	defer func() { // stay safe and not reassign an existing variable
// 		if ok {
// 			env.Set(fe.Identifier, obj)
// 		}
// 	}()
// 	val := Eval(fe.StarterValue, env)
// 	if isError(val) {
// 		return val
// 	}

// 	env.Set(fe.StarterName.Value, val)

// 	// err := Eval(fe.Starter, env)
// 	// if isError(err) {
// 	// 	return err
// 	// }
// 	for {
// 		evaluated := Eval(fe.Condition, env)
// 		if isError(evaluated) {
// 			return evaluated
// 		}
// 		if !isTruthy(evaluated) {
// 			break
// 		}
// 		res := Eval(fe.Block, env)
// 		if isError(res) {
// 			return res
// 		}
// 		if res.Type() == object.BREAK_OBJ {
// 			break
// 		}
// 		if res.Type() == object.CONTINUE_OBJ {
// 			err := Eval(fe.Closer, env)
// 			if isError(err) {
// 				return err
// 			}
// 			continue
// 		}
// 		if res.Type() == object.RETURN_VALUE_OBJ {
// 			return res
// 		}
// 		err := Eval(fe.Closer, env)
// 		if isError(err) {
// 			return err
// 		}
// 	}
// 	return NULL
// }

func loopIterable(next func() (object.Object, object.Object), env *object.Environment, fi *ast.ForIn) object.Object {
	k, v := next()
	for k != nil && v != nil {
		env.Set(fi.Key, k)
		env.Set(fi.Value, v)
		res := Eval(fi.Block, env)
		if isError(res) {
			return res
		}
		if res != nil {
			if res.Type() == object.BREAK_OBJ {
				break
			}
			if res.Type() == object.CONTINUE_OBJ {
				k, v = next()
				continue
			}
			if res.Type() == object.RETURN_VALUE_OBJ {
				return res
			}
		}
		k, v = next()
	}
	return NULL
}
