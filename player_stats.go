package main

const (
	rpgStatStr uint8 = iota
	rpgStatDex
	rpgStatInt
	rpgStatVit
	rpgStatsCount
)

func (p *playerStruct) setDefaultStats() {
	p.rpgStats = [rpgStatsCount]int{
		5,
		5,
		5,
		5,
	}
}

func (p *playerStruct) getStatsMaxHp() int {
	return 4 + p.getExperienceLevel() + p.rpgStats[rpgStatVit]
}

func (p *playerStruct) getStatsMovementTime() int {
	return 10 - p.rpgStats[rpgStatDex]/7
}
