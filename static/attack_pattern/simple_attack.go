package attackpattern

import (
	"diabloidrl/lib/calculations"
	// "diabloidrl/lib/calculations/primitives"
)

// import "diabloidrl/lib/calculations"

type SimpleAttack struct {
	// Simple attack is just an adjacent square equal to attacker's size.
}

func (SimpleAttack) CanBePerformedOn(attacker, target ActorForPattern) bool {
	attackerX, attackerY := attacker.GetCoords()
	attackerW := attacker.GetSize()
	targetX, targetY := target.GetCoords()
	targetW := target.GetSize()

	return calculations.AreRectsInTaxicabRange(attackerX, attackerY, attackerW, attackerW,
		targetX, targetY, targetW, targetW, attackerW)
}

func (SimpleAttack) GetAttackedCoords(attacker ActorForPattern, targetX, targetY int) [][2]int {
	// attackerX, attackerY := attacker.GetCoords()
	attackerW := attacker.GetSize()

	coords := [][2]int{}
	for x := targetX; x < targetX+attackerW; x++ {
		for y := targetY; y < targetY+attackerW; y++ {
			coords = append(coords, [2]int{x, y})
		}
	}
	return coords
}

func (SimpleAttack) GetAimAt(attacker, target ActorForPattern) (int, int) {
	attackerX, attackerY := attacker.GetCoords()
	attackerW := attacker.GetSize()
	targetX, targetY := target.GetCoords()
	targetW := target.GetSize()

	// var line []primitives.Point
	// if attackerW%2 == 1 {
	// 	line = primitives.GetLineOfLength(attackerX, attackerY, targetX, targetY, attackerW+1)
	// } else {
	// 	panic("Unimplemented")
	// }
	// last := line[len(line)-1]
	// return last.X, last.Y

	// Searching
	bestX, bestY := attackerX-attackerW, attackerY-attackerW
	bestDist := 10000
	targetCenterX, targetCenterY := targetX+targetW/2, targetY+targetW/2
	for x := attackerX - attackerW; x <= attackerX+attackerW; x++ {
		for y := attackerY - attackerW; y <= attackerY+attackerW; y++ {
			if calculations.AreTwoCellRectsOverlapping(x, y, attackerW, attackerW, attackerX, attackerY, attackerW, attackerW) {
				continue
			}
			acx, acy := x+attackerW/2, y+attackerW/2
			dist := calculations.SquareDistanceInt(targetCenterX, targetCenterY, acx, acy)
			if dist < bestDist {
				bestDist = dist
				bestX, bestY = x, y
			}
		}
	}
	return bestX, bestY
}
