package object

type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string {
	if b.Value {
		return "kweli"
	} else {
		return "sikweli"
	}
}
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

func (b *Boolean) HashKey() HashKey {
	var value uint64

	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return HashKey{Type: b.Type(), Value: value}
}
