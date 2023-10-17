package main

import (
	"diabloidrl/lib/calculations"
	"diabloidrl/static"
	"strconv"
)

const (
	pActionWait = iota
	pActionMove
	pActionAttack
)

type PawnAction struct {
	code                                int
	x, y                                int // target
	wasDone                             bool
	ticksBeforeAction, ticksAfterAction int
	attackData                          *static.AttackSkill
}

func (pa *PawnAction) set(code, delBefore, delAfter, vx, vy int) {
	pa.code = code
	pa.ticksBeforeAction = delBefore
	pa.ticksAfterAction = delAfter
	pa.x, pa.y = vx, vy
	pa.attackData = nil
	pa.wasDone = false
}

func (pa *PawnAction) setAttack(attacker *pawn, adata *static.AttackSkill, delBefore, delAfter int, targetPawn *pawn) {
	pa.code = pActionAttack
	pa.attackData = adata
	pa.ticksBeforeAction = calculations.IntPercentage(delBefore, adata.HitTimePercentage)
	pa.ticksAfterAction = calculations.IntPercentage(delAfter, adata.HitTimePercentage)
	pa.x, pa.y = pa.attackData.Pattern.GetAimAt(attacker, targetPawn)
	log.AppendMessagef("ATTACK SELECTED: %d, %d; at %+v", pa.x, pa.y, pa.attackData.Pattern)
	pa.wasDone = false
}

func (pa *PawnAction) reset() {
	pa.set(pActionWait, 0, 10, 0, 0)
}

func (pa *PawnAction) setVector(vx, vy int) {
	pa.x, pa.y = vx, vy
}

func (pa *PawnAction) getCoords() (int, int) {
	return pa.x, pa.y
}

func (pa *PawnAction) updateDelays() {
	if pa.code != pActionWait && pa.ticksBeforeAction == 0 && pa.ticksAfterAction == 0 {
		// just a failsafe to double-check the logic
		if false {
			panic("Non-updated action delay for code " + strconv.Itoa(pa.code))
		}
	}
	if pa.ticksBeforeAction > 0 {
		pa.ticksBeforeAction--
	} else if pa.ticksAfterAction > 0 {
		pa.ticksAfterAction--
	}
}

func (pa *PawnAction) markExecuted() {
	pa.wasDone = true
}

func (pa *PawnAction) isEmpty() bool {
	return pa.ticksBeforeAction == 0 && pa.ticksAfterAction == 0
}

func (pa *PawnAction) canActionOccurNow() bool {
	return pa.ticksBeforeAction == 0 && !pa.wasDone
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
