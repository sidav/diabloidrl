package main

func game(d *dungeon, pc *playerController) {
	for !stopGame {
		d.exploreAroundPlayer()
		for !stopGame && d.canPawnAct(player) {
			pc.act(d)
		}
		d.clearDeadPawns()
		for _, p := range d.pawns {
			if p.getRegenCooldown() != 0 && d.currentTick%p.getRegenCooldown() == 0 {
				p.regainHitpoints(1)
			}
			if p.canUseFlaskInTicks > 0 {
				p.canUseFlaskInTicks--
			}
			if p.canActInTicks > 0 {
				p.canActInTicks--
			} else if !p.isPlayer() {
				d.actForPawn(p)
			}
		}
		d.currentTick++
	}
}
