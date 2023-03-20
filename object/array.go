package object

import (
	"bytes"
	"strings"
)

type Array struct {
	Elements []Object
	offset   int
}

func (ao *Array) Type() ObjectType { return ARRAY_OBJ }
func (ao *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

func (ao *Array) Next() (Object, Object) {
	idx := ao.offset
	if len(ao.Elements) > idx {
		ao.offset = idx + 1
		return &Integer{Value: int64(idx)}, ao.Elements[idx]
	}
	return nil, nil
}

func (ao *Array) Reset() {
	ao.offset = 0
}

func (a *Array) Method(method string, args []Object) Object {
	switch method {
	case "idadi":
		return a.len(args)
	case "sukuma":
		return a.push(args)
	case "yamwisho":
		return a.last()
	default:
		return newError("Samahani, function hii haitumiki na Strings (Neno)")
	}
}

func (a *Array) len(args []Object) Object {
	if len(args) != 0 {
		return newError("Samahani, tunahitaji Hoja 0, wewe umeweka %d", len(args))
	}
	return &Integer{Value: int64(len(a.Elements))}
}

func (a *Array) last() Object {
	length := len(a.Elements)
	if length > 0 {
		return a.Elements[length-1]
	}
	return &Null{}
}

func (a *Array) push(args []Object) Object {
	a.Elements = append(a.Elements, args...)
	return a
}
