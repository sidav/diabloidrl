package attackpattern

import (
	"diabloidrl/lib/calculations"
	"diabloidrl/lib/calculations/primitives"
)

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

func areCoordsOnLine(x, y, lfx, lfy, ltx, lty int) bool {
	line := primitives.GetLine(lfx, lfy, ltx, lty)
	for _, c := range line {
		if c.X == x && c.Y == y {
			return true
		}
	}
	return false
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

func areCoordsInRect(x, y, rx, ry, rw, rh int) bool {
	return calculations.AreCoordsInTileRect(x, y, rx, ry, rw, rh)
}

func areActorsInTaxicabRange(a1, a2 ActorForPattern, dist int) bool {
	attackerX, attackerY := a1.GetCoords()
	attackerW := a1.GetSize()
	targetX, targetY := a2.GetCoords()
	targetW := a2.GetSize()
	return calculations.AreRectsInTaxicabRange(attackerX, attackerY, attackerW, attackerW,
		targetX, targetY, targetW, targetW, dist)
}
