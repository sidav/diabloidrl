package attackpattern

import (
	"diabloidrl/lib/calculations"
	"diabloidrl/lib/calculations/primitives"
)

// import "diabloidrl/lib/calculations"

type LineAttack struct {
	Size, Length int
}

func (l *LineAttack) CanBePerformedOn(attacker, target ActorForPattern) bool {
	attackerX, attackerY := attacker.GetCoords()
	attackerW := attacker.GetSize()
	targetX, targetY := target.GetCoords()
	targetW := target.GetSize()

	return calculations.AreRectsInTaxicabRange(attackerX, attackerY, attackerW, attackerW,
		targetX, targetY, targetW, targetW, l.Size*l.Length)
}

func (l *LineAttack) GetAttackedCoords(attacker ActorForPattern, targetX, targetY int) [][2]int {
	attackerX, attackerY := attacker.GetCoords()
	attackerW := attacker.GetSize()

	var line []primitives.Point
	line = primitives.GetLine(attackerX, attackerY, targetX, targetY)
	coords := [][2]int{}
	for i := 0; i < len(line); i += l.Size {
		lx, ly := line[i].X, line[i].Y
		for x := lx; x < lx+l.Size; x++ {
			for y := ly; y < ly+l.Size; y++ {
				if !calculations.AreCoordsInTileRect(x, y, attackerX, attackerY, attackerW, attackerW) {
					coords = append(coords, [2]int{x, y})
				}
			}
		}
	}
	return clearRepeatedCoords(coords)
}

func (l *LineAttack) GetAimAt(attacker, target ActorForPattern) (int, int) {
	attackerX, attackerY := attacker.GetCoords()
	attackerW := attacker.GetSize()
	targetX, targetY := target.GetCoords()
	targetW := target.GetSize()

	// Searching
	attackerCenterX, attackerCenterY := attackerX+attackerW/2, attackerY+attackerW/2
	targetCenterX, targetCenterY := targetX+targetW/2, targetY+targetW/2
	return findBestCoordsOnPerimeterByScore(attackerX-l.Length, attackerY-l.Length, attackerX+l.Length, attackerY+l.Length,
		func(x, y int) int {
			if calculations.AreTwoCellRectsOverlapping(x, y, l.Size, l.Size, attackerX, attackerY, attackerW, attackerW) ||
				!areCoordsOnLine(targetCenterX, targetCenterY, attackerCenterX, attackerCenterY, x, y) {
				return -1000000
			}
			acx, acy := x+l.Size/2, y+l.Size/2
			return -calculations.SquareDistanceInt(targetCenterX, targetCenterY, acx, acy)
		},
	)
}
