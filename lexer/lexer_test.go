package lexer

import (
	"testing"
	"github.com/khardi/gomonkey/token"
)

func Test_NextToken(t *testing.T) {
	wants := []*token.Token{
		{token.ASSIGN, "="},
		{token.PLUS,   "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.EOF,   ""},
	}

	lex := New("=+(){},;")

	for i, want := range wants {
		got := lex.NextToken()

		if want.Type != got.Type {
			t.Fatalf("\nType\ni:%v\nwant:%v\ngot:%v\n", i, want.Type, got.Type)
		}
		if want.Literal != got.Literal {
			t.Fatalf("\nLiteral\ni:%v\nwant:%v\ngot:%v\n", i, want.Literal, got.Literal)
		}
	}
}
