package static

import "strings"

type affixAdder struct {
	affixAdjective                   string // both needed if it's affix "of ..."
	affixName                        string // for escaping "item of X of Y" case (for suffixes only)
	incompatibleWithAffixOfAdjective string // I couldn't figure out a better way for this :(
	// for armor:
	armorFunc          func(*ArmorStats)
	bodyOnly, headOnly bool
	// for weapons:
	weaponFunc            func(*WeaponStats)
	meleeOnly, rangedOnly bool
	// for amulets:
	// TODO
	// for flasks:
	flaskFunc func(*FlaskStats)
	// universal:
	anyFunc func(ItemStatsInterface)
	// TODO (add item stats interface)
}

func (a *affixAdder) isCompatibleWith(a2 *affixAdder) bool {
	return strings.ToLower(a.incompatibleWithAffixOfAdjective) != strings.ToLower(a2.affixAdjective) &&
		strings.ToLower(a2.incompatibleWithAffixOfAdjective) != strings.ToLower(a.affixAdjective)
}

func addAffixesNamesToItemName(affixes []*affixAdder, name string) string {
	prefString := ""
	suffString := ""
	hasSuffix := false
	for _, a := range affixes {
		if a.affixName != "" {
			if hasSuffix {
				suffString = a.affixAdjective + " " + suffString
			} else {
				suffString = a.affixName
			}
			hasSuffix = true
		} else {
			prefString = a.affixAdjective + " " + prefString
		}
	}
	if len(prefString) > 0 {
		name = prefString + name
	}
	if len(suffString) > 0 {
		name = name + " of " + suffString
	}
	return name
}

