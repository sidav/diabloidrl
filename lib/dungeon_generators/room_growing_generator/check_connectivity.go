package roomgrowinggenerator

func (g *Generator) checkConnectivity() bool {
selectAnyPassableTile:
	for x := range g.tiles {
		for y := range g.tiles[x] {
			if g.tileAt(x, y).isConnective() {
				g.tileAt(x, y).Connected = true
				break selectAnyPassableTile
			}
		}
	}
	iterationSuccess := true
	for iterationSuccess {
		iterationSuccess = g.iterateConnectivity()
	}
	return !g.anyUncheckedTilesPresent()
}

func (g *Generator) iterateConnectivity() bool {
	change := false
	for x := 1; x < len(g.tiles)-1; x++ {
		for y := 1; y < len(g.tiles[x])-1; y++ {
			if g.tileAt(x, y).isConnective() && !g.tileAt(x, y).Connected {
				if g.tileAt(x-1, y).Connected || g.tileAt(x+1, y).Connected || g.tileAt(x, y-1).Connected || g.tileAt(x, y+1).Connected {
					g.tileAt(x, y).Connected = true
					change = true
				}
			}
		}
	}
	return change
}

func (g *Generator) anyUncheckedTilesPresent() bool {
	for x := range g.tiles {
		for y := range g.tiles[x] {
			if !g.tileAt(x, y).Connected && g.tileAt(x, y).isConnective() {
				return true
			}
		}
	}
	return false
}
