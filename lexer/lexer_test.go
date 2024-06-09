package lexer

import (
	"testing"

	"github.com/NuruProgramming/Nuru/token"
)

func TestNextToken(t *testing.T) {
	input := `
	// Testing kama lex luther iko sawa
	fanya tano = 5;
	fanya kumi = 10;

	fanya jumla = unda(x, y){
	x + y;
	};

	fanya jibu = jumla(tano, kumi);

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
	"ba ngi"
	[1, 2];
	{"mambo": "vipi"}
	. // test dot
	tumia muda
	
	badili (a) {
		ikiwa 2 {
			andika(2)
		}
		kawaida {
			andika(0)
		}
	}
	
	tupu
	
	kwa i, v ktk j`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "fanya"},
		{token.IDENT, "tano"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "fanya"},
		{token.IDENT, "kumi"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "fanya"},
		{token.IDENT, "jumla"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "unda"},
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
		{token.LET, "fanya"},
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
		{token.LBRACKET, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RBRACKET, "]"},
		{token.SEMICOLON, ";"},
		{token.LBRACE, "{"},
		{token.STRING, "mambo"},
		{token.COLON, ":"},
		{token.STRING, "vipi"},
		{token.RBRACE, "}"},
		{token.DOT, "."},
		{token.IMPORT, "tumia"},
		{token.IDENT, "muda"},
		{token.SWITCH, "badili"},
		{token.LPAREN, "("},
		{token.IDENT, "a"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.CASE, "ikiwa"},
		{token.INT, "2"},
		{token.LBRACE, "{"},
		{token.IDENT, "andika"},
		{token.LPAREN, "("},
		{token.INT, "2"},
		{token.RPAREN, ")"},
		{token.RBRACE, "}"},
		{token.DEFAULT, "kawaida"},
		{token.LBRACE, "{"},
		{token.IDENT, "andika"},
		{token.LPAREN, "("},
		{token.INT, "0"},
		{token.RPAREN, ")"},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},
		{token.NULL, "tupu"},
		{token.FOR, "kwa"},
		{token.IDENT, "i"},
		{token.COMMA, ","},
		{token.IDENT, "v"},
		{token.IN, "ktk"},
		{token.IDENT, "j"},
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
