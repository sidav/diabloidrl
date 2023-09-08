package static

func selectRandomAppropriateUniqueAffixesFor(stats ItemStatsInterface, affixCount int) []*affixAdder {
	var selectedAffixes []*affixAdder
	try := 0
selectAffix:
	for len(selectedAffixes) < affixCount {
		if try == 100 {
			panic("Can't select more affixes.")
		}
		currAff := selectRandomAppropriateAffixFor(stats)
		for _, a := range selectedAffixes {
			if a == currAff || !a.isCompatibleWith(currAff) {
				try++
				continue selectAffix
			}
		}
		selectedAffixes = append(selectedAffixes, currAff)
	}
	return selectedAffixes
}

func selectRandomAppropriateAffixFor(stats ItemStatsInterface) *affixAdder {
	affixIndex := rnd.SelectRandomIndexFromWeighted(
		len(allAffixes),
		func(i int) int {
			currAff := allAffixes[i]
			switch stats.(type) {
			case *ArmorStats:
				if currAff.armorFunc == nil && currAff.anyFunc == nil {
					return 0
				}
				if currAff.headOnly && stats.(*ArmorStats).Slot != ArmorSlotHead {
					return 0
				}
				if currAff.bodyOnly && stats.(*ArmorStats).Slot != ArmorSlotBody {
					return 0
				}
				return 1
			case *WeaponStats:
				if currAff.weaponFunc == nil && currAff.anyFunc == nil {
					return 0
				}
				if currAff.meleeOnly && stats.(*WeaponStats).Range > 1 {
					return 0
				}
				if currAff.rangedOnly && uint8(stats.(*WeaponStats).Range) < 1 {
					return 0
				}
				return 1
			case *AmuletStats:
				if currAff.anyFunc == nil {
					return 0
				}
				return 1
			default:
				panic("Type check error")
			}
		},
	)
	return allAffixes[affixIndex]
}
