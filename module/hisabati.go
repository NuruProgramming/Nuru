package module

import (
	"math"
	"math/rand"
	"time"

	"github.com/NuruProgramming/Nuru/object"
)

var MathFunctions = map[string]object.ModuleFunction{
	"PI":        pi,
	"e":         e,
	"phi":       phi,
	"ln10":      ln10,
	"ln2":       ln2,
	"log10e":    log10e,
	"log2e":     log2e,
	"log2":      log2,
	"sqrt1_2":   sqrt1_2,
	"sqrt2":     sqrt2,
	"sqrt3":     sqrt3,
	"sqrt5":     sqrt5,
	"EPSILON":   epsilon,
	"abs":       abs,
	"sign":      sign,
	"ceil":      ceil,
	"floor":     floor,
	"sqrt":      sqrt,
	"cbrt":      cbrt,
	"root":      root,
	"hypot":     hypot,
	"random":    random,
	"factorial": factorial,
	"round":     round,
	"max":       max,
	"min":       min,
	"exp":       exp,
	"expm1":     expm1,
	"log":       log,
	"log10":     log10,
	"log1p":     log1p,
	"cos":       cos,
	"sin":       sin,
	"tan":       tan,
	"acos":      acos,
	"asin":      asin,
	"atan":      atan,
	"cosh":      cosh,
	"sinh":      sinh,
	"tanh":      tanh,
	"acosh":     acosh,
	"asinh":     asinh,
	"atanh":     atanh,
	"atan2":     atan2,
}

var Constants = map[string]object.Object{
	"PI":      &object.Float{Value: math.Pi},
	"e":       &object.Float{Value: math.E},
	"phi":     &object.Float{Value: (1 + math.Sqrt(5)) / 2},
	"ln10":    &object.Float{Value: math.Log10E},
	"ln2":     &object.Float{Value: math.Ln2},
	"log10e":  &object.Float{Value: math.Log10E},
	"log2e":   &object.Float{Value: math.Log2E},
	"sqrt1_2": &object.Float{Value: 1 / math.Sqrt2},
	"sqrt2":   &object.Float{Value: math.Sqrt2},
	"sqrt3":   &object.Float{Value: math.Sqrt(3)},
	"sqrt5":   &object.Float{Value: math.Sqrt(5)},
	"EPSILON": &object.Float{Value: 2.220446049250313e-16},
}

func pi(args []object.Object, defs map[string]object.Object) object.Object {
	return &object.Float{Value: math.Pi}
}

func e(args []object.Object, defs map[string]object.Object) object.Object {
	return &object.Float{Value: math.E}
}

func phi(args []object.Object, defs map[string]object.Object) object.Object {
	return &object.Float{Value: (1 + math.Sqrt(5)) / 2}
}

func ln10(args []object.Object, defs map[string]object.Object) object.Object {
	return &object.Float{Value: math.Log10E}
}

func ln2(args []object.Object, defs map[string]object.Object) object.Object {
	return &object.Float{Value: math.Ln2}
}

func log10e(args []object.Object, defs map[string]object.Object) object.Object {
	return &object.Float{Value: math.Log10E}
}

func log2e(args []object.Object, defs map[string]object.Object) object.Object {
	return &object.Float{Value: math.Log2E}
}

func sqrt1_2(args []object.Object, defs map[string]object.Object) object.Object {
	return &object.Float{Value: 1 / math.Sqrt2}
}

func sqrt2(args []object.Object, defs map[string]object.Object) object.Object {
	return &object.Float{Value: math.Sqrt2}
}

func sqrt3(args []object.Object, defs map[string]object.Object) object.Object {
	return &object.Float{Value: math.Sqrt(3)}
}

func sqrt5(args []object.Object, defs map[string]object.Object) object.Object {
	return &object.Float{Value: math.Sqrt(5)}
}

func epsilon(args []object.Object, defs map[string]object.Object) object.Object {
	return &object.Float{Value: 2.220446049250313e-16}
}

func abs(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "Undo hili haliruhusu ufafanuzi."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "Undo hili linahitaji hoja moja tu"}
	}
	if args[0].Type() != object.INTEGER_OBJ && args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "Hoja lazima iwe namba"}
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		if arg.Value < 0 {
			return &object.Integer{Value: -arg.Value}
		}
		return arg
	case *object.Float:
		if arg.Value < 0 {
			return &object.Float{Value: -arg.Value}
		}
		return arg
	default:
		return &object.Error{Message: "Hoja lazima iwe namba"}
	}
}

