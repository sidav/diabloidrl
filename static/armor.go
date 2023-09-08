package static

import (
	"diabloidrl/lib/random"
	"fmt"
)

const (
	ArmorSlotHead uint8 = iota
	ArmorSlotBody
)

type ArmorStats struct {
	Name            string
	Slot            uint8
	Defense         int
	EvasionModifier int
	Rarity          int

	Egos              []*ego
	affixDescriptions []string

	weightForSelection int
}

func (as *ArmorStats) clone() *ArmorStats {
	newArmorStats := *as
	if as == &newArmorStats {
		panic("I must learn more")
	}
	return &newArmorStats
}

func (as *ArmorStats) GetName() string {
	return as.Name
}

func (as *ArmorStats) GetEgos() []*ego {
	return as.Egos
}

func (as *ArmorStats) GetAffixDescriptions() []string {
	return as.affixDescriptions
}

func (as *ArmorStats) addAffixDescription(s string, args ...interface{}) {
	as.affixDescriptions = append(as.affixDescriptions, fmt.Sprintf(s, args...))
}

func (as *ArmorStats) addEgo(e *ego) {
	as.Egos = append(as.Egos, e)
}

func GenerateRandomArmorStats(minRarity int) *ArmorStats {
	generators := []func() *ArmorStats{
		// common
		func() *ArmorStats {
			return getRandomArmorStats(rnd)
		},
		// rare
		func() *ArmorStats {
			return makeAffixedArmorStatsFromBase(getRandomArmorStats(rnd), 1)
		},
		// very rare (2 affixes)
		func() *ArmorStats {
			return makeAffixedArmorStatsFromBase(getRandomArmorStats(rnd), 2)
		},
		func() *ArmorStats {
			return makeAffixedArmorStatsFromBase(getRandomArmorStats(rnd), 3)
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

func makeAffixedArmorStatsFromBase(baseStats *ArmorStats, affixCount int) *ArmorStats {
	stats := baseStats.clone()
	selectedAffixes := selectRandomAppropriateUniqueAffixesFor(stats, affixCount)
	for _, aff := range selectedAffixes {
		if aff.armorFunc != nil {
			aff.armorFunc(stats)
		}
		if aff.anyFunc != nil {
			aff.anyFunc(stats)
		}
	}
	stats.Name = addAffixesNamesToItemName(selectedAffixes, stats.Name)
	stats.Rarity = affixCount
	return stats
}

func getRandomArmorStats(rnd random.PRNG) *ArmorStats {
	ind := rnd.SelectRandomIndexFromWeighted(len(sTableArmors), func(i int) int { return sTableArmors[i].weightForSelection })
	return sTableArmors[ind]
}

var sTableArmors = []*ArmorStats{
	{
		Name:               "Cap",
		Slot:               ArmorSlotHead,
		Defense:            1,
		EvasionModifier:    1,
		weightForSelection: 1,
	},
	{
		Name:               "Helmet",
		Slot:               ArmorSlotHead,
		Defense:            2,
		EvasionModifier:    -1,
		weightForSelection: 1,
	},
	{
		Name:               "Robe",
		Slot:               ArmorSlotBody,
		Defense:            1,
		EvasionModifier:    1,
		weightForSelection: 1,
	},
	{
		Name:               "Leather armor",
		Slot:               ArmorSlotBody,
		Defense:            3,
		EvasionModifier:    -2,
		weightForSelection: 1,
	},
}
