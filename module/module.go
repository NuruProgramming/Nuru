package module

import "github.com/NuruProgramming/Nuru/object"

var Mapper = map[string]*object.Module{}

func init() {
	Mapper["os"] = &object.Module{Name: "os", Functions: OsFunctions}
	Mapper["muda"] = &object.Module{Name: "time", Functions: TimeFunctions}
	Mapper["mtandao"] = &object.Module{Name: "net", Functions: NetFunctions}
	Mapper["jsoni"] = &object.Module{Name: "json", Functions: JsonFunctions}
	Mapper["hisabati"] = &object.Module{Name: "hisabati", Functions: MathFunctions}
}
