package object

import (
	"bytes"
	"strings"

	"github.com/NuruProgramming/Nuru/ast"
)

type Function struct {
	Name       string
	Parameters []*ast.Identifier
	Defaults   map[string]ast.Expression
	Body       *ast.BlockStatement
	Env        *Environment
	BaseRefCountable
}

// GetOutgoingReferences returns all references held by this function
func (f *Function) GetOutgoingReferences() []Object {
	var refs []Object

	// Add environment to references
	if f.Env != nil {
		refs = append(refs, f.Env.GetAllValues()...)
	}

	return refs
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("unda")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

// NewFunction creates a new reference-counted function
func NewFunction(name string, params []*ast.Identifier, defaults map[string]ast.Expression, body *ast.BlockStatement, env *Environment) *Function {
	fn := &Function{
		Name:       name,
		Parameters: params,
		Defaults:   defaults,
		Body:       body,
		Env:        env,
	}

	// Track the new object in the reference counter
	GlobalRefCounter.TrackObject(fn)

	// Keep reference to environment
	if env != nil {
		Retain(env)
	}

	return fn
}
