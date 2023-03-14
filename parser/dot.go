package parser

import (
	"github.com/AvicennaJr/Nuru/ast"
	"github.com/AvicennaJr/Nuru/token"
)

func (p *Parser) parseMethod(obj ast.Expression) ast.Expression {
	precedence := p.curPrecedence()
	exp := &ast.MethodExpression{Token: p.curToken, Object: obj}
	p.nextToken()
	exp.Method = p.parseExpression(precedence)
	if !p.expectPeek(token.LPAREN) {
		return nil
	}
	exp.Arguments = p.parseExpressionList(token.RPAREN)
	return exp
}
