package token

import (
	"fmt"
)

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN   = "ASSIGN"
	PLUS     = "PLUS"
	MINUS    = "MINUS"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"

	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	TRUE     = "TRUE"
	FALSE    = "FALSE"

	EQ = "EQ"
)

type Token struct {
	Literal   string
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

func NewMinus() *Token {
	return &Token{Literal: "-", TokenType: MINUS}
}

func NewAsterisk() *Token {
	return &Token{Literal: "*", TokenType: ASTERISK}
}

func NewSlash() *Token {
	return &Token{Literal: "/", TokenType: SLASH}
}

func NewLt() *Token {
	return &Token{Literal: "<", TokenType: LT}
}

func NewGt() *Token {
	return &Token{Literal: ">", TokenType: GT}
}

func NewLparen() *Token {
	return &Token{Literal: "(", TokenType: LPAREN}
}

func NewRparen() *Token {
	return &Token{Literal: ")", TokenType: RPAREN}
}

func NewLbrace() *Token {
	return &Token{Literal: "{", TokenType: LBRACE}
}

func NewRbrace() *Token {
	return &Token{Literal: "}", TokenType: RBRACE}
}

func NewComma() *Token {
	return &Token{Literal: ",", TokenType: COMMA}
}

func NewSemicolon() *Token {
	return &Token{Literal: ";", TokenType: SEMICOLON}
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

func NewEOF() *Token {
	return &Token{Literal: "EOF", TokenType: EOF}
}

func NewLet() *Token {
	return &Token{Literal: "let", TokenType: LET}
}

func NewFn() *Token {
	return &Token{Literal: "fn", TokenType: FUNCTION}
}
