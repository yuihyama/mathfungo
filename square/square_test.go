package square

import (
	"fmt"
	"testing"
)

func TestSquare(t *testing.T) {
	var tests = []struct {
		x, want float64
	}{
		{0, 0},
		{1, 1},
		{2, 4},
		{10, 100},
		{-1, 1},
		{-2, 4},
		{0.1, 0.010000000000000002},
	}
	for _, test := range tests {
		if got := Square(test.x); got != test.want {
			t.Errorf("got: Square(%v) = %v, want: %v", test.x, got, test.want)
		}
	}
}

func ExampleSquare() {
	fmt.Println(Square(0))
	fmt.Println(Square(1))
	fmt.Println(Square(2))
	fmt.Println(Square(-2))
	// Output:
	// 0
	// 1
	// 4
	// 4
}

func BenchmarkSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Square(2)
	}
}
