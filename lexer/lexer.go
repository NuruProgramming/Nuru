// This will convert the sequence of characters into a sequence of tokens

package lexer

import (
	"github.com/NuruProgramming/Nuru/token"
	"strings"
)

type Lexer struct {
	input        []rune
	position     int
	readPosition int
	ch           rune
	line         int
	col          int
}

var filename string

func New(file, input string) *Lexer {
	filename = file
	l := &Lexer{input: []rune(input), line: 1}
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

	if l.ch == rune('/') && l.peekChar() == rune('/') {
		l.skipSingleLineComment()
		return l.NextToken()
	}

	if l.ch == rune('/') && l.peekChar() == rune('*') {
		l.skipMultiLineComment()
		return l.NextToken()
	}

	switch l.ch {
	case rune('='):
		if l.peekChar() == rune('=') {
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.EQ, Literal: string(ch) + string(l.ch), Filename: filename, Line: dts(l.line, l.col)}
		} else {
			tok = token.Token{Type: token.ASSIGN, Literal: string(l.ch), Filename: filename, Line: sts(l.line, l.col)}
		}
	case rune(';'):
		l.col++
		tok = token.Token{Type: token.SEMICOLON, Literal: string(l.ch), Filename: filename, Line: sts(l.line, l.col)}
	case rune('('):
		l.col++
		tok = token.Token{Type: token.LPAREN, Literal: string(l.ch), Filename: filename, Line: sts(l.line, l.col)}
	case rune(')'):
		l.col++
		tok = token.Token{Type: token.RPAREN, Literal: string(l.ch), Filename: filename, Line: sts(l.line, l.col)}
	case rune('{'):
		l.col++
		tok = token.Token{Type: token.LBRACE, Literal: string(l.ch), Filename: filename, Line: sts(l.line, l.col)}
	case rune('}'):
		l.col++
		tok = token.Token{Type: token.RBRACE, Literal: string(l.ch), Filename: filename, Line: sts(l.line, l.col)}
	case rune(','):
		l.col++
		tok = token.Token{Type: token.COMMA, Literal: string(l.ch), Filename: filename, Line: sts(l.line, l.col)}
	case rune('+'):
		l.col++
		if l.peekChar() == rune('=') {
			l.col++
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.PLUS_ASSIGN, Literal: string(ch) + string(l.ch), Filename: filename, Line: dts(l.line, l.col)}
		} else if l.peekChar() == rune('+') {
			l.col++
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.PLUS_PLUS, Literal: string(ch) + string(l.ch), Filename: filename, Line: dts(l.line, l.col)}
		} else {
			tok = token.Token{Type: token.PLUS, Literal: string(l.ch), Filename: filename, Line: sts(l.line, l.col)}
		}
	case rune('-'):
		l.col++
		if l.peekChar() == rune('=') {
			l.col++
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.MINUS_ASSIGN, Literal: string(ch) + string(l.ch), Filename: filename, Line: dts(l.line, l.col)}
		} else if l.peekChar() == rune('-') {
			l.col++
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.MINUS_MINUS, Literal: string(ch) + string(l.ch), Filename: filename, Line: dts(l.line, l.col)}
		} else {
			tok = token.Token{Type: token.MINUS, Literal: string(l.ch), Filename: filename, Line: sts(l.line, l.col)}
		}
	case rune('!'):
		l.col++
		if l.peekChar() == rune('=') {
			l.col++
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.NOT_EQ, Literal: string(ch) + string(l.ch), Filename: filename, Line: dts(l.line, l.col)}
		} else {
			tok = token.Token{Type: token.BANG, Literal: string(l.ch), Filename: filename, Line: sts(l.line, l.col)}
		}
	case rune('/'):
		l.col++
		if l.peekChar() == rune('=') {
			l.col++
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.SLASH_ASSIGN, Literal: string(ch) + string(l.ch), Filename: filename, Line: dts(l.line, l.col)}
		} else {
			tok = token.Token{Type: token.SLASH, Literal: string(l.ch), Filename: filename, Line: sts(l.line, l.col)}
		}
	case rune('*'):
		l.col++
		if l.peekChar() == rune('=') {
			l.col++
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.ASTERISK_ASSIGN, Literal: string(ch) + string(l.ch), Filename: filename, Line: dts(l.line, l.col)}
		} else if l.peekChar() == rune('*') {
			l.col++
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.POW, Literal: string(ch) + string(l.ch), Filename: filename, Line: dts(l.line, l.col)}
		} else {
			tok = token.Token{Type: token.ASTERISK, Literal: string(l.ch), Filename: filename, Line: sts(l.line, l.col)}
		}
	case rune('<'):
		l.col++
		if l.peekChar() == rune('=') {
			l.col++
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.LTE, Literal: string(ch) + string(l.ch), Filename: filename, Line: dts(l.line, l.col)}
		} else {
			tok = token.Token{Type: token.LT, Literal: string(l.ch), Filename: filename, Line: sts(l.line, l.col)}
		}
	case rune('>'):
		l.col++
		if l.peekChar() == rune('=') {
			l.col++
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.GTE, Literal: string(ch) + string(l.ch), Filename: filename, Line: dts(l.line, l.col)}
		} else {
			tok = token.Token{Type: token.GT, Literal: string(l.ch), Filename: filename, Line: sts(l.line, l.col)}
		}
	case rune('"'), rune('\''):
		l.col++
		tok.Type = token.STRING
		tok.Line.Start.Line = l.line
		tok.Line.Start.Column = l.col
		tok.Literal = l.readString(l.ch)
		tok.Filename = filename
		tok.Line.End.Line = l.line
		tok.Line.End.Column = l.col
	case rune('['):
		tok = token.Token{Type: token.LBRACKET, Literal: string(l.ch), Filename: filename, Line: sts(l.line, l.col)}
	case rune(']'):
		tok = token.Token{Type: token.RBRACKET, Literal: string(l.ch), Filename: filename, Line: sts(l.line, l.col)}
	case rune(':'):
		tok = token.Token{Type: token.COLON, Literal: string(l.ch), Filename: filename, Line: sts(l.line, l.col)}
	case rune('@'):
		tok = token.Token{Type: token.AT, Literal: string(l.ch), Filename: filename, Line: sts(l.line, l.col)}
	case rune('.'):
		tok = token.Token{Type: token.DOT, Literal: string(l.ch), Filename: filename, Line: sts(l.line, l.col)}
	case rune('&'):
		l.col++
		if l.peekChar() == rune('&') {
			l.col++
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.AND, Literal: string(ch) + string(l.ch), Filename: filename, Line: dts(l.line, l.col)}
		}
	case rune('|'):
		l.col++
		if l.peekChar() == rune('|') {
			l.col++
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.OR, Literal: string(ch) + string(l.ch), Filename: filename, Line: dts(l.line, l.col)}
		}
	case rune('%'):
		l.col++
		if l.peekChar() == rune('=') {
			l.col++
			ch := l.ch
			l.readChar()
			tok = token.Token{Type: token.MODULUS_ASSIGN, Literal: string(ch) + string(l.ch), Filename: filename, Line: dts(l.line, l.col)}
		} else {
			tok = token.Token{Type: token.MODULUS, Literal: string(l.ch), Filename: filename, Line: sts(l.line, l.col)}
		}
	case 0:
		l.col++
		tok.Literal = ""
		tok.Type = token.EOF
		tok.Filename = filename
		tok.Line = sts(l.line, l.col)
	default:
		if isLetter(l.ch) {
			tok.Line.Start.Line = l.line
			tok.Line.Start.Column = l.col
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			tok.Filename = filename
			tok.Line.End.Line = l.line
			tok.Line.End.Column = l.col
			return tok
		} else if isDigit(l.ch) && isLetter(l.peekChar()) {
			tok.Line.Start.Line = l.line
			tok.Line.Start.Column = l.col
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			tok.Filename = filename
			tok.Line.End.Line = l.line
			tok.Line.End.Column = l.col
			return tok
		} else if isDigit(l.ch) {
			tok.Line.Start.Line = l.line
			tok.Line.Start.Column = l.col
			tok = l.readDecimal()
			tok.Filename = filename
			tok.Line.End.Line = l.line
			tok.Line.End.Column = l.col
			return tok
		} else {
			l.col++
			tok = token.Token{Type: token.ILLEGAL, Literal: string(l.ch), Filename: filename, Line: sts(l.line, l.col)}
		}
	}

	l.readChar()
	return tok
}

