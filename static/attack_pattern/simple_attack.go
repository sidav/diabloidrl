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
	return areActorsInTaxicabRange(attacker, target, attacker.GetSize())
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

	// Searching
	targetCenterX, targetCenterY := targetX+targetW/2, targetY+targetW/2
	return findBestCoordsOnPerimeterByScore(attackerX-attackerW, attackerY-attackerW, attackerX+attackerW, attackerY+attackerW,
		func(x, y int) int {
			if calculations.AreTwoCellRectsOverlapping(x, y, attackerW, attackerW, attackerX, attackerY, attackerW, attackerW) {
				return -1000000
			}
			acx, acy := x+attackerW/2, y+attackerW/2
			return -calculations.SquareDistanceInt(targetCenterX, targetCenterY, acx, acy)
		},
	)
}
