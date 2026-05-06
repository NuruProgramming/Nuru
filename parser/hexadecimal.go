package parser

import (
	"fmt"
	"math/big"

	"github.com/NuruProgramming/Nuru/ast"
)

func (p *Parser) parseHexadecimalLiteral() ast.Expression {
	lit := &ast.HexadecimalLiteral{Token: p.curToken}
	tlit := p.curToken.Literal

	for _, x := range tlit {
		if !(isDigit(x) || isLetter(x)) {
			msg := fmt.Sprintf("Mstari %d: Hatuwezi kuparse %q kama hexa desimali", p.curToken.Line, tlit)
			p.errors = append(p.errors, msg)
			return nil

		}
	}

	value := new(big.Int)
	value.SetString(tlit, 16)

	lit.Value = value.Int64()

	return lit
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'f' || 'A' <= ch && ch <= 'F'
}
