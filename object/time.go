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
func (t *Time) Method(method string, args []Object) Object {
	switch method {
	case "ongeza":
		return t.add(args)

	}
	return nil
}

func (t *Time) add(args []Object) Object {
	if len(args) != 1 {
		return newError("Samahani, tunahitaji Hoja 1, wewe umeweka %d", len(args))
	}

	cur_time, _ := time.Parse("2006-01-02 15:04:05", t.Inspect())

	objvalue := args[0].Inspect()
	inttime, err := strconv.Atoi(objvalue)

	if err != nil {
		return newError("namba tu zinaruhusiwa kwenye hoja")
	}

	next_time := cur_time.Add(time.Duration(inttime) * time.Hour)
	return &Time{TimeValue: string(next_time.Format("2006-01-02 15:04:05"))}
}
