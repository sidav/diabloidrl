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
	for x := targetX - sa.RadiusFromTarget; x <= targetX+sa.RadiusFromTarget; x++ {
		for y := targetY - sa.RadiusFromTarget; y <= targetY+sa.RadiusFromTarget; y++ {
			if calculations.AreCoordsInTileRect(x, y, attackerX, attackerY, attackerW, attackerW) {
				continue
			}
			if !calculations.AreRectsInTaxicabRange(x, y, 1, 1, attackerX, attackerY, attackerW, attackerW, sa.RadiusFromAttacker) {
				continue
			}
			if calculations.SquareDistanceInt(x, y, targetX, targetY) > sa.RadiusFromTarget*sa.RadiusFromTarget {
				continue
			}
			// if !calculations.AreCoordsInRange(x, y, targetX, targetY, sa.RadiusFromTarget) {
			// 	continue
			// }
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
	rectLeft, rectRight := attackerX-sa.RadiusFromAttacker, attackerX+attackerW+sa.RadiusFromAttacker-1
	rectTop, rectBottom := attackerY-sa.RadiusFromAttacker, attackerY+attackerW+sa.RadiusFromAttacker-1
	// return targetCenterX, targetCenterY
	return findBestCoordsOnPerimeterByScore(rectLeft, rectTop, rectRight, rectBottom,
		func(x, y int) int {
			if x != rectLeft && x != rectRight && x != (rectRight+rectLeft+1)/2 {
				return -1000
			}
			if y != rectTop && y != rectBottom && y != (rectBottom+rectTop+1)/2 {
				return -1000
			}
			// return -calculations.TaxicabDistance(targetCenterX, targetCenterY, x, y)
			return -calculations.SquareDistanceInt(x, y, targetCenterX, targetCenterY)
		},
	)
}
