package object

import (
	"hash/fnv"
	"strings"
)

type String struct {
	Value  string
	offset int
}

func (s *String) Inspect() string  { return s.Value }
func (s *String) Type() ObjectType { return STRING_OBJ }
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))

	return HashKey{Type: s.Type(), Value: h.Sum64()}
}
func (s *String) Next() (Object, Object) {
	offset := s.offset
	if len(s.Value) > offset {
		s.offset = offset + 1
		return &Integer{Value: int64(offset)}, &String{Value: string(s.Value[offset])}
	}
	return nil, nil
}
func (s *String) Reset() {
	s.offset = 0
}
func (s *String) Method(method string, args []Object) Object {
	switch method {
	case "idadi":
		return s.len(args)
	case "herufikubwa":
		return s.upper(args)
	case "herufindogo":
		return s.lower(args)
	case "gawa":
		return s.split(args)
	default:
		return newError("Samahani, kiendesha hiki hakitumiki na tungo (Neno)")
	}
}

func (s *String) len(args []Object) Object {
	if len(args) != 0 {
		return newError("Samahani, tunahitaji Hoja 0, wewe umeweka %d", len(args))
	}
	return &Integer{Value: int64(len(s.Value))}
}

func (s *String) upper(args []Object) Object {
	if len(args) != 0 {
		return newError("Samahani, tunahitaji Hoja 0, wewe umeweka %d", len(args))
	}
	return &String{Value: strings.ToUpper(s.Value)}
}

func (s *String) lower(args []Object) Object {
	if len(args) != 0 {
		return newError("Samahani, tunahitaji Hoja 0, wewe umeweka %d", len(args))
	}
	return &String{Value: strings.ToLower(s.Value)}
}

func (s *String) split(args []Object) Object {
	if len(args) > 1 {
		return newError("Samahani, tunahitaji Hoja 1 au 0, wewe umeweka %d", len(args))
	}
	sep := " "
	if len(args) == 1 {
		sep = args[0].(*String).Value
	}
	parts := strings.Split(s.Value, sep)
	length := len(parts)
	elements := make([]Object, length)
	for k, v := range parts {
		elements[k] = &String{Value: v}
	}
	return &Array{Elements: elements}
}
