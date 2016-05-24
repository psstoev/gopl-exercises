package main

import "testing"

var inputs [1000]string

func init() {
	for i := 0; i < len(inputs); i++ {
		inputs[i] = RandStringRunes(4)
	}
}

func BenchmarkSlowJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SlowJoin(inputs[:])
	}
}

func BenchmarkFastJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FastJoin(inputs[:])
	}
}
