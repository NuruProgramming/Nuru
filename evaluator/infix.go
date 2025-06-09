package evaluator

import (
	"math"
	"strings"

	"github.com/NuruProgramming/Nuru/object"
)

func evalInfixExpression(operator string, left, right object.Object, line int) object.Object {
	if right == nil {
		return newError("Mstari %d: Umekosea hapa", line)
	}
	if left == nil {
		return newError("Mstari %d: Umekosea hapa", line)
	}
	switch {

	case operator == "ktk":
		return evalInExpression(left, right, line)

	case left.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ:
		return evalStringInfixExpression(operator, left, right, line)

	case operator == "+" && left.Type() == object.DICT_OBJ && right.Type() == object.DICT_OBJ:
		leftVal := left.(*object.Dict).Pairs
		rightVal := right.(*object.Dict).Pairs

		// Create a new reference-counted dictionary
		result := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}
		object.GlobalRefCounter.TrackObject(result)

		// Copy pairs from left dictionary
		for k, v := range leftVal {
			// Retain references for the new dictionary
			object.Retain(v.Key)
			object.Retain(v.Value)
			result.Pairs[k] = v
		}

		// Copy pairs from right dictionary
		for k, v := range rightVal {
			// If key already exists in result, release old value
			if oldPair, exists := result.Pairs[k]; exists {
				object.Release(oldPair.Value)
			}

			// Retain references for the new dictionary
			object.Retain(v.Key)
			object.Retain(v.Value)
			result.Pairs[k] = v
		}

		return result

	case operator == "+" && left.Type() == object.ARRAY_OBJ && right.Type() == object.ARRAY_OBJ:
		leftVal := left.(*object.Array).Elements
		rightVal := right.(*object.Array).Elements

		// Create a new reference-counted array
		result := &object.Array{Elements: make([]object.Object, 0, len(leftVal)+len(rightVal))}
		object.GlobalRefCounter.TrackObject(result)

		// Copy elements from left array
		for _, elem := range leftVal {
			object.Retain(elem)
			result.Elements = append(result.Elements, elem)
		}

		// Copy elements from right array
		for _, elem := range rightVal {
			object.Retain(elem)
			result.Elements = append(result.Elements, elem)
		}

		return result

	case operator == "*" && left.Type() == object.ARRAY_OBJ && right.Type() == object.INTEGER_OBJ:
		leftVal := left.(*object.Array).Elements
		rightVal := int(right.(*object.Integer).Value)

		// If multiplier is zero or negative, return empty array
		if rightVal <= 0 {
			return &object.Array{Elements: []object.Object{}}
		}

		// Create a new reference-counted array with sufficient capacity
		result := &object.Array{Elements: make([]object.Object, 0, len(leftVal)*rightVal)}
		object.GlobalRefCounter.TrackObject(result)

		// Add repeated elements
		for i := 0; i < rightVal; i++ {
			for _, elem := range leftVal {
				object.Retain(elem)
				result.Elements = append(result.Elements, elem)
			}
		}

		return result

	case operator == "*" && left.Type() == object.INTEGER_OBJ && right.Type() == object.ARRAY_OBJ:
		leftVal := int(left.(*object.Integer).Value)
		rightVal := right.(*object.Array).Elements

		// If multiplier is zero or negative, return empty array
		if leftVal <= 0 {
			return &object.Array{Elements: []object.Object{}}
		}

		// Create a new reference-counted array with sufficient capacity
		result := &object.Array{Elements: make([]object.Object, 0, len(rightVal)*leftVal)}
		object.GlobalRefCounter.TrackObject(result)

		// Add repeated elements
		for i := 0; i < leftVal; i++ {
			for _, elem := range rightVal {
				object.Retain(elem)
				result.Elements = append(result.Elements, elem)
			}
		}

		return result

	case operator == "*" && left.Type() == object.STRING_OBJ && right.Type() == object.INTEGER_OBJ:
		leftVal := left.(*object.String).Value
		rightVal := right.(*object.Integer).Value
		return &object.String{Value: strings.Repeat(leftVal, int(rightVal))}

	case operator == "*" && left.Type() == object.INTEGER_OBJ && right.Type() == object.STRING_OBJ:
		leftVal := left.(*object.Integer).Value
		rightVal := right.(*object.String).Value
		return &object.String{Value: strings.Repeat(rightVal, int(leftVal))}

	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, left, right, line)

	case left.Type() == object.FLOAT_OBJ && right.Type() == object.FLOAT_OBJ:
		return evalFloatInfixExpression(operator, left, right, line)

	case left.Type() == object.INTEGER_OBJ && right.Type() == object.FLOAT_OBJ:
		return evalFloatIntegerInfixExpression(operator, left, right, line)

	case left.Type() == object.FLOAT_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalFloatIntegerInfixExpression(operator, left, right, line)

	case operator == "==":
		return nativeBoolToBooleanObject(left == right)

	case operator == "!=":
		return nativeBoolToBooleanObject(left != right)
	case left.Type() == object.BOOLEAN_OBJ && right.Type() == object.BOOLEAN_OBJ:
		return evalBooleanInfixExpression(operator, left, right, line)

	case left.Type() != right.Type():
		return newError("Mstari %d: Aina Hazilingani: %s %s %s",
			line, left.Type(), operator, right.Type())

	default:
		return newError("Mstari %d: Operesheni Haieleweki: %s %s %s",
			line, left.Type(), operator, right.Type())
	}
}

