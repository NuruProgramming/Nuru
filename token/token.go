package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "HARAMU"
	EOF     = "MWISHO"

	// Identifiers + literals
	IDENT  = "KITAMBULISHI"
	INT    = "NAMBA"
	STRING = "NENO"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	LBRACKET  = "["
	RBRACKET  = "]"
	COLON     = ":"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "ACHA"
	TRUE     = "KWELI"
	FALSE    = "SIKWELI"
	IF       = "KAMA"
	ELSE     = "SIVYO"
	RETURN   = "RUDISHA"
	WHILE    = "WAKATI"
)

var keywords = map[string]TokenType{
	"fn":      FUNCTION,
	"acha":    LET,
	"kweli":   TRUE,
	"sikweli": FALSE,
	"kama":    IF,
	"au":      ELSE,
	"sivyo":   ELSE,
	"wakati":  WHILE,
	"rudisha": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
