package main

const (
	rpgStatStr uint8 = iota
	rpgStatVgr
	rpgStatDex
	rpgStatInt
	rpgStatVit
	rpgStatsCount
)

func (p *playerStruct) setDefaultStats() {
	p.rpgStats = [rpgStatsCount]int{
		10,
		10,
		10,
		10,
		10,
	}
}

func (p *playerStruct) getStatsMaxHp() int {
	return p.rpgStats[rpgStatVit] + 10
}

func (p *playerStruct) getStatsMaxStm() int {
	return 5 + p.rpgStats[rpgStatVgr]/2
}

func (p *playerStruct) getStatsMovementTime() int {
	return 11 - p.rpgStats[rpgStatDex]/7
}
