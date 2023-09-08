package main

type tile struct {
	code            int
	isBloody        bool
	hasGibs         bool
	isOpened        bool
	wasSeenByPlayer bool
	wasOnPlayerPath bool
}

func (t *tile) isWalkable() bool {
	if t.code == tileDoor {
		return t.isOpened
	}
	return t.getStaticData().walkable
}

func (t *tile) isOpaque() bool {
	if t.code == tileDoor {
		return !t.isOpened
	}
	return t.getStaticData().opaque
}

func (t *tile) getStaticData() *tileStatic {
	return sTableTile[t.code]
}

const (
	tileFloor = iota
	tileEntrypoint
	tileWall
	tileCage
	tileDoor
	tileChest
)

type tileStatic struct {
	openable, walkable, opaque bool
	ascii                      rune
}

var sTableTile = map[int]*tileStatic{
	tileFloor: {
		ascii:    '.',
		walkable: true,
		opaque:   false,
	},
	tileEntrypoint: {
		ascii:    '<',
		walkable: true,
		opaque:   false,
	},
	tileWall: {
		ascii:    ' ',
		walkable: false,
		opaque:   true,
	},
	tileCage: {
		ascii:    '#',
		walkable: false,
		opaque:   false,
	},
	tileDoor: {
		ascii:    '?',
		openable: true,
		walkable: true,
	},
	tileChest: {
		ascii:    '=',
		openable: false, // it's not a door
		walkable: false,
	},
}
