package roomgrowinggenerator

func (g *Generator) placeRandomVault() {
	for try := 0; try < len(allVaults); try++ {
		randomVault := allVaults[rnd.Rand(len(allVaults))]
		place, x, y := g.selectCoordsToPlaceVault(randomVault)
		if !place {
			continue
		}
		g.placeVaultAt(randomVault, x, y)
	}
}

func (g *Generator) placeVaultAt(v []string, x, y int) {
	for i := 0; i < len(v); i++ {
		for j := 0; j < len(v[i]); j++ {
			g.tileAt(x+i, y+j).setByVaultChar(rune(v[i][j]))
		}
	}
}

func (g *Generator) selectCoordsToPlaceVault(v []string) (bool, int, int) {
	cands := make([][2]int, 0)
	for x := len(v); x < len(g.Tiles)-len(v); x++ {
		for y := len(v[0]); y < len(g.Tiles[x])-len(v[0]); y++ {
			if g.canVaultBePlacedAt(v, x, y) {
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

func (g *Generator) canVaultBePlacedAt(v []string, vx, vy int) bool {
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
				good = true || currentCode != TILE_DOOR
			case '#':
				good = currentCode == TILE_WALL || currentCode == TILE_FLOOR || currentCode == TILE_UNFILLED
			case '+':
				good = (currentCode == TILE_WALL || currentCode == TILE_DOOR) && g.isTileGoodForDoor(x, y, true)
			case '.':
				good = currentCode == TILE_FLOOR || currentCode == TILE_UNFILLED
			}
			if !good {
				return false
			}
		}
	}
	return true
}

func makeVaultsVarians() {
	newVaults := make([][]string, 0)
	for _, v := range allVaults {
		mirrorX := []string{}
		for x := range v {
			mirrorX = append(mirrorX, invertString(v[x]))
		}
		mirrorY := []string{}
		for x := range v {
			mirrorY = append(mirrorY, v[len(v)-1-x])
		}
		newVaults = append(newVaults, mirrorX, mirrorY)
	}
	allVaults = append(allVaults, newVaults...)
}

func invertString(s string) string {
	newS := make([]byte, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		newS = append(newS, s[i])
	}
	return string(newS)
}

var allVaults = [][]string{
	{
		"        ",
		" ###### ",
		" #....# ",
		" #....# ",
		" #....+ ",
		" #....# ",
		" ###### ",
		"        ",
	},
	{
		"        ",
		" ###### ",
		" #....# ",
		" #.##.# ",
		" #.##.+ ",
		" #....# ",
		" ###### ",
		"        ",
	},
	{
		"##+##    ",
		"#...#    ",
		"#...+    ",
		"#...#####",
		"#.......#",
		"#.......+",
		"#########",
	},
	{
		"           ",
		" ######### ",
		" ###...### ",
		" ##.....## ",
		" +.......+ ",
		" ##.....## ",
		" ###...### ",
		" ######### ",
		"           ",
	},
}
