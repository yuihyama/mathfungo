package add

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	var tests = []struct {
		x    float64
		y    float64
		want float64
	}{
		{1, 1, 2},
		{1.0, 1.0, 2},
		{1.1, 1.1, 2.2},
		{2.1, 2.1, 4.2},
		{-1, -1, -2},
		{-1.0, -1.0, -2.0},
	}
	for _, test := range tests {
		if got := Add(test.x, test.y); got != test.want {
			t.Errorf("got: Add(%v, %v) = %v, want: %v", test.x, test.y, got, test.want)
		}
	}
}

func ExampleAdd() {
	fmt.Println(Add(1, 1))
	// Output: 2
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1.1, 1.1)
	}
}