func sign(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "Undo hili haliruhusu ufafanuzi."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "Undo hili linahitaji hoja moja tu"}
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		if arg.Value == 0 {
			return &object.Integer{Value: 0}
		} else if arg.Value > 0 {
			return &object.Integer{Value: 1}
		} else {
			return &object.Integer{Value: -1}
		}
	case *object.Float:
		if arg.Value == 0 {
			return &object.Integer{Value: 0}
		} else if arg.Value > 0 {
			return &object.Integer{Value: 1}
		} else {
			return &object.Integer{Value: -1}
		}
	default:
		return &object.Error{Message: "Hoja lazima iwe namba"}
	}
}

func ceil(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "Undo hili haliruhusu ufafanuzi."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "Undo hili linahitaji hoja moja tu"}
	}
	if args[0].Type() != object.INTEGER_OBJ && args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "Hoja lazima iwe namba"}
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Integer{Value: arg.Value}
	case *object.Float:
		return &object.Integer{Value: int64(math.Ceil(arg.Value))}
	default:
		return &object.Error{Message: "Hoja lazima iwe namba"}
	}
}

func floor(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "Undo hili haliruhusu ufafanuzi."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "Undo hili linahitaji hoja moja tu"}
	}
	if args[0].Type() != object.INTEGER_OBJ && args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "Hoja lazima iwe namba"}
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Integer{Value: arg.Value}
	case *object.Float:
		return &object.Integer{Value: int64(math.Floor(arg.Value))}
	default:
		return &object.Error{Message: "Hoja lazima iwe namba"}
	}
}

func sqrt(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "Undo hili haliruhusu ufafanuzi."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "Undo hili linahitaji hoja moja tu"}
	}
	if args[0].Type() != object.INTEGER_OBJ && args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "Hoja lazima iwe namba"}
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Sqrt(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Sqrt(arg.Value)}
	default:
		return &object.Error{Message: "Hoja lazima iwe namba"}
	}
}

func cbrt(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "Undo hili haliruhusu ufafanuzi."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "Undo hili linahitaji hoja moja tu"}
	}
	if args[0].Type() != object.INTEGER_OBJ && args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "Hoja lazima iwe namba"}
	}
	switch arg := args[0].(type) {
	case *object.Integer:
		return &object.Float{Value: math.Cbrt(float64(arg.Value))}
	case *object.Float:
		return &object.Float{Value: math.Cbrt(arg.Value)}
	default:
		return &object.Error{Message: "Hoja lazima iwe namba"}
	}
}

func root(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "Undo hili haliruhusu ufafanuzi."}
	}
	if len(args) != 2 {
		return &object.Error{Message: "Undo hili linahitaji hoja mbili tu"}
	}
	if args[0].Type() != object.INTEGER_OBJ && args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "Hoja ya kwanza lazima iwe namba"}
	}
	if args[1].Type() != object.INTEGER_OBJ {
		return &object.Error{Message: "Hoja ya pili lazima iwe namba"}
	}
	base, ok := args[0].(*object.Float)
	if !ok {
		base = &object.Float{Value: float64(args[0].(*object.Integer).Value)}
	}
	exp := args[1].(*object.Integer).Value

	if exp == 0 {
		return &object.Float{Value: 1.0}
	} else if exp < 0 {
		return &object.Error{Message: "Second Hoja lazima iwe a non-negative integer"}
	}

	x := 1.0
	for i := 0; i < 10; i++ {
		x = x - (math.Pow(x, float64(exp))-base.Value)/(float64(exp)*math.Pow(x, float64(exp-1)))
	}

	return &object.Float{Value: x}
}

func hypot(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "Undo hili haliruhusu ufafanuzi."}
	}
	if len(args) < 2 {
		return &object.Error{Message: "Undo hili linahitaji hoja moja tu"}
	}
	var sumOfSquares float64
	for _, arg := range args {
		if arg.Type() != object.INTEGER_OBJ && arg.Type() != object.FLOAT_OBJ {
			return &object.Error{Message: "Hoja lazima ziwe namba"}
		}
		switch num := arg.(type) {
		case *object.Integer:
			sumOfSquares += float64(num.Value) * float64(num.Value)
		case *object.Float:
			sumOfSquares += num.Value * num.Value
		}
	}
	return &object.Float{Value: math.Sqrt(sumOfSquares)}
}

func factorial(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "Undo hili haliruhusu ufafanuzi."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "Undo hili linahitaji hoja moja tu"}
	}
	if args[0].Type() != object.INTEGER_OBJ {
		return &object.Error{Message: "Hoja lazima iwe namba"}
	}
	n := args[0].(*object.Integer).Value
	if n < 0 {
		return &object.Error{Message: "Hoja lazima iwe a non-negative integer"}
	}
	result := int64(1)
	for i := int64(2); i <= n; i++ {
		result *= i
	}
	return &object.Integer{Value: result}
}

