package static

import (
	"diabloidrl/lib/random"
	"fmt"
)

type ShieldStats struct {
	Name string

	BlockChancePerc  int
	BlockStaminaCost int
	MaxBlockedDamage int

	Rarity int

	Egos              []*ego
	affixDescriptions []string

	weightForSelection int
}

func (shs *ShieldStats) clone() *ShieldStats {
	newArmorStats := *shs
	if shs == &newArmorStats {
		panic("I must learn more")
	}
	return &newArmorStats
}

func (shs *ShieldStats) GetName() string {
	return shs.Name
}

func (shs *ShieldStats) GetEgos() []*ego {
	return shs.Egos
}

func (shs *ShieldStats) GetAffixDescriptions() []string {
	return shs.affixDescriptions
}

func (shs *ShieldStats) addAffixDescription(s string, args ...interface{}) {
	shs.affixDescriptions = append(shs.affixDescriptions, fmt.Sprintf(s, args...))
}

func (shs *ShieldStats) addEgo(e *ego) {
	shs.Egos = append(shs.Egos, e)
}

func GenerateRandomShieldStats(minRarity int) *ShieldStats {
	generators := []func() *ShieldStats{
		// common
		func() *ShieldStats {
			return getRandomShieldStats(rnd)
		},
		// rare
		func() *ShieldStats {
			return makeAffixedShieldStatsFromBase(getRandomShieldStats(rnd), 1)
		},
		// very rare (2 affixes)
		func() *ShieldStats {
			return makeAffixedShieldStatsFromBase(getRandomShieldStats(rnd), 2)
		},
		func() *ShieldStats {
			return makeAffixedShieldStatsFromBase(getRandomShieldStats(rnd), 3)
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

func makeAffixedShieldStatsFromBase(baseStats *ShieldStats, affixCount int) *ShieldStats {
	stats := baseStats.clone()
	selectedAffixes := selectRandomAppropriateUniqueAffixesFor(stats, affixCount)
	for _, aff := range selectedAffixes {
		if aff.shieldFunc != nil {
			aff.shieldFunc(stats)
		}
		if aff.anyFunc != nil {
			aff.anyFunc(stats)
		}
	}
	stats.Name = addAffixesNamesToItemName(selectedAffixes, stats.Name)
	stats.Rarity = affixCount
	return stats
}

func getRandomShieldStats(rnd random.PRNG) *ShieldStats {
	ind := rnd.SelectRandomIndexFromWeighted(len(sTableShields), func(i int) int { return sTableShields[i].weightForSelection })
	return sTableShields[ind]
}

var sTableShields = []*ShieldStats{
	{
		Name:               "Light shield",
		BlockChancePerc:    30,
		BlockStaminaCost:   0,
		MaxBlockedDamage:   2,
		weightForSelection: 1,
	},
	{
		Name:               "Tower shield",
		BlockChancePerc:    40,
		BlockStaminaCost:   1,
		MaxBlockedDamage:   5,
		weightForSelection: 1,
	},
}
