package attackpattern

type AttackPattern interface {
	CanBePerformedOn(attacker, target ActorForPattern) bool
	GetAttackedCoords(attacker ActorForPattern, targetX, targetY int) [][2]int
	GetAimAt(attacker, target ActorForPattern) (int, int)
}

type ActorForPattern interface {
	GetSize() int
	GetCoords() (int, int)
}
