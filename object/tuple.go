package object

import "strings"

// Tuple is an immutable sequence of elements. Supports indexing (read), idadi, kitanzi, and iteration.
type Tuple struct {
	Elements []Object
	offset   int
	BaseRefCountable
}

// NewTuple returns a new tuple from the given elements. Retains each element.
func NewTuple(elems []Object) *Tuple {
	t := &Tuple{Elements: elems}
	Retain(t)
	for _, e := range elems {
		Retain(e)
	}
	return t
}

// GetOutgoingReferences returns all references held by the tuple.
func (t *Tuple) GetOutgoingReferences() []Object {
	return t.Elements
}

func (t *Tuple) Type() ObjectType { return TUPLE_OBJ }

func (t *Tuple) Inspect() string {
	if len(t.Elements) == 0 {
		return "jozi()"
	}
	parts := make([]string, len(t.Elements))
	for i, e := range t.Elements {
		if e != nil {
			parts[i] = e.Inspect()
		}
	}
	return "jozi(" + strings.Join(parts, ", ") + ")"
}

// Next implements Iterable.
func (t *Tuple) Next() (Object, Object) {
	if t.offset < len(t.Elements) {
		idx := t.offset
		val := t.Elements[t.offset]
		t.offset++
		return &Integer{Value: int64(idx)}, val
	}
	return nil, nil
}

// Reset implements Iterable.
func (t *Tuple) Reset() {
	t.offset = 0
}

func (t *Tuple) Method(method string, args []Object) Object {
	switch method {
	case "idadi":
		if len(args) != 0 {
			return newError("Samahani, idadi() haipokei hoja, wewe umeweka %d", len(args))
		}
		return &Integer{Value: int64(len(t.Elements))}
	case "kitanzi":
		return NewTupleIterator(t)
	default:
		return newError("Samahani, jozi haina method '%s'", method)
	}
}

// NewTupleIterator returns an iterator over the tuple. Retains the tuple.
func NewTupleIterator(t *Tuple) *Iterator {
	Retain(t)
	it := &Iterator{collection: t, kind: iteratorKindTuple, idx: 0}
	Retain(it)
	return it
}
