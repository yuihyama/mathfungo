package sub

import (
	"fmt"
	"testing"
)

func TestSub(t *testing.T) {
	var tests = []struct {
		x, y, want float64
	}{
		{0, 0, 0},
		{1, 1, 0},
		{2, 1, 1},
		{1, 2, -1},
		{1.0, 1.0, 0},
		{2.1, 1.1, 1.0},
		{-1, -1, 0},
		{-1, 1, -2},
		{-2.2, -2.2, 0},
	}
	for _, test := range tests {
		if got := Sub(test.x, test.y); got != test.want {
			t.Errorf("got: Sub(%v, %v) = %v, want: %v", test.x, test.y, got, test.want)
		}
	}
}

func ExampleSub() {
	fmt.Println(Sub(2, 1))
	// Output: 1
}

func BenchmarkSub(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sub(2, 1)
	}
}
