package main

import "diabloidrl/static"

type statusEffect struct {
	code              static.StatusEffectCode
	strength          int
	triggersEachTicks int
	remainingDuration int
}

func (p *pawn) receiveStatusEffect(se *statusEffect) {
	p.statusEffects = append(p.statusEffects, se)
}

func (p *pawn) cleanupStatusEffects() {
	for i := len(p.statusEffects) - 1; i >= 0; i-- {
		if p.statusEffects[i].remainingDuration <= 0 {
			p.statusEffects = append(p.statusEffects[:i], p.statusEffects[i+1:]...)
		}
	}
}

func (p *pawn) getStatusEffectByCode(code static.StatusEffectCode) *statusEffect {
	for _, se := range p.statusEffects {
		if se.code == code && se.remainingDuration > 0 {
			return se
		}
	}
	return nil
}

func (d *dungeon) applyPassiveStatusEffects(p *pawn) *statusEffect {
	for _, se := range p.statusEffects {
		se.remainingDuration--
		if se.remainingDuration%se.triggersEachTicks != 0 {
			continue
		}
		switch se.code {
		case static.StatusEffectCodeHealing:
			p.regainHitpoints(se.strength)
		}
	}
	return nil
}
