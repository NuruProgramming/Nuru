package evaluator

import (
	"github.com/NuruProgramming/Nuru/ast"
	"github.com/NuruProgramming/Nuru/object"
)

func evalTryCatchExpression(tce *ast.TryCatchExpression, env *object.Environment) object.Object {
	// Create a new environment for the try block
	tryEnv := object.NewEnclosedEnvironment(env)

	// Evaluate the try block
	result := Eval(tce.TryBlock, tryEnv)

	// If there's no error, return the result
	if !isError(result) {
		return result
	}

	// Handle return values from within try block
	if result.Type() == object.RETURN_VALUE_OBJ {
		return result
	}

	// We have an error, so evaluate the catch block
	catchEnv := object.NewEnclosedEnvironment(env)

	// If there's an error identifier, set it in the catch environment
	if tce.Identifier != "" {
		// The error is already an object.Error, so we can use it directly
		catchEnv.Set(tce.Identifier, result)
	}

	// Evaluate the catch block
	catchResult := Eval(tce.CatchBlock, catchEnv)

	return catchResult
}
