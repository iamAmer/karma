package lexer

import "karma/token"

type Lexer struct {
    input        string
    position     int
    readPosition int
    ch           byte
}

func New(input string) *Lexer {
	l := &Lexer {input : input}
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
	l.readPosition++
}

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

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{
		Type: tokenType,
		Literal: string(ch),
	}
}

// what is allowed in identifiers
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool{
	return '0' <= ch && ch <= '9'

	// also valid for ASCII version
	// return 48 <= ch && ch <= 57
}

func (l *Lexer) readIdentifier() string{
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func (l *Lexer) readNumber() string{
	position := l.position;
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[position : l.position]
}

func (l *Lexer) peekChar() byte{
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) makeTwoCharToken(expectedChar byte, twoCharType, singleCharType token.TokenType) token.Token {
    if l.peekChar() == expectedChar {
        ch := l.ch
        l.readChar()
        return token.Token{Type: twoCharType, Literal: string(ch) + string(l.ch)}
    }
    return newToken(singleCharType, l.ch)
}