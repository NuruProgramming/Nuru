package lexer

import (
	"testing"

	"github.com/AvicennaJr/Nuru/token"
)

func TestNextToken(t *testing.T) {
	input := `
	// Testing kama lex luther iko sawa
	acha tano = 5;
	acha kumi = 10;

	acha jumla = fn(x, y){
	x + y;
	};

	acha jibu = jumla(tano, kumi);

	!-/5;
	5 < 10 > 5;

	kama (5 < 10) {
		rudisha kweli;
	} sivyo {
		rudisha sikweli;
	}

	10 == 10;
	10 != 9; // Hii ni comment
	// Comment nyingine

	/*
	multiline comment
	*/

	/* multiline comment number twooooooooooo */
	5
	"bangi"
	"ba ngi"`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "acha"},
		{token.IDENT, "tano"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "acha"},
		{token.IDENT, "kumi"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "acha"},
		{token.IDENT, "jumla"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "acha"},
		{token.IDENT, "jibu"},
		{token.ASSIGN, "="},
		{token.IDENT, "jumla"},
		{token.LPAREN, "("},
		{token.IDENT, "tano"},
		{token.COMMA, ","},
		{token.IDENT, "kumi"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "kama"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "rudisha"},
		{token.TRUE, "kweli"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "sivyo"},
		{token.LBRACE, "{"},
		{token.RETURN, "rudisha"},
		{token.FALSE, "sikweli"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOT_EQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.STRING, "bangi"},
		{token.STRING, "ba ngi"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
