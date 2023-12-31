package main

import (
	"diabloidrl/lib/calculations"
	"diabloidrl/static"
)

func (d *dungeon) doMeleeHit(attacker, defender *pawn) {
	if attacker == defender {
		panic("Something is very wrong (hitting)")
	}
	toHitRoll := attacker.getHitDice().Roll(rnd)
	toEvadeRoll := rnd.Rand(defender.getEvasion())
	// log.AppendMessagef("To-hit %d <-> EV %d", toHitRoll, toEvadeRoll)
	renderer.addAnimationAt(animTypePawnIsActing, attacker.x, attacker.y, false)
	if toHitRoll < toEvadeRoll {
		log.AppendMessagef("%s evaded the attack!", defender.getName())
	} else {
		damage := attacker.getDamageDice().Roll(rnd)
		if rnd.Rand(100) < attacker.getCriticalChancePercent() {
			perc := attacker.getCriticalDamagePercent()
			log.AppendMessagef("Critical hit - %d%% damage!", perc)
			damage = damage * perc / 100
		}
		defenseRoll := rnd.Rand(defender.getArmorClass())
		if defenseRoll >= damage {
			log.AppendMessagef("%s's attack did no damage.", attacker.getName())
		} else {
			defender.hitpoints -= damage
			log.AppendMessagef("%s hit %s for %d damage.", attacker.getName(), defender.getName(), damage)
		}
		if defender.inv != nil && defender.inv.getSumOfEgoValuesOfEquippedItems(static.EgoCodeThornsChance) > 0 {
			thorns := defender.inv.getSumOfEgoValuesOfEquippedItems(static.EgoCodeThornsChance)
			if rnd.Rand(100) < thorns {
				thornDamage := damage / 4
				if thornDamage == 0 {
					thornDamage = 1
				}
				attacker.hitpoints -= thornDamage
				log.AppendMessagef("%s got %d damage from thorns!", attacker.getName(), thornDamage)
			}
		}
	}
	renderer.addAnimationAt(animTypeHit, defender.x, defender.y, true)
	attacker.spendTime(attacker.getHitTime())
}

func (d *dungeon) doRangedAttack(attacker, defender *pawn) {
	if attacker == defender {
		panic("Something is very wrong (shooting)")
	}
	if calculations.GetApproxDistFromTo(attacker.x, attacker.y, defender.x, defender.y) > attacker.getAttackRange() {
		panic("Range failure")
	}
	toHitRoll := attacker.getHitDice().Roll(rnd)
	toEvadeRoll := rnd.Rand(defender.getEvasion())

	renderer.addAnimationAt(animTypePawnIsActing, attacker.x, attacker.y, false)
	if toHitRoll < toEvadeRoll {
		log.AppendMessagef("%s evaded the shot!", defender.getName())
	} else {
		damage := attacker.getDamageDice().Roll(rnd)
		if rnd.Rand(100) < attacker.getCriticalChancePercent() {
			perc := attacker.getCriticalDamagePercent()
			log.AppendMessagef("Critical hit - %d%% damage!", perc)
			damage = damage * perc / 100
		}
		defenseRoll := rnd.Rand(defender.getArmorClass())
		if defenseRoll >= damage {
			log.AppendMessagef("%s's shot did no damage.", attacker.getName())
		} else {
			defender.hitpoints -= damage
			log.AppendMessagef("%s shot %s for %d damage.", attacker.getName(), defender.getName(), damage)
		}
	}
	renderer.addTwoCoordAnimationAt(animTypeHitscanProjectile, attacker.x, attacker.y, defender.x, defender.y, true)
	renderer.addAnimationAt(animTypeShot, defender.x, defender.y, false)
	attacker.spendTime(attacker.getHitTime())
}
