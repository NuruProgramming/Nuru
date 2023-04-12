package module

import (
	"os"

	"github.com/AvicennaJr/Nuru/object"
)

var OsFunctions = map[string]object.ModuleFunction{}

func init() {
	OsFunctions["toka"] = exit
}

func exit(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 0 {
		return &object.Null{}
	}

	os.Exit(0)

	return nil
}
