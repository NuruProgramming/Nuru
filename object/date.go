package object

import (
	"fmt"
	"time"
)

// Date represents a calendar date (year, month, day) without time.
type Date struct {
	Year  int
	Month int
	Day   int
}

func (d *Date) Type() ObjectType { return DATE_OBJ }

func (d *Date) Inspect() string {
	return fmt.Sprintf("%04d-%02d-%02d", d.Year, d.Month, d.Day)
}

func (d *Date) Method(method string, args []Object) Object {
	switch method {
	case "panga":
		return d.panga(args)
	default:
		return newError("Samahani, tarehe haina method '%s'", method)
	}
}

func (d *Date) panga(args []Object) Object {
	if len(args) != 1 {
		return newError("Samahani, panga inahitaji hoja 1 (muundo), wewe umeweka %d", len(args))
	}
	layout, ok := args[0].(*String)
	if !ok {
		return newError("Samahani, muundo lazima uwe neno")
	}
	t := time.Date(d.Year, time.Month(d.Month), d.Day, 0, 0, 0, 0, time.UTC)
	return &String{Value: t.Format(layout.Value)}
}
