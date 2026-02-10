package object

import (
	"fmt"
)

type ObjectType string

const (
	INTEGER_OBJ      = "NAMBA"
	FLOAT_OBJ        = "DESIMALI"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "TUPU"
	RETURN_VALUE_OBJ = "RUDISHA"
	ERROR_OBJ        = "KOSA"
	FUNCTION_OBJ     = "UNDO (FUNCTION)"
	STRING_OBJ       = "NENO"
	BUILTIN_OBJ      = "YA_NDANI"
	ARRAY_OBJ        = "ORODHA"
	DICT_OBJ         = "KAMUSI"
	CONTINUE_OBJ     = "ENDELEA"
	BREAK_OBJ        = "VUNJA"
	FILE_OBJ         = "FAILI"
	TIME_OBJ         = "MUDA"
	JSON_OBJ         = "JSONI"
	MODULE_OBJ       = "MODULE"
	BYTE_OBJ         = "BYTE"
	PACKAGE_OBJ      = "PAKEJI"
	INSTANCE         = "PAKEJI"
	AT               = "@"
	WEAK_REF_OBJ     = "KUMBUKUMBU"
	ENVIRONMENT_OBJ  = "MAZINGIRA"
	BASE64_OBJ       = "BASE64"
	ITERATOR_OBJ     = "KITANZI"
	BIG_INTEGER_OBJ  = "NAMBA_KUBWA"
	SET_OBJ               = "SETI"
	COMPILED_REGEX_OBJ     = "RE_ILIYOTAYARISHWA"
	TUPLE_OBJ              = "JOZI"
	DATE_OBJ               = "TAREHE"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type HashKey struct {
	Type  ObjectType
	Value uint64
}

type Hashable interface {
	HashKey() HashKey
}

// Iterable interface for dicts, strings and arrays
type Iterable interface {
	Next() (Object, Object)
	Reset()
}

// WeakReference holds a reference to an object without preventing garbage collection
type WeakReference struct {
	Target Object // Target object
}

func (w *WeakReference) Type() ObjectType { return WEAK_REF_OBJ }
func (w *WeakReference) Inspect() string {
	if w.Target == nil {
		return "tupu (kumbukumbu dhaifu)"
	}
	return fmt.Sprintf("kumbukumbu dhaifu: %s", w.Target.Inspect())
}

// NewWeakReference creates a new weak reference to an object
func NewWeakReference(target Object) *WeakReference {
	return &WeakReference{Target: target}
}

// GetTarget safely retrieves the target object, returns NULL if the target has been collected
func (w *WeakReference) GetTarget() Object {
	if w.Target == nil {
		return &Null{}
	}
	return w.Target
}

// ClearTarget explicitly clears the target reference
func (w *WeakReference) ClearTarget() {
	w.Target = nil
}

func newError(format string, a ...interface{}) *Error {
	format = fmt.Sprintf("\x1b[%dm%s\x1b[0m", 31, format)
	return &Error{Message: fmt.Sprintf(format, a...)}
}
