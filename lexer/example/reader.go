package main

import (
	"fmt"
	"io"
	"strings"
	"log"
)

func main() {
	
	r := strings.NewReader("abcd")
	
	read(r)

}

func read(r io.Reader) {
	
	b := make([]byte, 1)

	n, err := r.Read(b)
	if err == io.EOF {
		log.Fatal("EOF")
	} else if err != nil {
		log.Fatal(err)
	}

	fmt.Println("result -> ", n, string(b))

}