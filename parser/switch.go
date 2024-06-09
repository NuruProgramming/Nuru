package parser

import (
	"fmt"

	"github.com/NuruProgramming/Nuru/ast"
	"github.com/NuruProgramming/Nuru/token"
)

func (p *Parser) parseSwitchStatement() ast.Expression {
	expression := &ast.SwitchExpression{Token: p.curToken}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	p.nextToken()
	expression.Value = p.parseExpression(LOWEST)

	if expression.Value == nil {
		return nil
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	if !p.expectPeek(token.LBRACE) {
		return nil
	}
	p.nextToken()

	for !p.curTokenIs(token.RBRACE) {

		if p.curTokenIs(token.EOF) {
			msg := fmt.Sprintf("Mstari %d: Haukufunga ENDAPO (SWITCH)", p.curToken.Line)
			p.errors = append(p.errors, msg)
			return nil
		}
		tmp := &ast.CaseExpression{Token: p.curToken}

		if p.curTokenIs(token.DEFAULT) {

			tmp.Default = true

		} else if p.curTokenIs(token.CASE) {

			p.nextToken()

			if p.curTokenIs(token.DEFAULT) {
				tmp.Default = true
			} else {
				tmp.Expr = append(tmp.Expr, p.parseExpression(LOWEST))
				for p.peekTokenIs(token.COMMA) {
					p.nextToken()
					p.nextToken()
					tmp.Expr = append(tmp.Expr, p.parseExpression(LOWEST))
				}
			}
		} else {
			msg := fmt.Sprintf("Mstari %d: Tulitegemea Kauli IKIWA (CASE) au KAWAIDA (DEFAULT) lakini tumepewa: %s", p.curToken.Line, p.curToken.Type)
			p.errors = append(p.errors, msg)
			return nil
		}

		if !p.expectPeek(token.LBRACE) {
			return nil
		}

		tmp.Block = p.parseBlockStatement()
		p.nextToken()
		expression.Choices = append(expression.Choices, tmp)
	}

	count := 0
	for _, c := range expression.Choices {
		if c.Default {
			count++
		}
	}
	if count > 1 {
		msg := fmt.Sprintf("Kauli ENDAPO (SWITCH) hua na kauli 'KAWAIDA' (DEFAULT) moja tu! Wewe umeweka %d", count)
		p.errors = append(p.errors, msg)
		return nil

	}
	return expression

}
