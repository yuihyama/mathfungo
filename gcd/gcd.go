package gcd

func GCD(x, y int64) int64 {
	if y == 0 {
		return x
	} else {
		return GCD(y, x%y)
	}
}
