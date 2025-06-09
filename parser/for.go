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
	
	// Ensure we're at an identifier
	if !p.curTokenIs(token.IDENT) {
		p.peekError(token.IDENT)
		return nil
	}
	
	// Get the first identifier (could be key or value)
	val := p.curToken.Literal
	var key string
	
	// Move to the next token
	p.nextToken()
	
	// Check if we have a comma, which indicates a key-value pair
	if p.curTokenIs(token.COMMA) {
		p.nextToken()
		if !p.curTokenIs(token.IDENT) {
			p.peekError(token.IDENT)
			return nil
		}
		// First identifier becomes the key, second becomes the value
		key = val
		val = p.curToken.Literal
		p.nextToken()
	} else {
		// In single variable mode, we set key to "_" (unused)
		// This allows iterations like: kwa item ktk array { ... }
		key = "_"
	}
	
	// Set the key and value in the expression
	expression.Key = key
	expression.Value = val
	
	// Check for the "ktk" (IN) token
	if !p.curTokenIs(token.IN) {
		p.peekError(token.IN)
		return nil
	}
	
	// Move past the IN token
	p.nextToken()
	
	// Parse the iterable expression
	expression.Iterable = p.parseExpression(LOWEST)
	
	// Check for the opening brace
	if !p.expectPeek(token.LBRACE) {
		return nil
	}
	
	// Parse the block statement
	expression.Block = p.parseBlockStatement()
	
	return expression
}
