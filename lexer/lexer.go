package lexer

import (
	// "fmt"

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
	prev int
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
	lex.prev = lex.pos
	lex.pos += 1	
}

func (lex *Lexer) scrollWhitespace() {
	for ;(' ' == lex.ch || '\n' == lex.ch || '\t' == lex.ch || '\r' == lex.ch); {
		lex.readCh()
	}
}

func (lex *Lexer) NextToken() (*token.Token) {
	tok := &token.Token{token.ILLEGAL, "ILLEGAL"}
	
	lex.scrollWhitespace()

	s := string(lex.ch)
	
	switch lex.ch {

	case '=':
		tok = &token.Token{token.ASSIGN, s}
		break

	case '+':
		tok = &token.Token{token.PLUS, s}
		break

	case '-':
		tok = &token.Token{token.MINUS, s}
		break

	case '*':
		tok = &token.Token{token.ASTERISK, s}
		break

	case '/':
		tok = &token.Token{token.SLASH, s}
		break

	case '<':
		tok = &token.Token{token.LT, s}
		break

	case '>':
		tok = &token.Token{token.GT, s}
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

	if lex.isLetter() {
		id := lex.readIdent()
		tt := lex.lookupIdent(id)
		return &token.Token{tt, id}
	}

	if lex.isDigit() {
		d := lex.readDigit()
		return &token.Token{token.INT, d}
	}

	lex.readCh()
	
	return tok
}

func (lex *Lexer) isLetter() bool {
	return ('a' <= lex.ch && lex.ch <= 'z' || 
		    'A' <= lex.ch && lex.ch <= 'Z' ||
		    '_' == lex.ch)
}

func (lex *Lexer) readIdent() string {
	start := lex.prev
	for lex.isLetter() {
		lex.readCh()
	}
	end := lex.prev
	return lex.input[start:end]
}

func (lex *Lexer) lookupIdent(literal string) token.TokenType {
	switch literal {

	case "let":
		return token.LET

	case "fn":
		return token.FUNCTION

	case "if":
		return token.IF

	case "else":
		return token.ELSE

	case "return":
		return token.RETURN

	case "true":
		return token.TRUE

	case "false":
		return token.FALSE
	}

	return token.IDENT
}

func (lex *Lexer) isDigit() bool {
	return ('0' <= lex.ch && lex.ch <= '9')
}

func (lex *Lexer) readDigit() string {
	start := lex.prev
	for lex.isDigit() {
		lex.readCh()
	}
	end := lex.prev
	return lex.input[start:end]
}
