package main

import "diabloidrl/static"

const (
	pActionWait = iota
	pActionAttack
)

type PawnAction struct {
	actionCode                          int
	actionDone                          bool
	ticksBeforeAction, ticksAfterAction int
}

func (pa *PawnAction) updateDelays() {
	if pa.ticksBeforeAction > 0 {
		pa.ticksBeforeAction--
	} else if pa.ticksAfterAction > 0 {
		pa.ticksAfterAction--
	}
}

func (pa *PawnAction) setWait(delay int) {
	pa.actionCode = pActionWait
	pa.ticksBeforeAction = 0
	pa.ticksAfterAction = delay
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
