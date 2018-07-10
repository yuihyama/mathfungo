package add

import (
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
