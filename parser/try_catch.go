package parser

import (
	"github.com/NuruProgramming/Nuru/ast"
	"github.com/NuruProgramming/Nuru/token"
)

func (p *Parser) parseTryCatchExpression() ast.Expression {
	expression := &ast.TryCatchExpression{Token: p.curToken}

	// Parse the try block
	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	expression.TryBlock = p.parseBlockStatement()

	// Make sure we have a catch/bila block
	if !p.expectPeek(token.CATCH) {
		return nil
	}

	// Check if there's an error identifier
	if p.peekTokenIs(token.IDENT) {
		p.nextToken()
		expression.Identifier = p.curToken.Literal
	}

	// Parse the catch block
	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	expression.CatchBlock = p.parseBlockStatement()

	return expression
}
