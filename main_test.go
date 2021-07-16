package main

import (
	"fmt"
	"testing"
)

// func TestJoinStrUseSprint(t *testing.T) {
// 	type args struct {
// 		a string
// 		b string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want string
// 	}{
// 		// TODO: Add test cases.
// 		{name: "1", args: args{"aaa", "bbb"}, want: "aaabbb"},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := JoinStrUseSprint(tt.args.a, tt.args.b); got != tt.want {
// 				t.Errorf("JoinStrUseSprint() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestJoinStrUseNor(t *testing.T) {
// 	type args struct {
// 		a string
// 		b string
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want string
// 	}{
// 		// TODO: Add test cases.
// 		{name: "1", args: args{"aaa", "bbb"}, want: "aaabbb"},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := JoinStrUseNor(tt.args.a, tt.args.b); got != tt.want {
// 				t.Errorf("JoinStrUseNor() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestJoinStrUseNor(t *testing.T) {
	s := JoinStrUseNor("aaa", "bbb")
	t.Log(s)
}

func TestJoinStrUseSprint(t *testing.T) {
	s := JoinStrUseSprint("aaa", "bbb")
	t.Log(s)
}

func BenchmarkJoinStrUseNor(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		JoinStrUseNor("aaa", "bbb")
	}

}

func BenchmarkJoinStrUseSprint(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// JoinStrUseSprint("aaa", "bbb")
		type args struct {
			a string
			b string
		}
		tests := []struct {
			name string
			args args
			want string
		}{
			// TODO: Add test cases.
			{name: "1", args: args{"aaa", "bbb"}, want: "aaabbb"},
			{name: "2", args: args{"ccc", "bbb"}, want: "cccbbb"},
			{name: "3", args: args{"ddd", "bbb"}, want: "dddbbb"},
		}
		for _, tt := range tests {
			b.Run(tt.name, func(t *testing.B) {
				if got := JoinStrUseSprint(tt.args.a, tt.args.b); got != tt.want {
					t.Errorf("JoinStrUseSprint() = %v, want %v", got, tt.want)
					t.Fail()
				}
			})
		}
	}
}

func Benchmark_Alloc(b *testing.B) {

	for i := 0; i < b.N; i++ {
		fmt.Println(fmt.Sprintf("%d", i))
	}
}
