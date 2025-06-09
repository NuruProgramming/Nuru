package evaluator

import (
	"github.com/NuruProgramming/Nuru/ast"
	"github.com/NuruProgramming/Nuru/object"
)

func evalDictLiteral(node *ast.DictLiteral, env *object.Environment) object.Object {
	pairs := make(map[object.HashKey]object.DictPair)

	for keyNode, valueNode := range node.Pairs {
		key := Eval(keyNode, env)
		if isError(key) {
			return key
		}

		hashKey, ok := key.(object.Hashable)
		if !ok {
			return newError("Mstari %d: Hashing imeshindikana: %s", node.Token.Line, key.Type())
		}

		value := Eval(valueNode, env)
		if isError(value) {
			return value
		}

		hashed := hashKey.HashKey()
		pairs[hashed] = object.DictPair{Key: key, Value: value}
	}

	// Create a new dictionary with reference counting
	dict := &object.Dict{
		Pairs: pairs,
	}

	// Track the dictionary in the reference counter
	object.GlobalRefCounter.TrackObject(dict)

	// Increment reference counts for all keys and values
	for _, pair := range pairs {
		object.Retain(pair.Key)
		object.Retain(pair.Value)
	}

	return dict
}
