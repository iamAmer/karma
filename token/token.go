package token

// TokenType is the category of a token (identifier, keyword, operator, ..)
type TokenType string

// Token represents a lexical token with its type and literal value
type Token struct {
	// Type is the category of the token.
	Type TokenType
	// Literal is the exact text from the source code.
	Literal string
}

// keywords maps language keywords to their TokenType.
var keywords = map[string]TokenType {
	"fun": FUNCTION,
	"karma": KARMA,
	"true": TRUE,
	"false": FALSE,
	"if": IF,
	"else": ELSE,
	"return": RETURN,
}

// Special tokens
const (
	ILLEGAL = "ILLEGAL" // unknown token/character
	EOF = "EOF"         // end of file

	// Identifiers + literals
	IDENT = "IDENT"
	INT = "INT"

	// Operators
	ASSIGN = "="
	PLUS = "+"
	MINUS = "-"
	ASTERISK = "*"
	SLASH = "/"
	BANG = "!"

	LT = "<"
	GT = ">"
	EQ = "=="
	NOT_EQ = "!="

	// Delimiters
	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	FUNCTION = "FUNCTION"
	KARMA = "KARMA"
	TRUE = "TRUE"
	FALSE = "FALSE"
	IF = "IF"
	ELSE = "ELSE"
	RETURN = "RETURN"
)

// LookupIdent checks if an identifier is a keyword, returning the proper TokenType.
func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}