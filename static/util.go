package static

import (
	"diabloidrl/lib/random"
)

var rnd random.PRNG

func SetRandom(r random.PRNG) {
	rnd = r
}

func intPercentage(base, percentage int) int {
	// integer rounded divide for A/B is (A + B/2) / B
	return (base*percentage + 50) / 100
}
