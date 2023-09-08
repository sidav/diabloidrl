package bspdung

func minForcedRandInRange(min, max int) int {
	max -= 1
	if max < min {
		return min
	}
	return rnd.RandInRange(min, max)
}

func (g *Generator) fill(x, y, w, h int, chr rune) {
	for i := x; i < x+w; i++ {
		for j := y; j < y+h; j++ {
			g.dung[i][j] = chr
		}
	}
}
