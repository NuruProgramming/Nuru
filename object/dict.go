package object

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

// Global map to track objects that have been visited during inspection
// This helps prevent infinite recursion due to circular references
var inspectionDepth = 0
var maxInspectionDepth = 30

type DictPair struct {
	Key   Object
	Value Object
}

type Dict struct {
	Pairs            map[HashKey]DictPair
	offset           int
	BaseRefCountable // Add the BaseRefCountable to inherit reference counting functionality
}

// GetOutgoingReferences returns all references held by this dict
func (d *Dict) GetOutgoingReferences() []Object {
	var refs []Object
	for _, pair := range d.Pairs {
		refs = append(refs, pair.Key)
		refs = append(refs, pair.Value)
	}
	return refs
}

func (d *Dict) Type() ObjectType { return DICT_OBJ }
func (d *Dict) Inspect() string {
	return d.InspectWithDepth(0)
}

func (d *Dict) InspectWithDepth(depth int) string {
	return d.inspectWithTracking(depth, make(map[*Dict]bool))
}

// inspectWithTracking uses a visited map to track dictionaries that have already been visited
// to prevent infinite recursion when there are circular references
func (d *Dict) inspectWithTracking(depth int, visited map[*Dict]bool) string {
	var out bytes.Buffer

	// Check if we've already visited this dictionary in this inspection call
	if visited[d] {
		return "{...circular...}" // Indicate that there's a circular reference
	}

	// Check if we've reached the maximum inspection depth
	if depth > maxInspectionDepth {
		return "{...}" // Return a placeholder to indicate truncated output
	}

	// Mark this dictionary as visited
	visited[d] = true

	pairs := []string{}

	for _, pair := range d.Pairs {
		var keyStr string
		var valueStr string

		// Safely inspect keys and values
		if key, ok := pair.Key.(*Dict); ok {
			keyStr = key.inspectWithTracking(depth+1, visited)
		} else if key, ok := pair.Key.(*Array); ok {
			keyStr = key.InspectWithDepth(depth + 1)
		} else {
			keyStr = pair.Key.Inspect()
		}

		if value, ok := pair.Value.(*Dict); ok {
			valueStr = value.inspectWithTracking(depth+1, visited)
		} else if value, ok := pair.Value.(*Array); ok {
			valueStr = value.InspectWithDepth(depth + 1)
		} else {
			valueStr = pair.Value.Inspect()
		}

		pairs = append(pairs, fmt.Sprintf("%s: %s", keyStr, valueStr))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	// Remove this dictionary from the visited set when we're done with it
	delete(visited, d)

	return out.String()
}

func (d *Dict) Next() (Object, Object) {
	idx := 0
	dict := make(map[string]DictPair)
	var keys []string
	for _, v := range d.Pairs {
		dict[v.Key.Inspect()] = v
		keys = append(keys, v.Key.Inspect())
	}

	sort.Strings(keys)

	for _, k := range keys {
		if d.offset == idx {
			d.offset += 1
			return dict[k].Key, dict[k].Value
		}
		idx += 1
	}
	return nil, nil
}

func (d *Dict) Reset() {
	d.offset = 0
}

// AddWeakReference adds a value with a weak reference to prevent circular reference memory leaks
func (d *Dict) AddWeakReference(key Object, value Object) {
	if hashKey, ok := key.(Hashable); ok {
		hashed := hashKey.HashKey()
		weakRef := NewWeakReference(value)
		d.Pairs[hashed] = DictPair{Key: key, Value: weakRef}
	}
}

// GetAndResolveValue gets a value, resolving weak references if necessary
func (d *Dict) GetAndResolveValue(key Object) (Object, bool) {
	if hashKey, ok := key.(Hashable); ok {
		hashed := hashKey.HashKey()
		if pair, ok := d.Pairs[hashed]; ok {
			// If this is a weak reference, resolve it
			if weakRef, ok := pair.Value.(*WeakReference); ok {
				return weakRef.GetTarget(), true
			}
			return pair.Value, true
		}
	}
	return nil, false
}

// NewDict creates a new reference-counted dictionary
func NewDict() *Dict {
	dict := &Dict{
		Pairs: make(map[HashKey]DictPair),
	}
	// Track the new object in the reference counter
	GlobalRefCounter.TrackObject(dict)
	return dict
}

// Set adds or updates a key-value pair, handling reference counting
func (d *Dict) Set(key Object, value Object) {
	if hashable, ok := key.(Hashable); ok {
		hash := hashable.HashKey()

		// If replacing an existing value, decrement its reference count
		if pair, exists := d.Pairs[hash]; exists {
			Release(pair.Value)
		}

		// Increment reference count for the new value
		Retain(value)

		d.Pairs[hash] = DictPair{Key: key, Value: value}
	}
}

// Get retrieves a value by key
func (d *Dict) Get(key Object) (Object, bool) {
	if hashable, ok := key.(Hashable); ok {
		hash := hashable.HashKey()
		if pair, ok := d.Pairs[hash]; ok {
			return pair.Value, true
		}
	}
	return nil, false
}

// Delete removes a key-value pair, handling reference counting
func (d *Dict) Delete(key Object) bool {
	if hashable, ok := key.(Hashable); ok {
		hash := hashable.HashKey()
		if pair, exists := d.Pairs[hash]; exists {
			// Decrement reference count for the removed value
			Release(pair.Value)
			delete(d.Pairs, hash)
			return true
		}
	}
	return false
}