var allAffixes = []*affixAdder{
	// ARMORS
	// bad affixes
	{
		affixAdjective:                   "Weared",
		incompatibleWithAffixOfAdjective: "Sturdy",
		armorFunc: func(stats *ArmorStats) {
			perc := rnd.RandInRange(15, 50)
			stats.Defense = intPercentage(stats.Defense, 100-perc)
			stats.addAffixDescription("Defense -%d%%", perc)
		},
	},
	{
		affixAdjective:                   "Cumbersome",
		incompatibleWithAffixOfAdjective: "cosy",
		armorFunc: func(stats *ArmorStats) {
			mod := rnd.RandInRange(1, 2)
			stats.EvasionModifier -= mod
			stats.addAffixDescription("-%d to evasion rolls", mod)
		},
	},
	// good affixes
	{
		affixAdjective: "Sturdy",
		armorFunc: func(stats *ArmorStats) {
			perc := rnd.RandInRange(15, 50)
			stats.Defense = intPercentage(stats.Defense, 100+perc)
			stats.addAffixDescription("Defense +%d%%", perc)
		},
	},
	{
		affixAdjective: "Cosy",
		armorFunc: func(stats *ArmorStats) {
			mod := rnd.RandInRange(1, 2)
			stats.EvasionModifier += mod
			stats.addAffixDescription("+%d to evasion rolls", mod)
		},
	},

	// WEAPONS
	// bad affixes
	{
		affixAdjective:                   "Cracked",
		incompatibleWithAffixOfAdjective: "sharp",
		weaponFunc: func(stats *WeaponStats) {
			mod := rnd.RandInRange(1, stats.DamageDice.Number)
			stats.DamageDice.Modifier -= mod
			stats.addAffixDescription("-%d to damage rolls", mod)
		},
	},
	{
		affixAdjective:                   "Crude",
		incompatibleWithAffixOfAdjective: "swift",
		weaponFunc: func(stats *WeaponStats) {
			perc := rnd.RandInRange(10, 50)
			stats.Delay = intPercentage(stats.Delay, 100+perc)
			stats.addAffixDescription("%d%% slower attack", perc)
		},
	},
	// good affixes
	{
		affixAdjective: "Sharp",
		weaponFunc: func(stats *WeaponStats) {
			mod := rnd.RandInRange(1, 2)
			stats.DamageDice.Modifier += mod
			stats.addAffixDescription("+%d to damage rolls", mod)
		},
	},
	{
		affixAdjective: "Balanced",
		weaponFunc: func(stats *WeaponStats) {
			mod := rnd.RandInRange(1, 2)
			stats.ToHitDice.Modifier += mod
			stats.addAffixDescription("+%d to to-hit rolls", mod)
		},
	},
	{
		affixAdjective: "Swift",
		weaponFunc: func(stats *WeaponStats) {
			perc := rnd.RandInRange(15, 40)
			stats.Delay = intPercentage(stats.Delay, 100-perc)
			stats.addAffixDescription("%d%% faster attack", perc)
		},
	},
	{
		affixAdjective: "chaotic",
		affixName:      "chaos",
		meleeOnly:      true,
		weaponFunc: func(stats *WeaponStats) {
			maxDmg := stats.DamageDice.GetMaximumPossible()
			selectedDmg := intPercentage(maxDmg, rnd.RandInRange(100, 150))
			stats.DamageDice.Alter(1, selectedDmg, -1)
			stats.addAffixDescription("Greater damage dispersion")
		},
	},

	// Universal affixes
	// Bad
	{
		affixAdjective:                   "corrupted",
		affixName:                        "corruption",
		incompatibleWithAffixOfAdjective: "healing",
		anyFunc: func(stats ItemStatsInterface) {
			ego := &ego{
				Code:  EgoCodeRegenerationPeriodPercent,
				Value: 110,
			}
			stats.addEgo(ego)
			stats.addAffixDescription("10%% slower regeneration")
		},
	},
	{
		affixAdjective:                   "Misguided",
		incompatibleWithAffixOfAdjective: "critical",
		anyFunc: func(stats ItemStatsInterface) {
			stats.addEgo(&ego{
				Code:  EgoCodeAdditionalCritChancePercent,
				Value: -5,
			})
			stats.addAffixDescription("-5%% to crit chance")
		},
	},
	{
		affixAdjective:                   "Dark",
		incompatibleWithAffixOfAdjective: "shining",
		anyFunc: func(stats ItemStatsInterface) {
			val := rnd.RandInRange(1, 2)
			stats.addEgo(&ego{
				Code:  EgoCodeAdditionalLightRadius,
				Value: -val,
			})
			stats.addAffixDescription("-%d to light radius", val)
		},
	},
	{
		affixAdjective:                   "crippling",
		affixName:                        "cripple",
		incompatibleWithAffixOfAdjective: "surviving",
		anyFunc: func(stats ItemStatsInterface) {
			val := rnd.RandInRange(1, 10)
			stats.addEgo(&ego{
				Code:  EgoCodeAdditionalMaxHP,
				Value: -val,
			})
			stats.addAffixDescription("-%d to maximum health", val)
		},
	},
	// Good
	{
		affixAdjective: "surviving",
		affixName:      "survival",
		anyFunc: func(stats ItemStatsInterface) {
			val := rnd.RandInRange(1, 15)
			stats.addEgo(&ego{
				Code:  EgoCodeAdditionalMaxHP,
				Value: val,
			})
			stats.addAffixDescription("+%d to maximum health", val)
		},
	},
	{
		affixAdjective: "Shining",
		anyFunc: func(stats ItemStatsInterface) {
			val := rnd.RandInRange(1, 2)
			stats.addEgo(&ego{
				Code:  EgoCodeAdditionalLightRadius,
				Value: val,
			})
			stats.addAffixDescription("+%d to light radius", val)
		},
	},
	{
		affixAdjective: "healing",
		affixName:      "regeneration",
		anyFunc: func(stats ItemStatsInterface) {
			ego := &ego{
				Code:  EgoCodeRegenerationPeriodPercent,
				Value: 85,
			}
			stats.addEgo(ego)
			stats.addAffixDescription("15%% faster regeneration")
		},
	},
	{
		affixAdjective: "critical",
		affixName:      "precision",
		anyFunc: func(stats ItemStatsInterface) {
			stats.addEgo(&ego{
				Code:  EgoCodeAdditionalCritChancePercent,
				Value: 5,
			})
			stats.addAffixDescription("+5%% to crit chance")
		},
	},
	{
		affixAdjective: "severe",
		affixName:      "doom",
		anyFunc: func(stats ItemStatsInterface) {
			stats.addEgo(&ego{
				Code:  EgoCodeAdditionalCritDamagePercent,
				Value: 15,
			})
			stats.addAffixDescription("+15%% to crit damage")
		},
	},
	{
		affixAdjective: "spiked",
		affixName:      "spikes",
		anyFunc: func(stats ItemStatsInterface) {
			perc := rnd.RandInRange(1, 10)
			stats.addEgo(&ego{
				Code:  EgoCodeThornsChance,
				Value: perc,
			})
			stats.addAffixDescription("%d%% chance to return 25%% of melee damage", perc)
		},
	},
}
