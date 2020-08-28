package token

import (
	"fmt"
)

type TokenType string

const (

	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	IDENT = "IDENT"
	INT = "INT"

	ASSIGN = "ASSIGN"
	PLUS = "PLUS"
	MINUS = "-"
	ASTERISK = "*"
	SLASH = "/"
	LT = "<"
	GT = ">"

	COMMA = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET = "LET"
	IF = "IF"
	ELSE = "ELSE"
	RETURN = "RETURN"
	TRUE = "TRUE"
	FALSE = "FALSE"

	EQ = "EQ"
)

type Token struct {
	Literal string
	TokenType TokenType
}

func (tok *Token) String() string {
	return fmt.Sprintf("Token(%v, %v)", tok.Literal, tok.TokenType)
}

func New(literal string, tokentype TokenType) *Token {
	return &Token{Literal: literal, TokenType: tokentype}
}

func (tok *Token) Equal(other *Token) bool {
	return tok.Literal == other.Literal && tok.TokenType == other.TokenType
}

func NewAssign() *Token {
	return &Token{Literal: "=", TokenType: ASSIGN}
}

func NewPlus() *Token {
	return &Token{Literal: "+", TokenType: PLUS}
}

func NewIllegal(literal string) *Token {
	return &Token{Literal: literal, TokenType: ILLEGAL}
}

func NewEq() *Token {
	return &Token{Literal: "==", TokenType: EQ}
}

func NewIdent(literal string) *Token {
	return &Token{Literal: literal, TokenType: IDENT}
}
