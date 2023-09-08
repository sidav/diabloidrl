package dijkstra_map

import "math"

const (
	maxRank            = math.MaxInt16
	rankDiffOrthogonal = 10
	rankDiffDiagonal   = 14
)

type DijkstraMap struct {
	cells            [][]cell
	isCellPassable   func(int, int) bool
	neighbouringFunc func(int, int) []neighbourCoord
}

func New(w, h int, neigbourFunc func(int, int) []neighbourCoord, passabilityFunc func(int, int) bool) *DijkstraMap {
	dm := &DijkstraMap{}
	dm.cells = make([][]cell, w)
	for i := range dm.cells {
		dm.cells[i] = make([]cell, h)
	}
	dm.neighbouringFunc = neigbourFunc
	dm.isCellPassable = passabilityFunc
	return dm
}

func (dm *DijkstraMap) Purge() {
	for x := range dm.cells {
		for y := range dm.cells[x] {
			dm.cells[x][y].isOrigin = false
			dm.cells[x][y].isTarget = false
			dm.cells[x][y].rank = math.MaxUint16
		}
	}
}

func (dm *DijkstraMap) clear() {
	for x := range dm.cells {
		for y := range dm.cells[x] {
			if !dm.cells[x][y].isTarget {
				dm.cells[x][y].rank = math.MaxUint16
			}
		}
	}
}

func (dm *DijkstraMap) GetRankAt(x, y int) int {
	return dm.cells[x][y].rank
}

func (dm *DijkstraMap) GetVectorToBestNeighbourFrom(x, y int) (int, int) {
	if dm.cells[x][y].isTarget {
		return 0, 0
	}
	n := dm.bestNeighbourTo(x, y)
	if dm.GetRankAt(x, y) == dm.GetRankAt(n.x, n.y) {
		return 0, 0
	}
	return n.x - x, n.y - y
}

func (dm *DijkstraMap) SetTarget(x, y, additionalPriority int) {
	dm.cells[x][y].isTarget = true
	dm.cells[x][y].rank = 0 - additionalPriority
}

func (dm *DijkstraMap) RemoveTarget(x, y int) {
	dm.cells[x][y].isTarget = false
}

func (dm *DijkstraMap) Calculate() {
	dm.clear()
	calcContinues := true
	for calcContinues {
		calcContinues = false
		for x := range dm.cells {
			for y := range dm.cells[x] {
				calcContinues = dm.tryRefreshRankAt(x, y) || calcContinues
				calcContinues = dm.tryRefreshRankAt(len(dm.cells)-x-1, len(dm.cells[0])-y-1) || calcContinues
			}
		}
	}
}

func (dm *DijkstraMap) tryRefreshRankAt(x, y int) bool {
	lowest := dm.bestNeighbourTo(x, y)
	if dm.isCellPassable(x, y) && dm.cells[x][y].rank > dm.cells[lowest.x][lowest.y].rank+lowest.getCostIncreaseFromHere() {
		dm.cells[x][y].rank = dm.cells[lowest.x][lowest.y].rank + lowest.getCostIncreaseFromHere()
		return true
	}
	return false
}

func (dm *DijkstraMap) isInBounds(x, y int) bool {
	return x >= 0 && y >= 0 && x < len(dm.cells) && y < len(dm.cells[x])
}

func (dm *DijkstraMap) bestNeighbourTo(x, y int) neighbourCoord {
	currRank := maxRank // dm.cells[x][y].rank
	neighbours := dm.neighbouringFunc(x, y)
	ret := neighbours[0]
	for _, n := range neighbours {
		if dm.isInBounds(n.x, n.y) && dm.cells[n.x][n.y].rank < currRank {
			currRank = dm.cells[n.x][n.y].rank
			ret = n
		}
	}
	return ret
}
