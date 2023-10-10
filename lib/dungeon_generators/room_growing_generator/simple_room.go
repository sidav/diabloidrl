package roomgrowinggenerator

func (g *Generator) tryPlaceRandomRoom() bool {
	w, h := rnd.RandInRange(g.MinRoomSide, g.MaxRoomSide), rnd.RandInRange(g.MinRoomSide, g.MaxRoomSide)
	foundRoom, x, y := g.selectRandomTileCoordsForAppendedRoom(w, h)
	foundDoor, dx, dy := g.selectRandomTileCoordsForRoomsDoor(x-1, y-1, w+2, h+2)
	if !(foundRoom && foundDoor) {
		return false
	}
	g.drawRoom(x-1, y-1, w+2, h+2)
	g.placeDoor(dx, dy)
	return false
}

func (g *Generator) selectRandomTileCoordsForAppendedRoom(w, h int) (bool, int, int) {
	return g.selectRandomCoordsFromRect(1, 1, len(g.tiles)-w-1, len(g.tiles[0])-h-1,
		func(x, y int) bool {
			return g.isTileRectOfCode(x, y, w, h, TILE_UNFILLED, true) &&
				(g.countTileCodesInRect(x-1, y-1, w+2, h+2, TILE_WALL) > w ||
					g.countTileCodesInRect(x-1, y-1, w+2, h+2, TILE_WALL) > h)
		},
	)
}

func (g *Generator) selectRandomTileCoordsForRoomsDoor(x, y, w, h int) (bool, int, int) {
	return g.selectRandomCoordsFromRect(x, y, w, h, func(i, j int) bool {
		// exclude corners of current room
		isCorner := (i == x || i == x+w-1) && (j == y || j == y+h-1)
		return !isCorner && g.isTileGoodForDoor(i, j, false)
	})
}

func (g *Generator) drawRoom(x, y, w, h int) {
	g.fillTileRect(x, y, w, h, TILE_WALL, -1)
	g.fillTileRect(x+1, y+1, w-2, h-2, TILE_FLOOR, g.roomsCount)
	g.roomsCount++
}
