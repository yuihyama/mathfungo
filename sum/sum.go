package sum

func Sum(nums ...float64) float64 {
	res := 0.0
	for _, num := range nums {
		res += num
	}
	return res
}
