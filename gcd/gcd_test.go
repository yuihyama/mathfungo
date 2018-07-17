package gcd

import (
	"fmt"
	"testing"
)

func TestGCD(t *testing.T) {
	var tests = []struct {
		x, y, want int64
	}{
		{2, 4, 2},
		{3, 4, 1},
		{3, 6, 3},
		{8, 12, 4},
	}
	for _, test := range tests {
		if got := GCD(test.x, test.y); got != test.want {
			t.Errorf("got: GCD(%v, %v) = %v, want: %v", test.x, test.y, got, test.want)
		}
	}
}

func ExampleGCD() {
	fmt.Println(GCD(2, 4))
	fmt.Println(GCD(3, 4))
	fmt.Println(GCD(3, 6))
	fmt.Println(GCD(8, 12))
	// Output:
	// 2
	// 1
	// 3
	// 4
}

func BenchmarkGCD(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GCD(8, 12)
	}
}
