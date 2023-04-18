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

func newError(format string, a ...interface{}) *Error {
	format = fmt.Sprintf("\x1b[%dm%s\x1b[0m", 31, format)
	return &Error{Message: fmt.Sprintf(format, a...)}
}
