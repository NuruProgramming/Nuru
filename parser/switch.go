package parser

import (
	"fmt"

	"github.com/NuruProgramming/Nuru/ast"
	"github.com/NuruProgramming/Nuru/errd"
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
			synErr := &errd.MakosaSintaksia{
				Ujumbe: "Haukufunga ENDAPO (SWITCH)",
				Info:   p.curToken,
			}
			synErr.Onyesha()
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
			synErr := &errd.MakosaSintaksia{
				Ujumbe: fmt.Sprintf("Tulitegemea Kauli IKIWA (CASE) au KAWAIDA (DEFAULT) lakini tumepewa: %s", p.curToken.Type),
				Info:   p.curToken,
			}
			synErr.Onyesha()
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
		synErr := &errd.MakosaSintaksia{
			Ujumbe: fmt.Sprintf("Kauli ENDAPO (SWITCH) hua na kauli 'KAWAIDA' (DEFAULT) moja tu! Wewe umeweka %d", count),
			Info:   p.curToken,
		}
		synErr.Onyesha()
		return nil

	}
	return expression

}
