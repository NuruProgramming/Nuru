package parser

import (
	"fmt"

	"github.com/NuruProgramming/Nuru/ast"
	"github.com/NuruProgramming/Nuru/errd"
)

func (p *Parser) parseAssignmentExpression(exp ast.Expression) ast.Expression {
	switch node := exp.(type) {
	case *ast.Identifier:
		e := &ast.Assign{
			Token: p.curToken,
			Name:  exp.(*ast.Identifier),
		}
		precedence := p.curPrecedence()
		p.nextToken()
		e.Value = p.parseExpression(precedence)
		return e

	case *ast.IndexExpression:
	case *ast.PropertyExpression:
		e := &ast.PropertyAssignment{
			Token: p.curToken,
			Name:  exp.(*ast.PropertyExpression),
		}
		precedence := p.curPrecedence()
		p.nextToken()
		e.Value = p.parseExpression(precedence)
		return e
	default:
		if node != nil {
			synErr := &errd.MakosaSintaksia{
				Ujumbe:   fmt.Sprintf("Tulitegemea kupata kitambulishi au array, badala yake tumepata: %s", node.TokenLiteral()),
				Muktadha: p.context(),
				Info:     p.curToken,
			}
			synErr.Onyesha()
		} else {
			synErr := &errd.MakosaSintaksia{
				Ujumbe:   "Tumeshindwa kupata kosa hapa!",
				Muktadha: p.context(),
				Info:     p.curToken,
			}
			synErr.Onyesha()
		}
		return nil
	}

	ae := &ast.AssignmentExpression{Token: p.curToken, Left: exp}

	p.nextToken()

	ae.Value = p.parseExpression(LOWEST)

	return ae
}
