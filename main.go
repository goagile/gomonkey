package main

import (
	"fmt"
	"github.com/khardi/gomonkey/token"
) 

func main() {

	t := &token.Token{0, "+"}

	fmt.Println("Hello", t)

}