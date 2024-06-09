package module

import (
	"fmt"
	"strconv"
	"time"

	"github.com/NuruProgramming/Nuru/object"
)

var TimeFunctions = map[string]object.ModuleFunction{}

func init() {
	TimeFunctions["hasahivi"] = now
	TimeFunctions["lala"] = sleep
	TimeFunctions["tangu"] = since
}

func now(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 0 || len(defs) != 0 {
		return &object.Error{Message: "hatuhitaji hoja kwenye hasahivi"}
	}

	tn := time.Now()
	time_string := tn.Format("15:04:05 02-01-2006")

	return &object.Time{TimeValue: time_string}
}

func sleep(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "Hoja hii hairuhusiwi"}
	}
	if len(args) != 1 {
		return &object.Error{Message: "tunahitaji hoja moja tu"}
	}

	objvalue := args[0].Inspect()
	inttime, err := strconv.Atoi(objvalue)

	if err != nil {
		return &object.Error{Message: "namba tu zinaruhusiwa kwenye hoja"}
	}

	time.Sleep(time.Duration(inttime) * time.Second)

	return nil
}

func since(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 {
		return &object.Error{Message: "Hoja hii hairuhusiwi"}
	}
	if len(args) != 1 {
		return &object.Error{Message: "tunahitaji hoja moja tu"}
	}

	var (
		t   time.Time
		err error
	)

	switch m := args[0].(type) {
	case *object.Time:
		t, _ = time.Parse("15:04:05 02-01-2006", m.TimeValue)
	case *object.String:
		t, err = time.Parse("15:04:05 02-01-2006", m.Value)
		if err != nil {
			return &object.Error{Message: fmt.Sprintf("Hoja %s sii sahihi", args[0].Inspect())}
		}
	default:
		return &object.Error{Message: fmt.Sprintf("Hoja %s sii sahihi", args[0].Inspect())}
	}

	current_time := time.Now().Format("15:04:05 02-01-2006")
	ct, _ := time.Parse("15:04:05 02-01-2006", current_time)

	diff := ct.Sub(t)
	durationInSeconds := diff.Seconds()

	return &object.Integer{Value: int64(durationInSeconds)}
}
