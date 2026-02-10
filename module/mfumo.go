package module

import (
	"github.com/NuruProgramming/Nuru/object"
)

// MfumoFunctions contains runtime/memory-related functions (Policy A: meta in module)
var MfumoFunctions = map[string]object.ModuleFunction{}

func init() {
	MfumoFunctions["safishaMemori"] = safishaMemori
	MfumoFunctions["takwimuMemori"] = takwimuMemori
	MfumoFunctions["takwimuMemoriKwa"] = takwimuMemoriKwa
	MfumoFunctions["kumbukumbaDhaifu"] = kumbukumbaDhaifu
}

func safishaMemori(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 0 {
		return newError("Samahani, safishaMemori haihitaji hoja, wewe umeweka %d", len(args))
	}
	object.RunGC()
	return &object.Null{}
}

func takwimuMemori(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 0 {
		return newError("Samahani, takwimuMemori haihitaji hoja, wewe umeweka %d", len(args))
	}
	object.PrintObjectStats()
	return &object.Integer{Value: int64(object.GetObjectCount())}
}

func kumbukumbaDhaifu(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return newError("Samahani, kumbukumbaDhaifu inahitaji hoja 1, wewe umeweka %d", len(args))
	}
	return object.NewWeakReference(args[0])
}

func takwimuMemoriKwa(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 1 {
		return newError("Samahani, takwimuMemoriKwa inahitaji hoja 1, wewe umeweka %d", len(args))
	}
	if args[0].Type() != object.DICT_OBJ {
		return newError("Samahani, takwimuMemoriKwa inahitaji kamusi")
	}
	result := &object.Dict{Pairs: make(map[object.HashKey]object.DictPair)}
	totalCount := object.GetObjectCount()
	totalKey := &object.String{Value: "jumla"}
	totalValue := &object.Integer{Value: int64(totalCount)}
	totalHash := totalKey.HashKey()
	result.Pairs[totalHash] = object.DictPair{Key: totalKey, Value: totalValue}
	stats := object.GetTypeStats()
	for objType, count := range stats {
		typeKey := &object.String{Value: string(objType)}
		typeValue := &object.Integer{Value: int64(count)}
		typeHash := typeKey.HashKey()
		result.Pairs[typeHash] = object.DictPair{Key: typeKey, Value: typeValue}
	}
	return result
}
