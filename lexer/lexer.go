package lexer

import (
	"io"
	"strings"
	"text/scanner"

	"github.com/khardi/gomonkey/token"
)

func NewFromString(s string) *Lexer {
	return New(strings.NewReader(s))
}

func New(r io.Reader) *Lexer {
	lex := &Lexer{}
	lex.scr.Init(r)
	return lex
}

type Lexer struct {
	scr scanner.Scanner
}

func (lex *Lexer) Token() *token.Token {

	tok := lex.scr.Scan()

	if tok == scanner.EOF {
		return token.NewEOF()
	}

	if tok == '=' && lex.scr.Peek() == '=' {
		lex.scr.Scan()
		return token.NewEq()
	}

	if tok == '=' {
		return token.NewAssign()
	}

	if tok == '+' {
		return token.NewPlus()
	}

	if tok == '-' {
		return token.NewMinus()
	}

	if tok == '*' {
		return token.NewAsterisk()
	}

	if tok == '/' {
		return token.NewSlash()
	}

	if tok == '<' {
		return token.NewLt()
	}

	if tok == '>' {
		return token.NewGt()
	}

	if tok == '(' {
		return token.NewLparen()
	}

	if tok == ')' {
		return token.NewRparen()
	}

	if tok == '{' {
		return token.NewLbrace()
	}

	if tok == '}' {
		return token.NewRbrace()
	}

	if tok == ',' {
		return token.NewComma()
	}

	if tok == ';' {
		return token.NewSemicolon()
	}

	literal := lex.scr.TokenText()

	if isInt(literal) {
		return token.NewInt(literal)
	}

	if tokentype, ok := token.Reserved[literal]; ok {
		return token.New(literal, tokentype)
	}

	return token.NewIdent(literal)
}

func isInt(literal string) bool {
	for _, tok := range literal {
		if tok < '0' || tok > '9' {
			return false
		}
	}
	return true
}
