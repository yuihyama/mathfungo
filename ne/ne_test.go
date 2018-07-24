package ne

import (
	"fmt"
	"testing"
)

func TestNE(t *testing.T) {
	var tests = []struct {
		x, y float64
		want bool
	}{
		{0, 0, false},
		{0.0, 0.0, false},
		{1, 2, true},
		{1, 1, false},
		{1.1, 2.1, true},
		{1.1, 1.1, false},
		{-1.1, 2.1, true},
		{-2.1, -1.1, true},
	}
	for _, test := range tests {
		if got := NE(test.x, test.y); got != test.want {
			t.Errorf("got: NE(%v, %v) = %v, want: %v", test.x, test.y, got, test.want)
		}
	}
}

func ExampleNE() {
	fmt.Println(NE(0, 0))
	fmt.Println(NE(0, 0))
	fmt.Println(NE(1, 2))
	fmt.Println(NE(1, 1))
	fmt.Println(NE(1.1, 2.1))
	fmt.Println(NE(1.1, 1.1))
	fmt.Println(NE(-1.1, 2.1))
	fmt.Println(NE(-2.1, -1.1))
	// Output:
	// false
	// false
	// true
	// false
	// true
	// false
	// true
	// true
}

func BenchmarkNE(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NE(1, 2)
	}
}
