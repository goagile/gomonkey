package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/khardi/gomonkey/repl"
)

func main() {
	u, err := user.Current()
	if err != nil {
		log.Fatal("cannot read the user")
	}
	fmt.Printf("Hello, %v!\n", u.Username)

	repl.Start(os.Stdin, os.Stdout)
}
