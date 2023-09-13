package static

import (
	"diabloidrl/lib/random"
	"fmt"
)

type FlaskStats struct {
	Name string

	HealStrength       int
	HealTicksPeriod    int
	HealEffectDuration int

	CooldownBetweenSips int
	MaxCharges          int

	Rarity int

	Egos              []*ego
	affixDescriptions []string

	weightForSelection int
}

func (as *FlaskStats) clone() *FlaskStats {
	newFlaskStats := *as
	return &newFlaskStats
}

func (as *FlaskStats) GetName() string {
	return as.Name
}

func (as *FlaskStats) GetEgos() []*ego {
	return as.Egos
}

func (as *FlaskStats) GetAffixDescriptions() []string {
	return as.affixDescriptions
}

func (as *FlaskStats) addAffixDescription(s string, args ...interface{}) {
	as.affixDescriptions = append(as.affixDescriptions, fmt.Sprintf(s, args...))
}

func (as *FlaskStats) addEgo(e *ego) {
	as.Egos = append(as.Egos, e)
}

func GenerateRandomFlaskStats(minRarity int) *FlaskStats {
	generators := []func() *FlaskStats{
		// common
		func() *FlaskStats {
			return getRandomFlaskStats(rnd)
		},
		// rare
		func() *FlaskStats {
			return makeAffixedFlaskStatsFromBase(getRandomFlaskStats(rnd), 1)
		},
		// very rare (2 affixes)
		func() *FlaskStats {
			return makeAffixedFlaskStatsFromBase(getRandomFlaskStats(rnd), 2)
		},
		func() *FlaskStats {
			return makeAffixedFlaskStatsFromBase(getRandomFlaskStats(rnd), 3)
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

func makeAffixedFlaskStatsFromBase(baseStats *FlaskStats, affixCount int) *FlaskStats {
	stats := baseStats.clone()
	selectedAffixes := selectRandomAppropriateUniqueAffixesFor(stats, affixCount)
	for _, aff := range selectedAffixes {
		if aff.flaskFunc != nil {
			aff.flaskFunc(stats)
		}
		if aff.anyFunc != nil {
			aff.anyFunc(stats)
		}
	}
	stats.Name = addAffixesNamesToItemName(selectedAffixes, stats.Name)
	stats.Rarity = affixCount
	return stats
}

func getRandomFlaskStats(rnd random.PRNG) *FlaskStats {
	ind := rnd.SelectRandomIndexFromWeighted(len(sTableFlasks), func(i int) int { return sTableFlasks[i].weightForSelection })
	return sTableFlasks[ind]
}

var sTableFlasks = []*FlaskStats{
	// balanced
	{
		Name:                "Flask",
		MaxCharges:          2,
		CooldownBetweenSips: 1000,
		HealStrength:        1,
		HealTicksPeriod:     10,
		HealEffectDuration:  100,

		weightForSelection: 1,
	},
	// heals fast, short cooldown
	{
		Name:                "Vial",
		MaxCharges:          3,
		CooldownBetweenSips: 500,
		HealStrength:        1,
		HealTicksPeriod:     8,
		HealEffectDuration:  50,

		weightForSelection: 1,
	},
	// heals slowly, many charges
	{
		Name:                "Bottle",
		MaxCharges:          5,
		CooldownBetweenSips: 1000,
		HealStrength:        2,
		HealTicksPeriod:     25,
		HealEffectDuration:  200,

		weightForSelection: 1,
	},
}
