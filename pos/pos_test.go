package pos

import (
	"fmt"
	"testing"
)

func TestPos(t *testing.T) {
	var tests = []struct {
		x, want float64
	}{
		{0, 0},
		{-0, 0},
		{1, 1},
		{1.1, 1.1},
		{-1, 1},
		{-1.1, 1.1},
		{0.01, 0.01},
		{-0.01, 0.01},
	}
	for _, test := range tests {
		if got := Pos(test.x); got != test.want {
			t.Errorf("got: Pos(%v) = %v, want: %v", test.x, got, test.want)
		}
	}
}

func ExamplePos() {
	fmt.Println(Pos(0))
	fmt.Println(Pos(-0))
	fmt.Println(Pos(1))
	fmt.Println(Pos(1.1))
	fmt.Println(Pos(-1))
	fmt.Println(Pos(-1.1))
	fmt.Println(Pos(-0.01))
	// Output:
	// 0
	// 0
	// 1
	// 1.1
	// 1
	// 1.1
	// 0.01
}

func BenchmarkPos(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Pos(-1)
	}
}
