package main

import intmath "diabloidrl/lib/calculations/int_math"

func (d *dungeon) executePawnAction(p *pawn) {
	spentStamina := 0
	switch p.action.code {
	case pActionWait:
		// do nothing
	case pActionAttack:
		d.performAttackAction(p)
		spentStamina = p.action.attackData.StaminaCost
	case pActionMove:
		d.performMoveActionWithPawn(p)
		spentStamina = 1
	default:
		panic("executePawnAction(p *pawn): No such action...")
	}
	p.action.markExecuted()
	p.stamina = intmath.Max(0, p.stamina-spentStamina)
	if p.action.isEmpty() {
		p.action.reset()
	}
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
	if d.isTilePassableForPawn(newX, newY, p) {
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

func (d *dungeon) performAttackAction(p *pawn) {
	if p.action.attackData == nil {
		panic("Attack data not set for attack action!")
	}
	acoords := p.action.attackData.Pattern.GetAttackedCoords(p, p.action.x, p.action.y)
	for i := range acoords {
		d.performHitOnCoords(p, acoords[i][0], acoords[i][1], i == 0)
	}
}
