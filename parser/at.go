package parser

import "github.com/AvicennaJr/Nuru/ast"

func (p *Parser) parseAt() ast.Expression {
	return &ast.At{Token: p.curToken}
}
