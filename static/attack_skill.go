package static

import attackpattern "diabloidrl/static/attack_pattern"

type AttackSkill struct {
	Pattern             attackpattern.AttackPattern
	StaminaCost         int
	HitTimePercentage   int
	DamagePercentage    int
	ToHitRollPercentage int
}
