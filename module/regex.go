package module

import (
	"fmt"
	"regexp"

	"github.com/NuruProgramming/Nuru/object"
)

var ReFunctions = map[string]object.ModuleFunction{}

func init() {
	ReFunctions["linganisha"] = reLinganisha   // match
	ReFunctions["tafuta"] = reTafuta           // find first
	ReFunctions["tafutaZote"] = reTafutaZote   // find all
	ReFunctions["vikundi"] = reVikundi         // submatch (groups)
	ReFunctions["badilisha"] = reBadilisha     // replace all
	ReFunctions["gawa"] = reGawa               // split by pattern
}

func reError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

func reLinganisha(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return reError("linganisha inahitaji hoja 2 (pattern, neno), umeweka %d", len(args))
	}
	pat, ok := args[0].(*object.String)
	if !ok {
		return reError("linganisha: pattern lazima iwe neno")
	}
	neno, ok := args[1].(*object.String)
	if !ok {
		return reError("linganisha: neno lazima iwe neno")
	}
	re, err := regexp.Compile(pat.Value)
	if err != nil {
		return reError("linganisha: pattern si sahihi: %s", err.Error())
	}
	return &object.Boolean{Value: re.MatchString(neno.Value)}
}

func reTafuta(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return reError("tafuta inahitaji hoja 2 (pattern, neno), umeweka %d", len(args))
	}
	pat, ok := args[0].(*object.String)
	if !ok {
		return reError("tafuta: pattern lazima iwe neno")
	}
	neno, ok := args[1].(*object.String)
	if !ok {
		return reError("tafuta: neno lazima iwe neno")
	}
	re, err := regexp.Compile(pat.Value)
	if err != nil {
		return reError("tafuta: pattern si sahihi: %s", err.Error())
	}
	m := re.FindString(neno.Value)
	if m == "" {
		return &object.Null{}
	}
	return &object.String{Value: m}
}

func reTafutaZote(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return reError("tafutaZote inahitaji hoja 2 (pattern, neno), umeweka %d", len(args))
	}
	pat, ok := args[0].(*object.String)
	if !ok {
		return reError("tafutaZote: pattern lazima iwe neno")
	}
	neno, ok := args[1].(*object.String)
	if !ok {
		return reError("tafutaZote: neno lazima iwe neno")
	}
	re, err := regexp.Compile(pat.Value)
	if err != nil {
		return reError("tafutaZote: pattern si sahihi: %s", err.Error())
	}
	matches := re.FindAllString(neno.Value, -1)
	elems := make([]object.Object, len(matches))
	for i, s := range matches {
		elems[i] = &object.String{Value: s}
	}
	return &object.Array{Elements: elems}
}

func reVikundi(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return reError("vikundi inahitaji hoja 2 (pattern, neno), umeweka %d", len(args))
	}
	pat, ok := args[0].(*object.String)
	if !ok {
		return reError("vikundi: pattern lazima iwe neno")
	}
	neno, ok := args[1].(*object.String)
	if !ok {
		return reError("vikundi: neno lazima iwe neno")
	}
	re, err := regexp.Compile(pat.Value)
	if err != nil {
		return reError("vikundi: pattern si sahihi: %s", err.Error())
	}
	sub := re.FindStringSubmatch(neno.Value)
	if sub == nil {
		return &object.Null{}
	}
	elems := make([]object.Object, len(sub))
	for i, s := range sub {
		elems[i] = &object.String{Value: s}
	}
	return &object.Array{Elements: elems}
}

func reBadilisha(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 3 {
		return reError("badilisha inahitaji hoja 3 (pattern, neno, badiliko), umeweka %d", len(args))
	}
	pat, ok := args[0].(*object.String)
	if !ok {
		return reError("badilisha: pattern lazima iwe neno")
	}
	neno, ok := args[1].(*object.String)
	if !ok {
		return reError("badilisha: neno lazima iwe neno")
	}
	repl, ok := args[2].(*object.String)
	if !ok {
		return reError("badilisha: badiliko lazima liwe neno")
	}
	re, err := regexp.Compile(pat.Value)
	if err != nil {
		return reError("badilisha: pattern si sahihi: %s", err.Error())
	}
	return &object.String{Value: re.ReplaceAllString(neno.Value, repl.Value)}
}

func reGawa(args []object.Object, defs map[string]object.Object) object.Object {
	if len(args) != 2 {
		return reError("gawa inahitaji hoja 2 (pattern, neno), umeweka %d", len(args))
	}
	pat, ok := args[0].(*object.String)
	if !ok {
		return reError("gawa: pattern lazima iwe neno")
	}
	neno, ok := args[1].(*object.String)
	if !ok {
		return reError("gawa: neno lazima iwe neno")
	}
	re, err := regexp.Compile(pat.Value)
	if err != nil {
		return reError("gawa: pattern si sahihi: %s", err.Error())
	}
	parts := re.Split(neno.Value, -1)
	elems := make([]object.Object, len(parts))
	for i, s := range parts {
		elems[i] = &object.String{Value: s}
	}
	return &object.Array{Elements: elems}
}
