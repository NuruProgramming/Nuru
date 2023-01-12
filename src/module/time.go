package module

import (
	"strconv"
	"time"

	"github.com/AvicennaJr/Nuru/object"
)

var TimeFunctions = map[string]object.ModuleFunction{}

func init() {
	TimeFunctions["hasahivi"] = now
	TimeFunctions["lala"] = sleep
}

func now(args []object.Object) object.Object {
	if len(args) != 0 {
		return &object.Error{Message: "hatuhitaji hoja kwenye hasahivi"}
	}

	tn := time.Now()

	return &object.String{Value: tn.Format("2006-01-02 15:04:05")}
}

func sleep(args []object.Object) object.Object {
	if len(args) != 1 {
		return &object.Error{Message: "tunahitaji hoja moja tu kwenye "}
	}

	objvalue := args[0].Inspect()
	inttime, err := strconv.Atoi(objvalue)

	if err != nil {
		return &object.Error{Message: "namba tu zinaruhusiwa kwenye hoja"}
	}

	time.Sleep(time.Duration(inttime) * time.Second)

	return nil
}
