package main

func (d *dungeon) executePawnAction(p *pawn) {
	switch p.action.actionCode {
	case pActionWait:
		// do nothing
	case pActionBasicMeleeAttack:
		d.performMeleeHitAction(p)
	case pActionMove:
		d.performMoveActionWithPawn(p)
	default:
		panic("executePawnAction(p *pawn): No such action...")
	}
	p.action.markExecuted()
}

func (d *dungeon) performMoveActionWithPawn(p *pawn) {
	newX, newY := p.x+p.action.x, p.y+p.action.y
	if d.getTileAt(newX, newY).code == tileChest && !d.getTileAt(newX, newY).isOpened {
		d.getTileAt(newX, newY).isOpened = true
		d.generateRandomDrop(newX, newY, rnd.RandInRange(1, 3))
		return
	}
	if d.getTileAt(newX, newY).code == tileDoor && !d.getTileAt(newX, newY).isOpened {
		d.getTileAt(newX, newY).isOpened = true
		return
	}
	if d.isInBounds(newX, newY) && d.isTilePassableAndEmpty(newX, newY) {
		p.x, p.y = newX, newY
	}
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
