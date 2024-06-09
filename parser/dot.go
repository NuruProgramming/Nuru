package parser

import (
	"github.com/NuruProgramming/Nuru/ast"
	"github.com/NuruProgramming/Nuru/token"
)

func (p *Parser) parseMethod(obj ast.Expression) ast.Expression {
	tok := p.curToken
	precedence := p.curPrecedence()
	p.nextToken()
	if p.peekTokenIs(token.LPAREN) {
		exp := &ast.MethodExpression{Token: tok, Object: obj}
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
	} else {
		exp := &ast.PropertyExpression{Token: tok, Object: obj}
		exp.Property = p.parseIdentifier()
		return exp
	}
}
