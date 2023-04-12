package object

type ModuleFunction func(args []Object, defs map[string]Object) Object

type Module struct {
	Name      string
	Functions map[string]ModuleFunction
}

func (m *Module) Type() ObjectType {
	switch m.Name {
	case "time":
		return TIME_OBJ
	case "json":
		return JSON_OBJ
	default:
		return MODULE_OBJ
	}
}
func (m *Module) Inspect() string { return "Module: " + m.Name }
