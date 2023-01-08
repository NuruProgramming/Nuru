package parser

import (
	"github.com/AvicennaJr/Nuru/ast"
	"github.com/AvicennaJr/Nuru/token"
)

func (p *Parser) parseTime() *ast.Time {
	exp := &ast.Time{Token: p.curToken}
	if !p.expectPeek(token.DOT) {
		return nil
	}
	p.nextToken()
	exp.Method = p.parseExpression(13)
	if !p.expectPeek(token.LPAREN) {
		return nil
	}
	exp.Arguments = p.parseExpressionList(token.RPAREN)

	return exp
}