func round(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "Undo hili haliruhusu ufafanuzi."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "Undo hili linahitaji hoja moja tu"}
	}
	if args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "Hoja lazima iwe desimali"}
	}

	num := args[0].(*object.Float).Value
	return &object.Integer{Value: int64(num + 0.5)}
}

func max(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "Undo hili haliruhusu ufafanuzi."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "Undo hili linahitaji hoja moja tu"}
	}

	arg, ok := args[0].(*object.Array)
	if !ok {
		return &object.Error{Message: "Hoja lazima iwe an array"}
	}

	if len(arg.Elements) == 0 {
		return &object.Error{Message: "Orodha haipaswi kuwa tupu"}
	}

	var maxNum float64

	for _, element := range arg.Elements {
		if element.Type() != object.INTEGER_OBJ && element.Type() != object.FLOAT_OBJ {
			return &object.Error{Message: "Vipengee vya orodha lazima viwe namba"}
		}

		switch num := element.(type) {
		case *object.Integer:
			if float64(num.Value) > maxNum {
				maxNum = float64(num.Value)
			}
		case *object.Float:
			if num.Value > maxNum {
				maxNum = num.Value
			}
		default:
			return &object.Error{Message: "Vipengee vya orodha lazima viwe namba"}
		}
	}

	return &object.Float{Value: maxNum}
}

func min(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "Undo hili haliruhusu ufafanuzi."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "Undo hili linahitaji hoja moja tu"}
	}

	arg, ok := args[0].(*object.Array)
	if !ok {
		return &object.Error{Message: "Hoja lazima iwe an array"}
	}

	if len(arg.Elements) == 0 {
		return &object.Error{Message: "Orodha haipaswi kuwa tupu"}
	}

	minNum := math.MaxFloat64

	for _, element := range arg.Elements {
		if element.Type() != object.INTEGER_OBJ && element.Type() != object.FLOAT_OBJ {
			return &object.Error{Message: "Vipengee vya orodha lazima viwe namba"}
		}

		switch num := element.(type) {
		case *object.Integer:
			if float64(num.Value) < minNum {
				minNum = float64(num.Value)
			}
		case *object.Float:
			if num.Value < minNum {
				minNum = num.Value
			}
		default:
			return &object.Error{Message: "Vipengee vya orodha lazima viwe namba"}
		}
	}

	return &object.Float{Value: minNum}
}

func exp(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "Undo hili haliruhusu ufafanuzi."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "Undo hili linahitaji hoja moja tu"}
	}
	if args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "Hoja lazima iwe desimali"}
	}
	num := args[0].(*object.Float).Value
	return &object.Float{Value: math.Exp(num)}
}

func expm1(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "Undo hili haliruhusu ufafanuzi."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "Undo hili linahitaji hoja moja tu"}
	}
	if args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "Hoja lazima iwe desimali"}
	}
	num := args[0].(*object.Float).Value
	return &object.Float{Value: math.Expm1(num)}
}

func log(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "Undo hili haliruhusu ufafanuzi."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "Undo hili linahitaji hoja moja tu"}
	}
	if args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "Hoja lazima iwe desimali"}
	}
	num := args[0].(*object.Float).Value
	return &object.Float{Value: math.Log(num)}
}

func log10(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "Undo hili haliruhusu ufafanuzi."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "Undo hili linahitaji hoja moja tu"}
	}
	if args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "Hoja lazima iwe desimali"}
	}
	num := args[0].(*object.Float).Value
	return &object.Float{Value: math.Log10(num)}
}

func log2(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "Undo hili haliruhusu ufafanuzi."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "Undo hili linahitaji hoja moja tu"}
	}
	if args[0].Type() != object.INTEGER_OBJ && args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "Hoja lazima iwe namba"}
	}

	arg := extractFloatValue(args[0])

	if arg <= 0 {
		return &object.Error{Message: "Hoja lazima iwe kubwa kuliko 0"}
	}

	return &object.Float{Value: math.Log2(arg)}
}

func extractFloatValue(obj object.Object) float64 {
	switch obj := obj.(type) {
	case *object.Integer:
		return float64(obj.Value)
	case *object.Float:
		return obj.Value
	default:
		return 0
	}
}

func log1p(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "Undo hili haliruhusu ufafanuzi."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "Undo hili linahitaji hoja moja tu"}
	}
	if args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "Hoja lazima iwe desimali"}
	}
	num := args[0].(*object.Float).Value
	return &object.Float{Value: math.Log1p(num)}
}

