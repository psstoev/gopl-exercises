package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	sep = ""

	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
