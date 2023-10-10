package main

func (d *dungeon) executePawnAction(p *pawn) {

}

func (d *dungeon) movePawn(p *pawn, vx, vy int) bool {
	newX, newY := p.x+vx, p.y+vy
	if d.isInBounds(newX, newY) && d.isTilePassableAndEmpty(newX, newY) {
		p.x += vx
		p.y += vy
		return true
	}
	return false
}

func (d *dungeon) DefaultMoveActionWithPawn(p *pawn, vx, vy int) bool {
	newX, newY := p.x+vx, p.y+vy
	if d.getTileAt(newX, newY).code == tileChest && !d.getTileAt(newX, newY).isOpened {
		d.getTileAt(newX, newY).isOpened = true
		d.generateRandomDrop(newX, newY, rnd.RandInRange(1, 3))
		return false
	}
	if d.getTileAt(newX, newY).code == tileDoor && !d.getTileAt(newX, newY).isOpened {
		d.getTileAt(newX, newY).isOpened = true
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
		p.inv.addItemToStash(items[0])
		d.removeItem(items[0])
		return
	}
	panic("Pick up failed")
}
