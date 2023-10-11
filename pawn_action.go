package main

import "diabloidrl/static"

const (
	pActionWait = iota
	pActionMove
	pActionAttack
)

type PawnAction struct {
	actionCode                          int
	x, y                                int // target
	actionDone                          bool
	ticksBeforeAction, ticksAfterAction int
}

func (pa *PawnAction) set(code, delBefore, delAfter, vx, vy int) {
	pa.actionCode = code
	pa.ticksBeforeAction = delBefore
	pa.ticksAfterAction = delAfter
	pa.x, pa.y = vx, vy
	pa.actionDone = false
}

func (pa *PawnAction) setVector(vx, vy int) {
	pa.x, pa.y = vx, vy
}

func (pa *PawnAction) updateDelays() {
	if pa.ticksBeforeAction > 0 {
		pa.ticksBeforeAction--
	} else if pa.ticksAfterAction > 0 {
		pa.ticksAfterAction--
	}
}

func (pa *PawnAction) markExecuted() {
	pa.actionDone = true
}

func (pa *PawnAction) ended() bool {
	return pa.ticksBeforeAction == 0 && pa.ticksAfterAction == 0
}

func (pa *PawnAction) canActionOccurNow() bool {
	return pa.ticksBeforeAction == 0 && !pa.actionDone
}

func (p *pawn) useFlask() {
	if p.flaskCharges == 0 {
		// panic("Drinking from empty flask")
	}
	if p.canUseFlaskInTicks > 0 {
		panic("Drinking in cooldown")
	}
	flask := p.inv.getItemInSlot(invSlotFlask)
	if flask == nil {
		return
	}
	p.receiveStatusEffect(&statusEffect{
		code:              static.StatusEffectCodeHealing,
		strength:          flask.asFlask.HealStrength,
		triggersEachTicks: flask.asFlask.HealTicksPeriod,
		remainingDuration: flask.asFlask.HealEffectDuration,
	})
	p.flaskCharges--
	p.canUseFlaskInTicks += flask.asFlask.CooldownBetweenSips
}
