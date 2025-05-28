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

func today(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 0 || len(defs) != 0 {
		return &object.Error{Message: "hatuhitaji hoja kwenye leo"}
	}

	dateStr := time.Now().Format("02-01-2006")
	return &object.String{Value: dateStr}
}

func after(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 || len(args) != 1 {
		return &object.Error{Message: "tunahitaji hoja moja tu kwenye baada_ya"}
	}

	secondsStr := args[0].Inspect()
	seconds, err := strconv.Atoi(secondsStr)
	if err != nil {
		return &object.Error{Message: "hoja lazima iwe namba"}
	}

	future := time.Now().Add(time.Duration(seconds) * time.Second)
	return &object.Time{TimeValue: future.Format("15:04:05 02-01-2006")}
}

func diff(args []object.Object, defs map[string]object.Object) object.Object {
	if len(defs) != 0 || len(args) != 2 {
		return &object.Error{Message: "tunahitaji hoja mbili kwenye tofauti"}
	}

	parseTime := func(o object.Object) (time.Time, error) {
		switch v := o.(type) {
		case *object.Time:
			return time.Parse("15:04:05 02-01-2006", v.TimeValue)
		case *object.String:
			return time.Parse("15:04:05 02-01-2006", v.Value)
		default:
			return time.Time{}, fmt.Errorf("aina batili")
		}
	}

	t1, err1 := parseTime(args[0])
	t2, err2 := parseTime(args[1])

	if err1 != nil || err2 != nil {
		return &object.Error{Message: "tofauti inahitaji nyakati halali mbili"}
	}

	diff := t1.Sub(t2).Seconds()
	return &object.Integer{Value: int64(diff)}
}

func init() {
	TimeFunctions["hasahivi"] = now
	TimeFunctions["lala"] = sleep
	TimeFunctions["tangu"] = since
	TimeFunctions["leo"] = today
	TimeFunctions["baada_ya"] = after
	TimeFunctions["tofauti"] = diff
}


