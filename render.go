package main

import (
	"diabloidrl/lib/dijkstra_map"
	"diabloidrl/lib/util"
	"fmt"
	"time"

	"github.com/gdamore/tcell/v2"
)

type rendererStruct struct {
	view         util.Viewport
	dung         *dungeon
	uiH          int
	currentFrame int
	animations   []*animation
}

func (r *rendererStruct) attachToDungeonStruct(dung *dungeon) {
	r.dung = dung
}

func (r *rendererStruct) renderGameMainScreen(dung *dungeon) {
	// update bounds
	wid, hei := cw.GetConsoleSize()
	r.uiH = 5
	if hei > 15+LOG_SIZE+r.uiH {
		r.uiH = (hei - 15 - LOG_SIZE)
	}
	r.view.SetViewportSize(wid, hei-LOG_SIZE-r.uiH)
	r.view.SetViewportRealCenter(player.x, player.y)
	forceDraw := true
	for forceDraw || len(r.animations) > 0 {
		forceDraw = false
		cw.ClearScreen()
		r.drawMap(dung)
		// r.renderDijkstraDebug(dung)
		r.renderItemsOnFloor(dung)
		for _, p := range dung.pawns {
			r.renderPawn(dung, p, false)
		}
		// render first (by order) animation(s)
		if len(r.animations) > 0 {
			order := r.animations[0].order
			for _, a := range r.animations {
				if a.order == order {
					r.renderAnimation(dung, a)
					a.framesRemaining--
				}
			}
		}
		r.renderLog()
		r.renderUI(dung)
		cw.FlushScreen()
		r.currentFrame++
		// cleanup cleared animations
		if len(r.animations) > 0 {
			for i := len(r.animations) - 1; i >= 0; i-- {
				if r.animations[i].framesRemaining == 0 {
					r.animations = append(r.animations[:i], r.animations[i+1:]...)
					forceDraw = true
				}
			}
			time.Sleep(50 * time.Millisecond)
		} else {
			time.Sleep(10 * time.Millisecond)
		}
	}
}

func (r *rendererStruct) drawMap(dung *dungeon) {
	vw, vh := r.view.GetViewportSize()
	for x := 0; x < vw; x++ {
		for y := 0; y < vh; y++ {
			r.drawTile(dung, x, y)
		}
	}
}

func (r *rendererStruct) drawTile(dung *dungeon, x, y int) {
	realX, realY := r.view.ViewportCoordsToRealCoords(x, y)
	if !dung.isInBounds(realX, realY) || !dung.dmap[realX][realY].wasSeenByPlayer {
		return
	}
	t := dung.dmap[realX][realY]
	tStatic := t.getStaticData()
	chr := tStatic.ascii
	color := tcell.ColorWhite
	inverse := false
	switch t.code {
	case tileFloor:
		color = tcell.ColorDarkGray
	case tileDoor:
		color = tcell.ColorBlue
		if t.isOpened {
			chr = '\''
		} else {
			chr = '+'
		}
	case tileWall:
		color = tcell.ColorDarkCyan
		inverse = true
	case tileCage:
		color = tcell.ColorDarkCyan
	case tileChest:
		if t.isOpened {
			color = tcell.ColorDarkGray
		} else {
			color = tcell.ColorYellow
			inverse = true
		}
	}
	// gore
	if dung.dmap[realX][realY].isBloody {
		color = tcell.ColorDarkRed
	}
	if dung.dmap[realX][realY].code == tileFloor && dung.dmap[realX][realY].hasGibs {
		chr = ';'
	}

	if !dung.playerFOVMap[realX][realY] {
		color = tcell.ColorDarkBlue
	}
	cw.SetStyle(color, tcell.ColorBlack)
	if inverse || t.wasOnPlayerPath {
		cw.InverseStyle()
	}
	cw.PutChar(chr, x, y)

}

// func (r *renderer) renderPlayer() {
// 	cw.SetStyle(tcell.ColorWhite, tcell.ColorBlack)
// 	x, y := r.view.RealCoordsToScreenCoords(player.x, player.y)
// 	cw.PutChar('@', x, y)
// }

