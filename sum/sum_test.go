package sum

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	var tests = []struct {
		nums1, want float64
	}{
		{0, 0},
		{1, 1},
		{1.0, 1.0},
		{-1, -1},
		{-1.0, -1.0},
	}
	for _, test := range tests {
		if got1 := Sum(test.nums1); got1 != test.want {
			t.Errorf("got: Sum(%v) = %v, want: %v", test.nums1, got1, test.want)
		}
	}

	nums2 := []float64{1.0, 2.0, 3.0}
	want2 := 6.0
	if got2 := Sum(nums2...); got2 != want2 {
		t.Errorf("got: Sum(%v) = %v, want: %v", nums2, got2, want2)
	}

	nums3 := Sum(-1, 2, 3.4)
	want3 := 4.4
	if got3 := nums3; got3 != want3 {
		t.Errorf("got: Sum(%v) = %v, want: %v", nums3, got3, want3)
	}

}

func ExampleSum() {
	fmt.Println(Sum(0))
	fmt.Println(Sum(1))
	fmt.Println(Sum(1.0, 2.0, 3.0))
	fmt.Println(Sum(1, 2, 3))
	fmt.Println(Sum(-1, 2, 3.4))
	nums4 := []float64{1, 2, 3, 4.4}
	fmt.Println(Sum(nums4...))
	// Output:
	// 0
	// 1
	// 6
	// 6
	// 4.4
	// 10.4
}

func BenchmarkSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sum(1, 2, 3)
	}
}
