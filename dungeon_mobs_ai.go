package main

import (
	"diabloidrl/static"
)

func (d *dungeon) aiActForPawn(p *pawn) {
	if d.isTileInPlayerFOV(p.x, p.y) {
		p.mob.aiState = mobStateAttacking
		p.mob.AiStateTimeout = 10
	} else if p.mob.AiStateTimeout > 0 {
		p.mob.AiStateTimeout--
		if p.mob.AiStateTimeout == 0 && p.mob.aiState != mobStateIdle {
			p.mob.aiState = mobStateIdle
			p.action.reset()
		}
	}
	switch p.mob.aiState {
	case mobStateIdle:
		return
	case mobStateAttacking:
		selectedAttack := d.selectMobAttackWhichLands(p)
		if selectedAttack != nil {
			targetX, targetY := selectedAttack.Pattern.GetAimCoords(p, player)
			p.action.set(pActionAttack, p.getHitTime(), 0, targetX, targetY)
			p.action.attackData = selectedAttack
			return
		}
		vx, vy := d.getStepForPawnToPawn(p, player)
		if vx == 0 && vy == 0 {
			p.action.set(pActionWait, 0, ticksInTurn, vx, vy)
			return
		}
		p.action.set(pActionMove, 0, p.getMovementTime(), vx, vy)
	}
}

func (d *dungeon) selectMobAttackWhichLands(m *pawn) *static.Attack {
	var candidates []int
	for i := range m.mob.stats.Attacks {
		if m.mob.stats.Attacks[i].Pattern.CanBePerformedOn(m, player) {
			candidates = append(candidates, i)
		}
	}
	if len(candidates) == 0 {
		return nil
	}
	return m.mob.stats.Attacks[candidates[rnd.Rand(len(candidates))]]
}
