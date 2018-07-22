package constant

import (
	"fmt"
	"testing"
)

func TestConstant(t *testing.T) {
	var tests = []struct {
		x, want float64
	}{
		{0, 0},
		{-0, 0},
		{1, 1},
		{0.1, 0.1},
		{-1, -1},
		{-0.1, -0.1},
	}
	for _, test := range tests {
		if got := Constant(test.x); got != test.want {
			t.Errorf("got: Constant(%v) = %v, want: %v", test.x, got, test.want)
		}
	}
}

func ExampleConstant() {
	fmt.Println(Constant(0))
	fmt.Println(Constant(-0))
	fmt.Println(Constant(1))
	fmt.Println(Constant(0.1))
	fmt.Println(Constant(-1))
	fmt.Println(Constant(-0.1))
	// Output:
	// 0
	// 0
	// 1
	// 0.1
	// -1
	// -0.1
}

func BenchmarkConstant(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Constant(1)
	}
}
