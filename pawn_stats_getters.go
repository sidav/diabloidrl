package main

import (
	"diabloidrl/lib/random"
	"diabloidrl/static"
)

func (p *pawn) GetSize() int {
	if p.isPlayer() {
		return 1
	}
	if p.mob.stats.Size == 0 {
		return 1
	}
	return p.mob.stats.Size
}

func (p *pawn) getMaxHitpoints() int {
	maxHp := 0
	if p.isPlayer() {
		maxHp += p.playerStats.getStatsMaxHp()
	} else {
		maxHp += p.mob.stats.MaxHitpoints
	}
	if p.inv != nil {
		maxHp += p.inv.getSumOfEgoValuesOfEquippedItems(static.EgoCodeAdditionalMaxHP)
	}
	return maxHp
}

func (p *pawn) getMaxStamina() int {
	maxStm := 0
	if p.isPlayer() {
		maxStm += p.playerStats.getStatsMaxStm()
	} else {
		maxStm += 1000 // p.mob.stats.MaxHitpoints
	}
	if p.inv != nil {
		// maxStm += p.inv.getSumOfEgoValuesOfEquippedItems(static.EgoCodeAdditionalMaxHP)
	}
	return maxStm
}

func (p *pawn) getRegenCooldown() int {
	regenPeriod := 0
	if p.isPlayer() {
		// it takes this much ticks to fully regen player HP regardless of its max amount.
		const tickForFullRegen = 5000
		regenPeriod = tickForFullRegen / p.getMaxHitpoints()
	} else {
		regenPeriod = p.mob.stats.RegenCooldown
	}
	if p.inv != nil {
		perc := p.inv.getMultiplicativePercentOfEgoValuesOfEquippedItems(static.EgoCodeRegenerationPeriodPercent)
		// log.AppendMessagef("dbg: perc %d; period %d -> %d", perc, regenPeriod, regenPeriod*perc/100)
		regenPeriod = regenPeriod * perc / 100
	}
	return regenPeriod
}

func (p *pawn) getVisionRadius() int {
	radius := 5
	if p.inv != nil {
		radius += p.inv.getSumOfEgoValuesOfEquippedItems(static.EgoCodeAdditionalLightRadius)
	}
	return radius
}

func (p *pawn) getMovementTime() int {
	if p.isPlayer() {
		return p.playerStats.getStatsMovementTime()
	}
	return p.mob.stats.MovementTime
}

func (p *pawn) getHitTime() int {
	if p.inv.getItemInSlot(invSlotWeapon) != nil {
		return p.inv.getItemInSlot(invSlotWeapon).asWeapon.Delay
	}
	if p.isPlayer() {
		return 10
	}
	return p.mob.stats.HitTime
}

func (p *pawn) getName() string {
	if p.isPlayer() {
		return "you"
	}
	return p.mob.stats.Name
}

func (p *pawn) getHitDice() *random.Dice {
	if p.inv.getItemInSlot(invSlotWeapon) != nil {
		return &p.inv.getItemInSlot(invSlotWeapon).asWeapon.ToHitDice
	}
	if p.isPlayer() {
		return random.NewDice(1, 6, 0)
	}
	return &p.mob.stats.ToHit
}

func (p *pawn) getDamageDice() *random.Dice {
	if p.inv.getItemInSlot(invSlotWeapon) != nil {
		return &p.inv.getItemInSlot(invSlotWeapon).asWeapon.DamageDice
	}
	if p.isPlayer() {
		return random.NewDice(1, 2, 0)
	}
	return &p.mob.stats.Damage
}

func (p *pawn) getArmorClass() int {
	armorClass := 0
	if p.inv.getItemInSlot(invSlotBody) != nil {
		armorClass += p.inv.getItemInSlot(invSlotBody).asArmor.Defense
	}
	if p.inv.getItemInSlot(invSlotHelmet) != nil {
		armorClass += p.inv.getItemInSlot(invSlotHelmet).asArmor.Defense
	}
	return armorClass
}

func (p *pawn) getEvasion() int {
	ev := 0
	if p.isPlayer() {
		ev = 6
	} else {
		ev = p.mob.stats.Evasion
	}
	if p.inv != nil {
		if p.inv.getItemInSlot(invSlotBody) != nil {
			ev += p.inv.getItemInSlot(invSlotBody).asArmor.EvasionModifier
		}
		if p.inv.getItemInSlot(invSlotHelmet) != nil {
			ev += p.inv.getItemInSlot(invSlotHelmet).asArmor.EvasionModifier
		}
	}
	return ev
}

// func (p *pawn) getAttackRange() int {
// 	if p.inv != nil && p.inv.getItemInSlot(invSlotWeapon) != nil {
// 		return p.inv.getItemInSlot(invSlotWeapon).asWeapon.Range
// 	}
// 	if p.isPlayer() {
// 		return 0
// 	}
// 	return p.mob.stats.AttackRange
// }

func (p *pawn) getMaxFlaskCharges() int {
	if p.inv.getItemInSlot(invSlotFlask) != nil {
		return p.inv.getItemInSlot(invSlotFlask).asFlask.MaxCharges
	}
	return 0
}

func (p *pawn) getCriticalChancePercent() int {
	perc := 5
	if p.isPlayer() {
		// TODO
	} else {
		perc = p.mob.stats.CritChancePercent
	}
	if p.inv != nil {
		perc += p.inv.getSumOfEgoValuesOfEquippedItems(static.EgoCodeAdditionalCritChancePercent)
	}
	return perc
}

func (p *pawn) getCriticalDamagePercent() int {
	perc := 150
	if p.isPlayer() {
		// TODO
	} else {
		perc = p.mob.stats.CritDamagePercent
	}
	if p.inv != nil {
		perc += p.inv.getSumOfEgoValuesOfEquippedItems(static.EgoCodeAdditionalCritDamagePercent)
	}
	return perc
}
