// doc.go
//
// Package token defines the tokens used by the Karma programming language lexer.
//
// Tokens are the smallest unit of meaning in the source code. Before parsing,
// the lexer converts raw source code into a stream of tokens. Tokens allow the
// parser to distinguish between numbers, variable names, keywords, and special
// characters without worrying about the exact text.
//
// Example source code:
//
//     karma five = 5;
//     karma ten = 10;
//     karma add = fn(x, y) {
//         x + y;
//     };
//     karma result = add(five, ten);
//
// This code contains the following types of tokens:
//   - Integers: 5, 10
//   - Identifiers: five, ten, add, x, y, result
//   - Keywords: karma, fn
//   - Special characters: (, ), {, }, =, ,, ;
//
// TokenType represents the category of the token, e.g. "IDENT" for identifiers,
// "INT" for integer literals, "KARMA" for the karma(let) keyword.
//
// Token has a Type (to classify the token), and a Literal (the exact string
// from the source code). Example: {"INT", 5}
//
// Special token types:
//   - ILLEGAL: token/character that does not exist
//   - EOF: end of file

package token
