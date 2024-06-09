package parser

import (
	"github.com/NuruProgramming/Nuru/ast"
	"github.com/NuruProgramming/Nuru/token"
)

func (p *Parser) parseImport() ast.Expression {
	exp := &ast.Import{Token: p.curToken}
	exp.Identifiers = make(map[string]*ast.Identifier)
	for p.curToken.Line == p.peekToken.Line {
		p.nextToken()
		identifier := &ast.Identifier{Value: p.curToken.Literal}
		exp.Identifiers[p.curToken.Literal] = identifier
		if p.peekTokenIs(token.COMMA) {
			p.nextToken()
		}
		if p.peekTokenIs(token.EOF) {
			break
		}
	}

	return exp
}
