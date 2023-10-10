package roomgrowinggenerator

import (
	"diabloidrl/lib/tcell_console_wrapper"
	"fmt"
	"strconv"

	"github.com/gdamore/tcell/v2"
)

var cw *tcell_console_wrapper.ConsoleWrapper

func SetDebugCw(cwr *tcell_console_wrapper.ConsoleWrapper) {
	cw = cwr
}

func dbgPanic(comment string, commentArgs ...interface{}) {
	panic(fmt.Sprintf(comment, commentArgs...))
}

func dbgFlush(awaitKeypress bool) {
	cw.FlushScreen()
	if awaitKeypress {
		cw.ReadKey()
	}
}

func dbgMessage(comment string, commentArgs ...interface{}) {
	cw.SetStyle(tcell.ColorBlack, tcell.ColorDarkRed)
	cw.PutString(fmt.Sprintf(comment, commentArgs...), 0, 0)
}

func (gen *Generator) dbgDrawCurrentState(showIds bool) {
	cw.ClearScreen()
	for x := range gen.tiles {
		for y := range gen.tiles[x] {
			char := '?'
			switch gen.tiles[x][y].Code {
			case TILE_UNFILLED:
				cw.SetStyle(tcell.ColorWhite, tcell.ColorBlack)
				char = '.'
			case TILE_FLOOR:
				if !showIds {
					continue
				}
				cw.SetStyle(tcell.ColorWhite, tcell.ColorBlack)
				idStr := strconv.Itoa(gen.tiles[x][y].roomId)
				char = rune(idStr[len(idStr)-1])
			case TILE_DOOR:
				cw.SetStyle(tcell.ColorWhite, tcell.ColorBlack)
				char = '+'
			case TILE_WALL:
				cw.SetStyle(tcell.ColorBlack, tcell.ColorDarkRed)
				char = '#'
			case TILE_FENCE:
				cw.SetStyle(tcell.ColorWhite, tcell.ColorBlack)
				char = '"'
			}
			cw.PutChar(char, x, y)
		}
	}
}

func (g *Generator) dbgHighlightTile(x, y int) {
	cw.SetStyle(tcell.ColorDarkRed, tcell.ColorDarkCyan)
	cw.PutChar('*', x, y)
}

func (g *Generator) dbgHighlightTileWithComment(x, y int, comment string, commentArgs ...interface{}) {
	cw.SetStyle(tcell.ColorDarkRed, tcell.ColorDarkCyan)
	cw.PutChar('*', x, y)
	cw.PutString(fmt.Sprintf(comment, commentArgs...), 0, 0)
}

func (g *Generator) dbgShowVault(v []string, hlx, hly int) {
	for x := range v {
		for y := range v[x] {
			cw.SetStyle(tcell.ColorDarkRed, tcell.ColorBlack)
			if hlx == x && hly == y {
				cw.InverseStyle()
			}
			cw.PutChar(rune(v[x][y]), x+len(g.tiles), y)
		}
	}
}
