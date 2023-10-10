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

func (pa *PawnAction) canActionOccurNow() bool {
	return pa.ticksBeforeAction == 0 && !pa.actionDone
}

func (p *pawn) regainHitpoints(hp int) {
	p.hitpoints += hp
	if p.hitpoints > p.getMaxHitpoints() {
		p.hitpoints = p.getMaxHitpoints()
	}
}

func (p *pawn) acquireExperience(exp int) {
	if p.isPlayer() {
		levelBefore := p.playerStats.getExperienceLevel()
		p.playerStats.experience += exp
		if p.playerStats.getExperienceLevel() > levelBefore {
			p.playerStats.skillPoints++
			p.hitpoints = p.getMaxHitpoints()
		}
	}
}

func (p *pawn) spendTime(ticks int) {
	if p.isPlayer() {
		p.playerStats.lastActionTicks += ticks
	}
	p.canActInTicks += ticks
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
	p.spendTime(10)
}
