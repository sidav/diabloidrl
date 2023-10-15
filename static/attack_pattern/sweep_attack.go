package attackpattern

import (
	"diabloidrl/lib/calculations"
	// "diabloidrl/lib/calculations/primitives"
)

// import "diabloidrl/lib/calculations"

type SweepAttack struct {
	RadiusFromAttacker int
	RadiusFromTarget   int
}

func (sa *SweepAttack) CanBePerformedOn(attacker, target ActorForPattern) bool {
	return areActorsInTaxicabRange(attacker, target, sa.RadiusFromAttacker)
}

func (sa *SweepAttack) GetAttackedCoords(attacker ActorForPattern, targetX, targetY int) [][2]int {
	attackerX, attackerY := attacker.GetCoords()
	attackerW := attacker.GetSize()

	coords := [][2]int{}
	for x := targetX - sa.RadiusFromTarget; x < targetX+sa.RadiusFromTarget; x++ {
		for y := targetY - sa.RadiusFromTarget; y < targetY+sa.RadiusFromTarget; y++ {
			if calculations.AreCoordsInTileRect(x, y, attackerX, attackerY, attackerW, attackerW) {
				continue
			}
			if calculations.DistanceBetweenSquares(x, y, 1, attackerX, attackerY, attackerW) > sa.RadiusFromAttacker {
				continue
			}
			if calculations.SquareDistanceInt(x, y, targetX, targetY) > sa.RadiusFromTarget*sa.RadiusFromTarget {
				continue
			}
			coords = append(coords, [2]int{x, y})
		}
	}
	return coords
}

func (sa *SweepAttack) GetAimAt(attacker, target ActorForPattern) (int, int) {
	attackerX, attackerY := attacker.GetCoords()
	attackerW := attacker.GetSize()
	targetX, targetY := target.GetCoords()
	targetW := target.GetSize()

	// Searching
	targetCenterX, targetCenterY := targetX+targetW/2, targetY+targetW/2
	// return targetCenterX, targetCenterY
	return findBestCoordsOnPerimeterByScore(attackerX-sa.RadiusFromAttacker, attackerY-sa.RadiusFromAttacker,
		attackerX+attackerW+sa.RadiusFromAttacker, attackerY+attackerW+sa.RadiusFromAttacker,
		func(x, y int) int {
			// acx, acy := x+attackerW/2, y+attackerW/2
			return -calculations.SquareDistanceInt(targetCenterX, targetCenterY, x, y)
		},
	)
}
