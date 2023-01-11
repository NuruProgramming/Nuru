package object

type ModuleFunction func(args []Object) Object

type Module struct {
	Name      string
	Functions map[string]ModuleFunction
}

func (m *Module) Type() ObjectType { return MODULE_OBJ }
func (m *Module) Inspect() string  { return "Module: " + m.Name }
