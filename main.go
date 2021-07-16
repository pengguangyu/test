package main

import (
	"fmt"
)

func JoinStrUseSprint(a, b string) string {
	return fmt.Sprintf("%s%s", a, b)
}

func JoinStrUseNor(a, b string) string {
	return a + b
}

// 运行所有的测试
// go test -v -bench=.  main_test.go main.go
// 运行所有的基准测试函数
// go test -bench=. -benchtime=1s -benchmem -count=10
func main() {
}
