package intseq

import "fmt"

func IntSeq(start, stop, step int) int {
	if step <= 0 {
		return 0
	}

	for i := start; i < stop; i += step {
		fmt.Print(i, " ")
	}

	return stop
}
