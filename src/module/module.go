package module

import "github.com/AvicennaJr/Nuru/object"

var Mapper = map[string]*object.Module{}

func init() {
	Mapper["os"] = &object.Module{Name: "os", Functions: OsFunctions}
}
