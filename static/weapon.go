package static

import (
	"diabloidrl/lib/random"
	attackpattern "diabloidrl/static/attack_pattern"
	"fmt"
)

type WeaponStats struct {
	Name         string
	ToHitDice    random.Dice
	DamageDice   random.Dice
	Delay        int
	AttackSkills []AttackSkill
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
		AttackSkills: []AttackSkill{
			{
				Pattern:             attackpattern.SimpleAttack{},
				HitTimePercentage:   100,
				DamagePercentage:    100,
				ToHitRollPercentage: 100,
			},
		},
	},
	{
		Name:               "Short Blade",
		ToHitDice:          *random.NewDice(2, 3, 0),
		DamageDice:         *random.NewDice(2, 6, 0),
		Delay:              12,
		weightForSelection: 4,
		AttackSkills: []AttackSkill{
			{
				Pattern:             attackpattern.SimpleAttack{},
				HitTimePercentage:   100,
				DamagePercentage:    100,
				ToHitRollPercentage: 100,
			},
		},
	},
	{
		Name:               "Spear",
		ToHitDice:          *random.NewDice(2, 3, 0),
		DamageDice:         *random.NewDice(1, 6, 0),
		Delay:              20,
		weightForSelection: 1,
		AttackSkills: []AttackSkill{
			{
				Pattern: &attackpattern.LineAttack{
					Size:   1,
					Length: 2,
				},
				HitTimePercentage:   100,
				DamagePercentage:    100,
				ToHitRollPercentage: 100,
			},
		},
	},
	{
		Name:               "Axe",
		ToHitDice:          *random.NewDice(2, 3, 0),
		DamageDice:         *random.NewDice(2, 3, 0),
		Delay:              20,
		weightForSelection: 1,
		AttackSkills: []AttackSkill{
			{
				Pattern: &attackpattern.SweepAttack{
					RadiusFromAttacker: 1,
					RadiusFromTarget:   1,
				},
				HitTimePercentage:   100,
				DamagePercentage:    100,
				ToHitRollPercentage: 100,
			},
		},
	},
	{
		Name:               "Crossbow",
		ToHitDice:          *random.NewDice(1, 3, 0),
		DamageDice:         *random.NewDice(2, 4, 0),
		Delay:              25,
		weightForSelection: 1,
		Range:              4,
		AttackSkills: []AttackSkill{
			{
				Pattern: &attackpattern.LineAttack{
					Size:   1,
					Length: 4,
				},
				HitTimePercentage:   100,
				DamagePercentage:    100,
				ToHitRollPercentage: 100,
			},
		},
	},
	{
		Name:               "Revolver",
		ToHitDice:          *random.NewDice(1, 3, 0),
		DamageDice:         *random.NewDice(1, 3, 0),
		Delay:              15,
		weightForSelection: 1,
		Range:              3,
		AttackSkills: []AttackSkill{
			{
				Pattern: &attackpattern.LineAttack{
					Size:   1,
					Length: 5,
				},
				HitTimePercentage:   100,
				DamagePercentage:    100,
				ToHitRollPercentage: 100,
			},
		},
	},
}
