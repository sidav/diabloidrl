package main

import (
	"diabloidrl/lib/dijkstra_map"
	"diabloidrl/lib/fov/permissive_fov"
	"diabloidrl/lib/pathfinding/astar"
)

type dungeon struct {
	dmap                [][]*tile
	pawns               []*pawn
	items               []*item
	playerExplorationDM *dijkstra_map.DijkstraMap
	playerFOVMap        [][]bool
	pathfinder          *astar.AStarPathfinder
	currentTick         int
}

func (d *dungeon) getTileAt(x, y int) *tile {
	return d.dmap[x][y]
}

func (d *dungeon) isInBounds(x, y int) bool {
	return x >= 0 && y >= 0 && x < len(d.dmap) && y < len(d.dmap[x])
}

func (d *dungeon) isTilePassable(x, y int) bool {
	return d.isInBounds(x, y) && d.dmap[x][y].isWalkable()
}

func (d *dungeon) isTilePassableAndEmpty(x, y int) bool {
	return d.isTilePassable(x, y) && d.getPawnAt(x, y) == nil
}

func (d *dungeon) isTilePassableForPawn(x, y int, p *pawn) bool {
	if p.isPlayer() || p.mob.stats.Size <= 1 {
		return d.isTilePassableAndEmpty(x, y)
	} else {
		for i := x - 1; i <= x+1; i++ {
			for j := y - 1; j <= y+1; j++ {
				if !d.isTilePassable(i, j) {
					return false
				}
				pawnAt := d.getPawnAt(i, j)
				if pawnAt != nil && pawnAt != p {
					return false
				}
			}
		}
		return true
	}
}

func (d *dungeon) isTileOpaque(x, y int) bool {
	return d.dmap[x][y].isOpaque()
}

func (d *dungeon) placeGoreAround(x, y int) {
	d.dmap[x][y].isBloody = true
	d.dmap[x][y].hasGibs = true
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if d.isInBounds(i, j) {
				if rnd.OneChanceFrom(2) {
					d.dmap[i][j].isBloody = true
					if rnd.OneChanceFrom(4) {
						d.dmap[i][j].hasGibs = true
					}
				}
			}
		}
	}
}

func (d *dungeon) isTileInPlayerFOV(x, y int) bool {
	return d.playerFOVMap[x][y]
}

func (d *dungeon) isAnyPawnInPlayerFOV(excludePlayer bool) bool {
	for _, m := range d.pawns {
		if d.playerFOVMap[m.x][m.y] && !(excludePlayer && m.isPlayer()) {
			return true
		}
	}
	return false
}

func (d *dungeon) getAllPawnsInPlayerFOV(excludePlayer bool) (ms []*pawn) {
	for _, m := range d.pawns {
		if d.playerFOVMap[m.x][m.y] && !(excludePlayer && m.isPlayer()) {
			ms = append(ms, m)
		}
	}
	return ms
}

func (d *dungeon) getEntrypointCoords() (int, int) {
	for x := range d.dmap {
		for y := range d.dmap[x] {
			if d.dmap[x][y].code == tileEntrypoint {
				return x, y
			}
		}
	}
	panic("No entrypoint found")
}

func (d *dungeon) resetPlayerPath() {
	for x := range d.dmap {
		for y := range d.dmap[x] {
			d.dmap[x][y].wasOnPlayerPath = false
		}
	}
}

func (d *dungeon) exploreAroundPlayer() bool {
	radius := player.getVisionRadius()
	d.playerFOVMap = permissive_fov.GetFovMapFrom(player.x, player.y, radius, len(d.dmap), len(d.dmap[0]), d.isTileOpaque)
	newTileExplored := false
	for x := range d.playerFOVMap {
		for y := range d.playerFOVMap[x] {
			if d.playerFOVMap[x][y] {
				if d.dmap[x][y].wasSeenByPlayer == false {
					newTileExplored = true
				}
				d.dmap[x][y].wasSeenByPlayer = true
			}
		}
	}
	return newTileExplored
}
