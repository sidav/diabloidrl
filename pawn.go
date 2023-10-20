package main

import (
	intmath "diabloidrl/lib/calculations/int_math"
)

type pawn struct {
	x, y          int
	hitpoints     int
	stamina       int
	mob           *mobStruct
	playerStats   *playerStruct
	inv           *inventory
	action        PawnAction
	statusEffects []*statusEffect

	flaskCharges       int
	canUseFlaskInTicks int
}

func (p *pawn) GetCoords() (int, int) {
	return p.x, p.y
}

func (p *pawn) isPlayer() bool {
	return p.playerStats != nil
}

func (p *pawn) regainHitpoints(hp int) {
	p.hitpoints += hp
	if p.hitpoints > p.getMaxHitpoints() {
		p.hitpoints = p.getMaxHitpoints()
	}
}

func (p *pawn) regainStamina(stm int) {
	p.stamina = intmath.Min(p.getMaxStamina(), p.stamina+stm)
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
