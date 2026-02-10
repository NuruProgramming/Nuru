package object

import (
	"fmt"
	"regexp"
)

// CompiledRegex is a compiled regular expression with methods that take the string only.
type CompiledRegex struct {
	Re *regexp.Regexp
}

func (c *CompiledRegex) Type() ObjectType { return COMPILED_REGEX_OBJ }
func (c *CompiledRegex) Inspect() string {
	if c.Re == nil {
		return "re_iliyotayarishwa(tupu)"
	}
	return fmt.Sprintf("re_iliyotayarishwa(%q)", c.Re.String())
}

func (c *CompiledRegex) Method(method string, args []Object) Object {
	if c.Re == nil {
		return newError("re iliyotayarishwa: regex si sahihi")
	}
	switch method {
	case "linganisha":
		return c.linganisha(args)
	case "tafuta":
		return c.tafuta(args)
	case "tafutaZote":
		return c.tafutaZote(args)
	case "vikundi":
		return c.vikundi(args)
	case "badilisha":
		return c.badilisha(args)
	case "gawa":
		return c.gawa(args)
	default:
		return newError("Samahani, re_iliyotayarishwa haina method '%s'", method)
	}
}

func (c *CompiledRegex) linganisha(args []Object) Object {
	if len(args) != 1 {
		return newError("linganisha inahitaji hoja 1 (neno), umeweka %d", len(args))
	}
	s, ok := args[0].(*String)
	if !ok {
		return newError("linganisha: neno lazima iwe neno")
	}
	return &Boolean{Value: c.Re.MatchString(s.Value)}
}

func (c *CompiledRegex) tafuta(args []Object) Object {
	if len(args) != 1 {
		return newError("tafuta inahitaji hoja 1 (neno), umeweka %d", len(args))
	}
	s, ok := args[0].(*String)
	if !ok {
		return newError("tafuta: neno lazima iwe neno")
	}
	m := c.Re.FindString(s.Value)
	if m == "" {
		return &Null{}
	}
	return &String{Value: m}
}

func (c *CompiledRegex) tafutaZote(args []Object) Object {
	if len(args) != 1 {
		return newError("tafutaZote inahitaji hoja 1 (neno), umeweka %d", len(args))
	}
	s, ok := args[0].(*String)
	if !ok {
		return newError("tafutaZote: neno lazima iwe neno")
	}
	matches := c.Re.FindAllString(s.Value, -1)
	elems := make([]Object, len(matches))
	for i, m := range matches {
		elems[i] = &String{Value: m}
	}
	return &Array{Elements: elems}
}

func (c *CompiledRegex) vikundi(args []Object) Object {
	if len(args) != 1 {
		return newError("vikundi inahitaji hoja 1 (neno), umeweka %d", len(args))
	}
	s, ok := args[0].(*String)
	if !ok {
		return newError("vikundi: neno lazima iwe neno")
	}
	sub := c.Re.FindStringSubmatch(s.Value)
	if sub == nil {
		return &Null{}
	}
	elems := make([]Object, len(sub))
	for i, m := range sub {
		elems[i] = &String{Value: m}
	}
	return &Array{Elements: elems}
}

func (c *CompiledRegex) badilisha(args []Object) Object {
	if len(args) != 2 {
		return newError("badilisha inahitaji hoja 2 (neno, badiliko), umeweka %d", len(args))
	}
	s, ok := args[0].(*String)
	if !ok {
		return newError("badilisha: neno lazima iwe neno")
	}
	repl, ok := args[1].(*String)
	if !ok {
		return newError("badilisha: badiliko lazima liwe neno")
	}
	return &String{Value: c.Re.ReplaceAllString(s.Value, repl.Value)}
}

func (c *CompiledRegex) gawa(args []Object) Object {
	if len(args) != 1 {
		return newError("gawa inahitaji hoja 1 (neno), umeweka %d", len(args))
	}
	s, ok := args[0].(*String)
	if !ok {
		return newError("gawa: neno lazima iwe neno")
	}
	parts := c.Re.Split(s.Value, -1)
	elems := make([]Object, len(parts))
	for i, p := range parts {
		elems[i] = &String{Value: p}
	}
	return &Array{Elements: elems}
}
