package bspdung

import (
	"time"

	"github.com/gdamore/tcell/v2"
)

const (
	debugOutputEnabled = true
)

var disableWait = false

func (g *Generator) drawCurrentState() {
	if !debugOutputEnabled {
		return
	}
	g.Cw.ClearScreen()
	g.Cw.SetStyle(tcell.ColorWhite, tcell.ColorBlack)
	for x := range g.dung {
		for y := range g.dung[x] {
			g.Cw.PutChar(g.dung[x][y], x, y)
		}
	}
	g.Cw.FlushScreen()
	g.debugWait()
}

func (g *Generator) drawCurrentLeaf(x, y, w, h int) {
	if !debugOutputEnabled {
		return
	}
	g.debugPrintf(0, "SPLIT: x %d, y %d, w %d, h %d", x, y, w, h)
	g.Cw.SetStyle(tcell.ColorGreen, tcell.ColorBlack)
	for i := x; i < x+w; i++ {
		for j := y; j < y+h; j++ {
			g.Cw.PutChar('.', i, j)
		}
	}
	g.Cw.FlushScreen()
	g.debugWait()
}

func (g *Generator) drawCurrentFill(x, y, w, h int, varName string, varValue int) {
	if !debugOutputEnabled {
		return
	}
	g.debugPrintf(1, "%s: %d", varName, varValue)
	g.Cw.SetStyle(tcell.ColorRed, tcell.ColorBlack)
	for i := x; i < x+w; i++ {
		for j := y; j < y+h; j++ {
			g.Cw.PutChar('*', i, j)
		}
	}
	g.Cw.FlushScreen()
	g.debugWait()
}

func (g *Generator) debugPrintf(y int, str string, args ...interface{}) {
	g.Cw.SetStyle(tcell.ColorWhite, tcell.ColorBlack)
	g.Cw.PutStringf(len(g.dung)+1, y, str+"                           ", args...)
}

func (g *Generator) debugWait() {
	const key = true
	if disableWait {
		return
	}
	if key {
		if g.Cw.ReadKey() == "ESCAPE" {
			disableWait = true
		}
	} else {
		time.Sleep(500 * time.Millisecond)
	}
}