func (r *rendererStruct) renderPawn(d *dungeon, p *pawn, inverse bool) {
	if d.playerFOVMap[p.x][p.y] && r.view.AreRealCoordsInViewport(p.x, p.y) {
		if p.isPlayer() {
			cw.SetStyle(tcell.ColorWhite, tcell.ColorBlack)
			x, y := r.view.RealCoordsToScreenCoords(p.x, p.y)
			cw.PutChar('@', x, y)
		} else {
			switch p.mob.stats.Rarity {
			case 0:
				cw.SetStyle(tcell.ColorBrown, tcell.ColorBlack)
			case 1:
				cw.SetStyle(tcell.ColorYellow, tcell.ColorBlack)
			default:
				cw.SetStyle(tcell.ColorDarkMagenta, tcell.ColorBlack)
			}
			x, y := r.view.RealCoordsToScreenCoords(p.x, p.y)
			if inverse {
				cw.InverseStyle()
			}
			for i := range p.mob.stats.AsciiPic {
				for j := range p.mob.stats.AsciiPic[i] {
					chr := rune(p.mob.stats.AsciiPic[i][j])
					cw.PutChar(chr, x+j-p.mob.stats.Size/2, y+i-p.mob.stats.Size/2)
				}
			}
		}
	}
}

func (r *rendererStruct) renderItemsOnFloor(d *dungeon) {
	cw.ResetStyle()
	for _, itm := range d.items {
		if d.playerFOVMap[itm.x][itm.y] && r.view.AreRealCoordsInViewport(itm.x, itm.y) {
			switch itm.getRarity() {
			case 0:
				cw.SetStyle(tcell.ColorWhite, tcell.ColorBlack)
			case 1:
				cw.SetStyle(tcell.ColorDarkCyan, tcell.ColorBlack)
			case 2:
				cw.SetStyle(tcell.ColorYellow, tcell.ColorBlack)
			case 3:
				cw.SetStyle(tcell.ColorDarkMagenta, tcell.ColorBlack)
			default:
				cw.SetStyle(tcell.ColorBlack, tcell.ColorWhite)
			}
			x, y := r.view.RealCoordsToScreenCoords(itm.x, itm.y)
			cw.PutChar(itm.getAscii(), x, y)
		}
	}
}

func (r *rendererStruct) renderLog() {
	cw.ResetStyle()
	_, h := r.view.GetViewportSize()
	for i := 0; i < len(log.Last_msgs); i++ {
		cw.PutString(log.Last_msgs[i].GetText(), 0, h+i)
	}
}

func (r *rendererStruct) renderDijkstraDebug(dung *dungeon) {
	cw.SetStyle(tcell.ColorWhite, tcell.ColorBlack)
	dm := dijkstra_map.New(len(dung.dmap), len(dung.dmap[0]), dijkstra_map.AllNeighbours, dung.isTilePassable)
	dm.SetTarget(1, 1, 0)
	dm.SetTarget(1, 5, 0)
	dm.SetTarget(20, 25, 0)
	dm.SetTarget(40, 15, 0)
	dm.SetTarget(65, 15, 0)
	dm.SetTarget(len(dung.dmap)-1, len(dung.dmap[0])-1, 0)
	start := time.Now()
	dm.Calculate()
	cw.PutStringf(len(dung.dmap), 3, "%d mcs", time.Since(start).Microseconds())
	for x := range dung.dmap {
		for y := range dung.dmap[x] {
			val := dm.GetRankAt(x, y) / 10
			strval := fmt.Sprintf("%d", val)
			cw.SetStyle(tcell.ColorGreen, tcell.ColorBlack)
			if val >= 10 {
				cw.SetStyle(tcell.ColorYellow, tcell.ColorBlack)
				strval = string(rune('A' + (val-10)/2))
			}
			if val > 19 {
				cw.SetStyle(tcell.ColorGray, tcell.ColorBlack)
			}
			if val >= 62 {
				strval = "   "
			}
			if len(strval) < 2 {
				cw.PutString(strval, x, y)
			}
		}
	}
}
