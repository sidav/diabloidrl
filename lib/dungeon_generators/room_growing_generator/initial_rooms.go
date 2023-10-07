package roomgrowinggenerator

func (g *Generator) setInitialRooms() {
	w := len(g.Tiles) / 3
	h := len(g.Tiles[0]) / 3
	g.drawRoom(w, h, w, h)
}
