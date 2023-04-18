package object

import (
	"fmt"
	"strconv"
	"time"
)

type Time struct {
	TimeValue string
}

func (t *Time) Type() ObjectType { return TIME_OBJ }
func (t *Time) Inspect() string  { return t.TimeValue }
func (t *Time) Method(method string, args []Object, defs map[string]Object) Object {
	switch method {
	case "ongeza":
		return t.add(args, defs)
	case "tangu":
		return t.since(args, defs)
	}
	return nil
}

func (t *Time) add(args []Object, defs map[string]Object) Object {
	if len(defs) != 0 {
		var sec, min, hr, d, m, y int
		for k, v := range defs {
			objvalue := v.Inspect()
			inttime, err := strconv.Atoi(objvalue)
			if err != nil {
				return newError("namba tu zinaruhusiwa kwenye hoja")
			}
			switch k {
			case "sekunde":
				sec = inttime
			case "dakika":
				min = inttime
			case "saa":
				hr = inttime
			case "siku":
				d = inttime
			case "miezi":
				m = inttime
			case "miaka":
				y = inttime
			default:
				return newError("Hukuweka muda sahihi")
			}
		}
		cur_time, _ := time.Parse("15:04:05 02-01-2006", t.Inspect())
		next_time := cur_time.
			Add(time.Duration(sec)*time.Second).
			Add(time.Duration(min)*time.Minute).
			Add(time.Duration(hr)*time.Hour).
			AddDate(y, m, d)
		return &Time{TimeValue: string(next_time.Format("15:04:05 02-01-2006"))}
	}

	if len(args) != 1 {
		return newError("Samahani, tunahitaji Hoja 1, wewe umeweka %d", len(args))
	}

	cur_time, _ := time.Parse("15:04:05 02-01-2006", t.Inspect())

	objvalue := args[0].Inspect()
	inttime, err := strconv.Atoi(objvalue)

	if err != nil {
		return newError("namba tu zinaruhusiwa kwenye hoja")
	}

	next_time := cur_time.Add(time.Duration(inttime) * time.Hour)
	return &Time{TimeValue: string(next_time.Format("15:04:05 02-01-2006"))}
}

func (t *Time) since(args []Object, defs map[string]Object) Object {
	if len(defs) != 0 {
		return &Error{Message: "Hoja hii hairuhusiwi"}
	}
	if len(args) != 1 {
		return &Error{Message: "tunahitaji hoja moja tu"}
	}

	var (
		o   time.Time
		err error
	)

	switch m := args[0].(type) {
	case *Time:
		o, _ = time.Parse("15:04:05 02-01-2006", m.TimeValue)
	case *String:
		o, err = time.Parse("15:04:05 02-01-2006", m.Value)
		if err != nil {
			return &Error{Message: fmt.Sprintf("Hoja %s sii sahihi", args[0].Inspect())}
		}
	default:
		return &Error{Message: fmt.Sprintf("Hoja %s sii sahihi", args[0].Inspect())}
	}

	ct, _ := time.Parse("15:04:05 02-01-2006", t.TimeValue)

	diff := ct.Sub(o)
	durationInSeconds := diff.Seconds()

	return &Integer{Value: int64(durationInSeconds)}
}
