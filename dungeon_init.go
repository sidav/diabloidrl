package main

import (
	"diabloidrl/lib/dijkstra_map"
	roomgrowinggenerator "diabloidrl/lib/dungeon_generators/room_growing_generator"
	"diabloidrl/lib/pathfinding/astar"
	"diabloidrl/static"
)

func (d *dungeon) init(generatedMap [][]roomgrowinggenerator.Tile) {
	d.initFromCharMap(generatedMap)
	for i := 0; i < 3; i++ {
		d.placeChestAtRandom()
	}
	player.x, player.y = d.getEntrypointCoords()
	d.addPawnAt(player, player.x, player.y)
	totalMobs := len(generatedMap) * len(generatedMap[0]) / 75
	log.AppendMessagef("Total %d mobs.", totalMobs)
	for i := 0; i < totalMobs; i++ {
		rarity := 0
		if i < totalMobs/6 {
			rarity = 1
		}
		d.initNewPawnByStats(static.GenerateRandomMobBase(rnd, rarity))
	}
	d.playerExplorationDM = dijkstra_map.New(len(d.dmap), len(d.dmap[0]), dijkstra_map.AllNeighbours,
		func(x, y int) bool {
			return d.isTilePassable(x, y) || d.getTileAt(x, y).code == tileDoor
		})
	d.pathfinder = &astar.AStarPathfinder{
		DiagonalMoveAllowed:       true,
		ForceGetPath:              true,
		ForceIncludeFinish:        true,
		AutoAdjustDefaultMaxSteps: true,
		MapWidth:                  len(d.dmap),
		MapHeight:                 len(d.dmap[0]),
	}
}

func (d *dungeon) initNewPawnByStats(stats *static.MobStats) {
	x, y := 0, 0
	for !d.isTilePassableAndEmpty(x, y) {
		x, y = rnd.RandInRange(1, len(d.dmap)-2), rnd.RandInRange(1, len(d.dmap[0])-2)
	}
	m := &pawn{
		mob: &mobStruct{},
	}
	m.mob.initFromStatic(stats)
	d.addPawnAt(m, x, y)
}

func (d *dungeon) initFromCharMap(genMap [][]roomgrowinggenerator.Tile) {
	d.dmap = make([][]*tile, len(genMap))
	for i := range d.dmap {
		d.dmap[i] = make([]*tile, len(genMap[i]))
	}
	for x := range d.dmap {
		for y := range d.dmap[x] {
			d.dmap[x][y] = &tile{}
			switch genMap[x][y].Code {
			case roomgrowinggenerator.TILE_WALL, roomgrowinggenerator.TILE_UNFILLED:
				d.dmap[x][y].code = tileWall
			case roomgrowinggenerator.TILE_FENCE:
				d.dmap[x][y].code = tileCage
			case roomgrowinggenerator.TILE_FLOOR:
				d.dmap[x][y].code = tileFloor
			case roomgrowinggenerator.TILE_DOOR:
				d.dmap[x][y].code = tileDoor
			case roomgrowinggenerator.TILE_ENTRYPOINT:
				d.dmap[x][y].code = tileEntrypoint
			default:
				panic("No code for code " + string(genMap[x][y].Code))
			}
		}
	}
}

func (d *dungeon) placeChestAtRandom() {
	x, y := -1, -1
	for !(d.isInBounds(x, y)) {
		x, y := rnd.RandInRange(1, len(d.dmap)-2), rnd.RandInRange(1, len(d.dmap[0])-2)
		if d.getTileAt(x, y).code == tileFloor {
			d.getTileAt(x, y).code = tileChest
			return
		} else {
			x, y = -1, -1
		}
	}
}
