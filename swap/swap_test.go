package swap

import (
	"fmt"
	"testing"
)

func ExampleSwap() {
	fmt.Println(Swap(1, 2))
	fmt.Println(Swap(1.1, 2.2))
	// Output:
	// 2 1
	// 2.2 1.1
}

func BenchmarkSwap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Swap(1, 2)
	}
}
