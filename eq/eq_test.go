package eq

import (
	"fmt"
	"testing"
)

func TestEQ(t *testing.T) {
	var tests = []struct {
		x, y float64
		want bool
	}{
		{0, 0, true},
		{0.0, 0.0, true},
		{1, 1, true},
		{-1, -1, true},
		{-1.1, -1.1, true},
		{0, 1, false},
		{0.0, 1.1, false},
		{-1, -2, false},
		{-1.1, -2.1, false},
	}
	for _, test := range tests {
		if got := EQ(test.x, test.y); got != test.want {
			t.Errorf("got: EQ(%v, %v) = %v, want: %v", test.x, test.y, got, test.want)
		}
	}
}

func ExampleEQ() {
	fmt.Println(EQ(0, 0))
	fmt.Println(EQ(0.0, 0.0))
	fmt.Println(EQ(1, 1))
	fmt.Println(EQ(-1, -1))
	fmt.Println(EQ(-1.1, -1.1))
	fmt.Println(EQ(0, 1))
	fmt.Println(EQ(0.0, 1.1))
	fmt.Println(EQ(-1, -2))
	fmt.Println(EQ(-1.1, -2.1))
	// Output:
	// true
	// true
	// true
	// true
	// true
	// false
	// false
	// false
	// false
}

func BenchmarkEQ(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EQ(1, 1)
	}
}
