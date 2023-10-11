package main

import "diabloidrl/lib/calculations"

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
		if p.getAttackRange() > 1 && d.isTileInPlayerFOV(p.x, p.y) &&
			p.getAttackRange() >= calculations.GetApproxDistFromTo(p.x, p.y, player.x, player.y) {
			p.action.set(pActionWait, 0, ticksInTurn, 0, 0)
			// d.doRangedAttack(p, player)
		} else {
			if d.arePawnsTouching(p, player) {
				p.action.set(pActionBasicMeleeAttack, p.getHitTime(), 0, player.x, player.y)
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
}
