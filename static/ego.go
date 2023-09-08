package static

type EgoCodeType uint8

// "Ego" is a property that doesn't directly alter item's own stats
// For example, "+2 to hit" on a weapon is NOT an ego
// Egos should be taken into account in the game's own logic, not in this package (here egos are only assigned)
type ego struct {
	Code  EgoCodeType
	Value int
}

const (
	EgoCodeRegenerationPeriodPercent EgoCodeType = iota
	EgoCodeAdditionalCritChancePercent
	EgoCodeAdditionalCritDamagePercent
	EgoCodeAdditionalLightRadius
	EgoCodeAdditionalMaxHP
	EgoCodeThornsChance
)
