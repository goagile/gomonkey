package lexer

import (
	// "fmt"
	"io"
	"strings"
	// "github.com/khardi/gomonkey/token"
)

func NewFromString(s string) *Lexer {
	lex := &Lexer{
		r: strings.NewReader(s),
		buf: make([]byte, 0),
	}
	return lex 
}

func New(r io.Reader) *Lexer {
	lex := &Lexer{
		r: r,
		buf: make([]byte, 0),
	}
	return lex 
}

type Lexer struct {
	r io.Reader
	ch byte
	buf []byte
}

func (lex *Lexer) read() {
	b := make([]byte, 1)
	if _, err := lex.r.Read(b); err == io.EOF {
		lex.ch = 0
		return
	}
	lex.ch = b[0]
	lex.buf = append(lex.buf, lex.ch)
}

func (lex *Lexer) literal() string {
	return string(lex.buf)
}

// func New(r io.Reader) *Lexer {
// 	lex := &Lexer{r: r}
// 	lex.readCh()
// 	return lex 
// }

// type Lexer struct {
// 	r io.Reader
// 	ch byte
// 	pos int
// 	prev int
// }

// func (lex *Lexer) readCh() {
// 	// if lex.pos >= len(lex.input) {
// 	// 	// fmt.Println("lex.pos >= len(lex.input)", lex.pos >= len(lex.input))
// 	// 	lex.ch = 0
// 	// 	return
// 	// }
// 	b := make([]byte, 1)
// 	n, err := io.Read(b)
// 	if err == io.EOF {
// 		return
// 	}

// 	lex.ch = b
// 	// lex.prev = lex.pos
// 	// lex.pos += 1
// }

// func (lex *Lexer) scrollWhitespace() {
// 	for ;(' ' == lex.ch || '\n' == lex.ch || '\t' == lex.ch || '\r' == lex.ch); {
// 		lex.readCh()
// 	}
// }

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

// func (lex *Lexer) lookupIdent(literal string) token.TokenType {
// 	switch literal {

// 	case "let":
// 		return token.LET

// 	case "fn":
// 		return token.FUNCTION

// 	case "if":
// 		return token.IF

// 	case "else":
// 		return token.ELSE

// 	case "return":
// 		return token.RETURN

// 	case "true":
// 		return token.TRUE

// 	case "false":
// 		return token.FALSE
// 	}

// 	return token.IDENT
// }
