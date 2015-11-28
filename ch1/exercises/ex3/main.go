package main

import (
	"fmt"
	"os"
	"strings"
)

func concatArgs() string {
	var s, sep string

	sep = ""

	for _, arg := range os.Args[1:] {
		s = sep + arg
		sep = " "
	}

	return s
}

func joinArgs() string {
	return strings.Join(os.Args[1:], " ")
}

func main() {
	fmt.Println(joinArgs())
}
