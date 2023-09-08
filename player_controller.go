package main

const (
	pcModeDefault uint8 = iota
	pcModeAutoexplore
	pcModeAutoattack
)

type playerController struct {
	mode uint8
}

func (pc *playerController) act(dung *dungeon) {
	switch pc.mode {
	case pcModeDefault:
		pc.defaultMode(dung)
		player.playerStats.lastActionTicks = 0
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
		dung.DefaultMoveActionWithPawn(player, vx, vy)
		itms := dung.getItemsAt(player.x, player.y)
		if len(itms) == 1 {
			log.AppendMessagef("You see here %s.", itms[0].getName())
		}
		if len(itms) > 1 {
			log.AppendMessagef("You see here %d items", len(itms))
		}
		return
	}
	switch key {
	case "ESCAPE":
		stopGame = true
	case "s":
		player.spendTime(player.getMovementTime())
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
	case "TAB":
		pc.mode = pcModeAutoattack
	}
}

func (pc *playerController) pickUp(d *dungeon) {
	itms := d.getItemsAt(player.getCoords())
	if len(itms) > 0 {
		log.AppendMessagef("You pick up %s.", itms[0].getName())
		d.pickUpItemWithPawn(player)
	} else {
		log.AppendMessage("There is nothing here to pick up.")
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
