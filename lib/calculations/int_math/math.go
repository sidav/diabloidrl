package intmath

func IntAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
