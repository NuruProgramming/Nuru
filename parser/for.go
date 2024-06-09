package parser

import (
	"github.com/NuruProgramming/Nuru/ast"
	"github.com/NuruProgramming/Nuru/token"
)

func (p *Parser) parseForExpression() ast.Expression {
	expression := &ast.For{Token: p.curToken}
	p.nextToken()
	if !p.curTokenIs(token.IDENT) {
		return nil
	}
	if !p.peekTokenIs(token.ASSIGN) {
		return p.parseForInExpression(expression)
	}

	// In future will allow: kwa i = 0; i<10; i++ {andika(i)}
	// expression.Identifier = p.curToken.Literal
	// expression.StarterName = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
	// if expression.StarterName == nil {
	// 	return nil
	// }
	// if !p.expectPeek(token.ASSIGN) {
	// 	return nil
	// }

	// p.nextToken()

	// expression.StarterValue = p.parseExpression(LOWEST)
	// // expression.Starter = p.parseExpression(LOWEST)
	// if expression.StarterValue == nil {
	// 	return nil
	// }
	// p.nextToken()
	// for p.curTokenIs(token.SEMICOLON) {
	// 	p.nextToken()
	// }
	// expression.Condition = p.parseExpression(LOWEST)
	// if expression.Condition == nil {
	// 	return nil
	// }
	// p.nextToken()
	// for p.curTokenIs(token.SEMICOLON) {
	// 	p.nextToken()
	// }
	// expression.Closer = p.parseExpression(LOWEST)
	// if expression.Closer == nil {
	// 	return nil
	// }
	// p.nextToken()
	// for p.curTokenIs(token.SEMICOLON) {
	// 	p.nextToken()
	// }
	// if !p.curTokenIs(token.LBRACE) {
	// 	return nil
	// }
	// expression.Block = p.parseBlockStatement()
	// return expression
	return nil
}

func (p *Parser) parseForInExpression(initialExpression *ast.For) ast.Expression {
	expression := &ast.ForIn{Token: initialExpression.Token}
	if !p.curTokenIs(token.IDENT) {
		return nil
	}
	val := p.curToken.Literal
	var key string
	p.nextToken()
	if p.curTokenIs(token.COMMA) {
		p.nextToken()
		if !p.curTokenIs(token.IDENT) {
			return nil
		}
		key = val
		val = p.curToken.Literal
		p.nextToken()
	}
	expression.Key = key
	expression.Value = val
	if !p.curTokenIs(token.IN) {
		return nil
	}
	p.nextToken()
	expression.Iterable = p.parseExpression(LOWEST)
	if !p.expectPeek(token.LBRACE) {
		return nil
	}
	expression.Block = p.parseBlockStatement()
	return expression
}
