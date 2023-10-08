package roomgrowinggenerator

import (
	"diabloidrl/lib/random"
)

var rnd random.PRNG

type Generator struct {
	Tiles       [][]tile
	MinRoomSide int

	placedDoorsBetweenRoomIds [][2]int
	roomsCount                int
}

func (g *Generator) Init() {
	// makeVaultsVarians()
	// for i := range allVaults {
	// 	g.dbgShowVault(allVaults[i])
	// 	g.dbgFlush()
	// }
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
	for rooms := 0; rooms < 25; rooms++ {
		if rnd.Rand(2) > 0 {
			g.placeRandomRoom()
		} else {
			g.placeRandomVault(false)
		}
	}
	for vaults := 0; vaults < 10; vaults++ {
		g.placeRandomVault(true)
	}
	for doors := 0; doors < g.roomsCount/4; doors++ {
		g.addRandomDoor()
	}
	if !g.checkConnectivity() {
		g.tileAt(0, 0).Code = TILE_DOOR
	}
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
