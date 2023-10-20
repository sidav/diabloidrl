package main

import (
	"diabloidrl/lib/tcell_console_wrapper"
	"diabloidrl/static"
	"fmt"

	"github.com/gdamore/tcell/v2"
)

func (pc *playerController) showPlayerStats() {
	var lines []string
	// TODO: move this to the console wrapper
	cw.ClearScreen()
	w, h := cw.GetConsoleSize()
	cw.SetStyle(tcell.ColorBlack, tcell.ColorBlue)
	cw.DrawRect(0, 0, w-1, h-1)
	cw.PutStringCenteredAt("PLAYER STATS", w/2, 0)
	exp, remExp := player.playerStats.getNormalizedCurrAndRemainingLevelExp()
	lines = append(lines,
		fmt.Sprintf("Level %d - exp %d (%d/%d exp for next level)", player.playerStats.getExperienceLevel(), player.playerStats.experience, exp, remExp),
		"",
		fmt.Sprintf("Health %d/%d - regen each %d ticks", player.hitpoints, player.getMaxHitpoints(), player.getRegenCooldown()),
		fmt.Sprintf("Movement time %d, Light radius: %d", player.getMovementTime(), player.getVisionRadius()),
		"",
		fmt.Sprintf("To Hit %s, DMG %d-%d, attack delay %d",
			player.getHitDice().GetShortDescriptionString(),
			player.getDamageDice().GetMinimumPossible(), player.getHitDice().GetMaximumPossible(),
			player.getHitTime(),
		),
		fmt.Sprintf("Crit chance %d%%, crit damage percentage %d%%", player.getCriticalChancePercent(), player.getCriticalDamagePercent()),
		"",
		fmt.Sprintf("Evasion %d, Armor Class %d", player.getEvasion(), player.getArmorClass()),
	)
	cw.ResetStyle()
	for i, l := range lines {
		cw.PutString(l, 1, i+1)
	}
	cw.FlushScreen()
	cw.ReadKey()
}

func (pc *playerController) callLevelUpMenu() {
	menu := tcell_console_wrapper.DescriptionHeavySelectMenu{Title: "Level up a stat"}
	menu.AddMenuItem("Strength    ", []string{
		fmt.Sprintf("Current: %d", player.playerStats.rpgStats[rpgStatStr]),
		"Useless for now"},
	)
	menu.AddMenuItem("Vigor       ", []string{
		fmt.Sprintf("Current: %d", player.playerStats.rpgStats[rpgStatVgr]),
		fmt.Sprintf(" Max stamina: %d", player.getMaxStamina())},
	)
	menu.AddMenuItem("Dexterity   ", []string{
		fmt.Sprintf("Current: %d", player.playerStats.rpgStats[rpgStatDex]),
		fmt.Sprintf(" Movement time: %d", player.getMovementTime())},
	)
	menu.AddMenuItem("Intelligence", []string{
		fmt.Sprintf("Current: %d", player.playerStats.rpgStats[rpgStatInt]),
		"Useless for now"},
	)
	menu.AddMenuItem("Vitality    ", []string{
		fmt.Sprintf("Current: %d", player.playerStats.rpgStats[rpgStatVit]),
		fmt.Sprintf("Max HP: %d", player.getMaxHitpoints()),
	})
	selected := menu.Call(cw)
	switch selected {
	case -1:
		return
	default:
		player.playerStats.rpgStats[selected]++
		player.playerStats.skillPoints--
	}
}

func (pc *playerController) callInventoryMenu() {
	menu := tcell_console_wrapper.DescriptionHeavySelectMenu{Title: "Inventory"}
	menu.AddMenuItem("See the backpack", []string{fmt.Sprintf("%d items inside", len(player.inv.stash)), ""})
	menu.AddMenuItem("Weapon", pc.getItemDescriptionsArrayForInv(player.inv.getItemInSlot(invSlotWeapon), true))
	menu.AddMenuItem("Body armor", pc.getItemDescriptionsArrayForInv(player.inv.getItemInSlot(invSlotBody), true))
	menu.AddMenuItem("Helmet", pc.getItemDescriptionsArrayForInv(player.inv.getItemInSlot(invSlotHelmet), true))
	menu.AddMenuItem("Amulet", pc.getItemDescriptionsArrayForInv(player.inv.getItemInSlot(invSlotAmulet), true))
	menu.AddMenuItem("Flask", pc.getItemDescriptionsArrayForInv(player.inv.getItemInSlot(invSlotFlask), true))
	selected := menu.Call(cw)
	switch selected {
	case -1:
		return
	case 0:
		pc.callBackpackMenu()
	case 1:
		pc.callChangeItemMenu(invSlotWeapon)
	case 2:
		pc.callChangeItemMenu(invSlotBody)
	case 3:
		pc.callChangeItemMenu(invSlotHelmet)
	case 4:
		pc.callChangeItemMenu(invSlotAmulet)
	case 5:
		pc.callChangeItemMenu(invSlotFlask)
	}
}

func (pc *playerController) getItemDescriptionsArrayForInv(i *item, includeName bool) (descs []string) {
	if i == nil {
		descs = append(descs, "(Nothing)")
	} else {
		if includeName {
			descs = append(descs, i.getName())
		}
		descs = append(descs, i.getDescription())
		descs = append(descs, i.getStatic().GetAffixDescriptions()...)
	}
	return
}

func (pc *playerController) callBackpackMenu() {
	menu := tcell_console_wrapper.DescriptionHeavySelectMenu{Title: "Backpack"}
	for _, itm := range player.inv.stash {
		menu.AddMenuItem(itm.getName(), pc.getItemDescriptionsArrayForInv(itm, false))
	}
	selected := menu.Call(cw)
	switch selected {
	case -1:
		return
	}
}

func (pc *playerController) callChangeItemMenu(slot uint8) {
	var title string
	undertitles := []string{"Current: "}
	undertitles = append(undertitles, pc.getItemDescriptionsArrayForInv(player.inv.getItemInSlot(slot), true)...)
	var filter func(*item) bool
	switch slot {
	case invSlotWeapon:
		title = "Select weapon to wield"
		filter = func(i *item) bool { return i.isWeapon() }
	case invSlotBody:
		title = "Select armor to wear"
		filter = func(i *item) bool { return i.isArmor() && i.asArmor.Slot == static.ArmorSlotBody }
	case invSlotHelmet:
		title = "Select headpiece to put on"
		filter = func(i *item) bool { return i.isArmor() && i.asArmor.Slot == static.ArmorSlotHead }
	case invSlotAmulet:
		title = "Select amulet to wear"
		filter = func(i *item) bool { return i.isAmulet() }
	case invSlotFlask:
		title = "Select flask to ready"
		filter = func(i *item) bool { return i.isFlask() }
	default:
		panic("Unimplemented")
	}
	menu := tcell_console_wrapper.DescriptionHeavySelectMenu{
		Title:           title,
		UndertitleLines: undertitles,
	}
	itms := player.inv.selectItemsFromStash(filter)
	for _, itm := range itms {
		menu.AddMenuItem(itm.getName(), pc.getItemDescriptionsArrayForInv(itm, false))
	}
	menu.AddMenuItem("Nothing", []string{})
	index := menu.Call(cw)
	if index == -1 {
		return
	}
	var selectedItem *item = nil
	if index < len(itms) {
		selectedItem = itms[index]
	}
	player.inv.swapItemInSlotWithItemFromStash(slot, selectedItem)
}
