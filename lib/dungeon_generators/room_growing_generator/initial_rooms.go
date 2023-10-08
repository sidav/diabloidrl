package roomgrowinggenerator

func (g *Generator) setInitialRooms() {
	// place vault
	if rnd.Rand(2) == 0 {
		vlt := initialVaults[rnd.Rand(len(initialVaults))]
		vlt = makeRandomTransofrmationForVault(vlt)
		w, h := len(vlt), len(vlt[0])
		g.placeVaultAt(vlt, len(g.Tiles)/2-w/2, len(g.Tiles[0])/2-h/2)
		return
	}
	// place other room
	initialSetters := []func(){
		// room in a center
		func() {
			w := len(g.Tiles) / 3
			h := len(g.Tiles[0]) / 3
			g.drawRoom(w, h, w, h)
		},
		// room the size of whole level
		func() {
			w := len(g.Tiles)
			h := len(g.Tiles[0])
			g.drawRoom(0, 0, w, h)
		},
		// narrow room the width of whole level
		func() {
			w := 4 * len(g.Tiles) / 5
			x := (len(g.Tiles) - w) / 2
			h := len(g.Tiles[0]) / 4
			g.drawRoom(x, len(g.Tiles[0])/2-h/2, w, h)
		},
	}
	index := rnd.Rand(len(initialSetters))
	initialSetters[index]()
}
