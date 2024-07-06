package parser

import (
	"fmt"
	"strconv"

	"github.com/NuruProgramming/Nuru/ast"
	"github.com/NuruProgramming/Nuru/errd"
)

func (p *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		synErr := &errd.MakosaSintaksia{
			Ujumbe:   fmt.Sprintf("Mstari %d: Hatuwezi kuparse %q kama namba", p.curToken.Line, p.curToken.Literal),
			Muktadha: p.context(),
			Info:     p.curToken,
		}
		synErr.Onyesha()
		return nil
	}
	lit.Value = value

	return lit
}
