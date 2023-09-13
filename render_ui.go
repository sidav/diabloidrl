package main

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
)

func (r *rendererStruct) renderUI(dung *dungeon) {
	w, h := r.view.GetViewportSize()
	uiTopYCoord := h + LOG_SIZE
	sphereW := 130 * r.uiH / 100
	if sphereW%2 == 0 {
		sphereW--
	}
	cw.ResetStyle()
	cw.PutStringf(sphereW, uiTopYCoord, "T: %d (%d)", dung.currentTick, player.playerStats.lastActionTicks)
	if player.inv.getItemInSlot(invSlotFlask) != nil {
		if player.canUseFlaskInTicks == 0 {
			cw.PutStringf(sphereW, uiTopYCoord, "Q: use flask (%d/%d charges)", player.flaskCharges, player.getMaxFlaskCharges())
		} else {
			cw.SetStyle(tcell.ColorDarkGray, tcell.ColorBlack)
			cw.PutStringf(sphereW, uiTopYCoord, "Flask (%d/%d charges) on cooldown: %d",
				player.flaskCharges, player.getMaxFlaskCharges(), player.canUseFlaskInTicks)
		}
	}
	// hp sphere
	cw.SetStyle(tcell.ColorDarkRed, tcell.ColorBlack)
	cw.DrawPartiallyFilledEllipse(0, uiTopYCoord, sphereW, r.uiH, player.hitpoints, player.getMaxHitpoints(), '.', '#')
	cw.PutStringCenteredAt(fmt.Sprintf("%d/%d", player.hitpoints, player.getMaxHitpoints()), sphereW/2, uiTopYCoord+r.uiH/2)
	// mana sphere
	cw.SetStyle(tcell.ColorDarkBlue, tcell.ColorBlack)
	cw.DrawPartiallyFilledEllipse(w-sphereW, uiTopYCoord, sphereW, r.uiH, 7, 10, '.', '#')
	// exp bar
	exp, remExp := player.playerStats.getNormalizedCurrAndRemainingLevelExp()
	cw.RenderProgressBar("", sphereW-2, uiTopYCoord+r.uiH-1, w-sphereW*2+3, exp, remExp, tcell.ColorWhite, tcell.ColorWhite, tcell.ColorDarkBlue)
	if player.playerStats.skillPoints > 0 {
		cw.SetStyle(tcell.ColorBlack, tcell.ColorYellow)
		cw.PutStringCenteredAt(fmt.Sprintf(" %d STATS POINTS AVAILABLE! ", player.playerStats.skillPoints), w/2, uiTopYCoord+r.uiH-2)
	} else {
		cw.PutStringCenteredAt(fmt.Sprintf("LVL %d (exp %d - %d/%d exp)",
			player.playerStats.getExperienceLevel(), player.playerStats.experience, exp, remExp), w/2, uiTopYCoord+r.uiH-2)
	}
}
