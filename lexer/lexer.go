package lexer

import (
	// "fmt"
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

	if tokentype, ok := reserved[literal]; ok {
		return token.New(literal, tokentype)
	}

	return token.NewIdent(literal)
}

// func (lex *Lexer) NextToken() (*token.Token) {
// 	tok := &token.Token{token.ILLEGAL, "ILLEGAL"}

// 	lex.scrollWhitespace()

// 	s := string(lex.ch)

// 	switch lex.ch {

// 	case '=':
// 		tok = &token.Token{token.ASSIGN, s}
// 		break

// 	case '+':
// 		tok = &token.Token{token.PLUS, s}
// 		break

// 	case '-':
// 		tok = &token.Token{token.MINUS, s}
// 		break

// 	case '*':
// 		tok = &token.Token{token.ASTERISK, s}
// 		break

// 	case '/':
// 		tok = &token.Token{token.SLASH, s}
// 		break

// 	case '<':
// 		tok = &token.Token{token.LT, s}
// 		break

// 	case '>':
// 		tok = &token.Token{token.GT, s}
// 		break

// 	case '(':
// 		tok = &token.Token{token.LPAREN, s}
// 		break

// 	case ')':
// 		tok = &token.Token{token.RPAREN, s}
// 		break

// 	case '{':
// 		tok = &token.Token{token.LBRACE, s}
// 		break

// 	case '}':
// 		tok = &token.Token{token.RBRACE, s}
// 		break

// 	case ',':
// 		tok = &token.Token{token.COMMA, s}
// 		break

// 	case ';':
// 		tok = &token.Token{token.SEMICOLON, s}
// 		break

// 	case 0:
// 		tok = &token.Token{token.EOF, ""}
// 		break
// 	}

// 	if lex.isLetter() {
// 		id := lex.readIdent()
// 		tt := lex.lookupIdent(id)
// 		return &token.Token{tt, id}
// 	}

// 	if lex.isDigit() {
// 		d := lex.readDigit()
// 		return &token.Token{token.INT, d}
// 	}

// 	lex.readCh()

// 	return tok
// }

// func (lex *Lexer) isLetter() bool {
// 	return ('a' <= lex.ch && lex.ch <= 'z' ||
// 		    'A' <= lex.ch && lex.ch <= 'Z' ||
// 		    '_' == lex.ch)
// }

// func (lex *Lexer) isDigit() bool {
// 	return ('0' <= lex.ch && lex.ch <= '9')
// }

// func (lex *Lexer) readIdent() string {
// 	start := lex.prev
// 	for lex.isLetter() {
// 		lex.readCh()
// 	}
// 	end := lex.prev
// 	return lex.input[start:end]
// }

// func (lex *Lexer) readDigit() string {
// 	start := lex.prev
// 	// fmt.Println("lex.isDigit()", lex.isDigit())
// 	for lex.isDigit() {
// 		lex.readCh()
// 	}
// 	end := lex.prev
// 	return lex.input[start:end]
// }

var reserved = map[string]token.TokenType{
	"let":    token.LET,
	"fn":     token.FUNCTION,
	"if":     token.IF,
	"else":   token.ELSE,
	"return": token.RETURN,
	"true":   token.TRUE,
	"false":  token.FALSE,
}
