package main

import (
	"diabloidrl/static"
)

func (d *dungeon) aiActForPawn(p *pawn) {
	if d.isPawnInPlayerFov(p) {
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

		if d.isPawnInPlayerFov(p) && (d.arePawnsTouching(p, player) || rnd.Rand(3) != 0) {
			selectedAttack := d.selectMobAttackWhichLands(p)
			if selectedAttack != nil {
				p.action.setAttack(p, selectedAttack, p.getHitTime(), 0, player)
				return
			}
		}
		vx, vy := d.getStepForPawnToPawn(p, player)
		if vx == 0 && vy == 0 {
			p.action.set(pActionWait, 0, ticksInTurn, vx, vy)
			return
		}
		p.action.set(pActionMove, 0, p.getMovementTime(), vx, vy)
	}
}

func (d *dungeon) selectMobAttackWhichLands(m *pawn) *static.AttackSkill {
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