func evalFloatIntegerInfixExpression(operator string, left, right object.Object, line int) object.Object {
	var leftVal, rightVal float64
	if left.Type() == object.FLOAT_OBJ {
		leftVal = left.(*object.Float).Value
		rightVal = float64(right.(*object.Integer).Value)
	} else {
		leftVal = float64(left.(*object.Integer).Value)
		rightVal = right.(*object.Float).Value
	}

	var val float64
	switch operator {
	case "+":
		val = leftVal + rightVal
	case "-":
		val = leftVal - rightVal
	case "*":
		val = leftVal * rightVal
	case "**":
		val = math.Pow(float64(leftVal), float64(rightVal))
	case "/":
		val = leftVal / rightVal
	case "%":
		val = math.Mod(leftVal, rightVal)
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("Mstari %d: Operesheni Haieleweki: %s %s %s",
			line, left.Type(), operator, right.Type())
	}

	if math.Mod(val, 1) == 0 {
		return &object.Integer{Value: int64(val)}
	} else {
		return &object.Float{Value: val}
	}
}

func evalStringInfixExpression(operator string, left, right object.Object, line int) object.Object {

	leftVal := left.(*object.String).Value
	rightVal := right.(*object.String).Value

	switch operator {
	case "+":
		return &object.String{Value: leftVal + rightVal}
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("Mstari %d: Operesheni Haieleweki: %s %s %s", line, left.Type(), operator, right.Type())
	}
}

func evalBooleanInfixExpression(operator string, left, right object.Object, line int) object.Object {
	leftVal := left.(*object.Boolean).Value
	rightVal := right.(*object.Boolean).Value

	switch operator {
	case "&&":
		return nativeBoolToBooleanObject(leftVal && rightVal)
	case "||":
		return nativeBoolToBooleanObject(leftVal || rightVal)
	default:
		return newError("Mstari %d: Operesheni Haieleweki: %s %s %s", line, left.Type(), operator, right.Type())
	}
}

func evalFloatInfixExpression(operator string, left, right object.Object, line int) object.Object {
	leftVal := left.(*object.Float).Value
	rightVal := right.(*object.Float).Value

	switch operator {
	case "+":
		return &object.Float{Value: leftVal + rightVal}
	case "-":
		return &object.Float{Value: leftVal - rightVal}
	case "*":
		return &object.Float{Value: leftVal * rightVal}
	case "**":
		return &object.Float{Value: math.Pow(float64(leftVal), float64(rightVal))}
	case "/":
		return &object.Float{Value: leftVal / rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("Mstari %d: Operesheni Haieleweki: %s %s %s",
			line, left.Type(), operator, right.Type())
	}
}

func evalIntegerInfixExpression(operator string, left, right object.Object, line int) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "**":
		return &object.Float{Value: float64(math.Pow(float64(leftVal), float64(rightVal)))}
	case "/":
		x := float64(leftVal) / float64(rightVal)
		if math.Mod(x, 1) == 0 {
			return &object.Integer{Value: int64(x)}
		} else {
			return &object.Float{Value: x}
		}
	case "%":
		return &object.Integer{Value: leftVal % rightVal}
	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case "<=":
		return nativeBoolToBooleanObject(leftVal <= rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case ">=":
		return nativeBoolToBooleanObject(leftVal >= rightVal)
	case "==":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("Mstari %d: Operesheni Haieleweki: %s %s %s",
			line, left.Type(), operator, right.Type())
	}
}
