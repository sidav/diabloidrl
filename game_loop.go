package main

func game(d *dungeon, pc *playerController) {
	for !stopGame {
		d.exploreAroundPlayer()
		for !stopGame && player.action.ended() {
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
			p.cleanupStatusEffects()
			d.applyPassiveStatusEffects(p)

			p.action.updateDelays()
			if !p.isPlayer() && p.action.ended() {
				d.aiActForPawn(p)
			}
		}
		d.currentTick++
	}
}
