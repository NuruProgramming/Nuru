package object

import (
	"fmt"
	"sync"
)

// VerboseGC controls whether garbage collection messages are displayed
var VerboseGC bool = false

// RefCountable is an interface for objects that can be reference counted
type RefCountable interface {
	Object
	IncRefCount()
	DecRefCount() int
	GetRefCount() int
	MarkColor(color Color)
	GetColor() Color
	GetOutgoingReferences() []Object
}

// Color represents the mark state for cycle detection
type Color int

const (
	White Color = iota // Not visited
	Gray               // Being visited
	Black              // Visited completely
)

// RefCounter manages reference counting and garbage collection
type RefCounter struct {
	objects     map[Object]bool
	mutex       sync.Mutex
	gcThreshold int
}

// Global reference counter
var GlobalRefCounter = NewRefCounter()

// NewRefCounter creates a new reference counter
func NewRefCounter() *RefCounter {
	return &RefCounter{
		objects:     make(map[Object]bool),
		gcThreshold: 1000, // Run GC after this many objects are tracked
	}
}

// TrackObject adds an object to the reference counter
func (rc *RefCounter) TrackObject(obj Object) {
	// Skip basic types like integers, booleans, etc.
	switch obj.Type() {
	case INTEGER_OBJ, BOOLEAN_OBJ, FLOAT_OBJ, NULL_OBJ, STRING_OBJ:
		return
	}

	rc.mutex.Lock()
	defer rc.mutex.Unlock()

	rc.objects[obj] = true

	// Check if we need to run a garbage collection cycle
	if len(rc.objects) > rc.gcThreshold {
		go rc.CollectGarbage()
	}
}

// CollectGarbage runs a garbage collection cycle
func (rc *RefCounter) CollectGarbage() {
	rc.mutex.Lock()
	defer rc.mutex.Unlock()

	// Reset colors of all objects
	for obj := range rc.objects {
		if refObj, ok := obj.(RefCountable); ok {
			refObj.MarkColor(White)
		}
	}

	// Find potential roots of cycles (objects with refCount > 0)
	var roots []RefCountable
	for obj := range rc.objects {
		if refObj, ok := obj.(RefCountable); ok {
			if refObj.GetRefCount() > 0 {
				roots = append(roots, refObj)
			}
		}
	}

	// Mark phase - find unreachable cycles
	for _, root := range roots {
		rc.markCycles(root)
	}

	// Sweep phase - collect objects in cycles
	var toRemove []Object
	for obj := range rc.objects {
		if refObj, ok := obj.(RefCountable); ok {
			// White objects are in unreachable cycles
			if refObj.GetColor() == White && refObj.GetRefCount() > 0 {
				// Break the cycle
				for _, ref := range refObj.GetOutgoingReferences() {
					if refChild, ok := ref.(RefCountable); ok {
						refChild.DecRefCount()
					}
				}

				// Mark for removal if refCount is now 0
				if refObj.GetRefCount() == 0 {
					toRemove = append(toRemove, obj)
				}
			}
		}
	}

	// Remove collected objects
	for _, obj := range toRemove {
		delete(rc.objects, obj)
	}

	// Only print garbage collection messages when verbose mode is enabled
	if VerboseGC {
		fmt.Printf("Ukusanyaji wa takataka umekamilika: %d vitu kuondolewa\n", len(toRemove))
	}
}

// markCycles performs DFS to detect cycles
func (rc *RefCounter) markCycles(obj RefCountable) {
	obj.MarkColor(Gray) // Mark as being visited

	// Visit all children
	for _, child := range obj.GetOutgoingReferences() {
		if refChild, ok := child.(RefCountable); ok {
			switch refChild.GetColor() {
			case White:
				rc.markCycles(refChild)
			case Gray:
				// Found a cycle
			case Black:
				// Already fully visited
			}
		}
	}

	obj.MarkColor(Black) // Mark as fully visited
}

// BaseRefCountable provides base implementation for RefCountable objects
type BaseRefCountable struct {
	refCount int
	color    Color
}

// IncRefCount increments the reference count
func (b *BaseRefCountable) IncRefCount() {
	b.refCount++
}

// DecRefCount decrements the reference count and returns the new count
func (b *BaseRefCountable) DecRefCount() int {
	b.refCount--
	return b.refCount
}

// GetRefCount returns the current reference count
func (b *BaseRefCountable) GetRefCount() int {
	return b.refCount
}

// MarkColor sets the mark color for cycle detection
func (b *BaseRefCountable) MarkColor(c Color) {
	b.color = c
}

// GetColor returns the current mark color
func (b *BaseRefCountable) GetColor() Color {
	return b.color
}

// Release decrements the reference count and checks if the object should be collected
func Release(obj Object) {
	if refObj, ok := obj.(RefCountable); ok {
		count := refObj.DecRefCount()
		if count == 0 {
			// The object is no longer referenced, remove it from tracking
			GlobalRefCounter.mutex.Lock()
			delete(GlobalRefCounter.objects, obj)
			GlobalRefCounter.mutex.Unlock()
		}
	}
}

// Retain increments the reference count of an object
func Retain(obj Object) Object {
	if refObj, ok := obj.(RefCountable); ok {
		refObj.IncRefCount()
		GlobalRefCounter.TrackObject(obj)
	}
	return obj
}

// GetTypeStats returns a map of object types to their counts
func GetTypeStats() map[ObjectType]int {
	GlobalRefCounter.mutex.Lock()
	defer GlobalRefCounter.mutex.Unlock()

	typeCounts := make(map[ObjectType]int)

	for obj := range GlobalRefCounter.objects {
		typeCounts[obj.Type()]++
	}

	return typeCounts
}
