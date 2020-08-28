package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

func main() {

	var s scanner.Scanner

	line := `

		if ten != 10 {
			return true
		} else if ten == 22 {
			return false
		} else if a >= b {
			return true
		}

	`

	// line := `let x = 234;`

	s.Init(strings.NewReader(line))

	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {

		if tok == '!' && s.Peek() == '=' {
			fmt.Println("NOT_EQ")

		} else if tok == '=' && s.Peek() == '=' {
			fmt.Println("EQ")

		} else if tok == '>' && s.Peek() == '=' {
			fmt.Println("GTE")

		}
	}

}
