package lexer

import (
	"github.com/AvicennaJr/Nuru/token"
)

type Lexer struct {
	input        string // string or runes. Runes should be better.
	position     int
	readPosition int
	ch           byte // make this a rune too
	line         int
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token
	l.skipWhitespace()
	if l.ch == '/' && l.peekChar() == '/' {
		l.skipSingleLineComment()
		return l.NextToken()
	}
	if l.ch == '/' && l.peekChar() == '*' {
		l.skipMultiLineComment()
		return l.NextToken()
	}

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch), Line: l.line}
		} else {
			tok = newToken(token.ASSIGN, l.line, l.ch)
		}
	case ';':
		tok = newToken(token.SEMICOLON, l.line, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.line, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.line, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.line, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.line, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.line, l.ch)
	case '+':
		tok = newToken(token.PLUS, l.line, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.line, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch), Line: l.line}
		} else {
			tok = newToken(token.BANG, l.line, l.ch)
		}
	case '/':
		tok = newToken(token.SLASH, l.line, l.ch)
	case '*':
		if l.peekChar() == '*' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.POW, Literal: string(ch) + string(l.ch), Line: l.line}
		} else {
			tok = newToken(token.ASTERISK, l.line, l.ch)
		}
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.LTE, Literal: string(ch) + string(l.ch), Line: l.line}
		} else {
			tok = newToken(token.LT, l.line, l.ch)
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.GTE, Literal: string(ch) + string(l.ch), Line: l.line}
		} else {
			tok = newToken(token.GT, l.line, l.ch)
		}
	case '"':
		tok.Type = token.STRING
		tok.Literal = l.readString()
		tok.Line = l.line
	case '\'':
		tok = token.Token{Type: token.STRING, Literal: l.readSingleQuoteString(), Line: l.line}
	case '[':
		tok = newToken(token.LBRACKET, l.line, l.ch)
	case ']':
		tok = newToken(token.RBRACKET, l.line, l.ch)
	case ':':
		tok = newToken(token.COLON, l.line, l.ch)
	case '&':
		if l.peekChar() == '&' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.AND, Literal: string(ch) + string(l.ch), Line: l.line}
		}
	case '|':
		if l.peekChar() == '|' {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.OR, Literal: string(ch) + string(l.ch), Line: l.line}
		}
	case '%':
		tok = newToken(token.MODULUS, l.line, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
		tok.Line = l.line
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			tok.Line = l.line
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			tok.Line = l.line
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.line, l.ch)
		}
	}

	l.readChar()
	return tok
}

func newToken(tokenType token.TokenType, line int, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch), Line: line}
}

func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		if l.ch == '\n' {
			l.line++
		}
		l.readChar()
	}
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) skipSingleLineComment() {
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
	l.skipWhitespace()
}

func (l *Lexer) skipMultiLineComment() {
	endFound := false

	for !endFound {
		if l.ch == 0 {
			endFound = true
		}

		if l.ch == '*' && l.peekChar() == '/' {
			endFound = true
			l.readChar()
		}

		l.readChar()
		l.skipWhitespace()
	}

}

func (l *Lexer) readString() string {
	position := l.position + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}

	return l.input[position:l.position]
}

func (l *Lexer) readSingleQuoteString() string {
	var str string
	for {
		l.readChar()
		if l.ch == '\'' || l.ch == 0 {
			break
		} else if l.ch == '\\' && l.peekChar() == '\'' {
			l.readChar()
		}
		str += string(l.ch)
	}
	return str
}
