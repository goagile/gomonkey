package main

import (
	"fmt"
	"github.com/khardi/gomonkey/token"
) 

func main() {

	t := &token.Token{token.PLUS, "+"}

	fmt.Println("Hello", t)

}