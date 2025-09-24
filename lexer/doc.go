// Package lexer implements the lexical scanner for the Karma programming language.
//
// Package lexer takes raw source code as input and converts it into a stream of tokens.
// Each token represents a meaningful element of the source code, such as identifiers,
// keywords, literals, or operators.
//
// This package provides a simple, stateful lexer. It is initialized with a source
// code string and exposes a single method, NextToken, which returns the next token
// each time it is called. The lexer does not store or buffer tokens — it simply scans
// character by character and produces tokens on demand.
//
// Note: For simplicity, this lexer works with an in-memory string and does not track
// filenames or line numbers. In a production environment, using an io.Reader and
// recording position metadata would be preferable.
//
// Enhancements:
//   - [ ] Support filename and line number in tokens
//   - [ ] Unicode support
package lexer
