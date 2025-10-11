package ast

import (
	"karma/token"
	"testing"
)

// type LetStatement struct {
// 	Token token.Token
// 	Name  *Identifier
// type Identifier struct {
// Token token.Token
// Value string
// }
// 	Value Expression
// }

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			// karma myVar = anotherVar;
			&LetStatement{
				Token: token.Token{Type: token.KARMA, Literal: "karma"},
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}
	if program.String() != "karma myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
