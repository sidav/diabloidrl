package roomgrowinggenerator

func (g *Generator) isTileGoodForDoor(x, y int, shouldConnectRooms bool) bool {
	if x <= 0 || x >= len(g.Tiles)-1 || y <= 0 || y >= len(g.Tiles[x])-1 {
		return false
	}
	if !(g.areCoordsInBounds(x, y) && g.tileAt(x, y).Code == TILE_WALL) {
		return false
	}
	if g.countTileCodesAround(x, y, TILE_WALL) != 3 {
		return false
	}
	// check if this tile is not on a corner
	isNonCornerTile := g.tileCodeAt(x-1, y) == TILE_WALL && g.tileCodeAt(x+1, y) == TILE_WALL ||
		g.tileCodeAt(x, y-1) == TILE_WALL && g.tileCodeAt(x, y+1) == TILE_WALL
	connectsRooms := g.tileCodeAt(x-1, y) == TILE_FLOOR && g.tileCodeAt(x+1, y) == TILE_FLOOR ||
		g.tileCodeAt(x, y-1) == TILE_FLOOR && g.tileCodeAt(x, y+1) == TILE_FLOOR
	return isNonCornerTile && (connectsRooms || !shouldConnectRooms)
}

func (g *Generator) addRandomDoor() {
	cands := make([][2]int, 0)
	for x := g.MinRoomSide + 1; x < len(g.Tiles)-g.MinRoomSide; x++ {
		for y := g.MinRoomSide + 1; y < len(g.Tiles[x])-g.MinRoomSide; y++ {
			if g.isTileGoodForDoor(x, y, true) {
				cands = append(cands, [2]int{x, y})
			}
		}
	}
	if len(cands) > 0 {
		index := rnd.Rand(len(cands))
		x, y := cands[index][0], cands[index][1]
		g.tileAt(x, y).Code = TILE_DOOR
	}
}
