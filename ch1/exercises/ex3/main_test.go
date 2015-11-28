package main

import "testing"

func Benchmark_concatArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		concatArgs()
	}
}

func Benchmark_joinArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		joinArgs()
	}
}
