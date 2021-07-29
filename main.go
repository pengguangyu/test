package main

import (
	"github.com/pkg/profile"
	"math/rand"
	"time"
)

//////////////////////
// CPU 性能分析 假设我们实现了这么一个程序，随机生成了 5 组数据，并且使用冒泡排序法排序。

func generate(n int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, 0)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Int())
	}
	return nums
}
func bubbleSort(nums []int) {
	for i := 0; i < len(nums); i++ {
		for j := 1; j < len(nums)-i; j++ {
			if nums[j] < nums[j-1] {
				nums[j], nums[j-1] = nums[j-1], nums[j]
			}
		}
	}
}

//////////////////
// 内存性能分析 假设我们实现了这么一个程序，生成长度为 N 的随机字符串，拼接在一起。

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func concat(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		s += randomString(n)
	}
	return s
}

func main() {
	// 为了简单，直接将数据输出到标准输出 os.Stdout。运行该程序，将输出定向到文件 cpu.pprof 中
	// setp1: go run main.go > cpu.pprof
	// pprof.StartCPUProfile(os.Stdout)
	// defer pprof.StopCPUProfile()

	// 如果程序本身有输出，则会相互干扰，直接记录到一个文件中
	// step1: go run main.go
	// f, _ := os.OpenFile("cpu.pprof", os.O_CREATE|os.O_RDWR, 0644)
	// defer f.Close()
	// pprof.StartCPUProfile(f)
	// defer pprof.StopCPUProfile()

	// step2: go tool pprof -http=:9999 cpu.pprof
	// 除了在网页中查看分析数据外，我们也可以在命令行中使用交互模式查看
	// step2: go tool pprof cpu.pprof 进入debug，然后 top 最后 exit退出
	/*
		还可以按照 cum (累计消耗)排序
		$ go tool pprof cpu.pprof
		(pprof) top --cum
		Showing nodes accounting for 14.14s, 99.16% of 14.26s total
		Dropped 34 nodes (cum <= 0.07s)
			flat  flat%   sum%        cum   cum%
			14.14s 99.16% 99.16%     14.17s 99.37%  main.bubbleSort
				0     0% 99.16%     14.17s 99.37%  main.main
				0     0% 99.16%     14.17s 99.37%  runtime.main
	*/

	n := 10
	for i := 0; i < 5; i++ {
		nums := generate(n)
		bubbleSort(nums)
		n *= 10
	}

	///////////////////

	// defer profile.Start().Stop()                                              // cpu
	defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop() // mem
	concat(100)
}
