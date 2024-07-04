package parser

import (
	"fmt"
	"strconv"

	"github.com/NuruProgramming/Nuru/ast"
	"github.com/NuruProgramming/Nuru/errd"
)

func (p *Parser) parseFloatLiteral() ast.Expression {
	fl := &ast.FloatLiteral{Token: p.curToken}
	value, err := strconv.ParseFloat(p.curToken.Literal, 64)
	if err != nil {
		synErr := &errd.MakosaSintaksia{
			Ujumbe: fmt.Sprintf("Hatuwezi kuparse %q kama desimali", p.curToken.Line, p.curToken.Literal),
			Info:   p.curToken,
		}
		synErr.Onyesha()

		return nil
	}
	fl.Value = value
	return fl
}
