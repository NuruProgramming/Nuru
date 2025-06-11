package object

import "fmt"

type Hexadecimal struct {
	Value int64
}

func (i *Hexadecimal) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Hexadecimal) Type() ObjectType { return INTEGER_OBJ }

func (i *Hexadecimal) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}
