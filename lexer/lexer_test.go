package lexer

import (
	"reflect"
	"testing"

	"github.com/khardi/gomonkey/token"
)

//
// Token
//
func Test_Token_EOF(t *testing.T) {
	want := token.NewEOF()
	lex := NewFromString("")

	got := lex.Token()

	if !want.Equal(got) {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_Token_ASSIGN(t *testing.T) {
	want := token.NewAssign()
	lex := NewFromString("=")

	got := lex.Token()

	if !want.Equal(got) {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_Token_PLUS(t *testing.T) {
	want := token.NewPlus()
	lex := NewFromString("+")

	got := lex.Token()

	if !want.Equal(got) {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_Token_EQ(t *testing.T) {
	want := token.NewEq()
	lex := NewFromString("==")

	got := lex.Token()

	if !want.Equal(got) {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_Token_EQ_First(t *testing.T) {
	want := token.NewEq()
	lex := NewFromString("==+")

	got := lex.Token()

	if !want.Equal(got) {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_Token_EQ_Second(t *testing.T) {
	want := token.NewPlus()
	lex := NewFromString("==+")
	lex.Token()

	got := lex.Token()

	if !want.Equal(got) {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_Token_EQ_EQ(t *testing.T) {
	want := []*token.Token{
		token.NewEq(),
		token.NewEq(),
	}
	lex := NewFromString("====")

	got := []*token.Token{
		lex.Token(),
		lex.Token(),
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_Token_IDENT(t *testing.T) {
	want := token.NewIdent("variable")
	lex := NewFromString("variable")

	got := lex.Token()

	if !want.Equal(got) {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

//
// Let
//
func Test_Token_LET(t *testing.T) {
	want := token.NewLet()
	lex := NewFromString(" let x = 2; ")

	got := lex.Token()

	if !want.Equal(got) {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_Token_Operation(t *testing.T) {

	lex := NewFromString("=+-*/<>(){},;")

	for i, want := range []*token.Token{
		token.NewAssign(),
		token.NewPlus(),
		token.NewMinus(),
		token.NewAsterisk(),
		token.NewSlash(),
		token.NewLt(),
		token.NewGt(),
		token.NewLparen(),
		token.NewRparen(),
		token.NewLbrace(),
		token.NewRbrace(),
		token.NewComma(),
		token.NewSemicolon(),
	} {
		got := lex.Token()

		if !want.Equal(got) {
			t.Fatalf("\nexample:%v\nwant:%v\ngot:%v\n",
				i+1, want, got,
			)
		}
	}
}

func Test_Token_Let(t *testing.T) {

	lex := NewFromString("let ten = 10;")

	for i, want := range []*token.Token{
		token.NewLet(),
		token.NewIdent("ten"),
		token.NewAssign(),
		token.NewIdent("10"),
		token.NewSemicolon(),
	} {
		got := lex.Token()

		if !want.Equal(got) {
			t.Fatalf("\nexample:%v\nwant:%v\ngot:%v\n",
				i+1, want, got,
			)
		}
	}
}

func Test_Token_Fn(t *testing.T) {

	lex := NewFromString("let f = fn(a, b) { a + b };")

	for i, want := range []*token.Token{
		token.NewLet(),
		token.NewIdent("f"),
		token.NewAssign(),
		token.NewFn(),
		token.NewLparen(),
		token.NewIdent("a"),
		token.NewComma(),
		token.NewIdent("b"),
		token.NewRparen(),
		token.NewLbrace(),
		token.NewIdent("a"),
		token.NewPlus(),
		token.NewIdent("b"),
		token.NewRbrace(),
		token.NewSemicolon(),
	} {
		got := lex.Token()

		if !want.Equal(got) {
			t.Fatalf("\nexample:%v\nwant:%v\ngot:%v\n",
				i+1, want, got,
			)
		}
	}
}

// func Test_NextToken_If_Else_Return_True_False(t *testing.T) {

// 	lex := New(strings.NewReader(
// 		"    if (5 < 10) {       " +
// 		"        return true;    " +
// 		"    } else {            " +
// 		"        return false;   " +
// 		"    }                   ",
// 	))

// 	wants := []*token.Token{
// 		{token.IF, "if"},
// 		{token.LPAREN, "("},
// 		{token.INT, "5"},
// 		{token.LT, "<"},
// 		{token.INT, "10"},
// 		{token.RPAREN, ")"},
// 		{token.LBRACE, "{"},
// 		{token.RETURN, "return"},
// 		{token.TRUE, "true"},
// 		{token.SEMICOLON, ";"},
// 		{token.RBRACE, "}"},
// 		{token.ELSE, "else"},
// 		{token.LBRACE, "{"},
// 		{token.RETURN, "return"},
// 		{token.FALSE, "false"},
// 		{token.SEMICOLON, ";"},
// 		{token.RBRACE, "}"},
// 		{token.EOF,   ""},
// 	}

// 	for i, want := range wants {
// 		got := lex.NextToken()

// 		if want.Type != got.Type {
// 			t.Fatalf("\nType\ni:%v\nwant:%v\ngot:%v\n", i, want, got)
// 		}
// 		if want.Literal != got.Literal {
// 			t.Fatalf("\nLiteral\ni:%v\nwant:%v\ngot:%v\n", i, want, got)
// 		}
// 	}
// }
