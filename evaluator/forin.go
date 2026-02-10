package evaluator

import (
	"github.com/NuruProgramming/Nuru/ast"
	"github.com/NuruProgramming/Nuru/object"
)

func evalForInExpression(fie *ast.ForIn, env *object.Environment, line int) object.Object {
	// Evaluate the iterable expression
	iterable := Eval(fie.Iterable, env)
	if isError(iterable) {
		return iterable
	}

	// Check if the iterable is nil
	if iterable == nil {
		return newError("Mstari %d: Hakuna kitu cha kutengeneza loop nalo", line)
	}

	// Save existing variables with same names to restore them later
	var existingKeyIdentifier object.Object
	var okk bool

	// Only save and restore the key variable if it's not the special "_" case
	if fie.Key != "_" {
		existingKeyIdentifier, okk = env.Get(fie.Key)
	}

	existingValueIdentifier, okv := env.Get(fie.Value)

	// Set up deferred function to restore environment
	defer func() {
		if okk && fie.Key != "_" {
			env.Set(fie.Key, existingKeyIdentifier)
		}
		if okv {
			env.Set(fie.Value, existingValueIdentifier)
		}
	}()

	// Check if the object implements Iterable interface and process accordingly
	switch i := iterable.(type) {
	case object.Iterable:
		// If iterable is an Iterator (from .kitanzi()), release it when done so it can release the collection
		if _, ok := iterable.(*object.Iterator); ok {
			defer object.Release(iterable)
		}
		// Reset the iterator and set up deferred reset
		i.Reset()
		defer func() {
			i.Reset()
		}()
		return loopIterable(i.Next, env, fie)
	default:
		// Return a more descriptive error message
		return newError("Mstari %d: Samahani, '%s' haiwezi kutumiwa na 'kwa...ktk'. Unahitaji kutumia orodha, kamusi, neno, seti, jozi, au kitanzi",
			line, iterable.Type())
	}
}
