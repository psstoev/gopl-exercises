package main

import (
	"fmt"

	"./tempconv"
)

func main() {
	fmt.Println(tempconv.KToF(tempconv.CToK(tempconv.AbsoluteZeroC)))
	fmt.Println(tempconv.CToF(tempconv.AbsoluteZeroC))
}
