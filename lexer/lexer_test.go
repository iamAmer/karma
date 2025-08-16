package lexer

import (
	"testing"
	"karma/token"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		t.Logf("Test %d: got token={Type:%s, Literal:%q}", i, tok.Type, tok.Literal)

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - token type mismatch: expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - token literal mismatch: expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}