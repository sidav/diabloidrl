package roomgrowinggenerator

func (g *Generator) tryPlaceRandomVault(inside bool) bool {
	var randomVault []string
	if inside {
		randomVault = insideVaults[rnd.Rand(len(insideVaults))]
	} else {
		randomVault = outsideVaults[rnd.Rand(len(outsideVaults))]
	}
	randomVault = makeRandomTransofrmationForVault(randomVault)
	place, x, y := g.selectCoordsToPlaceVault(randomVault, inside)
	if !place {
		return false
	}
	g.placeVaultAt(randomVault, x, y)
	return true
}

func (g *Generator) placeVaultAt(v []string, x, y int) {
	updateId := doesVaultContainDoor(v)
	roomPlaced := false
	for i := 0; i < len(v); i++ {
		for j := 0; j < len(v[i]); j++ {
			rx, ry := x+i, y+j
			g.tileAt(rx, ry).setByVaultChar(rune(v[i][j]))
			if updateId && rune(v[i][j]) == '.' {
				g.tileAt(rx, ry).roomId = g.roomsCount
				roomPlaced = true
			}
		}
	}
	if updateId {
		for i := 0; i < len(v); i++ {
			for j := 0; j < len(v[i]); j++ {
				if rune(v[i][j]) == '+' {
					g.placeDoor(x+i, y+j)
				}
			}
		}
	}
	if roomPlaced {
		g.roomsCount++
	}
}

func (g *Generator) selectCoordsToPlaceVault(v []string, inside bool) (bool, int, int) {
	return g.selectRandomCoordsFromRect(0, 0, len(g.Tiles)-len(v), len(g.Tiles[0])-len(v[0]),
		func(x, y int) bool {
			if inside {
				return g.canVaultBePlacedInsideAt(v, x, y)
			} else {
				return g.canVaultBePlacedOutsideAt(v, x, y)
			}
		},
	)
}

func (g *Generator) canVaultBePlacedInsideAt(v []string, vx, vy int) bool {
	if !g.isRectInBounds(vx-1, vy-1, len(v)+2, len(v[0])+2) {
		return false
	}
	if !g.doesRectBoundContainOnlyTile(vx-1, vy-1, len(v)+2, len(v[0])+2, TILE_FLOOR) {
		return false
	}
	for i := 0; i < len(v); i++ {
		for j := 0; j < len(v[i]); j++ {
			x, y := vx+i, vy+j
			if !g.areCoordsInBounds(x, y) {
				return false
			}
			currentCode := g.tileCodeAt(x, y)
			good := false
			switch rune(v[i][j]) {
			case ' ':
				good = currentCode != TILE_DOOR
			case '#', '\'':
				good = currentCode == TILE_WALL || currentCode == TILE_FLOOR
			case '+':
				good = currentCode == TILE_DOOR || currentCode == TILE_FLOOR //  && g.isTileGoodForDoor(x, y, true)
			case '.':
				good = currentCode == TILE_FLOOR
			}
			if !good {
				return false
			}
		}
	}
	return true
}

func (g *Generator) canVaultBePlacedOutsideAt(v []string, vx, vy int) bool {
	doorIntersections := 0
	for i := 0; i < len(v); i++ {
		for j := 0; j < len(v[i]); j++ {
			x, y := vx+i, vy+j
			if !g.areCoordsInBounds(x, y) {
				return false
			}
			currentCode := g.tileCodeAt(x, y)
			good := false
			switch rune(v[i][j]) {
			case ' ':
				good = currentCode != TILE_DOOR
			case '#', '\'':
				good = currentCode == TILE_WALL || currentCode == TILE_UNFILLED
			case '+':
				good = currentCode == TILE_DOOR || currentCode == TILE_UNFILLED //  && g.isTileGoodForDoor(x, y, true)
				if currentCode == TILE_WALL && g.isTileGoodForDoor(x, y, false) {
					good = true
					doorIntersections++
				}
			case '.':
				good = currentCode == TILE_FLOOR || currentCode == TILE_UNFILLED
			}
			if !good {
				return false
			}
		}
	}
	return true && doorIntersections > 0
}
