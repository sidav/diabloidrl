package roomgrowinggenerator

func (g *Generator) placeRandomRoom() {
	for try := 0; try < 100; try++ {
		w, h := rnd.RandInRange(g.MinRoomSide, 10), rnd.RandInRange(g.MinRoomSide, 10)
		foundRoom, x, y := g.selectRandomTileCoordsForAppendedRoom(w, h)
		foundDoor, dx, dy := g.selectRandomTileCoordsForDoor(x-1, y-1, w+2, h+2)
		if !(foundRoom && foundDoor) {
			continue
		}
		g.drawRoom(x-1, y-1, w+2, h+2)
		g.Tiles[dx][dy].Code = TILE_DOOR
		break
	}
}

func (g *Generator) selectRandomTileCoordsForAppendedRoom(w, h int) (bool, int, int) {
	cands := make([][2]int, 0)
	for x := w + 1; x < len(g.Tiles)-w-1; x++ {
		for y := h + 1; y < len(g.Tiles[x])-h-1; y++ {
			if g.isTileRectOfCode(x, y, w, h, TILE_UNFILLED, true) {
				if g.countTileCodesInRect(x-1, y-1, w+2, h+2, TILE_WALL) > w {
					cands = append(cands, [2]int{x, y})
				}
			}
		}
	}
	if len(cands) > 0 {
		index := rnd.Rand(len(cands))
		return true, cands[index][0], cands[index][1]
	} else {
		return false, 0, 0
	}
}

func (g *Generator) selectRandomTileCoordsForDoor(x, y, w, h int) (bool, int, int) {
	cands := make([][2]int, 0)
	for i := x; i < x+w; i++ {
		// exclude corners of current room
		cornerSkipModifier := 0
		if i == x || i == x+w-1 {
			cornerSkipModifier = 1
		}
		for j := y + cornerSkipModifier; j < y+h-cornerSkipModifier; j++ {
			if g.isTileGoodForDoor(i, j, false) {
				cands = append(cands, [2]int{i, j})
			}
		}
	}
	if len(cands) > 0 {
		index := rnd.Rand(len(cands))
		return true, cands[index][0], cands[index][1]
	} else {
		return false, 0, 0
	}
}

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

func (g *Generator) drawRoom(x, y, w, h int) {
	g.fillTileRect(x, y, w, h, TILE_WALL)
	g.fillTileRect(x+1, y+1, w-2, h-2, TILE_FLOOR)
}
