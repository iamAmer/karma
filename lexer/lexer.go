package lexer

import "karma/token"

// Lexer turns Karma source code into a stream of tokens.
type Lexer struct {
    input        string
    position     int  // index of current char in input
    readPosition int  // index of the next char to read
    ch           byte // current char
}

// New creates and initializes a new Lexer for the given input string.
func New(input string) *Lexer {
	l := &Lexer {input : input}
	l.readChar()
	return l
}

// readChar advances the lexer by one character, updating l.ch, l.position,
// and l.readPosition.
func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

// NextToken scans the next token from the input and returns it.
// It skips over whitespace and handles single-character operators,
// multi-character operators (like ==, !=), identifiers, keywords, and numbers.
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
		case '=': 
			tok = l.makeTwoCharToken('=', token.EQ, token.ASSIGN)
		case '+': 
			tok = newToken(token.PLUS, l.ch)
		case '-':
			tok = newToken(token.MINUS, l.ch)
		case '*':
			tok = newToken(token.ASTERISK, l.ch)
		case '/':
			tok = newToken(token.SLASH, l.ch)
		case '!':
			tok = l.makeTwoCharToken('=', token.NOT_EQ, token.BANG)
		case '<':
			tok = newToken(token.LT, l.ch)
		case '>':
			tok = newToken(token.GT, l.ch)
		case '(': 
			tok = newToken(token.LPAREN, l.ch)
		case ')': 
			tok = newToken(token.RPAREN, l.ch)
		case '{': 
			tok = newToken(token.LBRACE, l.ch)
		case '}': 
			tok = newToken(token.RBRACE, l.ch)
		case ',': 
			tok = newToken(token.COMMA, l.ch)
		case ';': 
			tok = newToken(token.SEMICOLON, l.ch)
		case 0: 
			tok.Type = token.EOF
			tok.Literal = ""
		default:
			if isLetter(l.ch) {
				tok.Literal = l.readIdentifier()
				tok.Type = token.LookupIdent(tok.Literal)
				return tok
			} else if isDigit(l.ch) {
				tok.Type = token.INT
				tok.Literal = l.readNumber()
				return tok
			} else {
				tok = newToken(token.ILLEGAL, l.ch)
			}
	}
	
	l.readChar()
	return tok
}

// newToken creates a new token.Token with the given type and single-character literal.
//
// It’s a helper for single-character tokens such as '+', '-', '{', '}'.
// Multi-character tokens like identifiers or numbers are handled separately.
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type: tokenType,
		Literal: string(ch),
	}
}

// isLetter reports whether ch is a valid identifier letter in Karma.
// Letters include A–Z, a–z, and underscore (_).
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// isDigit reports whether ch is an ASCII digit 0–9.
func isDigit(ch byte) bool{
	return '0' <= ch && ch <= '9'

	// also valid for ASCII version
	// return 48 <= ch && ch <= 57
}

// readIdentifier consumes an identifier from the input starting at l.position.
// Identifiers consist of letters and underscores. It returns the identifier string.
func (l *Lexer) readIdentifier() string{
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// skipWhitespace advances the lexer past any whitespace characters:
// spaces, tabs, carriage returns, and newlines.
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// readNumber consumes a contiguous run of digits from the input starting at l.position.
// It returns the number literal as a string.
func (l *Lexer) readNumber() string{
	position := l.position;
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position : l.position]
}

// peekChar returns the next character without advancing the lexer.
// If the end of input is reached, it returns 0.
func (l *Lexer) peekChar() byte{
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

// makeTwoCharToken checks if the next character matches expectedChar to form
// a two-character operator (like == or !=). If so, it consumes the second
// character and returns the combined token. Otherwise it returns the single-char token.
func (l *Lexer) makeTwoCharToken(expectedChar byte, twoCharType, singleCharType token.TokenType) token.Token {
    if l.peekChar() == expectedChar {
        ch := l.ch
        l.readChar()
        return token.Token{Type: twoCharType, Literal: string(ch) + string(l.ch)}
    }
    return newToken(singleCharType, l.ch)
}