package roomgrowinggenerator

func (g *Generator) selectCoordsToPlaceVault(v []string, inside bool) (bool, int, int) {
	cands := make([][2]int, 0)
	for x := 0; x < len(g.Tiles)-len(v); x++ {
		for y := 0; y < len(g.Tiles[x])-len(v[0]); y++ {
			if inside && g.canVaultBePlacedInsideAt(v, x, y) {
				cands = append(cands, [2]int{x, y})
			}
			if !inside && g.canVaultBePlacedOutsideAt(v, x, y) {
				cands = append(cands, [2]int{x, y})
			}
		}
	}
	if len(cands) == 0 {
		return false, 0, 0
	}
	index := rnd.Rand(len(cands))
	return true, cands[index][0], cands[index][1]
}

func (g *Generator) canVaultBePlacedInsideAt(v []string, vx, vy int) bool {
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
			case '#':
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
	wallIntersections := 0
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
			case '#':
				good = currentCode == TILE_WALL || currentCode == TILE_UNFILLED
				if currentCode == TILE_WALL {
					wallIntersections++
				}
			case '+':
				good = currentCode == TILE_DOOR || currentCode == TILE_FLOOR || currentCode == TILE_UNFILLED //  && g.isTileGoodForDoor(x, y, true)
			case '.':
				good = currentCode == TILE_FLOOR || currentCode == TILE_UNFILLED
			}
			if !good {
				// g.dbgDrawCurrentState(false)
				// g.dbgShowVault(v, i, j)
				// g.dbgHighlightTile(vx, vy)
				// g.dbgHighlightTileWithComment(x, y, "Tile '%v' rejected", string(v[i][j]))
				// g.dbgFlush()
				return false
			} else {
				// g.dbgDrawCurrentState(false)
				// g.dbgShowVault(v, i, j)
				// g.dbgHighlightTile(vx, vy)
				// g.dbgHighlightTileWithComment(x, y, "Tile '%v' passes", string(v[i][j]))
				// g.dbgFlush()
			}
		}
	}
	return true && wallIntersections > 0
}
