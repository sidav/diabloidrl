package astar

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func areCoordsInRect(x, y, rx, ry, w, h int) bool {
	if w == 0 {
		w = 1
	}
	if h == 0 {
		h = 1
	}
	return x >= rx && x < rx+w && y >= ry && y < ry+h
}

func getCellWithCoordsFromList(list *[]*Cell, x, y int) *Cell {
	for _, c := range *list {
		if c.X == x && c.Y == y {
			return c
		}
	}
	return nil
}

func getCellWithLowestHeuristicFromList(list *[]*Cell) *Cell {
	lowest := (*list)[0]
	for _, c := range *list {
		if c.h < lowest.h {
			lowest = c
		}
	}
	return lowest
}
