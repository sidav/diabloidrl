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

func (gen *Generator) dbgDrawCurrentState(showIds bool) {
	cw.ClearScreen()
	for x := range gen.Tiles {
		for y := range gen.Tiles[x] {
			char := '?'
			switch gen.Tiles[x][y].Code {
			case TILE_UNFILLED:
				cw.SetStyle(tcell.ColorWhite, tcell.ColorBlack)
				char = '.'
			case TILE_FLOOR:
				if !showIds {
					continue
				}
				cw.SetStyle(tcell.ColorWhite, tcell.ColorBlack)
				char = rune(strconv.Itoa(gen.Tiles[x][y].roomId)[0])
			case TILE_DOOR:
				cw.SetStyle(tcell.ColorWhite, tcell.ColorBlack)
				char = '+'
			case TILE_WALL:
				cw.SetStyle(tcell.ColorBlack, tcell.ColorDarkRed)
				char = '#'
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

func (g *Generator) dbgFlush() {
	cw.FlushScreen()
	cw.ReadKey()
}

func (g *Generator) dbgShowVault(v []string, hlx, hly int) {
	for x := range v {
		for y := range v[x] {
			cw.SetStyle(tcell.ColorDarkRed, tcell.ColorBlack)
			if hlx == x && hly == y {
				cw.InverseStyle()
			}
			cw.PutChar(rune(v[x][y]), x+len(g.Tiles), y)
		}
	}
}
