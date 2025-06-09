package object

import (
	"fmt"
)

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer

	// Retain reference to outer environment
	if outer != nil {
		Retain(outer)
	}

	return env
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	env := &Environment{store: s, outer: nil}

	// Track the new environment
	GlobalRefCounter.TrackObject(env)

	return env
}

type Environment struct {
	store            map[string]Object
	outer            *Environment
	BaseRefCountable // Add the BaseRefCountable to inherit reference counting functionality
}

// Implement Object interface
func (e *Environment) Type() ObjectType { return ENVIRONMENT_OBJ }
func (e *Environment) Inspect() string {
	return fmt.Sprintf("environment: %d variables", len(e.store))
}

// GetOutgoingReferences returns all references held by this environment
func (e *Environment) GetOutgoingReferences() []Object {
	var refs []Object

	// Add all stored values
	for _, obj := range e.store {
		refs = append(refs, obj)
	}

	// Add outer environment
	if e.outer != nil {
		refs = append(refs, e.outer)
	}

	return refs
}

// GetAllValues returns all values stored in this environment
func (e *Environment) GetAllValues() []Object {
	values := make([]Object, 0, len(e.store))
	for _, obj := range e.store {
		values = append(values, obj)
	}
	return values
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]

	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	// If replacing an existing value, release the old one
	if oldVal, exists := e.store[name]; exists {
		Release(oldVal)
	}

	// Retain the new value
	Retain(val)

	e.store[name] = val
	return val
}

func (e *Environment) Del(name string) bool {
	_, ok := e.store[name]
	if ok {
		// Release the value being deleted
		Release(e.store[name])
		delete(e.store, name)
	}
	return true
}
