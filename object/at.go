package object

import "fmt"

type At struct {
	Instance *Instance
}

func (a *At) Type() ObjectType { return AT }
func (a *At) Inspect() string {
	return fmt.Sprintf("@.%s", a.Instance.Package.Name.Value)
}
