package roomgrowinggenerator

func (g *Generator) placeEntrypoint() {
	placed, dx, dy := g.selectRandomCoordsFromRect(1, 1, len(g.tiles)-2, len(g.tiles[0])-2,
		func(x, y int) bool {
			return g.countTileCodesInPlusShapeAround(x, y, TILE_FLOOR) > 2
		})
	if placed {
		g.tiles[dx][dy].Code = TILE_ENTRYPOINT
		return
	}
	panic("Stairs can't be placed!")
}