func cos(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "Undo hili haliruhusu ufafanuzi."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "Undo hili linahitaji hoja moja tu"}
	}
	if args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "Hoja lazima iwe desimali"}
	}
	num := args[0].(*object.Float).Value
	return &object.Float{Value: math.Cos(num)}
}

func sin(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "Undo hili haliruhusu ufafanuzi."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "Undo hili linahitaji hoja moja tu"}
	}
	if args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "Hoja lazima iwe desimali"}
	}
	num := args[0].(*object.Float).Value
	return &object.Float{Value: math.Sin(num)}
}

func tan(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "Undo hili haliruhusu ufafanuzi."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "Undo hili linahitaji hoja moja tu"}
	}
	if args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "Hoja lazima iwe desimali"}
	}
	num := args[0].(*object.Float).Value
	return &object.Float{Value: math.Tan(num)}
}

func acos(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "Undo hili haliruhusu ufafanuzi."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "Undo hili linahitaji hoja moja tu"}
	}
	if args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "Hoja lazima iwe desimali"}
	}
	num := args[0].(*object.Float).Value
	return &object.Float{Value: math.Acos(num)}
}

func asin(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "Undo hili haliruhusu ufafanuzi."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "Undo hili linahitaji hoja moja tu"}
	}
	if args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "Hoja lazima iwe desimali"}
	}
	num := args[0].(*object.Float).Value
	return &object.Float{Value: math.Asin(num)}
}

func atan(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "Undo hili haliruhusu ufafanuzi."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "Undo hili linahitaji hoja moja tu"}
	}
	if args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "Hoja lazima iwe desimali"}
	}
	num := args[0].(*object.Float).Value
	return &object.Float{Value: math.Atan(num)}
}

func cosh(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "Undo hili haliruhusu ufafanuzi."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "Undo hili linahitaji hoja moja tu"}
	}
	if args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "Hoja lazima iwe desimali"}
	}
	num := args[0].(*object.Float).Value
	return &object.Float{Value: math.Cosh(num)}
}

func sinh(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "Undo hili haliruhusu ufafanuzi."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "Undo hili linahitaji hoja moja tu"}
	}
	if args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "Hoja lazima iwe desimali"}
	}
	num := args[0].(*object.Float).Value
	return &object.Float{Value: math.Sinh(num)}
}

func tanh(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "Undo hili haliruhusu ufafanuzi."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "Undo hili linahitaji hoja moja tu"}
	}
	if args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "Hoja lazima iwe desimali"}
	}
	num := args[0].(*object.Float).Value
	return &object.Float{Value: math.Tanh(num)}
}

func acosh(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "Undo hili haliruhusu ufafanuzi."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "Undo hili linahitaji hoja moja tu"}
	}
	if args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "Hoja lazima iwe desimali"}
	}
	num := args[0].(*object.Float).Value
	return &object.Float{Value: math.Acosh(num)}
}

func asinh(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "Undo hili haliruhusu ufafanuzi."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "Undo hili linahitaji hoja moja tu"}
	}
	if args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "Hoja lazima iwe desimali"}
	}
	num := args[0].(*object.Float).Value
	return &object.Float{Value: math.Asinh(num)}
}

func atan2(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "Undo hili haliruhusu ufafanuzi."}
	}
	if len(args) != 2 {
		return &object.Error{Message: "Undo hili linahitaji hoja mbili tu."}
	}
	if args[0].Type() != object.INTEGER_OBJ && args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "Hoja lazima ziwe namba"}
	}
	if args[1].Type() != object.INTEGER_OBJ && args[1].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "Hoja lazima ziwe namba"}
	}

	y := extractFloatValue(args[0])
	x := extractFloatValue(args[1])

	return &object.Float{Value: math.Atan2(y, x)}
}

func atanh(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "Undo hili haliruhusu ufafanuzi."}
	}
	if len(args) != 1 {
		return &object.Error{Message: "Undo hili linahitaji hoja moja tu"}
	}
	if args[0].Type() != object.FLOAT_OBJ {
		return &object.Error{Message: "Hoja lazima iwe desimali"}
	}
	num := args[0].(*object.Float).Value
	return &object.Float{Value: math.Atanh(num)}
}

func random(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "Undo hili haliruhusu ufafanuzi."}
	}

	if len(args) != 0 {
		return &object.Error{Message: "Undo hili halipaswi kupokea hoja."}
	}

	rand.Seed(time.Now().UnixNano())
	value := rand.Float64()

	return &object.Float{Value: value}
}
