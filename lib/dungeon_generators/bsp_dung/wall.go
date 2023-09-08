package bspdung

// returns true if the wall needa a door
func (g *Generator) drawRandomWall(x1, y1, x2, y2 int) bool {
	wallRune := '#'
	dotted := false
	if rnd.OneChanceFrom(8) {
		wallRune = '"'
	} else if rnd.OneChanceFrom(12) {
		dotted = true
	}
	for x := x1; x <= x2; x++ {
		for y := y1; y <= y2; y++ {
			if dotted && ((x2-x)%2 == 1 || (y2-y)%2 == 1) {
				continue
			}
			g.dung[x][y] = wallRune
		}
	}
	if dotted {
		return false
	}
	return true
}
