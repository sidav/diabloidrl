package main

import (
	"diabloidrl/static"
	"math"
)

const (
	invSlotWeapon uint8 = iota
	invSlotBody
	invSlotHelmet
	invSlotAmulet
	invSlotFlask
	invSlotsTotal
)

type inventory struct {
	equipped [invSlotsTotal]*item
	stash    []*item
}

func (inv *inventory) getSumOfEgoValuesOfEquippedItems(code static.EgoCodeType) int {
	sum := 0
	for _, itm := range inv.equipped {
		if itm == nil {
			continue
		}
		for _, ego := range itm.getStatic().GetEgos() {
			if ego.Code == code {
				sum += int(ego.Value)
			}
		}
	}
	return sum
}

func (inv *inventory) getMultiplicativePercentOfEgoValuesOfEquippedItems(code static.EgoCodeType) int {
	mul := 100.0
	for _, itm := range inv.equipped {
		if itm == nil {
			continue
		}
		for _, ego := range itm.getStatic().GetEgos() {
			if ego.Code == code {
				mul *= (float64(ego.Value) / 100.0)
			}
		}
	}
	return int(math.Round(mul))
}

func (inv *inventory) getItemInSlot(slot uint8) *item {
	if inv == nil {
		return nil
	}
	return inv.equipped[slot]
}

func (inv *inventory) putItemInSlot(slot uint8, i *item) {
	if inv.equipped[slot] != nil {
		panic("Item overwrite!")
	}
	inv.equipped[slot] = i
}

func (inv *inventory) selectItemsFromStash(filter func(*item) bool) (itms []*item) {
	for _, i := range inv.stash {
		if filter(i) {
			itms = append(itms, i)
		}
	}
	return
}

func (inv *inventory) addItemToStash(i *item) {
	if i == nil {
		panic("Nil item!")
	}
	inv.stash = append(inv.stash, i)
}

func (inv *inventory) removeItemFromStash(itm *item) {
	for i := range inv.stash {
		if inv.stash[i] == itm {
			inv.stash = append(inv.stash[:i], inv.stash[i+1:]...)
			return
		}
	}
	panic("No such item!")
}

func (inv *inventory) swapItemInSlotWithItemFromStash(slot uint8, itemFromStash *item) {
	if itemFromStash != nil {
		if inv.getItemInSlot(slot) != nil {
			inv.addItemToStash(inv.getItemInSlot(slot))
			inv.equipped[slot] = nil
		}
		inv.putItemInSlot(slot, itemFromStash)
		player.inv.removeItemFromStash(itemFromStash)
	} else {
		if inv.getItemInSlot(slot) != nil {
			player.inv.addItemToStash(inv.getItemInSlot(slot))
			inv.equipped[slot] = nil
		}
		inv.putItemInSlot(slot, nil)
	}
}
