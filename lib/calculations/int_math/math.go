package intmath

func DivideRounding(dividing, divisor int) int {
	if dividing < 0 || divisor < 0 {
		panic("This works only on positives")
	}
	return (dividing + divisor/2) / divisor
}

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
