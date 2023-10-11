package main

func (d *dungeon) executePawnAction(p *pawn) {
	switch p.action.actionCode {
	case pActionAttack:
	case pActionMove:
		d.DefaultMoveActionWithPawn(p)
	case pActionWait:
	default:
		panic("executePawnAction(p *pawn): No such action...")
	}
	p.action.markExecuted()
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

func (d *dungeon) DefaultMoveActionWithPawn(p *pawn) bool {
	newX, newY := p.x+p.action.x, p.y+p.action.y
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
	return d.movePawn(p, p.action.x, p.action.y)
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
