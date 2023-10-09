package roomgrowinggenerator

type tileCode uint8

const (
	TILE_UNFILLED tileCode = iota
	TILE_WALL
	TILE_FLOOR
	TILE_DOOR
	TILE_FENCE
	TILE_ENTRYPOINT
)

type Tile struct {
	Code      tileCode
	roomId    int
	Connected bool // for connectivity check
}

func (t *Tile) isConnective() bool {
	return t.Code == TILE_FLOOR || t.Code == TILE_DOOR || t.Code == TILE_ENTRYPOINT
}

func (t *Tile) setByVaultChar(vc rune) {
	switch vc {
	case charAny:
		// do nothing
	case charWall:
		t.Code = TILE_WALL
	case charDoor:
		t.Code = TILE_DOOR
	case charFloor, charFloorOldId:
		t.Code = TILE_FLOOR
	case charFence:
		t.Code = TILE_FENCE
	default:
		dbgPanic("No such char: %s", string(vc))
	}
}

// Safe to delete. Used only for debug output
func (t *Tile) getRune() rune {
	switch t.Code {
	case TILE_UNFILLED:
		return '.'
	case TILE_WALL:
		return '#'
	case TILE_FLOOR:
		return ' '
	case TILE_DOOR:
		return '+'
	}
	return '?'
}
