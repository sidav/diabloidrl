package static

import "fmt"

func PerformAffixSanityCheck() {
	for _, a1 := range allAffixes {
		if a1.incompatibleWithAffixOfAdjective != "" {
			crash := true
			for _, a2 := range allAffixes {
				if !a1.isCompatibleWith(a2) {
					if a2.isCompatibleWith(a1) {
						panic(fmt.Sprintf("Affix '%s' is incompatible with '%s', but '%s' seems to be compatible with '%s'",
							a1.affixAdjective, a2.affixAdjective, a2.affixAdjective, a1.affixAdjective))
					}
					crash = false
					break
				}
			}
			if crash {
				panic(fmt.Sprintf("Affix '%s' is incompatible with '%s', but no such affix exists",
					a1.affixAdjective, a1.incompatibleWithAffixOfAdjective))
			}
		}
	}
}
