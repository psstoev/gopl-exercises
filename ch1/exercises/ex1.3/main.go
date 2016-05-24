//!+

// Benchmark += against strings.Join()
package main

import (
	"math/rand"
	"strings"
	"time"
)

func SlowJoin(args []string) {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
}

func FastJoin(args []string) {
	strings.Join(args, " ")
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

//!-
