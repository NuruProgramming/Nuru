package evaluator

import (
	"github.com/NuruProgramming/Nuru/ast"
	"github.com/NuruProgramming/Nuru/object"
)

func evalWhileExpression(we *ast.WhileExpression, env *object.Environment) object.Object {
	var evaluated object.Object = NULL
	
	// Use iterative loop instead of recursion to prevent stack overflow
	for {
		// Evaluate the condition
		condition := Eval(we.Condition, env)
		if isError(condition) {
			return condition
		}
		
		// If condition is false, exit the loop
		if !isTruthy(condition) {
			break
		}
		
		// Execute the loop body
		evaluated = Eval(we.Consequence, env)
		if isError(evaluated) {
			return evaluated
		}
		
		// Handle control flow statements
		if evaluated != nil {
			switch evaluated.Type() {
			case object.BREAK_OBJ:
				// Exit the loop immediately
				return NULL
			case object.CONTINUE_OBJ:
				// Skip to the next iteration
				continue
			case object.RETURN_VALUE_OBJ:
				// Propagate the return value
				return evaluated
			}
		}
	}
	
	return evaluated
}
