package main

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
				d.placeGoreAround(p.getCoords())
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

func (d *dungeon) canPawnAct(p *pawn) bool {
	if p.canActInTicks < 0 {
		panic("wat")
	}
	return p.canActInTicks == 0
}

func (d *dungeon) getPawnAt(x, y int) *pawn {
	if player.x == x && player.y == y {
		return player
	}
	for _, m := range d.pawns {
		if m.x == x && m.y == y {
			return m
		}
	}
	return nil
}

func (d *dungeon) movePawn(p *pawn, vx, vy int) bool {
	newX, newY := p.x+vx, p.y+vy
	if d.isInBounds(newX, newY) && d.isTilePassableAndEmpty(newX, newY) {
		p.x += vx
		p.y += vy
		p.spendTime(p.getMovementTime())
		return true
	}
	return false
}

func (d *dungeon) DefaultMoveActionWithPawn(p *pawn, vx, vy int) bool {
	newX, newY := p.x+vx, p.y+vy
	if d.getTileAt(newX, newY).code == tileChest && !d.getTileAt(newX, newY).isOpened {
		d.getTileAt(newX, newY).isOpened = true
		d.generateRandomDrop(newX, newY, rnd.RandInRange(1, 3))
		p.spendTime(10)
		return false
	}
	if d.getTileAt(newX, newY).code == tileDoor && !d.getTileAt(newX, newY).isOpened {
		d.getTileAt(newX, newY).isOpened = true
		p.spendTime(10)
		return false
	}
	pawnAtCoords := d.getPawnAt(newX, newY)
	if pawnAtCoords != nil {
		d.doMeleeHit(p, pawnAtCoords)
		return false
	}
	return d.movePawn(p, vx, vy)
}

func (d *dungeon) pickUpItemWithPawn(p *pawn) {
	items := d.getItemsAt(p.x, p.y)
	if len(items) > 0 && p.inv != nil {
		p.spendTime(10)
		p.inv.addItemToStash(items[0])
		d.removeItem(items[0])
		return
	}
	panic("Pick up failed")
}
