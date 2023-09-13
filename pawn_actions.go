package main

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
	p.regainHitpoints(flask.asFlask.EachSipHeals)
	p.flaskCharges--
	p.canUseFlaskInTicks += flask.asFlask.CooldownBetweenSips
	p.spendTime(10)
}