func xts(line, col, x, y int) token.Span {
	var span token.Span

	span.Start.Line = line
	span.Start.Column = col
	span.End.Line = line + x
	span.End.Column = col + y

	return span

}

func sts(line, col int) token.Span { return xts(line, col, 0, 1) }
func dts(line, col int) token.Span { return xts(line, col, 0, 2) }

func newToken(tokenInfo token.Token) token.Token {
	tokenInfo.Filename = filename
	return tokenInfo
}

func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) || isDigit(l.ch) {
		l.col++
		if l.ch == '\n' {
			l.line++
			l.col = 1
		}
		l.readChar()
	}
	return string(l.input[position:l.position])
}

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '@'
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		if l.ch == '\n' {
			l.line++
			l.col = 0
		}
		l.readChar()
	}
}

func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.col++
		l.readChar()
	}
	return string(l.input[position:l.position])
}

func (l *Lexer) readDecimal() token.Token {
	integer := l.readNumber()
	if l.ch == '.' && isDigit(l.peekChar()) {
		l.readChar()
		l.col++
		fraction := l.readNumber()
		fl := len(fraction)
		return token.Token{Type: token.FLOAT, Literal: integer + "." + fraction, Line: xts(l.line, l.col-fl, 0, fl)}
	}
	il := len(integer)
	return token.Token{Type: token.INT, Literal: integer, Line: xts(l.line, l.col-il, 0, il)}
}

