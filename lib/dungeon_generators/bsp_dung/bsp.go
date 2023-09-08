package bspdung

import (
	"diabloidrl/lib/random"
	"diabloidrl/lib/tcell_console_wrapper"
)

var rnd random.PRNG

type Generator struct {
	dung                   [][]rune
	MinRoomSide            int
	RoomWForRandomDropping int
	RoomHForRandomDropping int
	MaxRoomRatio           float64

	Cw *tcell_console_wrapper.ConsoleWrapper // TODO: remove
}

func (g *Generator) Generate(r random.PRNG, w, h int) [][]rune {
	rnd = r
	//
	g.MinRoomSide = 5
	g.RoomWForRandomDropping = w / 2
	g.RoomHForRandomDropping = h / 2
	g.MaxRoomRatio = 0.65

	g.dung = make([][]rune, w)
	for i := range g.dung {
		g.dung[i] = make([]rune, h)
	}
	g.fill(0, 0, w, h, '#')
	g.fill(1, 1, w-2, h-2, ' ')
	g.splitSubDungeonRecursively(1, 1, w-2, h-2, rnd.OneChanceFrom(2), -1, -1)
	g.placeEntrypoint()
	g.drawCurrentState()
	return g.dung
}

func (g *Generator) splitSubDungeonRecursively(x, y, w, h int, preferHorizontal bool, restrictX, restrictY int) {
	// fmt.Printf("CALLED WITH %d,%d,%d,%d --- ", x, y, w, h)
	if g.shouldSplittingBreak(w, h) {
		return
	}
	g.drawCurrentState()
	g.drawCurrentLeaf(x, y, w, h)
	if rnd.OneChanceFrom(4) {
		preferHorizontal = !preferHorizontal
	}
	if w < 2*g.MinRoomSide || float64(w)/float64(h) < g.MaxRoomRatio {
		preferHorizontal = true
	} else if h < 2*g.MinRoomSide || float64(h)/float64(w) < g.MaxRoomRatio {
		preferHorizontal = false
	}

	if preferHorizontal {
		splitHeight := minForcedRandInRange(g.MinRoomSide, h-g.MinRoomSide)
		if y+splitHeight == restrictY {
			return
		}
		// fmt.Printf("DECISION: HOR, sY %d \n", splitY)
		// draw wall
		doorNeeded := g.drawRandomWall(x, y+splitHeight, x+w-1, y+splitHeight)
		g.drawCurrentFill(x, y+splitHeight, w, 1, "SplitHeight", splitHeight)
		// door
		doorX, doorY := 0, 0
		if doorNeeded {
			doorX, doorY = g.addDoorOnLine(x, y+splitHeight, x+w-1, y+splitHeight)
		}
		// up
		g.splitSubDungeonRecursively(x, y, w, splitHeight, !preferHorizontal, doorX, doorY)
		// bottom
		g.splitSubDungeonRecursively(x, y+splitHeight+1, w, h-splitHeight-1, !preferHorizontal, doorX, doorY)
	} else {
		splitWidth := minForcedRandInRange(g.MinRoomSide, w-g.MinRoomSide)
		if x+splitWidth == restrictX {
			return
		}
		// fmt.Printf("DECISION: VER, sX %d \n", splitX)
		doorNeeded := g.drawRandomWall(x+splitWidth, y, x+splitWidth, y+h-1)
		g.drawCurrentFill(x+splitWidth, y, 1, h, "SplitWidth", splitWidth)
		// door
		doorX, doorY := 0, 0
		if doorNeeded {
			doorX, doorY = g.addDoorOnLine(x+splitWidth, y, x+splitWidth, y+h-1)
		}
		// left
		g.splitSubDungeonRecursively(x, y, splitWidth, h, !preferHorizontal, doorX, doorY)
		// right
		g.splitSubDungeonRecursively(x+splitWidth+1, y, w-splitWidth-1, h, !preferHorizontal, doorX, doorY)
	}
}

func (g *Generator) placeEntrypoint() {
	x, y := 0, 0
	for g.dung[x][y] != ' ' {
		x, y = rnd.RandInRange(1, len(g.dung)-2), rnd.RandInRange(1, len(g.dung[0])-2)
	}
	g.dung[x][y] = '<'
}

func (g *Generator) shouldSplittingBreak(w, h int) bool {
	if w <= g.RoomWForRandomDropping && h <= g.RoomHForRandomDropping && rnd.OneChanceFrom(5) {
		return true
	}

	return w <= 2*g.MinRoomSide && h <= 2*g.MinRoomSide
}

func (g *Generator) addDoorOnLine(x1, y1, x2, y2 int) (int, int) {
	var dx, dy int
	if y1 == y2 {
		dx, dy = rnd.RandInRange(x1+1, x2-1), y1
	} else {
		dx, dy = x1, rnd.RandInRange(y1+1, y2-1)
	}
	g.dung[dx][dy] = '+'
	return dx, dy
}
