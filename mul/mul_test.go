package mul

import (
	"fmt"
	"testing"
)

func TestMul(t *testing.T) {
	var tests = []struct {
		x, y, want float64
	}{
		{0, 0, 0},
		{1, 1, 1},
		{2, 1, 2},
		{1, 2, 2},
		{1.0, 1.0, 1.0},
		{2.1, 1, 2.1},
		{-1, -1, 1},
		{-1, 1, -1},
		{-2.2, -2.1, 4.620000000000001},
	}
	for _, test := range tests {
		if got := Mul(test.x, test.y); got != test.want {
			t.Errorf("got: Mul(%v, %v) = %v, want: %v", test.x, test.y, got, test.want)
		}
	}
}

func ExampleMul() {
	fmt.Println(Mul(2, 1))
	// Output: 2
}

func BenchmarkMul(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Mul(2, 1)
	}
}
