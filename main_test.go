package main

import "testing"

func fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return fib(n-2) + fib(n-1)
}

func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		fib(30) // run fib(30) b.N times
	}
}

// go test -bench="Fib$" -cpuprofile=cpu.pprof .
// 测试使用pprof
// go tool pprof -http=:9999 cpu.pprof
// go tool pprof cpu.pprof      top top --sum
