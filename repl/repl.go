package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/khardi/gomonkey/lexer"
	"github.com/khardi/gomonkey/token"
)

const Prompt = ">> "

func Start(in io.Reader, out io.Writer) {
	s := bufio.NewScanner(in)
	for s.Scan() {
		fmt.Print(Prompt)
		lex := lexer.NewFromString(s.Text())
		for tok := lex.Token(); !tok.Equal(token.NewEOF()); tok = lex.Token() {
			fmt.Println(tok)
		}
	}
}
