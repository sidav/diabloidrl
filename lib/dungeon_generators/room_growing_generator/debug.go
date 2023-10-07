package roomgrowinggenerator

import (
	"diabloidrl/lib/tcell_console_wrapper"
	"strconv"

	"github.com/gdamore/tcell/v2"
)

var cw *tcell_console_wrapper.ConsoleWrapper

func SetDebugCw(cwr *tcell_console_wrapper.ConsoleWrapper) {
	cw = cwr
}

func (gen *Generator) dbgDrawCurrentState() {
	cw.ClearScreen()
	for x := range gen.Tiles {
		for y := range gen.Tiles[x] {
			char := '?'
			switch gen.Tiles[x][y].Code {
			case TILE_UNFILLED:
				cw.SetStyle(tcell.ColorWhite, tcell.ColorBlack)
				char = '.'
			case TILE_FLOOR:
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
	cw.SetStyle(tcell.ColorDarkRed, tcell.ColorBlack)
	cw.PutChar('*', x, y)
}

func (g *Generator) dbgFlush() {
	cw.FlushScreen()
	cw.ReadKey()
}
