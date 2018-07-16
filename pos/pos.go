package pos

func Pos(x float64) float64 {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
