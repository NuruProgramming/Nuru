package object

import (
	"sort"
	"strings"
)

// Set is a collection of unique hashable elements. Iteration order is by element Inspect().
type Set struct {
	items  map[HashKey]Object
	order  []Object // sorted by Inspect() for deterministic iteration
	offset int
	BaseRefCountable
}

// NewSet returns a new empty set.
func NewSet() *Set {
	s := &Set{
		items: make(map[HashKey]Object),
		order: nil,
	}
	Retain(s)
	return s
}

// GetOutgoingReferences returns all references held by the set.
func (s *Set) GetOutgoingReferences() []Object {
	return s.order
}

func (s *Set) Type() ObjectType { return SET_OBJ }

func (s *Set) Inspect() string {
	if len(s.order) == 0 {
		return "seti()"
	}
	parts := make([]string, len(s.order))
	for i, o := range s.order {
		parts[i] = o.Inspect()
	}
	return "seti(" + strings.Join(parts, ", ") + ")"
}

// Next implements Iterable. Returns (index, value) or (nil, nil) when done.
func (s *Set) Next() (Object, Object) {
	if s.offset < len(s.order) {
		idx := s.offset
		val := s.order[s.offset]
		s.offset++
		return &Integer{Value: int64(idx)}, val
	}
	return nil, nil
}

// Reset implements Iterable.
func (s *Set) Reset() {
	s.offset = 0
}

// Add adds a hashable element to the set. Returns the set for chaining.
func (s *Set) Add(obj Object) Object {
	hashable, ok := obj.(Hashable)
	if !ok {
		return newError("Samahani, kipengele cha seti kinapaswa kuwa kinachoweza kukokotolewa (namna: namba, neno, booleani), umeweka %s", obj.Type())
	}
	key := hashable.HashKey()
	if _, exists := s.items[key]; exists {
		return s
	}
	s.items[key] = obj
	Retain(obj)
	s.order = append(s.order, obj)
	sort.Slice(s.order, func(i, j int) bool {
		return s.order[i].Inspect() < s.order[j].Inspect()
	})
	return s
}

// Has returns kweli if the set contains the given key (left side of "ktk").
func (s *Set) Has(obj Object) bool {
	hashable, ok := obj.(Hashable)
	if !ok {
		return false
	}
	_, ok = s.items[hashable.HashKey()]
	return ok
}

// Remove removes the element from the set. No-op if not present.
func (s *Set) Remove(obj Object) Object {
	hashable, ok := obj.(Hashable)
	if !ok {
		return newError("Samahani, kipengele cha seti kinapaswa kuwa kinachoweza kukokotolewa")
	}
	key := hashable.HashKey()
	if v, exists := s.items[key]; exists {
		delete(s.items, key)
		Release(v)
		// rebuild order without v
		newOrder := make([]Object, 0, len(s.order)-1)
		for _, o := range s.order {
			if o != v {
				newOrder = append(newOrder, o)
			}
		}
		s.order = newOrder
	}
	return s
}

// Len returns the number of elements.
func (s *Set) Len() int {
	return len(s.order)
}

// Method dispatches set methods.
func (s *Set) Method(method string, args []Object) Object {
	switch method {
	case "idadi":
		return &Integer{Value: int64(len(s.order))}
	case "ona":
		return s.ona(args)
	case "ongeza":
		return s.ongeza(args)
	case "ondoa":
		return s.ondoa(args)
	case "kitanzi":
		return NewSetIterator(s)
	default:
		return newError("Samahani, seti haina method '%s'", method)
	}
}

func (s *Set) ona(args []Object) Object {
	if len(args) != 1 {
		return newError("Samahani, ona() inahitaji hoja 1, wewe umeweka %d", len(args))
	}
	return &Boolean{Value: s.Has(args[0])}
}

func (s *Set) ongeza(args []Object) Object {
	for _, arg := range args {
		if res := s.Add(arg); isError(res) {
			return res
		}
	}
	return s
}

func (s *Set) ondoa(args []Object) Object {
	if len(args) != 1 {
		return newError("Samahani, ondoa() inahitaji hoja 1, wewe umeweka %d", len(args))
	}
	s.Remove(args[0])
	return s
}

// NewSetIterator returns an iterator over the set. Retains the set.
func NewSetIterator(s *Set) *Iterator {
	Retain(s)
	it := &Iterator{collection: s, kind: iteratorKindSet, idx: 0}
	Retain(it)
	return it
}
