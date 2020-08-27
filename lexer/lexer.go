package lexer

import (
	"github.com/khardi/gomonkey/token"
)

func New(s string) *Lexer {
	lex := &Lexer{input: s}
	lex.readCh()
	return lex 
}

type Lexer struct {
	input string
	ch byte
	pos int
}

func (lex *Lexer) eof() bool {
	return lex.pos >= len(lex.input)
}

func (lex *Lexer) readCh() {
	if lex.eof() {
		lex.ch = 0
		return
	}
	lex.ch = lex.input[lex.pos]
	lex.pos += 1	
}

func (lex *Lexer) NextToken() (*token.Token) {
	tok := &token.Token{token.ILLEGAL, "ILLEGAL"}
	
	s := string(lex.ch)
	
	switch lex.ch {

	case '=':
		tok = &token.Token{token.ASSIGN, s}
		break

	case '+':
		tok = &token.Token{token.PLUS, s}
		break

	case '(':
		tok = &token.Token{token.LPAREN, s}
		break

	case ')':
		tok = &token.Token{token.RPAREN, s}
		break

	case '{':
		tok = &token.Token{token.LBRACE, s}
		break

	case '}':
		tok = &token.Token{token.RBRACE, s}
		break

	case ',':
		tok = &token.Token{token.COMMA, s}
		break

	case ';':
		tok = &token.Token{token.SEMICOLON, s}
		break

	case 0:
		tok = &token.Token{token.EOF, ""}
		break
	}

	lex.readCh()
	
	return tok
}
