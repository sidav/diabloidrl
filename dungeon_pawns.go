package main

import (
	"diabloidrl/lib/calculations"
)

func (d *dungeon) addPawnAt(p *pawn, x, y int) {
	p.hitpoints = p.getMaxHitpoints()
	p.x, p.y = x, y
	d.pawns = append(d.pawns, p)
}

func (d *dungeon) clearDeadPawns() {
	for i := len(d.pawns) - 1; i >= 0; i-- {
		p := d.pawns[i]
		if !p.isPlayer() && p.hitpoints <= 0 {
			// gibs
			if p.hitpoints <= -p.getMaxHitpoints()/4 {
				log.AppendMessagef("%s explodes into gore!", p.getName())
				player.acquireExperience(2 * p.mob.stats.GivesExperience)
				d.placeGoreAround(p.GetCoords())
				renderer.addAnimationAt(animTypeGibs, p.x, p.y, false)
			} else {
				log.AppendMessagef("%s drops dead.", p.getName())
				player.acquireExperience(p.mob.stats.GivesExperience)
			}
			d.generateDropFromPawn(p)
			d.pawns = append(d.pawns[:i], d.pawns[i+1:]...)
		}
	}
}

func (d *dungeon) arePawnsTouching(p1, p2 *pawn) bool {
	s1 := p1.GetSize()
	s2 := p2.GetSize()
	return calculations.AreRectsInTaxicabRange(p1.x, p1.y, s1, s1, p2.x, p2.y, s2, s2, 1)
}

func (d *dungeon) getPawnAt(x, y int) *pawn {
	// if player.x == x && player.y == y {
	// 	return player
	// }
	for _, m := range d.pawns {
		if m.GetSize() <= 1 {
			if m.x == x && m.y == y {
				return m
			}
		} else {
			xDist := x - m.x
			yDist := y - m.y
			if xDist >= 0 && xDist < m.mob.stats.Size && yDist >= 0 && yDist < m.mob.stats.Size {
				return m
			}
		}
	}
	return nil
}
