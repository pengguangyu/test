package main

import "testing"

//           后缀是Fib的    每次1s         3轮      内存
// go test -bench='Fib$' -benchtime=1s -count=3 -benchmem .
func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(30) // run fib(30) b.N times
	}
}