func (l *Lexer) peekChar() rune {
	if l.readPosition >= len(l.input) {
		return rune(0)
	} else {
		return l.input[l.readPosition]
	}
}

// func (l *Lexer) peekTwoChar() rune {
// 	if l.readPosition+1 >= len(l.input) {
// 		return rune(0)
// 	} else {
// 		return l.input[l.readPosition+1]
// 	}
// }

func (l *Lexer) skipSingleLineComment() {
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
	l.skipWhitespace()
}

func (l *Lexer) skipMultiLineComment() {
	endFound := false

	for !endFound {
		l.col++
		if l.ch == 0 {
			endFound = true
		}

		if l.ch == '*' && l.peekChar() == '/' {
			endFound = true
			l.col++
			l.readChar()
		}

		l.readChar()
		l.skipWhitespace()
	}

}

func (l *Lexer) readString(delim rune) string {
	var str strings.Builder
	for {
		l.col++
		l.readChar()
		if l.ch == delim || l.ch == 0 {
			break
		} else if l.ch == '\n' {
			l.line++
			l.col = 0
		} else if l.ch == '\\' {
			l.col++
			switch l.peekChar() {
			case 'n':
				l.readChar()
				l.ch = '\n'
			case 'r':
				l.readChar()
				l.ch = '\r'
			case 't':
				l.readChar()
				l.ch = '\t'
			case '"':
				l.readChar()
				l.ch = '"'
			case '\\':
				l.readChar()
				l.ch = '\\'
			case '\'':
				l.readChar()
				l.ch = '\''
			}
		}
		str.WriteRune(l.ch)
	}
	return str.String()
}
