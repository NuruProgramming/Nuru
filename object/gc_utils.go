package object

import (
	"fmt"
)

// SafeAdd adds a value to a dictionary safely, using weak references for potential cycles
func SafeAdd(dict *Dict, key Object, value Object) {
	// Check if adding this value would create a cycle
	if isCyclicReference(dict, value) {
		// Use a weak reference instead
		dict.AddWeakReference(key, value)
		if VerboseGC {
			fmt.Printf("Onyo: Imegunduliwa marejeleo ya mduara yanayoweza kutokea. Kutumia rejeleo dhaifu kwa ufunguo %s\n", key.Inspect())
		}
	} else {
		// Otherwise use normal reference
		dict.Set(key, value)
	}
}

// isCyclicReference checks if adding target to container would create a cycle
func isCyclicReference(container Object, target Object) bool {
	// Simple case: direct self-reference
	if container == target {
		return true
	}

	// Use DFS to check for cycles
	visited := make(map[Object]bool)
	return detectCycle(target, container, visited)
}

// detectCycle performs DFS to find a path from target to container
func detectCycle(current Object, target Object, visited map[Object]bool) bool {
	if visited[current] {
		return false // Already visited, no cycle through this path
	}

	visited[current] = true

	// Get references from the current object
	var refs []Object

	switch obj := current.(type) {
	case *Dict:
		for _, pair := range obj.Pairs {
			refs = append(refs, pair.Value)
		}
	case *Array:
		refs = obj.Elements
	case *Function:
		// Functions can reference through their environment
		if obj.Env != nil {
			for _, value := range obj.Env.GetAllValues() {
				refs = append(refs, value)
			}
		}
	}

	// Check if any reference is the target
	for _, ref := range refs {
		if ref == target {
			return true // Found a cycle
		}

		// Recursively check this reference
		if detectCycle(ref, target, visited) {
			return true
		}
	}

	return false
}

// RunGC forces a garbage collection cycle
func RunGC() {
	GlobalRefCounter.CollectGarbage()
}

// GetObjectCount returns the number of tracked objects
func GetObjectCount() int {
	GlobalRefCounter.mutex.Lock()
	defer GlobalRefCounter.mutex.Unlock()
	return len(GlobalRefCounter.objects)
}

// PrintObjectStats prints statistics about tracked objects
func PrintObjectStats() {
	GlobalRefCounter.mutex.Lock()
	defer GlobalRefCounter.mutex.Unlock()

	typeCounts := make(map[ObjectType]int)

	for obj := range GlobalRefCounter.objects {
		typeCounts[obj.Type()]++
	}

	fmt.Println("=== Takwimu za vitu ===")
	fmt.Printf("Jumla ya vitu: %d\n", len(GlobalRefCounter.objects))

	for objType, count := range typeCounts {
		fmt.Printf("  %s: %d\n", objType, count)
	}

	fmt.Println("========================")
}
