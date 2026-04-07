package module

import (
	"github.com/NuruProgramming/Nuru/object"
)

type ReporterFunc func(pass bool, message string)

var TestReporter ReporterFunc

var TestFunctions = map[string]object.ModuleFunction{}

func init() {
	TestFunctions["hakiki"] = assert
}

func assert(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) < 1 {
		return &object.Error{Message: "Hakiki inahitaji condition."}
	}

	condition := args[0]
	isTrue := true

	switch val := condition.(type) {
	case *object.Boolean:
		isTrue = val.Value
	case *object.Null:
		isTrue = false
	default:
		isTrue = true
	}

	msg := "Jaribio"
	if len(args) > 1 {
		if m, ok := args[1].(*object.String); ok {
			msg = m.Value
		}
	}

	if TestReporter != nil {
		TestReporter(isTrue, msg)
		return &object.Boolean{Value: true}
	}

	if isTrue {
		return &object.Boolean{Value: true}
	}
	return &object.Error{Message: msg}
}
