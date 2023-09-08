package static

import (
	"diabloidrl/lib/random"
	"fmt"
)

type AmuletStats struct {
	Name   string
	Rarity int

	Egos              []*ego
	affixDescriptions []string

	weightForSelection int
}

func (as *AmuletStats) clone() *AmuletStats {
	newAmuletStats := *as
	if as == &newAmuletStats {
		panic("I must learn more")
	}
	return &newAmuletStats
}

func (as *AmuletStats) GetName() string {
	return as.Name
}

func (as *AmuletStats) GetEgos() []*ego {
	return as.Egos
}

func (as *AmuletStats) GetAffixDescriptions() []string {
	return as.affixDescriptions
}

func (as *AmuletStats) addAffixDescription(s string, args ...interface{}) {
	as.affixDescriptions = append(as.affixDescriptions, fmt.Sprintf(s, args...))
}

func (as *AmuletStats) addEgo(e *ego) {
	as.Egos = append(as.Egos, e)
}

func GenerateRandomAmuletStats(minRarity int) *AmuletStats {
	generators := []func() *AmuletStats{
		// common
		func() *AmuletStats {
			return getRandomAmuletStats(rnd)
		},
		// rare
		func() *AmuletStats {
			return makeAffixedAmuletStatsFromBase(getRandomAmuletStats(rnd), 1)
		},
		// very rare (2 affixes)
		func() *AmuletStats {
			return makeAffixedAmuletStatsFromBase(getRandomAmuletStats(rnd), 2)
		},
		func() *AmuletStats {
			return makeAffixedAmuletStatsFromBase(getRandomAmuletStats(rnd), 3)
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

func makeAffixedAmuletStatsFromBase(baseStats *AmuletStats, affixCount int) *AmuletStats {
	stats := baseStats.clone()
	selectedAffixes := selectRandomAppropriateUniqueAffixesFor(stats, affixCount)
	for _, aff := range selectedAffixes {
		// if aff.AmuletFunc != nil {
		// 	aff.AmuletFunc(stats)
		// }
		if aff.anyFunc != nil {
			aff.anyFunc(stats)
		}
	}
	stats.Name = addAffixesNamesToItemName(selectedAffixes, stats.Name)
	stats.Rarity = affixCount
	return stats
}

func getRandomAmuletStats(rnd random.PRNG) *AmuletStats {
	ind := rnd.SelectRandomIndexFromWeighted(len(sTableAmulets), func(i int) int { return sTableAmulets[i].weightForSelection })
	return sTableAmulets[ind]
}

var sTableAmulets = []*AmuletStats{
	{
		Name:               "Amulet",
		weightForSelection: 1,
	},
	{
		Name:               "Charm",
		weightForSelection: 1,
	},
	{
		Name:               "Talisman",
		weightForSelection: 1,
	},
}
