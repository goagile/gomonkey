package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

func main() {

	var s scanner.Scanner

	code := `
	
		let x = 2;

		if x > 1 {
			return true
		} else {
			return false
		}
	
	`

	s.Init(strings.NewReader(code))

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		fmt.Println(s.TokenText())
	}

}
