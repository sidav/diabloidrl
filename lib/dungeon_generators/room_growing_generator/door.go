package roomgrowinggenerator

func (g *Generator) placeDoor(x, y int) {
	g.tileAt(x, y).Code = TILE_DOOR
	id1, id2 := g.getRoomIdsNear(x, y)
	if !g.wereRoomsAlreadyConnected(id1, id2) {
		g.placedDoorsBetweenRoomIds = append(g.placedDoorsBetweenRoomIds,
			[2]int{id1, id2})
	}
}

func (g *Generator) getRoomIdsNear(x, y int) (int, int) {
	id1, id2 := -1, -1
	if g.tileAt(x-1, y).isConnective() && g.tileAt(x+1, y).isConnective() {
		id1 = g.tileAt(x-1, y).roomId
		id2 = g.tileAt(x+1, y).roomId
	}
	if id1 == id2 && g.tileAt(x, y-1).isConnective() && g.tileAt(x, y+1).isConnective() {
		id1 = g.tileAt(x, y-1).roomId
		id2 = g.tileAt(x, y+1).roomId
	}
	// if id1 == id2 {
	// 	panic(fmt.Sprintf("Wat was connected?! IDs are %d and %d, LRI is %d", id1, id2, g.roomsCount))
	// }
	return id1, id2
}

func (g *Generator) wereRoomsAlreadyConnected(id1, id2 int) bool {
	for _, arr := range g.placedDoorsBetweenRoomIds {
		if arr[0] == id1 && arr[1] == id2 || arr[1] == id1 && arr[0] == id2 {
			return true
		}
	}
	return false
}

func (g *Generator) isTileGoodForDoor(x, y int, shouldConnectRooms bool) bool {
	if x <= 0 || x >= len(g.tiles)-1 || y <= 0 || y >= len(g.tiles[x])-1 {
		return false
	}
	if !(g.areCoordsInBounds(x, y) && g.tileAt(x, y).Code == TILE_WALL) {
		return false
	}
	if g.countTileCodesInPlusShapeAround(x, y, TILE_DOOR) > 0 {
		return false
	}
	if g.countTileCodesInPlusShapeAround(x, y, TILE_WALL) > 2 {
		return false
	}

	// check if this tile is not on a L-shaped corner
	isNonCornerTile := g.tileCodeAt(x-1, y) == TILE_WALL && g.tileCodeAt(x+1, y) == TILE_WALL ||
		g.tileCodeAt(x, y-1) == TILE_WALL && g.tileCodeAt(x, y+1) == TILE_WALL
	// check if there are floor tiles on front and back in any non-diagonal direction
	connectsRooms := g.tileCodeAt(x-1, y) == TILE_FLOOR && g.tileCodeAt(x+1, y) == TILE_FLOOR ||
		g.tileCodeAt(x, y-1) == TILE_FLOOR && g.tileCodeAt(x, y+1) == TILE_FLOOR
	return isNonCornerTile && (connectsRooms || !shouldConnectRooms)
}

func (g *Generator) addRandomDoor() {
	placed, dx, dy := g.selectRandomCoordsFromRect(1, 1, len(g.tiles)-2, len(g.tiles[0])-2,
		func(x, y int) bool {
			if g.isTileGoodForDoor(x, y, true) {
				id1, id2 := g.getRoomIdsNear(x, y)
				if id1 == id2 || g.wereRoomsAlreadyConnected(id1, id2) {
					return false
				}
				return true
			}
			return false
		},
	)
	if placed {
		g.placeDoor(dx, dy)
	}
}
