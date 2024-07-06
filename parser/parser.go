package parser

import (
	"fmt"
	"os"
	"strings"

	"github.com/NuruProgramming/Nuru/ast"
	"github.com/NuruProgramming/Nuru/errd"
	"github.com/NuruProgramming/Nuru/lexer"
	"github.com/NuruProgramming/Nuru/token"
)

const (
	_ int = iota
	LOWEST
	ASSIGN      // =
	COND        // OR or AND
	EQUALS      // ==
	LESSGREATER // > OR <
	SUM         // +
	PRODUCT     // *
	POWER       // ** we got the power XD
	MODULUS     // %
	PREFIX      //  -X OR !X
	CALL        // myFunction(X)
	INDEX       // Arrays
	DOT         // For methods
)

var precedences = map[token.TokenType]int{
	token.AND:             COND,
	token.OR:              COND,
	token.IN:              COND,
	token.ASSIGN:          ASSIGN,
	token.EQ:              EQUALS,
	token.NOT_EQ:          EQUALS,
	token.LT:              LESSGREATER,
	token.LTE:             LESSGREATER,
	token.GT:              LESSGREATER,
	token.GTE:             LESSGREATER,
	token.PLUS:            SUM,
	token.PLUS_ASSIGN:     SUM,
	token.MINUS:           SUM,
	token.MINUS_ASSIGN:    SUM,
	token.SLASH:           PRODUCT,
	token.SLASH_ASSIGN:    PRODUCT,
	token.ASTERISK:        PRODUCT,
	token.ASTERISK_ASSIGN: PRODUCT,
	token.POW:             POWER,
	token.MODULUS:         MODULUS,
	token.MODULUS_ASSIGN:  MODULUS,
	// token.BANG:     PREFIX,
	token.LPAREN:   CALL,
	token.LBRACKET: INDEX,
	token.DOT:      DOT, // Highest priority
}

type (
	prefixParseFn  func() ast.Expression
	infixParseFn   func(ast.Expression) ast.Expression
	postfixParseFn func() ast.Expression
)

type Parser struct {
	content []string
	l       *lexer.Lexer

	curToken  token.Token
	peekToken token.Token
	prevToken token.Token

	prefixParseFns  map[token.TokenType]prefixParseFn
	infixParseFns   map[token.TokenType]infixParseFn
	postfixParseFns map[token.TokenType]postfixParseFn
}

func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func (p *Parser) registerPostfix(tokenType token.TokenType, fn postfixParseFn) {
	p.postfixParseFns[tokenType] = fn
}

func New(context string, l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.content = strings.Split(context, "\n")

	p.nextToken()
	p.nextToken()

	p.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	p.registerPrefix(token.STRING, p.parseStringLiteral)
	p.registerPrefix(token.IDENT, p.parseIdentifier)
	p.registerPrefix(token.INT, p.parseIntegerLiteral)
	p.registerPrefix(token.FLOAT, p.parseFloatLiteral)
	p.registerPrefix(token.BANG, p.parsePrefixExpression)
	p.registerPrefix(token.MINUS, p.parsePrefixExpression)
	p.registerPrefix(token.PLUS, p.parsePrefixExpression)
	p.registerPrefix(token.TRUE, p.parseBoolean)
	p.registerPrefix(token.FALSE, p.parseBoolean)
	p.registerPrefix(token.LPAREN, p.parseGroupedExpression)
	p.registerPrefix(token.IF, p.parseIfExpression)
	p.registerPrefix(token.FUNCTION, p.parseFunctionLiteral)
	p.registerPrefix(token.LBRACKET, p.parseArrayLiteral)
	p.registerPrefix(token.LBRACE, p.parseDictLiteral)
	p.registerPrefix(token.WHILE, p.parseWhileExpression)
	p.registerPrefix(token.NULL, p.parseNull)
	p.registerPrefix(token.FOR, p.parseForExpression)
	p.registerPrefix(token.SWITCH, p.parseSwitchStatement)
	p.registerPrefix(token.IMPORT, p.parseImport)
	p.registerPrefix(token.PACKAGE, p.parsePackage)
	p.registerPrefix(token.AT, p.parseAt)

	p.infixParseFns = make(map[token.TokenType]infixParseFn)
	p.registerInfix(token.AND, p.parseInfixExpression)
	p.registerInfix(token.OR, p.parseInfixExpression)
	p.registerInfix(token.PLUS, p.parseInfixExpression)
	p.registerInfix(token.PLUS_ASSIGN, p.parseAssignEqualExpression)
	p.registerInfix(token.MINUS, p.parseInfixExpression)
	p.registerInfix(token.MINUS_ASSIGN, p.parseAssignEqualExpression)
	p.registerInfix(token.SLASH, p.parseInfixExpression)
	p.registerInfix(token.SLASH_ASSIGN, p.parseAssignEqualExpression)
	p.registerInfix(token.ASTERISK, p.parseInfixExpression)
	p.registerInfix(token.ASTERISK_ASSIGN, p.parseAssignEqualExpression)
	p.registerInfix(token.POW, p.parseInfixExpression)
	p.registerInfix(token.MODULUS, p.parseInfixExpression)
	p.registerInfix(token.MODULUS_ASSIGN, p.parseAssignmentExpression)
	p.registerInfix(token.EQ, p.parseInfixExpression)
	p.registerInfix(token.NOT_EQ, p.parseInfixExpression)
	p.registerInfix(token.LT, p.parseInfixExpression)
	p.registerInfix(token.LTE, p.parseInfixExpression)
	p.registerInfix(token.GT, p.parseInfixExpression)
	p.registerInfix(token.GTE, p.parseInfixExpression)
	p.registerInfix(token.LPAREN, p.parseCallExpression)
	p.registerInfix(token.LBRACKET, p.parseIndexExpression)
	p.registerInfix(token.ASSIGN, p.parseAssignmentExpression)
	p.registerInfix(token.IN, p.parseInfixExpression)
	p.registerInfix(token.DOT, p.parseMethod)

	p.postfixParseFns = make(map[token.TokenType]postfixParseFn)
	p.registerPostfix(token.PLUS_PLUS, p.parsePostfixExpression)
	p.registerPostfix(token.MINUS_MINUS, p.parsePostfixExpression)

	return p
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		program.Statements = append(program.Statements, stmt)

		p.nextToken()
	}
	return program
}

