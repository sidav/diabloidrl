package main

type playerStruct struct {
	experience  int
	skillPoints int
	rpgStats    [rpgStatsCount]int

	// just for UI
	lastActionTicks int
}

func (ps *playerStruct) getExperienceLevel() int {
	for i := range experienceForLevels {
		if ps.experience < experienceForLevels[i] {
			return i + 1
		}
	}
	return len(experienceForLevels) + 1
}

// returns from 0
func (ps *playerStruct) getNormalizedCurrAndRemainingLevelExp() (int, int) {
	lvl := ps.getExperienceLevel()
	if lvl == 1 {
		return ps.experience, experienceForLevels[0]
	} else if lvl <= len(experienceForLevels) {
		return ps.experience - experienceForLevels[lvl-2], experienceForLevels[lvl-1] - experienceForLevels[lvl-2]
	} else {
		return 0, 100000
	}
}

var experienceForLevels = []int{
	5, 11, 20, 35, 60, 100, 161, 249, 371, 534, 747, 1018, 1357, 1774, 2279, 2883,
	3598, 4436, 5410, 6533, 7818, 9279, 10931, 12789, 14869, 17186, 19757, 22599,
	25729, 29165, 32925, 37028, 41493, 46339, 51586, 57254, 63364, 69937, 76994,
	84557, 92649, 101292, 110509, 120324, 130761, 141844, 153597, 166045, 179214,
	193130, 207819, 223307, 239621, 256788, 274836, 293793, 313687, 334547, 356402,
	379281, 403214, 428231, 454362, 481638, 510090, 539749, 570647, 602816, 636288,
	671096, 707273, 744853, 783870, 824358, 866351, 909884, 954992, 1001710, 1050074,
	1100120, 1151884, 1205403, 1260714, 1317854, 1376861, 1437773, 1500628, 1565465,
	1632323, 1701241, 1772258, 1845414, 1920749, 1998304, 2078120, 2160238, 2244699,
	2331545, 2420818, 2512560,
}
