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
	TimeFunctions["tangu"] = since
}

func now(args []object.Object) object.Object {
	if len(args) != 0 {
		return &object.Error{Message: "hatuhitaji hoja kwenye hasahivi"}
	}

	tn := time.Now()
	time_string := tn.Format("2006-01-02 15:04:05")

	return &object.Time{TimeValue: time_string}
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

func since(args []object.Object) object.Object {

	if len(args) != 1 {
		return &object.Error{Message: "tunahitaji hoja moja tu kwenye "}
	}

	t, err := time.Parse("2006-01-02 15:04:05", args[0].Inspect())

	if err != nil {

		return &object.Error{Message: "tumeshindwa kuparse hoja zako"}
	}

	current_time := time.Now().Format("2006-01-02 15:04:05")
	ct, _ := time.Parse("2006-01-02 15:04:05", current_time)

	diff := ct.Sub(t)
	durationInSeconds := diff.Seconds()

	return &object.Float{Value: durationInSeconds}
}
