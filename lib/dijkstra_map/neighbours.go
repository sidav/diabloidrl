package dijkstra_map

type neighbourCoord struct {
	x, y       int
	isDiagonal bool
}

func (c *neighbourCoord) getCostIncreaseFromHere() int {
	if c.isDiagonal {
		return rankDiffDiagonal
	}
	return rankDiffOrthogonal
}

func AllNeighbours(x, y int) []neighbourCoord {
	return []neighbourCoord{
		{x - 1, y, false},
		{x, y - 1, false},
		{x + 1, y, false},
		{x, y + 1, false},
		{x + 1, y - 1, true},
		{x - 1, y - 1, true},
		{x - 1, y + 1, true},
		{x + 1, y + 1, true},
	}
}

func OrthogonalNeighbours(x, y int) []neighbourCoord {
	return []neighbourCoord{
		{x, y - 1, false},
		{x - 1, y, false},
		{x + 1, y, false},
		{x, y + 1, false},
	}
}
