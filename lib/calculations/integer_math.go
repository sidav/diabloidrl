package calculations

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func IntDivideRounding(dividing, divisor int) int {
	if dividing < 0 || divisor < 0 {
		panic("This works only on positives")
	}
	return (dividing + divisor/2) / divisor
}

func IntPercentage(whole, percent int) int {
	return IntDivideRounding(percent*whole, 100)
}
