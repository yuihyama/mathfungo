package cube

import (
	"fmt"
	"testing"
)

func TestCube(t *testing.T) {
	var tests = []struct {
		x, want float64
	}{
		{0, 0},
		{1, 1},
		{2, 8},
		{10, 1000},
		{-1, -1},
		{-2, -8},
		{0.1, 0.0010000000000000002},
		{-0.1, -0.0010000000000000002},
	}
	for _, test := range tests {
		if got := Cube(test.x); got != test.want {
			t.Errorf("got: Cube(%v) = %v, want: %v", test.x, got, test.want)
		}
	}
}

func ExampleCube() {
	fmt.Println(Cube(0))
	fmt.Println(Cube(1))
	fmt.Println(Cube(2))
	fmt.Println(Cube(10))
	fmt.Println(Cube(-1))
	fmt.Println(Cube(-2))
	fmt.Println(Cube(0.1))
	fmt.Println(Cube(-0.1))
	// Output: 0
	// 1
	// 8
	// 1000
	// -1
	// -8
	// 0.0010000000000000002
	// -0.0010000000000000002
}

func BenchmarkCube(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Cube(3)
	}
}
