package attackpattern

// import "diabloidrl/lib/calculations/primitives"

func clearRepeatedCoords(arr [][2]int) [][2]int {
	for i := len(arr) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if arr[i] == arr[j] {
				arr = append(arr[:i], arr[i+1:]...)
				break
			}
		}
	}
	return arr
}

func findBestCoordsOnPerimeterByScore(leftx, topy, rightx, bottomy int, score func(x, y int) int) (int, int) {
	bestX, bestY := leftx, topy
	bestScore := score(bestX, bestY)
	for x := leftx; x <= rightx; x++ {
		y := topy
		newScore := score(x, y)
		if newScore > bestScore {
			bestScore = newScore
			bestX, bestY = x, y
		}
		y = bottomy
		newScore = score(x, y)
		if newScore > bestScore {
			bestScore = newScore
			bestX, bestY = x, y
		}
	}
	for y := topy; y <= bottomy; y++ {
		x := leftx
		newScore := score(x, y)
		if newScore > bestScore {
			bestScore = newScore
			bestX, bestY = x, y
		}
		x = rightx
		newScore = score(x, y)
		if newScore > bestScore {
			bestScore = newScore
			bestX, bestY = x, y
		}
	}
	return bestX, bestY
}
