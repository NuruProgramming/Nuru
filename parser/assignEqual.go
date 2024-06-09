package parser

import (
	"fmt"

	"github.com/NuruProgramming/Nuru/ast"
)

func (p *Parser) parseAssignEqualExpression(exp ast.Expression) ast.Expression {
	switch node := exp.(type) {
	case *ast.Identifier:
		e := &ast.AssignEqual{
			Token: p.curToken,
			Left:  exp.(*ast.Identifier),
		}
		precendence := p.curPrecedence()
		p.nextToken()
		e.Value = p.parseExpression(precendence)
		return e
	case *ast.IndexExpression:
		ae := &ast.AssignmentExpression{Token: p.curToken, Left: exp}

		p.nextToken()

		ae.Value = p.parseExpression(LOWEST)

		return ae
	default:
		if node != nil {
			msg := fmt.Sprintf("Mstari %d:Tulitegemea kupata kitambulishi au array, badala yake tumepata: %s", p.curToken.Line, node.TokenLiteral())
			p.errors = append(p.errors, msg)
		} else {
			msg := fmt.Sprintf("Mstari %d: Umekosea mkuu", p.curToken.Line)
			p.errors = append(p.errors, msg)
		}
		return nil
	}
}
