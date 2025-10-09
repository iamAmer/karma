// Package parser implements the parsing stage of the Karma programming language.
//
// The parser takes a sequence of tokens produced by the lexer and transforms
// them into an Abstract Syntax Tree (AST). It provides mechanisms for reading
// tokens, detecting syntax errors, and building structured program
// representations that can later be interpreted or compiled.
package parser

import (
	"fmt"
	"karma/ast"
	"karma/lexer"
	"karma/token"
)

// Parser represents the syntactic analyzer for the Karma language.
type Parser struct {
	l *lexer.Lexer
	curToken token.Token
	peekToken token.Token

	errors []string
}

// New creates and returns a new Parser instance initialized with a given lexer.
func New(l *lexer.Lexer) *Parser {
	p := &Parser {
		l : l,
		errors: []string{},
	}
	
	p.nextToken()
	p.nextToken()

	return p
}

// nextToken advances the parser’s tokens by one position.
func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

// curTokenIs checks whether the current token’s type matches the given type.
func(p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

// peekTokenIs checks whether the next token’s type matches the given type.
func(p *Parser) peekTokenIS(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// expectedPeek checks whether the next token matches the expected type.
func(p *Parser) expectedPeek(t token.TokenType) bool {
	if p.peekTokenIS(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}


// parseLetStatement parses a `let` statement of the form:
//	karma <identifier> = <expression>;
func(p *Parser) parseLetStatement() *ast.LetStatement{
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectedPeek(token.IDENT) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectedPeek(token.ASSIGN) {
		return nil
	}

	// we skipped the expression part for now
	if !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parseReturnStatement parses a return statement of the form:
//	return <expression>;
func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	p.nextToken()

	//TODO: parse the expression

	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

// parseStatement determines which type of statement the current token represents
// and delegates to the appropriate parsing function.
func (p *Parser) parseStatement() ast.Statement{
	switch p.curToken.Type {
	case token.KARMA:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default: 
		return nil
	}
}

// ParseProgram parses a complete Karma program.
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()

		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		p.nextToken()
	}
	return program
}

// Errors returns all syntax errors collected during parsing.func (p *Parser) Errors() []string {
func (p *Parser) Errors() []string {
	return p.errors
}

// peekError records an error when the next token does not match the expected type.
func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}