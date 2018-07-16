package factorial

import (
	"fmt"
	"testing"
)

func TestFactorial(t *testing.T) {
	var tests = []struct {
		x, want float64
	}{
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
		{6, 720},
		{6.0, 720},
		{-1, 0},
		{-1.0, 0},
	}
	for _, test := range tests {
		if got := Factorial(test.x); got != test.want {
			t.Errorf("got: Factorial(%v) = %v, want: %v", test.x, got, test.want)
		}
	}
}

func ExampleFactorial() {
	fmt.Println(Factorial(0))
	fmt.Println(Factorial(1))
	fmt.Println(Factorial(2))
	fmt.Println(Factorial(3))
	fmt.Println(Factorial(4))
	fmt.Println(Factorial(5))
	fmt.Println(Factorial(6))
	fmt.Println(Factorial(6.0))
	fmt.Println(Factorial(-1))
	fmt.Println(Factorial(-1.0))
	// Output:
	// 1
	// 1
	// 2
	// 6
	// 24
	// 120
	// 720
	// 720
	// 0
	// 0
}

func BenchmarkFactorial(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Factorial(6)
	}
}
