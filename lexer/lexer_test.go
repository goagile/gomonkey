package lexer

import (
	"testing"
	"github.com/khardi/gomonkey/token"
)

func Test_NextToken(t *testing.T) {
	wants := []*token.Token{
		{token.ASSIGN,   "="},
		{token.PLUS,     "+"},
		{token.MINUS,    "-"},
		{token.ASTERISK, "*"},
		{token.SLASH, 	 "/"},
		{token.LT,       "<"},
		{token.GT,       ">"},
		{token.LPAREN,   "("},
		{token.RPAREN,   ")"},
		{token.LBRACE,   "{"},
		{token.RBRACE,   "}"},
		{token.COMMA,    ","},
		{token.SEMICOLON,";"},
		{token.EOF,      ""},
	}

	lex := New("=+-*/<>(){},;")

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

func Test_NextToken_Let_10(t *testing.T) {
	wants := []*token.Token{
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.EOF,   ""},
	}

	lex := New("let ten = 10;")

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

func Test_NextToken_Fn(t *testing.T) {
	wants := []*token.Token{
		{token.LET,      "let"},
		{token.IDENT,    "f"},
		{token.ASSIGN,   "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN,   "("},
		{token.IDENT,    "a"},
		{token.COMMA,    ","},
		{token.IDENT,    "b"},
		{token.RPAREN,   ")"},
		{token.LBRACE,   "{"},
		{token.IDENT,    "a"},
		{token.PLUS,     "+"},
		{token.IDENT,    "b"},
		{token.RBRACE,   "}"},
		{token.SEMICOLON,";"},
		{token.EOF,      ""},
	}

	lex := New("let f = fn(a, b) { a + b };")

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
