package object

import (
	"fmt"
	"sort"
)

// Iterator wraps a collection (Array, Dict, or String) with its own index
// so multiple iterators can be used independently without mutating the collection.
const (
	iteratorKindArray = iota
	iteratorKindDict
	iteratorKindString
)

type Iterator struct {
	collection Object
	kind      int
	idx       int
	// For dict we need stable order: sorted key strings, built once at creation
	dictKeys []string
	dictMap  map[string]DictPair // key Inspect() -> pair
	BaseRefCountable
}

// NewArrayIterator returns a new iterator over the array. Retains the array.
func NewArrayIterator(a *Array) *Iterator {
	Retain(a)
	it := &Iterator{collection: a, kind: iteratorKindArray, idx: 0}
	Retain(it)
	return it
}

// NewDictIterator returns a new iterator over the dict. Retains the dict.
func NewDictIterator(d *Dict) *Iterator {
	Retain(d)
	keys := make([]string, 0, len(d.Pairs))
	dictMap := make(map[string]DictPair)
	for _, pair := range d.Pairs {
		k := pair.Key.Inspect()
		dictMap[k] = pair
		keys = append(keys, k)
	}
	sort.Strings(keys)
	it := &Iterator{collection: d, kind: iteratorKindDict, idx: 0, dictKeys: keys, dictMap: dictMap}
	Retain(it)
	return it
}

// NewStringIterator returns a new iterator over the string. Retains the string.
func NewStringIterator(s *String) *Iterator {
	Retain(s)
	it := &Iterator{collection: s, kind: iteratorKindString, idx: 0}
	Retain(it)
	return it
}

func (it *Iterator) Type() ObjectType { return ITERATOR_OBJ }

func (it *Iterator) Inspect() string {
	return fmt.Sprintf("kitanzi(%s)", it.collection.Inspect())
}

// Next implements Iterable. Returns (keyOrIndex, value) or (nil, nil) when done.
func (it *Iterator) Next() (Object, Object) {
	switch it.kind {
	case iteratorKindArray:
		a := it.collection.(*Array)
		if it.idx < len(a.Elements) {
			key := &Integer{Value: int64(it.idx)}
			val := a.Elements[it.idx]
			it.idx++
			return key, val
		}
		return nil, nil
	case iteratorKindString:
		s := it.collection.(*String)
		if it.idx < len(s.Value) {
			key := &Integer{Value: int64(it.idx)}
			val := &String{Value: string(s.Value[it.idx])}
			it.idx++
			return key, val
		}
		return nil, nil
	case iteratorKindDict:
		if it.idx < len(it.dictKeys) {
			k := it.dictKeys[it.idx]
			pair := it.dictMap[k]
			it.idx++
			return pair.Key, pair.Value
		}
		return nil, nil
	}
	return nil, nil
}

// Reset implements Iterable.
func (it *Iterator) Reset() {
	it.idx = 0
}

// GetOutgoingReferences returns the wrapped collection for refcounting.
func (it *Iterator) GetOutgoingReferences() []Object {
	if it.collection == nil {
		return nil
	}
	return []Object{it.collection}
}

// ReleaseCollection releases the held collection. Called when iterator refcount hits 0.
func (it *Iterator) ReleaseCollection() {
	if it.collection != nil {
		Release(it.collection)
		it.collection = nil
	}
}
