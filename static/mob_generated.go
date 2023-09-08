package static

import "diabloidrl/lib/random"

func generateRareMobStats(rnd random.PRNG) *MobStats {
	stats := getWeigtedRandomMobBase(rnd).clone()
	namePrefix := ""
	modFuncs := []func(){
		func() {
			namePrefix = "Raging"
			stats.Damage.Modifier += rnd.RandInRange(1, 3)
		},
		func() {
			namePrefix = "Elusive"
			stats.Evasion += rnd.RandInRange(1, 3)
		},
		func() {
			namePrefix = "Deadeye"
			stats.ToHit.Modifier += rnd.RandInRange(1, 3)
		},
		func() {
			namePrefix = "Fast"
			stats.MovementTime -= rnd.RandInRange(1, 5)
		},
		func() {
			namePrefix = "Furious"
			stats.HitTime -= rnd.RandInRange(1, 4)
		},
	}
	epicType := rnd.Rand(len(modFuncs))
	modFuncs[epicType]()

	percentage := rnd.RandInRange(125, 200)
	stats.MaxHitpoints = percentage * stats.MaxHitpoints / 100
	stats.GivesExperience = percentage * 2 * (stats.GivesExperience + 1) / 100
	stats.Name = namePrefix + " " + stats.Name
	stats.Rarity = 1
	return stats
}
