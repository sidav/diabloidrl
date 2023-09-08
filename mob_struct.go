package main

import (
	"diabloidrl/static"
)

const (
	mobStateIdle uint8 = iota
	mobStateAttacking
)

type mobStruct struct {
	stats        *static.MobStats
	currentState uint8
	stateTimeout uint8 // counted in actions, not in ticks
}

func (ms *mobStruct) initFromStatic(s *static.MobStats) {
	if ms == nil {
		panic("Mob is nil")
	}
	ms.stats = s
}
