package object

import (
	"fmt"

	"github.com/NuruProgramming/Nuru/ast"
)

type Package struct {
	Name  *ast.Identifier
	Env   *Environment
	Scope *Environment
}

func (p *Package) Type() ObjectType { return PACKAGE_OBJ }
func (p *Package) Inspect() string {
	return fmt.Sprintf("Pakeji: %s", p.Name.Value)
}
