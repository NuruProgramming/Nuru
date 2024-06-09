package parser

import (
	"github.com/NuruProgramming/Nuru/ast"
	"github.com/NuruProgramming/Nuru/token"
)

func (p *Parser) parseFunctionLiteral() ast.Expression {
	lit := &ast.FunctionLiteral{Token: p.curToken}

	if p.peekTokenIs(token.IDENT) {
		p.nextToken()
		lit.Name = p.curToken.Literal
	}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	if !p.parseFunctionParameters(lit) {
		return nil
	}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	lit.Body = p.parseBlockStatement()

	return lit
}

func (p *Parser) parseFunctionParameters(lit *ast.FunctionLiteral) bool {
	lit.Defaults = make(map[string]ast.Expression)
	for !p.peekTokenIs(token.RPAREN) {
		p.nextToken()

		if p.curTokenIs(token.COMMA) {
			continue
		}

		ident := &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
		lit.Parameters = append(lit.Parameters, ident)

		if p.peekTokenIs(token.ASSIGN) {
			p.nextToken()
			p.nextToken()
			lit.Defaults[ident.Value] = p.parseExpression(LOWEST)
		} else {
			if len(lit.Defaults) > 0 {
				return false
			}
		}

		if !(p.peekTokenIs(token.COMMA) || p.peekTokenIs(token.RPAREN)) {
			return false
		}
	}

	return p.expectPeek(token.RPAREN)
}

func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	exp := &ast.CallExpression{Token: p.curToken, Function: function}
	exp.Arguments = p.parseExpressionList(token.RPAREN)
	return exp
}
