package module

import (
	"time"

	"github.com/AvicennaJr/Nuru/object"
)

var TimeFunctions = map[string]object.ModuleFunction{}

func init() {
	TimeFunctions["hasahivi"] = now
}

func now(args []object.Object) object.Object {
	if len(args) != 0 {
		return &object.Error{Message : "hatuhitaji hoja kwenye hasahivi"}
	}
 
	tn := time.Now()

	return &object.String{Value: tn.Format("2006-01-02 15:04:05")}
}
