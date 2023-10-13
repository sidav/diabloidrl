package main

func (pc *playerController) autoexploreMode(dung *dungeon) {
	if pc.shouldAutoexplorePause(dung) {
		pc.mode = pcModeDefault
		return
	}
	if dung.currentTick%50 == 0 {
		renderer.renderGameMainScreen(dung)
	}
	vx, vy := pc.getNextAutoexploreStep(dung)
	if vx == 0 && vy == 0 {
		log.AppendMessage("Auto-explore finished.")
		pc.mode = pcModeDefault
		return
	}
	if cw.ReadKeyAsync(25) != "" {
		log.AppendMessage("Auto-explore interrupted.")
		pc.mode = pcModeDefault
		return
	}
	dung.dmap[player.x][player.y].wasOnPlayerPath = true
	player.action.set(pActionMove, 0, player.getMovementTime(), vx, vy)
}

func (pc *playerController) getNextAutoexploreStep(dung *dungeon) (int, int) {
	dung.playerExplorationDM.Purge()
	for x := 0; x < len(dung.dmap); x++ {
		for y := 0; y < len(dung.dmap[x]); y++ {
			if !dung.dmap[x][y].wasSeenByPlayer {
				dung.playerExplorationDM.SetTarget(x, y, 0)
			}
		}
	}
	dung.playerExplorationDM.Calculate()
	return dung.playerExplorationDM.GetVectorToBestNeighbourFrom(player.x, player.y)
}

func (pc *playerController) shouldAutoexplorePause(dung *dungeon) bool {
	if dung.isAnyPawnInPlayerFOV(true) {
		log.AppendMessagef("%s spotted!", dung.getAllPawnsInPlayerFOV(true)[0].getName())
		return true
	}
	px, py := player.GetCoords()
	sight := player.getVisionRadius()
	for x := px - sight; x <= px+sight; x++ {
		for y := py - sight; y <= py+sight; y++ {
			if dung.isInBounds(x, y) && dung.playerFOVMap[x][y] &&
				dung.getTileAt(x, y).code == tileChest && !dung.getTileAt(x, y).isOpened {

				log.AppendMessage("Unopened chest spotted!")
				return true
			}
		}
	}
	return false
}
