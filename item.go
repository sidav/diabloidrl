package main

import (
	"diabloidrl/lib/strings"
	"diabloidrl/static"
	"fmt"
)

type item struct {
	x, y     int
	asWeapon *static.WeaponStats
	asArmor  *static.ArmorStats
	asAmulet *static.AmuletStats
}

func (i *item) isWeapon() bool {
	return i.asWeapon != nil
}

func (i *item) isArmor() bool {
	return i.asArmor != nil
}

func (i *item) isAmulet() bool {
	return i.asAmulet != nil
}

func (i *item) getRarity() int {
	if i.isWeapon() {
		return i.asWeapon.Rarity
	}
	if i.isArmor() {
		return i.asArmor.Rarity
	}
	if i.isAmulet() {
		return i.asAmulet.Rarity
	}
	panic("No rarity")
}

func (i *item) getAscii() rune {
	if i.isWeapon() {
		if i.asWeapon.Range > 1 {
			return '{'
		}
		return '('
	}
	if i.isArmor() {
		if i.asArmor.Slot == static.ArmorSlotHead {
			return '^'
		}
		return '['
	}
	if i.isAmulet() {
		return '"'
	}
	return '?'
}

func (i *item) initAsRandomItem(minRarity int) {
	itemGenerators := []func(){
		func() {
			i.asWeapon = static.GenerateRandomWeaponStats(minRarity)
		},
		func() {
			i.asArmor = static.GenerateRandomArmorStats(minRarity)
		},
		func() {
			i.asAmulet = static.GenerateRandomAmuletStats(minRarity)
		},
	}
	itype := rnd.Rand(len(itemGenerators))
	itemGenerators[itype]()
}

func (i *item) getName() string {
	return i.getStatic().GetName()
}

func (i *item) getStatic() static.ItemStatsInterface {
	if i.isWeapon() {
		return i.asWeapon
	}
	if i.isArmor() {
		return i.asArmor
	}
	if i.isAmulet() {
		return i.asAmulet
	}
	panic("Item static retrieval failure")
}

func (i *item) getDescription() string {
	// if i == nil {
	// 	return "(Empty)"
	// }
	if i.isWeapon() {
		return fmt.Sprintf(
			"To-Hit %s, Damage %d-%d Delay %d",
			i.asWeapon.ToHitDice.GetShortDescriptionString(),
			i.asWeapon.DamageDice.GetMinimumPossible(), i.asWeapon.DamageDice.GetMaximumPossible(),
			i.asWeapon.Delay,
		)
	}
	if i.isArmor() {
		return fmt.Sprintf(
			"Armor class %d, evasion %s",
			i.asArmor.Defense,
			strings.StringifyIntegerAsModifier(i.asArmor.EvasionModifier),
		)
	}
	if i.isAmulet() {
		return "Amulet"
	}
	return "Undefined item"
}
