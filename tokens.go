package main

import "fmt"

type TokenType int8

const (
	// Single-character tokens.
	TT_LEFT_PAREN = iota
	TT_RIGHT_PAREN
	TT_LEFT_BRACE
	TT_RIGHT_BRACE
	TT_COMMA
	TT_DOT
	TT_MINUS
	TT_PLUS
	TT_SEMICOLON
	TT_SLASH
	TT_STAR

	// One or two character tokens.
	TT_BANG
	TT_BANG_EQUAL
	TT_EQUAL
	TT_EQUAL_EQUAL
	TT_GREATER
	TT_GREATER_EQUAL
	TT_LESS
	TT_LESS_EQUAL

	// Literals.
	TT_IDENTIFIER
	TT_STRING
	TT_NUMBER

	// Keywords.
	TT_AND
	TT_CLASS
	TT_ELSE
	TT_FALSE
	TT_FUN
	TT_FOR
	TT_IF
	TT_NIL
	TT_OR
	TT_PRINT
	TT_RETURN
	TT_SUPER
	TT_THIS
	TT_TRUE
	TT_VAR
	TT_WHILE

	// Other
	TT_EOF
)

type Token struct {
	TokenType TokenType
	Lexeme    string
	Literal   interface{}
	Line      int
}

func (t *Token) String() string {
	return fmt.Sprintf("%d %s %s", t.TokenType, t.Lexeme, t.Literal)
}
