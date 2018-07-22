package lt

import (
	"fmt"
	"testing"
)

func TestLT(t *testing.T) {
	var tests = []struct {
		x, y float64
		want bool
	}{
		{0, 0, false},
		{1, 2, true},
		{1.0, 2.0, true},
		{-1, -2, false},
		{-1.0, -2.0, false},
		{-1.0, 2, true},
	}
	for _, test := range tests {
		if got := LT(test.x, test.y); got != test.want {
			t.Errorf("got: LT(%v, %v) = %v, want: %v", test.x, test.y, got, test.want)
		}
	}
}

func ExampleLT() {
	fmt.Println(LT(0, 0))
	fmt.Println(LT(1, 2))
	fmt.Println(LT(1.0, 2.0))
	fmt.Println(LT(-1, -2))
	fmt.Println(LT(-1.0, -2.0))
	fmt.Println(LT(-1.0, 2))
	// Output:
	// false
	// true
	// true
	// false
	// false
	// true
}

func BenchmarkLT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LT(1, 2)
	}
}
