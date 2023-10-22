package main

import (
	"diabloidrl/lib/calculations"
	"math"
)

func (pc *playerController) doAutoAttackTurn(dung *dungeon) {
	// find closest one
	distanceToTarget := math.MaxUint32
	mobs := dung.getAllPawnsInPlayerFOV(true)
	var selectedMob *pawn
	for _, m := range mobs {
		dist := calculations.GetApproxDistFromTo(player.x, player.y, m.x, m.y)
		if dist < distanceToTarget {
			distanceToTarget = dist
			selectedMob = m
		}
	}

	pc.mode = pcModeDefault
	if selectedMob == nil {
		return
	}

	if pc.getAttackPattern().Pattern.CanBePerformedOn(player, selectedMob) {
		player.action.setAttack(player, pc.getAttackPattern(), 0, player.getHitTime(), selectedMob)
	} else {
		vx, vy := dung.getStepForPawnToPawn(player, selectedMob)
		player.action.set(pActionMove, 0, player.getMovementTime(), vx, vy)
	}
	// }
}
