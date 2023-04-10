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

	exp.Defaults = make(map[string]ast.Expression)

	for !p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		if p.curTokenIs(token.COMMA) {
			continue
		}
		if p.peekTokenIs(token.ASSIGN) {
			name := p.curToken.Literal
			p.nextToken()
			p.nextToken()
			val := p.parseExpression(LOWEST)
			exp.Defaults[name] = val
		} else {
			exp.Arguments = append(exp.Arguments, p.parseExpression(LOWEST))
		}
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return exp
}
