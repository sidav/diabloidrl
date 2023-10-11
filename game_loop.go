package main

func game(d *dungeon, pc *playerController) {
	for !stopGame {
		d.exploreAroundPlayer()
		for !stopGame && player.action.isEmpty() {
			pc.act(d)
		}
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
		}
		for _, p := range d.pawns {
			if !p.isPlayer() && p.hitpoints <= 0 {
				continue
			}
			if p.action.canActionOccurNow() {
				d.executePawnAction(p)
			}
			if !p.isPlayer() && p.action.isEmpty() {
				d.aiActForPawn(p)
			}
		}
		d.clearDeadPawns()
		d.currentTick++
	}
}
