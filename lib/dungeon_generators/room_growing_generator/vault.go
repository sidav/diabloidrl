package roomgrowinggenerator

func (g *Generator) placeRandomVault(inside bool) {
	for try := 0; try < 10; try++ {
		var randomVault []string
		if inside {
			randomVault = insideVaults[rnd.Rand(len(insideVaults))]
		} else {
			randomVault = outsideVaults[rnd.Rand(len(outsideVaults))]
		}
		randomVault = makeRandomTransofrmationForVault(randomVault)
		// g.dbgShowVault(randomVault)
		// g.dbgFlush()

		place, x, y := g.selectCoordsToPlaceVault(randomVault, inside)
		if !place {
			continue
		}
		// g.dbgDrawCurrentState(false)
		// g.dbgHighlightTile(x, y)
		// g.dbgFlush()

		if inside {
			g.placeInsideVaultAt(randomVault, x, y)
		} else {
			g.placeOutsideVaultAt(randomVault, x, y)
		}

		// g.dbgDrawCurrentState(false)
		// g.dbgHighlightTile(x, y)
		// g.dbgFlush()

		return
	}
}

func (g *Generator) placeInsideVaultAt(v []string, x, y int) {
	roomPlaced := false
	for i := 0; i < len(v); i++ {
		for j := 0; j < len(v[i]); j++ {
			g.tileAt(x+i, y+j).setByVaultChar(rune(v[i][j]))
			if rune(v[i][j]) == '.' {
				g.tileAt(x+i, y+j).roomId = g.roomsCount
				roomPlaced = true
			}
		}
	}
	for i := 0; i < len(v); i++ {
		for j := 0; j < len(v[i]); j++ {
			if rune(v[i][j]) == '+' {
				g.placeDoor(x+i, y+j)
			}
		}
	}
	if roomPlaced {
		g.roomsCount++
	}
}

func (g *Generator) placeOutsideVaultAt(v []string, x, y int) {
	roomPlaced := false
	for i := 0; i < len(v); i++ {
		for j := 0; j < len(v[i]); j++ {
			rx, ry := x+i, y+j
			g.tileAt(rx, ry).setByVaultChar(rune(v[i][j]))
			if rune(v[i][j]) == '.' {
				g.tileAt(rx, ry).roomId = g.roomsCount
				roomPlaced = true
			}
		}
	}
	for i := 0; i < len(v); i++ {
		for j := 0; j < len(v[i]); j++ {
			if rune(v[i][j]) == '+' {
				g.placeDoor(x+i, y+j)
			}
		}
	}
	if roomPlaced {
		g.roomsCount++
	}
}

// Outside walls are FORCED to be connected by a door.
var outsideVaults = [][]string{
	{
		"##########",
		"#........+",
		"##########",
	},
	{
		"#####    ",
		"#...#    ",
		"#...#    ",
		"#...#####",
		"#.......#",
		"#.......+",
		"#########",
	},
	{
		"  #####  ",
		" ##...## ",
		"##.....##",
		"#.......+",
		"##.....##",
		" ##...## ",
		"  #####  ",
	},
}
var insideVaults = [][]string{
	{
		" ... ",
		"..#..",
		".###.",
		"..#..",
		" ... ",
	},
	{
		"........",
		".######.",
		".#....#.",
		".#.##.#.",
		".#.##.+.",
		".#....#.",
		".######.",
		"........",
	},
}
