package main

func (d *dungeon) generateDropFromPawn(p *pawn) {
	const dropPercentageChance = 20
	tries := 1
	switch p.mob.stats.Rarity {
	case 1:
		tries = 10
	}
	count := 0
	for i := 0; i < tries; i++ {
		if rnd.Rand(100) < dropPercentageChance {
			count++
		}
	}
	d.generateRandomDrop(p.x, p.y, count)
}

func (d *dungeon) generateRandomDrop(x, y int, count int) {
	for i := 0; i < count; i++ {
		itm := &item{}
		itm.initAsRandomItem(0)
		d.addItemAt(itm, x, y)
	}
}
