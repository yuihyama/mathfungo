package intseq

import (
	"fmt"
	"testing"
)

func ExampleIntSeq() {
	fmt.Println(IntSeq(0, 0, 0))
	fmt.Println(IntSeq(0, 1, 1))
	fmt.Println(IntSeq(0, 2, 1))
	fmt.Println(IntSeq(0, 10, 1))
	fmt.Println(IntSeq(-3, 10, 1))
	fmt.Println(IntSeq(0, 10, 2))
	fmt.Println(IntSeq(0, 10, 0))
	// Output:
	// 0
	// 0 1
	// 0 1 2
	// 0 1 2 3 4 5 6 7 8 9 10
	// -3 -2 -1 0 1 2 3 4 5 6 7 8 9 10
	// 0 2 4 6 8 10
	// 0
}

func BenchmarkIntSeq(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IntSeq(0, 10, 1)
	}
}
