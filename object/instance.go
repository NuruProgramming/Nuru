package object

import "fmt"

type Instance struct {
	Package *Package
	Env     *Environment
}

func (i *Instance) Type() ObjectType { return INSTANCE }
func (i *Instance) Inspect() string {
	return fmt.Sprintf("Pakeji: %s", i.Package.Name.Value)
}
