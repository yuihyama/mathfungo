package add

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	expect1 := 2.0
	actual1 := Add(1.0, 1.0)

	if expect1 != actual1 {
		t.Errorf("%v != %v", expect1, actual1)
	}

	expect2 := 2.2
	actual2 := Add(1.1, 1.1)

	if expect2 != actual2 {
		t.Errorf("%v != %v", expect2, actual2)
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