// manage token literals:

func (p *Parser) nextToken() {
	p.prevToken = p.curToken
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(p.curToken)
		return false
	}
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}
	return LOWEST
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}

	return LOWEST
}

func (p *Parser) peekError(t token.Token) {
	synErr := &errd.MakosaSintaksia{
		Ujumbe:   fmt.Sprintf("Tulitegemea kupata '%s', badala yake tumepata '%s'", t.Type, p.peekToken.Type),
		Muktadha: p.context(),
		Info:     t,
	}
	synErr.Onyesha()
}

func (p *Parser) context() string {
	var muk string
	var caret string
	lineContent := p.content[p.curToken.Line.Start.Line-1]
	if os.Getenv("RANGI") == "0" {
		caret = "^"
	} else {
		caret = fmt.Sprintf("\x1b[%dm%s\x1b[0m", 31, "^")
	}

	muk = fmt.Sprintf("\t%s\n\t%s%s", lineContent, strings.Repeat(" ", p.curToken.Line.Start.Column-1), caret)

	return muk
}

// parse expressions

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.curToken}

	stmt.Expression = p.parseExpression(LOWEST)

	if p.peekTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	postfix := p.postfixParseFns[p.curToken.Type]
	if postfix != nil {
		return (postfix())
	}
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken)
		return nil
	}
	leftExp := prefix()

	for !p.peekTokenIs(token.SEMICOLON) && precedence < p.peekPrecedence() {
		infix := p.infixParseFns[p.peekToken.Type]
		if infix == nil {
			p.noInfixParseFnError(p.peekToken)
			return nil
		}

		p.nextToken()
		leftExp = infix(leftExp)
	}
	return leftExp

}

// prefix expressions

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.nextToken()

	expression.Right = p.parseExpression(PREFIX)

	return expression
}

func (p *Parser) noPrefixParseFnError(t token.Token) {
	synErr := &errd.MakosaSintaksia{
		Ujumbe: fmt.Sprintf("Tumeshindwa kuparse '%s'", t.Literal),
		Muktadha: p.context(),
		Info:   t,
	}
	synErr.Onyesha()
}

// infix expressions

func (p *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
		Left:     left,
	}

	precedence := p.curPrecedence()
	p.nextToken()
	expression.Right = p.parseExpression(precedence)
	return expression
}

func (p *Parser) noInfixParseFnError(t token.Token) {
	synErr := &errd.MakosaSintaksia{
		Ujumbe: fmt.Sprintf("Tumeshindwa kuparse '%s'", t.Literal),
		Muktadha: p.context(),
		Info:   t,
	}
	synErr.Onyesha()
}

func (p *Parser) parseGroupedExpression() ast.Expression {
	p.nextToken()

	exp := p.parseExpression(LOWEST)

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return exp
}

// postfix expressions

func (p *Parser) parsePostfixExpression() ast.Expression {
	expression := &ast.PostfixExpression{
		Token:    p.prevToken,
		Operator: p.curToken.Literal,
	}
	return expression
}
