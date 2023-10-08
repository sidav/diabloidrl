package roomgrowinggenerator

type tileCode uint8

const (
	TILE_UNFILLED tileCode = iota
	TILE_WALL
	TILE_FLOOR
	TILE_DOOR
	TILE_FENCE
)

type tile struct {
	Code      tileCode
	roomId    int
	Connected bool // for interconnectedness check
}

func (t *tile) isConnective() bool {
	return t.Code == TILE_FLOOR || t.Code == TILE_DOOR
}

func (t *tile) setByVaultChar(vc rune) {
	switch vc {
	case ' ':
		// do nothing
	case '#':
		t.Code = TILE_WALL
	case '+':
		t.Code = TILE_DOOR
	case '.':
		t.Code = TILE_FLOOR
	case '\'':
		t.Code = TILE_FENCE
	default:
		dbgPanic("No such char: %v", vc)
	}
}

func (t *tile) getRune() rune {
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
