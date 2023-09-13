package main

type pawn struct {
	x, y          int
	hitpoints     int
	mob           *mobStruct
	playerStats   *playerStruct
	inv           *inventory
	canActInTicks int
	statusEffects []*statusEffect

	flaskCharges       int
	canUseFlaskInTicks int
}

func (p *pawn) getCoords() (int, int) {
	return p.x, p.y
}

func (p *pawn) isPlayer() bool {
	return p.playerStats != nil
}
