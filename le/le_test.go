package le

import (
	"fmt"
	"testing"
)

func TestLE(t *testing.T) {
	var tests = []struct {
		x, y float64
		want bool
	}{
		{0, 0, true},
		{-0, -0, true},
		{0.0, 0.0, true},
		{1, 2, true},
		{2, 1, false},
		{2.1, 1.1, false},
		{-1, -1, true},
		{-1.1, -2.1, false},
	}
	for _, test := range tests {
		if got := LE(test.x, test.y); got != test.want {
			t.Errorf("got: LE(%v, %v) = %v, want: %v", test.x, test.y, got, test.want)
		}
	}
}

func ExampleLE() {
	fmt.Println(LE(0, 0))
	fmt.Println(LE(-0, -0))
	fmt.Println(LE(0.0, 0.0))
	fmt.Println(LE(1, 2))
	fmt.Println(LE(2, 1))
	fmt.Println(LE(2.1, 1.1))
	fmt.Println(LE(-1, -1))
	fmt.Println(LE(-1.1, -2.1))
	// Output:
	// true
	// true
	// true
	// true
	// false
	// false
	// true
	// false
}

func BenchmarkLE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LE(1, 2)
	}
}
