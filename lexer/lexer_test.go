package lexer

import (
	"testing"
	"reflect"
	"github.com/khardi/gomonkey/token"
)

//
// read ch
//
func Test_read_ch_a(t *testing.T) {
	want := "a"

	s := "ab"
	lex := NewFromString(s)
	lex.read()

	got := string(lex.ch)

	if want != got {
		t.Fatalf("\nwant:%q\ngot:%q\n", want, got)
	}
}

func Test_read_ch_b(t *testing.T) {
	want := "b"

	s := "ab"
	lex := NewFromString(s)
	lex.read()
	lex.read()

	got := string(lex.ch)

	if want != got {
		t.Fatalf("\nwant:%q\ngot:%q\n", want, got)
	}
}

func Test_read_ch_0(t *testing.T) {
	want := "\x00"

	s := "ab"
	lex := NewFromString(s)
	lex.read()
	lex.read()
	lex.read()

	got := string(lex.ch)

	if want != got {
		t.Fatalf("\nwant:%q\ngot:%q\n", want, got)
	}
}

//
// readbuf literal
//
func Test_readbuf_literal_a(t *testing.T) {
	want := "a"

	s := "ab"
	lex := NewFromString(s)
	lex.readbuf()

	got := lex.literal()

	if want != got {
		t.Fatalf("\nwant:%q\ngot:%q\n", want, got)
	}
}

func Test_read_literal_b(t *testing.T) {
	want := "ab"

	s := "ab"
	lex := NewFromString(s)
	lex.readbuf()
	lex.readbuf()

	got := lex.literal()

	if want != got {
		t.Fatalf("\nwant:%q\ngot:%q\n", want, got)
	}
}

func Test_read_literal_0(t *testing.T) {
	want := "ab\x00"

	s := "ab"
	lex := NewFromString(s)
	lex.readbuf()
	lex.readbuf()
	lex.readbuf()

	got := lex.literal()

	if want != got {
		t.Fatalf("\nwant:%q\ngot:%q\n", want, got)
	}
}

//
// clear buf
//
func Test_clearbuf(t *testing.T) {
	want := "b"

	s := "ab"
	lex := NewFromString(s)
	lex.readbuf()
	lex.clearbuf()
	lex.readbuf()

	got := lex.literal()

	if want != got {
		t.Fatalf("\nwant:%q\ngot:%q\n", want, got)
	}
}

// 
// scrollWhitespace
//
func Test_scrollWhitespace_a(t *testing.T) {
	want := "a"

	s := "  \n   a  \t   b   "
	lex := NewFromString(s)

	lex.read()
	lex.scrollWhitespace()

	got := string(lex.ch)

	if want != got {
		t.Fatalf("\nwant:%q\ngot:%q\n", want, got)
	}
}

func Test_scrollWhitespace_b(t *testing.T) {
	want := "b"

	s := "  \n   a  \t   b   "
	lex := NewFromString(s)

	lex.read()
	lex.scrollWhitespace()
	
	lex.read()
	lex.scrollWhitespace()

	got := string(lex.ch)

	if want != got {
		t.Fatalf("\nwant:%q\ngot:%q\n", want, got)
	}
}

//
// NextToken
//
func Test_Token_ASSIGN(t *testing.T) {
	want := token.NewAssign()
	lex := NewFromString("=")
	
	got := lex.Token()

	if !want.Equal(got) {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}

func Test_Token_ASSIGN_scrollWhitespace(t *testing.T) {
	want := token.NewAssign()
	lex := NewFromString("  \t  \n     =   ")

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
	lex := NewFromString("==")

	got := []*token.Token{
		lex.Token(),
		lex.Token(),
	}

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
	}
}


// func Test_Token_IDENT(t *testing.T) {
// 	want := token.NewIdent("ten")
// 	lex := NewFromString("ten")

// 	got := lex.Token()

// 	if !want.Equal(got) {
// 		t.Fatalf("\nwant:%v\ngot:%v\n", want, got)
// 	}
// }

// func Test_NextToken(t *testing.T) {
	
// 	lex := New(strings.NewReader("=+-*/<>(){},;"))

// 	wants := []*token.Token{
// 		{token.ASSIGN,   "="},
// 		{token.PLUS,     "+"},
// 		{token.MINUS,    "-"},
// 		{token.ASTERISK, "*"},
// 		{token.SLASH, 	 "/"},
// 		{token.LT,       "<"},
// 		{token.GT,       ">"},
// 		{token.LPAREN,   "("},
// 		{token.RPAREN,   ")"},
// 		{token.LBRACE,   "{"},
// 		{token.RBRACE,   "}"},
// 		{token.COMMA,    ","},
// 		{token.SEMICOLON,";"},
// 		{token.EOF,      ""},
// 	}

// 	for i, want := range wants {
// 		got := lex.NextToken()

// 		if want.Type != got.Type {
// 			t.Fatalf("\nType\ni:%v\nwant:%v\ngot:%v\n", i, want.Type, got.Type)
// 		}
// 		if want.Literal != got.Literal {
// 			t.Fatalf("\nLiteral\ni:%v\nwant:%v\ngot:%v\n", i, want.Literal, got.Literal)
// 		}
// 	}
// }

// func Test_NextToken_Let_10(t *testing.T) {
	
// 	lex := New(strings.NewReader("let ten = 10;"))

// 	wants := []*token.Token{
// 		{token.LET, "let"},
// 		{token.IDENT, "ten"},
// 		{token.ASSIGN, "="},
// 		{token.INT, "10"},
// 		{token.SEMICOLON, ";"},
// 		{token.EOF,   ""},
// 	}

// 	for i, want := range wants {
// 		got := lex.NextToken()

// 		if want.Type != got.Type {
// 			t.Fatalf("\nType\ni:%v\nwant:%v\ngot:%v\n", i, want.Type, got.Type)
// 		}
// 		if want.Literal != got.Literal {
// 			t.Fatalf("\nLiteral\ni:%v\nwant:%v\ngot:%v\n", i, want.Literal, got.Literal)
// 		}
// 	}
// }

// func Test_NextToken_Fn(t *testing.T) {

// 	lex := New(strings.NewReader("let f = fn(a, b) { a + b };"))

// 	wants := []*token.Token{
// 		{token.LET,      "let"},
// 		{token.IDENT,    "f"},
// 		{token.ASSIGN,   "="},
// 		{token.FUNCTION, "fn"},
// 		{token.LPAREN,   "("},
// 		{token.IDENT,    "a"},
// 		{token.COMMA,    ","},
// 		{token.IDENT,    "b"},
// 		{token.RPAREN,   ")"},
// 		{token.LBRACE,   "{"},
// 		{token.IDENT,    "a"},
// 		{token.PLUS,     "+"},
// 		{token.IDENT,    "b"},
// 		{token.RBRACE,   "}"},
// 		{token.SEMICOLON,";"},
// 		{token.EOF,      ""},
// 	}

// 	for i, want := range wants {
// 		got := lex.NextToken()

// 		if want.Type != got.Type {
// 			t.Fatalf("\nType\ni:%v\nwant:%v\ngot:%v\n", i, want.Type, got.Type)
// 		}
// 		if want.Literal != got.Literal {
// 			t.Fatalf("\nLiteral\ni:%v\nwant:%v\ngot:%v\n", i, want.Literal, got.Literal)
// 		}
// 	}
// }

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