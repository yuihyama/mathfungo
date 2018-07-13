package div

import (
	"fmt"
	"testing"
)

func TestDiv(t *testing.T) {
	var tests = []struct {
		x, y, want float64
	}{
		{0, 1, 0},
		{1, 1, 1},
		{10, 2, 5},
		{1.0, 1.0, 1.0},
		{10.0, 2.0, 5.0},
		{-10, 2.0, -5},
	}
	for _, test := range tests {
		if got := Div(test.x, test.y); got != test.want {
			t.Errorf("got: Div(%v, %v) = %v, want: %v", test.x, test.y, got, test.want)
		}
	}
}

func ExampleDiv() {
	fmt.Println(Div(0, 0))
	fmt.Println(Div(0.0, 0.0))
	fmt.Println(Div(1, 0))
	fmt.Println(Div(1.0, 0.0))
	fmt.Println(Div(-1, 0))
	fmt.Println(Div(-1.0, 0))
	// Output:
	// NaN
	// NaN
	// +Inf
	// +Inf
	// -Inf
	// -Inf
}

func BenchmarkDiv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Div(10, 2)
	}
}
