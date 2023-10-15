package attackpattern

type RoundAttack struct {
	Size int
}

func (r *RoundAttack) CanBePerformedOn(attacker, target ActorForPattern) bool {
	return areActorsInTaxicabRange(attacker, target, r.Size)
}

func (r *RoundAttack) GetAttackedCoords(attacker ActorForPattern, targetX, targetY int) [][2]int {
	attackerX, attackerY := attacker.GetCoords()
	attackerW := attacker.GetSize()

	coords := [][2]int{}
	for x := attackerX - r.Size; x < attackerX+attackerW+r.Size; x++ {
		for y := attackerY - r.Size; y < attackerY+attackerW+r.Size; y++ {
			if !areCoordsInRect(x, y, attackerX, attackerY, attackerW, attackerW) {
				coords = append(coords, [2]int{x, y})
			}
		}
	}
	return coords
}

func (r *RoundAttack) GetAimAt(attacker, target ActorForPattern) (int, int) {
	return 0, 0 // it doesn't matter
}
