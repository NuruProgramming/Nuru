package object

import (
	"bytes"
	"strings"
)

// Helper functions that need to be defined in this package
func isError(obj Object) bool {
	if obj != nil {
		return obj.Type() == ERROR_OBJ
	}
	return false
}

func ApplyFunction(fn *Function, args []Object) Object {
	// This is a placeholder - the actual implementation should be imported or defined elsewhere
	// For now, we'll return an error to indicate this needs to be properly implemented
	return &Error{Message: "ApplyFunction not implemented in array.go"}
}

// Array struct represents an array in Nuru
type Array struct {
	Elements         []Object
	offset           int
	BaseRefCountable // Add the BaseRefCountable to inherit reference counting functionality
}

// GetOutgoingReferences returns all references held by this array
func (ao *Array) GetOutgoingReferences() []Object {
	return ao.Elements
}

func (ao *Array) Type() ObjectType { return ARRAY_OBJ }
func (ao *Array) Inspect() string {
	return ao.InspectWithDepth(0)
}

func (ao *Array) InspectWithDepth(depth int) string {
	return ao.inspectWithTracking(depth, make(map[*Array]bool), make(map[*Dict]bool))
}

// inspectWithTracking uses visited maps to track arrays and dictionaries that have already been visited
// to prevent infinite recursion when there are circular references
func (ao *Array) inspectWithTracking(depth int, visitedArrays map[*Array]bool, visitedDicts map[*Dict]bool) string {
	var out bytes.Buffer

	// Check if we've already visited this array in this inspection call
	if visitedArrays[ao] {
		return "[...circular...]" // Indicate that there's a circular reference
	}

	// Check if we've reached the maximum inspection depth
	if depth > maxInspectionDepth {
		return "[...]" // Return a placeholder to indicate truncated output
	}

	// Mark this array as visited
	visitedArrays[ao] = true

	elements := []string{}
	if len(ao.Elements) != 0 {
		for _, e := range ao.Elements {
			if e == nil {
				continue
			}

			var elemStr string
			if dict, ok := e.(*Dict); ok {
				// Use the Dict's tracking method with our visitedDicts map
				elemStr = dict.inspectWithTracking(depth+1, visitedDicts)
			} else if arr, ok := e.(*Array); ok {
				// Recursively call our tracking method
				elemStr = arr.inspectWithTracking(depth+1, visitedArrays, visitedDicts)
			} else {
				elemStr = e.Inspect()
			}

			if elemStr != "" {
				elements = append(elements, elemStr)
			}
		}
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	// Remove this array from the visited set when we're done with it
	delete(visitedArrays, ao)

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

// AddWeakReference adds an element as a weak reference to prevent circular reference memory leaks
func (ao *Array) AddWeakReference(index int, value Object) bool {
	if index < 0 || index > len(ao.Elements) {
		return false
	}

	// If the index is at the end, append the weak reference
	if index == len(ao.Elements) {
		ao.Elements = append(ao.Elements, NewWeakReference(value))
	} else {
		// Otherwise, replace the element at the given index
		ao.Elements[index] = NewWeakReference(value)
	}
	return true
}

// GetAndResolveElement gets an element, resolving weak references if necessary
func (ao *Array) GetAndResolveElement(index int) (Object, bool) {
	if index < 0 || index >= len(ao.Elements) {
		return nil, false
	}

	element := ao.Elements[index]

	// If this is a weak reference, resolve it
	if weakRef, ok := element.(*WeakReference); ok {
		return weakRef.GetTarget(), true
	}

	return element, true
}

func (a *Array) Method(method string, args []Object) Object {
	switch method {
	case "idadi":
		return a.len(args)
	case "sukuma":
		return a.push(args)
	case "yamwisho":
		return a.last()
	case "unga":
		return a.join(args)
	case "chuja":
		return a.filter(args)
	case "tafuta":
		return a.find(args)
	case "kitanzi":
		return NewArrayIterator(a)
	case "sukumaKamaRef": // New method to add an element as a weak reference
		return a.pushAsWeakRef(args)
	// Methods like ramani and punguza require access to the evaluator environment
	// and will be handled in the evaluator package when processing method calls
	default:
		return newError("Samahani, kiendesha hiki hakitumiki na tungo (Neno)")
	}
}

// New method to push elements as weak references
func (a *Array) pushAsWeakRef(args []Object) Object {
	for _, arg := range args {
		a.Elements = append(a.Elements, NewWeakReference(arg))
	}
	return a
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
		// If the last element is a weak reference, resolve it
		if weakRef, ok := a.Elements[length-1].(*WeakReference); ok {
			return weakRef.GetTarget()
		}
		return a.Elements[length-1]
	}
	return &Null{}
}

func (a *Array) push(args []Object) Object {
	a.Elements = append(a.Elements, args...)
	return a
}

func (a *Array) join(args []Object) Object {
	if len(args) > 1 {
		return newError("Samahani, tunahitaji Hoja 1 au 0, wewe umeweka %d", len(args))
	}
	if len(a.Elements) > 0 {
		glue := ""
		if len(args) == 1 {
			glue = args[0].(*String).Value
		}
		length := len(a.Elements)
		newElements := make([]string, length)
		for k, v := range a.Elements {
			newElements[k] = v.Inspect()
		}
		return &String{Value: strings.Join(newElements, glue)}
	} else {
		return &String{Value: ""}
	}
}

func (a *Array) filter(args []Object) Object {
	if len(args) != 1 {
		return newError("Samahani, idadi ya hoja sii sahihi")
	}

	dummy := []Object{}
	filteredArr := Array{Elements: dummy}
	for _, obj := range a.Elements {
		if obj.Inspect() == args[0].Inspect() && obj.Type() == args[0].Type() {
			filteredArr.Elements = append(filteredArr.Elements, obj)
		}
	}
	return &filteredArr
}

func (a *Array) find(args []Object) Object {
	if len(args) != 1 {
		return newError("Samahani, idadi ya hoja sii sahihi")
	}

	for _, obj := range a.Elements {
		if obj.Inspect() == args[0].Inspect() && obj.Type() == args[0].Type() {
			return obj
		}
	}
	return &Null{}
}

// NewArray creates a new reference-counted array
func NewArray(elements []Object) *Array {
	array := &Array{
		Elements: elements,
	}

	// Increment reference count for all elements
	for _, elem := range elements {
		Retain(elem)
	}

	// Track the new object in the reference counter
	GlobalRefCounter.TrackObject(array)
	return array
}

// Append adds elements to the array, handling reference counting
func (ao *Array) Append(elements ...Object) {
	for _, elem := range elements {
		Retain(elem)
		ao.Elements = append(ao.Elements, elem)
	}
}

// Set updates an element at an index, handling reference counting
func (ao *Array) Set(index int, value Object) {
	if index >= 0 && index < len(ao.Elements) {
		// Release the old element
		Release(ao.Elements[index])

		// Retain the new element
		Retain(value)

		ao.Elements[index] = value
	}
}

// Get retrieves an element at an index
func (ao *Array) Get(index int) Object {
	if index >= 0 && index < len(ao.Elements) {
		return ao.Elements[index]
	}
	return &Null{}
}

// Remove removes an element at an index, handling reference counting
func (ao *Array) Remove(index int) {
	if index >= 0 && index < len(ao.Elements) {
		// Release the element being removed
		Release(ao.Elements[index])

		// Remove the element
		ao.Elements = append(ao.Elements[:index], ao.Elements[index+1:]...)
	}
}
