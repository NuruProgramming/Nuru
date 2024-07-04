package parser

import (
	"fmt"

	"github.com/NuruProgramming/Nuru/ast"
	"github.com/NuruProgramming/Nuru/errd"
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

			synErr := &errd.MakosaSintaksia{
				Ujumbe: fmt.Sprintf("Tulitegemea kupata kitambulishi au array, badala yake tumepata: %s", node.TokenLiteral()),
				Info:   p.curToken,
			}
			synErr.Onyesha()
		} else {
			synErr := &errd.MakosaSintaksia{
				Ujumbe: "Tumejaribu nakini hatujaweza kupata shida",
				Info:   p.curToken,
			}
			synErr.Onyesha()
		}
		return nil
	}
}
