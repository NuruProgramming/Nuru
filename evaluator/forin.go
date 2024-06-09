package evaluator

import (
	"github.com/NuruProgramming/Nuru/ast"
	"github.com/NuruProgramming/Nuru/object"
)

func evalForInExpression(fie *ast.ForIn, env *object.Environment, line int) object.Object {
	iterable := Eval(fie.Iterable, env)
	existingKeyIdentifier, okk := env.Get(fie.Key) // again, stay safe
	existingValueIdentifier, okv := env.Get(fie.Value)
	defer func() { // restore them later on
		if okk {
			env.Set(fie.Key, existingKeyIdentifier)
		}
		if okv {
			env.Set(fie.Value, existingValueIdentifier)
		}
	}()
	switch i := iterable.(type) {
	case object.Iterable:
		defer func() {
			i.Reset()
		}()
		return loopIterable(i.Next, env, fie)
	default:
		return newError("Mstari %d: Huwezi kufanya operesheni hii na %s", line, i.Type())
	}
}
