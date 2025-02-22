package parser

import (
	"log"
	"strings"

	"github.com/NuruProgramming/Nuru/ast"
	"github.com/NuruProgramming/Nuru/token"
)

func (p *Parser) parseImport() ast.Expression {
	exp := &ast.Import{Token: p.curToken}
	exp.Identifiers = make(map[string]*ast.Identifier)
	// Eka mahali hapa kwa mfano kawaida::hisabati
	var mahali strings.Builder
	// jina la maktaba hili kwa mafano hisabati
	var jina strings.Builder

	for p.curToken.Line == p.peekToken.Line {
		if p.curTokenIs(token.IMPORT) {
			p.nextToken()
			continue
		}

		// kwa kuruka ::
		if p.curTokenIs(token.COLON) && p.peekTokenIs(token.COLON) {
			p.nextToken()
			p.nextToken()
			mahali.WriteRune('/')
		}

		if p.curTokenIs(token.IDENT) {
			mahali.WriteString(p.curToken.Literal)
			p.nextToken()
			continue
		}

		if p.curTokenIs(token.ASSIGN) {
			p.nextToken()

			// Check the value
			if p.curTokenIs(token.IDENT) {
				jina.WriteString(p.curToken.Literal)
				p.nextToken()
			}

			continue
		}

		if p.curTokenIs(token.COMMA) {
			mhs := mahali.String()
			if jina.Len() <= 0 {
				mhss := strings.Split(mhs, "/")
				jina.WriteString(mhss[len(mhss)-1])
			}
			exp.Identifiers[jina.String()] = &ast.Identifier{Value: mhs}
			exp.Filename = p.filename

			mahali.Reset()
			p.nextToken()
			jina.Reset()
			print(p.curToken.Literal + "\n")
			continue
		}

		if p.peekTokenIs(token.EOF) || p.peekTokenIs(token.NEWLINE) {
			break
		}

		// Just log it and continue, should have a better way to indicate errors
		log.Printf("%s haijulikani", p.curToken.Literal)
		p.nextToken()
	}

	mhs := mahali.String()
	if jina.Len() == 0 {
		mhss := strings.Split(mhs, "/")
		jina.WriteString(mhss[len(mhss)-1])
	}
	exp.Identifiers[jina.String()] = &ast.Identifier{Value: mhs}
	exp.Filename = p.filename

	return exp
}
