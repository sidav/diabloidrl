package static

import (
	"diabloidrl/lib/random"
	attackpattern "diabloidrl/static/attack_pattern"
)

type MobStats struct {
	Name                                 string
	MaxHitpoints                         int
	ToHit, Damage                        random.Dice
	Evasion                              int
	CritChancePercent, CritDamagePercent int
	Attacks                              []*AttackSkill
	AttackRange                          int // 0 for melee
	RegenCooldown                        int
	MovementTime                         int
	HitTime                              int
	GivesExperience                      int
	Rarity                               int
	AsciiPic                             []string
	Size                                 int // Mob is size*size cells; 0 is threated like 1

	weightForSelection int
}

func (ms *MobStats) clone() *MobStats {
	newMobStats := *ms
	if ms == &newMobStats {
		panic("I must learn more")
	}
	return &newMobStats
}

func GenerateRandomMobBase(rnd random.PRNG, rarity int) *MobStats {
	switch rarity {
	case 0:
		return getWeigtedRandomMobBase(rnd)
	case 1:
		return generateRareMobStats(rnd)
	}
	panic("No such rarity!")
}

func getWeigtedRandomMobBase(rnd random.PRNG) *MobStats {
	ind := rnd.SelectRandomIndexFromWeighted(len(STableMobs), func(i int) int { return STableMobs[i].weightForSelection })
	return STableMobs[ind]
}

var STableMobs = []*MobStats{
	{
		Name:         "Zombie",
		AsciiPic:     []string{"z"},
		MaxHitpoints: 15,
		ToHit:        *random.NewDice(1, 6, 0),
		Attacks: []*AttackSkill{
			{
				Pattern:             attackpattern.SimpleAttack{},
				HitTimePercentage:   100,
				DamagePercentage:    100,
				ToHitRollPercentage: 100,
			},
		},
		Evasion:            1,
		Damage:             *random.NewDice(1, 2, 0),
		CritChancePercent:  1,
		CritDamagePercent:  150,
		RegenCooldown:      200,
		MovementTime:       15,
		HitTime:            10,
		GivesExperience:    1,
		weightForSelection: 10,
	},
	{
		Name:         "Skeleton",
		AsciiPic:     []string{"k"},
		MaxHitpoints: 5,
		ToHit:        *random.NewDice(2, 4, 0),
		Attacks: []*AttackSkill{
			{
				Pattern:             attackpattern.SimpleAttack{},
				HitTimePercentage:   50,
				DamagePercentage:    100,
				ToHitRollPercentage: 100,
			},
			{
				Pattern: &attackpattern.SweepAttack{
					RadiusFromAttacker: 1,
					RadiusFromTarget:   1,
				},
				HitTimePercentage:   200,
				DamagePercentage:    100,
				ToHitRollPercentage: 75,
			},
		},
		Evasion:            3,
		Damage:             *random.NewDice(1, 2, 0),
		CritChancePercent:  1,
		CritDamagePercent:  150,
		RegenCooldown:      200,
		MovementTime:       11,
		HitTime:            10,
		GivesExperience:    2,
		weightForSelection: 5,
	},
	{
		Name:         "Wraith",
		AsciiPic:     []string{"w"},
		MaxHitpoints: 5,
		ToHit:        *random.NewDice(1, 6, 0),
		Attacks: []*AttackSkill{
			{
				Pattern: &attackpattern.RoundAttack{
					Size: 1,
				},
				HitTimePercentage:   150,
				DamagePercentage:    100,
				ToHitRollPercentage: 100,
			},
		},
		Evasion:            5,
		Damage:             *random.NewDice(1, 2, 0),
		CritChancePercent:  1,
		CritDamagePercent:  150,
		MovementTime:       11,
		HitTime:            10,
		RegenCooldown:      50,
		GivesExperience:    5,
		weightForSelection: 3,
	},
	{
		Name:         "Cultist",
		AsciiPic:     []string{"c"},
		MaxHitpoints: 5,
		Attacks: []*AttackSkill{
			{
				Pattern:             attackpattern.SimpleAttack{},
				HitTimePercentage:   50,
				DamagePercentage:    100,
				ToHitRollPercentage: 100,
			},
			{
				Pattern: &attackpattern.LineAttack{
					Size:   1,
					Length: 5,
				},
				HitTimePercentage:   150,
				DamagePercentage:    100,
				ToHitRollPercentage: 100,
			},
		},
		ToHit:              *random.NewDice(1, 6, 0),
		Evasion:            2,
		Damage:             *random.NewDice(1, 4, 0),
		CritChancePercent:  1,
		CritDamagePercent:  150,
		AttackRange:        4,
		MovementTime:       11,
		HitTime:            10,
		RegenCooldown:      50,
		GivesExperience:    2,
		weightForSelection: 4,
	},
	{
		Name: "Giant Swordsman",
		AsciiPic: []string{
			" o ",
			"\\=0",
			"/ \\",
		},
		Size:         3,
		MaxHitpoints: 15,
		ToHit:        *random.NewDice(1, 6, 0),
		Attacks: []*AttackSkill{
			{
				Pattern:             attackpattern.SimpleAttack{},
				HitTimePercentage:   100,
				DamagePercentage:    100,
				ToHitRollPercentage: 100,
			},
			{
				Pattern: &attackpattern.SweepAttack{
					RadiusFromAttacker: 3,
					RadiusFromTarget:   4,
				},
				HitTimePercentage:   200,
				DamagePercentage:    100,
				ToHitRollPercentage: 100,
			},
		},
		Evasion:            2,
		Damage:             *random.NewDice(1, 4, 0),
		CritChancePercent:  1,
		CritDamagePercent:  150,
		MovementTime:       10,
		HitTime:            25,
		RegenCooldown:      50,
		GivesExperience:    2,
		weightForSelection: 1,
	},
	{
		Name: "Giant Pikeman",
		AsciiPic: []string{
			"^o ",
			"|=0",
			"/ \\",
		},
		Size:         3,
		MaxHitpoints: 15,
		ToHit:        *random.NewDice(1, 6, 0),
		Attacks: []*AttackSkill{
			{
				Pattern: &attackpattern.LineAttack{
					Size:   2,
					Length: 5,
				},
				HitTimePercentage:   150,
				DamagePercentage:    100,
				ToHitRollPercentage: 100,
			},
			{
				Pattern: &attackpattern.RoundAttack{
					Size: 2,
				},
				HitTimePercentage:   200,
				DamagePercentage:    100,
				ToHitRollPercentage: 100,
			},
		},
		Evasion:            2,
		Damage:             *random.NewDice(1, 4, 0),
		CritChancePercent:  1,
		CritDamagePercent:  150,
		MovementTime:       10,
		HitTime:            25,
		RegenCooldown:      50,
		GivesExperience:    2,
		weightForSelection: 1,
	},
}
