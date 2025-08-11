package token

type TokenType String

type Token struct {
	Type TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EFO = "EOF"

	IDENT = "IDENT"
	INT = "INT"

	ASSIGN = "="
	PLUS = "+"

	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET = "LET"
)