package object

import (
	"fmt"
	"strings"

	"github.com/NuruProgramming/Nuru/ast"
)

type Function struct {
	Name       string
	Parameters []*ast.Identifier
	Defaults   map[string]ast.Expression
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType { return FUNCTION_OBJ }
func (f *Function) Inspect() string {
	var out strings.Builder

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(fmt.Sprintf("[undo: %s", f.Name))

	if len(params) > 0 {
		out.WriteString(fmt.Sprintf("; inahitaji: %s", strings.Join(params, ", ")))
	}

	out.WriteRune(']')

	return out.String()
}
