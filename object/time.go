package object

import (
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
