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

		place, x, y := g.selectCoordsToPlaceVault(randomVault, inside)
		if !place {
			continue
		}

		g.placeVaultAt(randomVault, x, y)

		return
	}
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

func doesVaultContainDoor(v []string) bool {
	for x := range v {
		for y := range v[x] {
			if v[x][y] == '+' {
				return true
			}
		}
	}
	return false
}

// Outside vaults SHOULD have at least one outer door.
var outsideVaults = [][]string{
	{
		"##########",
		"#........+",
		"#.########",
		"#........#",
		"########.#",
		"#........#",
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
	{
		"   ####   ",
		"   #..#   ",
		"   #..#   ",
		"####..####",
		"#........+",
		"#........+",
		"####..####",
		"   #..#   ",
		"   #..#   ",
		"   ####   ",
	},
}

// Inside vaults SHOULD be surrounded by floors. TODO: rework this?..
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
	{
		"........",
		".######.",
		".#..#...",
		".#..#...",
		".#......",
		".#..#...",
		".#..#...",
		".######.",
		"........",
	},
}
