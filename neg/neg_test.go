package neg

import (
	"fmt"
	"testing"
)

func TestNeg(t *testing.T) {
	var tests = []struct {
		x, want float64
	}{
		{0, 0},
		{-0, 0},
		{1, -1},
		{1.1, -1.1},
		{-1, 1},
		{-1.1, 1.1},
		{0.01, -0.01},
		{-0.01, 0.01},
	}
	for _, test := range tests {
		if got := Neg(test.x); got != test.want {
			t.Errorf("got: Neg(%v) = %v, want: %v", test.x, got, test.want)
		}
	}
}

func ExampleNeg() {
	fmt.Println(Neg(0))
	fmt.Println(Neg(-0))
	fmt.Println(Neg(1))
	fmt.Println(Neg(1.1))
	fmt.Println(Neg(-1))
	fmt.Println(Neg(0.01))
	fmt.Println(Neg(-0.01))
	// Output:
	// 0
	// 0
	// -1
	// -1.1
	// 1
	// -0.01
	// 0.01
}

func BenchmarkNeg(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Neg(1)
	}
}
