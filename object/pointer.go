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
		return "Pointer(anwani=nil, thamani=nil)"
	}
	return fmt.Sprintf("Pointer(anwani=%p, thamani=%s)", p.Ref, p.Ref.Inspect())
}
