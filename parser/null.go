package parser

import (
	"github.com/NuruProgramming/Nuru/ast"
)

func (p *Parser) parseNull() ast.Expression {
	return &ast.Null{Token: p.curToken}
}
