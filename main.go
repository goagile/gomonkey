package main

import (
	"fmt"
	// "github.com/khardi/gomonkey/token"
	"github.com/khardi/gomonkey/lexer"
	"strings"
) 

func main() {


	s := "343;"

	lex := lexer.New(strings.NewReader(s))

	for i := range len(s) {
		lex.read()
		fmt.Println(i, t)
	}

}
