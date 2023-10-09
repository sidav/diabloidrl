package roomgrowinggenerator

import (
	"diabloidrl/lib/random"
)

var rnd random.PRNG

type Generator struct {
	Tiles                    [][]tile
	MinRoomSide, MaxRoomSide int

	placedDoorsBetweenRoomIds [][2]int
	roomsCount                int
}

func (g *Generator) Generate(w, h int, r random.PRNG) {
	g.roomsCount = 0
	g.placedDoorsBetweenRoomIds = make([][2]int, 0)
	rnd = r
	g.Tiles = make([][]tile, w)
	for i := range g.Tiles {
		g.Tiles[i] = make([]tile, h)
	}
	g.setInitialRooms()
	for g.calculateTileFillPercentage(TILE_UNFILLED) > 30 {
		if !tryFuncNTimes(
			func() bool {
				if rnd.Rand(3) == 0 {
					return g.tryPlaceRandomRoom()
				} else {
					return g.tryPlaceRandomVault(false)
				}
			},
			1000,
		) {
			break
		}
	}
	insideVaults := 0
	for insideVaults < 50 {
		if !tryFuncNTimes(func() bool { return g.tryPlaceRandomVault(true) }, 100) {
			break
		}
		insideVaults++
	}
	for doors := 0; doors < g.roomsCount/4; doors++ {
		g.addRandomDoor()
	}
	g.placeEntrypoint()
	if !g.checkConnectivity() {
		g.tileAt(0, 0).Code = TILE_DOOR
	}
}

func tryFuncNTimes(fnc func() bool, times int) bool {
	for i := 0; i < times; i++ {
		if fnc() {
			return true
		}
	}
	return false
}

func (g *Generator) calculateTileFillPercentage(code tileCode) int {
	w, h := len(g.Tiles), len(g.Tiles[0])
	square := w * h
	filledTiles := 0
	for x := range g.Tiles {
		for y := range g.Tiles[x] {
			if g.tileAt(x, y).Code == code {
				filledTiles++
			}
		}
	}
	return (100*(filledTiles) + square - 1) / square
}

func (g *Generator) countTileCodesInPlusShapeAround(x, y int, codeToCount tileCode) int {
	sum := 0
	if g.Tiles[x-1][y].Code == codeToCount {
		sum++
	}
	if g.Tiles[x+1][y].Code == codeToCount {
		sum++
	}
	if g.Tiles[x][y-1].Code == codeToCount {
		sum++
	}
	if g.Tiles[x][y+1].Code == codeToCount {
		sum++
	}
	return sum
}

func (g *Generator) tileAt(x, y int) *tile {
	return &g.Tiles[x][y]
}

func (g *Generator) tileCodeAt(x, y int) tileCode {
	return g.Tiles[x][y].Code
}

func (g *Generator) areCoordsInBounds(x, y int) bool {
	return x >= 0 && x < len(g.Tiles) && y >= 0 && y < len(g.Tiles[0])
}

func (g *Generator) isRectInBounds(x, y, w, h int) bool {
	return x >= 0 && x < len(g.Tiles)-w && y >= 0 && y < len(g.Tiles[0])-h
}

func (g *Generator) doesRectBoundContainOnlyTile(x, y, w, h int, codeToCount tileCode) bool {
	for i := x; i < x+w; i++ {
		if g.tileAt(i, y).Code != codeToCount || g.tileAt(i, y+h-1).Code != codeToCount {
			return false
		}
	}
	for j := y; j < y+h; j++ {
		if g.tileAt(x, j).Code != codeToCount || g.tileAt(x+w-1, j).Code != codeToCount {
			return false
		}
	}
	return true
}

func (g *Generator) countTileCodesInRect(x, y, w, h int, codeToCount tileCode) int {
	count := 0
	for i := x; i < x+w; i++ {
		for j := y; j < y+h; j++ {
			currCode := g.Tiles[i][j].Code
			if currCode == codeToCount {
				count++
			}
		}
	}
	return count
}

func (g *Generator) selectRandomCoordsFromRect(x, y, w, h int, selectionFunc func(int, int) bool) (bool, int, int) {
	cands := make([][2]int, 0)
	for i := x; i < x+w; i++ {
		for j := y; j < y+h; j++ {
			if selectionFunc(i, j) {
				cands = append(cands, [2]int{i, j})
			}
		}
	}
	if len(cands) > 0 {
		index := rnd.Rand(len(cands))
		return true, cands[index][0], cands[index][1]
	}
	return false, 0, 0
}

func (g *Generator) isTileRectOfCode(x, y, w, h int, tcode tileCode, allowEmpty bool) bool {
	requiredFound := false
	for i := x; i < x+w; i++ {
		for j := y; j < y+h; j++ {
			currCode := g.Tiles[i][j].Code
			if currCode == tcode {
				requiredFound = true
			}
			if currCode != tcode && !(allowEmpty && currCode == TILE_UNFILLED) {
				return false
			}
		}
	}
	return requiredFound
}

func (g *Generator) fillTileRect(x, y, w, h int, tcode tileCode, roomId int) {
	for i := x; i < x+w; i++ {
		for j := y; j < y+h; j++ {
			g.Tiles[i][j].Code = tcode
			if roomId != -1 {
				g.Tiles[i][j].roomId = roomId
			}
		}
	}
}
