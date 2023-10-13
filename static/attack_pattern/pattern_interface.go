package attackpattern

type AttackPattern interface {
	CanBePerformedOn(ActorForPattern, ActorForPattern) bool
	GetAttackCoords(attacker ActorForPattern, targetX, targetY int) [][2]int
	GetAimCoords(ActorForPattern, ActorForPattern) (int, int)
}

type ActorForPattern interface {
	GetSize() int
	GetCoords() (int, int)
}
