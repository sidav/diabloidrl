package main

const (
	pcModeDefault uint8 = iota
	pcModeAutoexplore
	pcModeAutoattack
)

type playerController struct {
	mode uint8

	prevX, prevY int
}

func (pc *playerController) act(dung *dungeon) {
	switch pc.mode {
	case pcModeDefault:
		pc.showItemsHereIfNeeded(dung)
		pc.defaultMode(dung)
		if !player.action.isEmpty() {
			player.playerStats.lastActionTicks = player.action.ticksBeforeAction + player.action.ticksAfterAction
		}
		dung.resetPlayerPath()
	case pcModeAutoexplore:
		pc.autoexploreMode(dung)
	case pcModeAutoattack:
		pc.doAutoAttackTurn(dung)
	}
}

func (pc *playerController) defaultMode(dung *dungeon) {
	renderer.renderGameMainScreen(dung)
	key := cw.ReadKey()
	vx, vy := pc.keyToStep(key)
	// log.AppendMessagef("vx, vy %d, %d", vx, vy)
	if vx != 0 || vy != 0 {
		mobAtCoords := dung.getPawnAt(player.x+vx, player.y+vy)
		if mobAtCoords != nil && pc.getAttackPattern().Pattern.CanBePerformedOn(player, mobAtCoords) {
			player.action.setAttack(player, pc.getAttackPattern(), 0, player.getHitTime(), mobAtCoords)
		} else {
			player.action.set(pActionMove, 0, player.getMovementTime(), vx, vy)
		}
		return
	}
	switch key {
	case "EXIT": // ctrl+c
		stopGame = true
	case "s":
		player.action.set(pActionWait, 0, ticksInTurn, 0, 0)
	case "g":
		pc.pickUp(dung)
	case "o":
		pc.mode = pcModeAutoexplore
	case "i":
		pc.callInventoryMenu()
	case "p":
		if player.playerStats.skillPoints > 0 {
			pc.callLevelUpMenu()
		} else {
			pc.showPlayerStats()
		}
	case "Q":
		if player.canUseFlaskInTicks == 0 { // player.flaskCharges > 0 {
			player.useFlask()
			log.AppendMessagef("You sip healing liquid.")
		} else {
			log.AppendMessagef("You can't drink yet!")
		}
	case "TAB":
		pc.mode = pcModeAutoattack
	default:
		log.AppendMessagef("Key '%s' does nothing.", key)
	}
}

func (pc *playerController) pickUp(d *dungeon) {
	itms := d.getItemsAt(player.GetCoords())
	if len(itms) > 0 {
		log.AppendMessagef("You pick up %s.", itms[0].getName())
		d.pickUpItemWithPawn(player)
	} else {
		log.AppendMessage("There is nothing here to pick up.")
	}
}

func (pc *playerController) showItemsHereIfNeeded(dung *dungeon) {
	px, py := player.GetCoords()
	if pc.prevX == px && pc.prevY == py {
		return
	}
	pc.prevX = px
	pc.prevY = py
	itms := dung.getItemsAt(px, py)
	if len(itms) == 1 {
		log.AppendMessagef("You see here %s.", itms[0].getName())
	}
	if len(itms) > 1 {
		log.AppendMessagef("You see here %d items", len(itms))
	}
}

func (pc *playerController) keyToStep(k string) (int, int) {
	switch k {
	case "q":
		return -1, -1
	case "w":
		return 0, -1
	case "e":
		return 1, -1
	case "a":
		return -1, 0
	case "d":
		return 1, 0
	case "z":
		return -1, 1
	case "x":
		return 0, 1
	case "c":
		return 1, 1
	default:
		return 0, 0
	}
}
