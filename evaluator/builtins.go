package evaluator

import (
	infunc "github.com/AvicennaJr/Nuru/builtins"
	"github.com/AvicennaJr/Nuru/object"
)

// An array of the buitin functions
// There is definitely a better way but
// not for now
var builtins = map[string]*object.Builtin{
	"andika": {Fn: infunc.Andika},
	"aina":   {Fn: infunc.Aina},
	"jaza":   {Fn: infunc.Jaza},
	"fungua": {Fn: infunc.Fungua},
}
