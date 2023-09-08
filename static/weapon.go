package static

import (
	"diabloidrl/lib/random"
	"fmt"
)

type WeaponStats struct {
	Name       string
	ToHitDice  random.Dice
	DamageDice random.Dice
	Delay      int
	// Ranged     bool
	Range             int // 0 for melee
	Rarity            int
	Egos              []*ego
	affixDescriptions []string

	weightForSelection int
}

func (ws *WeaponStats) clone() *WeaponStats {
	newWeaponStats := *ws
	if ws == &newWeaponStats {
		panic("I must learn more")
	}
	return &newWeaponStats
}

func (ws *WeaponStats) GetName() string {
	return ws.Name
}

func (ws *WeaponStats) GetEgos() []*ego {
	return ws.Egos
}

func (ws *WeaponStats) GetAffixDescriptions() []string {
	return ws.affixDescriptions
}

func (ws *WeaponStats) addAffixDescription(s string, args ...interface{}) {
	ws.affixDescriptions = append(ws.affixDescriptions, fmt.Sprintf(s, args...))
}

func (ws *WeaponStats) addEgo(e *ego) {
	ws.Egos = append(ws.Egos, e)
}

func GenerateRandomWeaponStats(minRarity int) *WeaponStats {
	generators := []func() *WeaponStats{
		// common
		func() *WeaponStats {
			return getRandomWeaponStats(rnd)
		},
		// rare
		func() *WeaponStats {
			return makePrefixWeaponStatsFromBase(getRandomWeaponStats(rnd), 1)
		},
		// very rare (2 affixes)
		func() *WeaponStats {
			return makePrefixWeaponStatsFromBase(getRandomWeaponStats(rnd), 2)
		},
		// very rare (3 affixes)
		func() *WeaponStats {
			return makePrefixWeaponStatsFromBase(getRandomWeaponStats(rnd), 3)
		},
	}
	rarity := rnd.SelectRandomIndexFromWeighted(len(generators),
		func(i int) int {
			if i < minRarity {
				return 0
			}
			return rarityProbabilities[i]
		},
	)
	return generators[rarity]()
}

func getRandomWeaponStats(rnd random.PRNG) *WeaponStats {
	ind := rnd.SelectRandomIndexFromWeighted(len(sTableWeapons), func(i int) int { return sTableWeapons[i].weightForSelection })
	return sTableWeapons[ind]
}

func makePrefixWeaponStatsFromBase(baseStats *WeaponStats, affixCount int) *WeaponStats {
	stats := baseStats.clone()
	selectedAffixes := selectRandomAppropriateUniqueAffixesFor(stats, affixCount)
	for _, aff := range selectedAffixes {
		if aff.weaponFunc != nil {
			aff.weaponFunc(stats)
		}
		if aff.anyFunc != nil {
			aff.anyFunc(stats)
		}
	}
	stats.Name = addAffixesNamesToItemName(selectedAffixes, stats.Name)
	stats.Rarity = affixCount
	return stats
}

var sTableWeapons = []*WeaponStats{
	{
		Name:               "Dagger",
		ToHitDice:          *random.NewDice(2, 3, 0),
		DamageDice:         *random.NewDice(2, 3, 0),
		Delay:              10,
		weightForSelection: 10,
	},
	{
		Name:               "Short Blade",
		ToHitDice:          *random.NewDice(2, 3, 0),
		DamageDice:         *random.NewDice(2, 6, 0),
		Delay:              12,
		weightForSelection: 4,
	},
	{
		Name:               "Rapier",
		ToHitDice:          *random.NewDice(2, 3, 0),
		DamageDice:         *random.NewDice(3, 6, 0),
		Delay:              20,
		weightForSelection: 1,
	},
	{
		Name:               "Crossbow",
		ToHitDice:          *random.NewDice(1, 3, 0),
		DamageDice:         *random.NewDice(2, 4, 0),
		Delay:              25,
		weightForSelection: 1,
		Range:              4,
	},
	{
		Name:               "Revolver",
		ToHitDice:          *random.NewDice(1, 3, 0),
		DamageDice:         *random.NewDice(1, 3, 0),
		Delay:              15,
		weightForSelection: 1,
		Range:              3,
	},
}
