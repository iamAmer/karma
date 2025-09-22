# Karma

Karma is my own interpreted programming language written in Go.  

## Structure
- **lexer/** – breaks input into tokens
- **token/** – defines the token types
- **ast/** – abstract syntax tree nodes
- **parser/** – builds AST from tokens
- **repl/** – interactive read-eval-print loop

## Current Progress
- Tokens defined
- Lexer implemented
- Parser supports `karma` and `return` statements so far