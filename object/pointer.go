package object

import (
	"fmt"
)

type Pointer struct {
	Ref Object
}

func (p *Pointer) Type() ObjectType {
	return POINTER_OBJ
}

func (p *Pointer) Inspect() string {
	if p.Ref == nil {
		return "Pointer(addr=nil, value=nil)"
	}
	return fmt.Sprintf("Pointer(addr=%p, value=%s)", p.Ref, p.Ref.Inspect())
}
