package object

import (
	"time"
)

type Time struct {
	time string
}

func (t *Time) Type() ObjectType { return TIME_OBJ }
func (t *Time) Inspect() string  { return t.time }
func (t *Time) Method(method string, args []Object) Object {
	switch method {
	case "hasahivi":
		return t.now(args)

	}
	return nil
}

func (t *Time) now(args []Object) Object {
	tn := time.Now()

	return &String{Value: tn.Format("2006-01-02 15:04:05")}
}
