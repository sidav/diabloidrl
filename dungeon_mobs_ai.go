package main

import "diabloidrl/lib/calculations"

func (d *dungeon) actForPawn(p *pawn) {
	if d.isTileInPlayerFOV(p.x, p.y) {
		p.mob.currentState = mobStateAttacking
		p.mob.stateTimeout = 10
	} else if p.mob.stateTimeout > 0 {
		p.mob.stateTimeout--
		if p.mob.stateTimeout == 0 {
			p.mob.currentState = mobStateIdle
		}
	}
	switch p.mob.currentState {
	case mobStateIdle:
		return
	case mobStateAttacking:
		if p.getAttackRange() > 1 && d.isTileInPlayerFOV(p.x, p.y) &&
			p.getAttackRange() >= calculations.GetApproxDistFromTo(p.x, p.y, player.x, player.y) {

			d.doRangedAttack(p, player)
		} else {
			vx, vy := d.getStepForPawnToPawn(p, player)
			if vx == 0 && vy == 0 {
				p.spendTime(10)
				return
			}
			d.DefaultMoveActionWithPawn(p, vx, vy)
		}
	}
}
