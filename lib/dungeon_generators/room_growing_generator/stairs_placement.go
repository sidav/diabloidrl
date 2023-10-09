package roomgrowinggenerator

func (g *Generator) placeEntrypoint() {
	placed, dx, dy := g.selectRandomCoordsFromRect(1, 1, len(g.Tiles)-2, len(g.Tiles[0])-2,
		func(x, y int) bool {
			return g.countTileCodesInPlusShapeAround(x, y, TILE_FLOOR) > 2
		})
	if placed {
		g.Tiles[dx][dy].Code = TILE_ENTRYPOINT
		return
	}
	panic("Stairs can't be placed!")
}
