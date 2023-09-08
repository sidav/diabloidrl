package main

import (
	"github.com/gdamore/tcell/v2"
)

const (
	animTypeHit uint8 = iota
	animTypePawnIsActing
	animTypeShot
	animTypeHitscanProjectile
	animTypeGibs
)

type animation struct {
	realX, realY    int
	realX2, realY2  int // additional, for two-coorded animations
	animType        uint8
	framesRemaining int8
	order           int
}

func (r *rendererStruct) addAnimationAt(atype uint8, x, y int, appendToPrevious bool) {
	var frames int
	switch atype {
	case animTypeHit:
		frames = 4
	case animTypeGibs:
		frames = 7
	case animTypeShot:
		frames = 4
	case animTypePawnIsActing:
		frames = 4
	default:
		panic("No animation timeout set")
	}
	order := len(r.animations)
	if appendToPrevious {
		order = r.animations[len(r.animations)-1].order
	}
	r.animations = append(r.animations, &animation{
		animType: atype,
		realX:    x, realY: y,
		framesRemaining: int8(frames),
		order:           order,
	})
}

func (r *rendererStruct) addTwoCoordAnimationAt(atype uint8, x1, y1, x2, y2 int, appendToPrevious bool) {
	var frames int
	switch atype {
	case animTypeHitscanProjectile:
		frames = len(r.dung.getDigitalLineOfSightBetween(x1, y1, x2, y2))
		if frames == 0 {
			log.AppendMessagef("DBG: Got empty animation line")
			return // add nothing
		}
	default:
		panic("No animation timeout set")
	}
	order := len(r.animations)
	if appendToPrevious {
		order = r.animations[len(r.animations)-1].order
	}
	r.animations = append(r.animations, &animation{
		animType: atype,
		realX:    x1, realY: y1,
		realX2: x2, realY2: y2,
		framesRemaining: int8(frames),
		order:           order,
	})
}

func (r *rendererStruct) renderAnimation(d *dungeon, a *animation) {
	if a.framesRemaining <= 0 {
		return
	}
	x, y := r.view.RealCoordsToScreenCoords(a.realX, a.realY)
	switch a.animType {
	case animTypeHit:
		cw.SetStyle(tcell.ColorDarkRed, tcell.ColorBlack)
		cw.PutChar(
			[]rune{'|', '/', '-', '\\'}[a.framesRemaining-1],
			x, y,
		)
	case animTypeShot:
		cw.SetStyle(tcell.ColorYellow, tcell.ColorBlack)
		cw.PutChar(
			[]rune{'X', 'x', '*', '+'}[a.framesRemaining-1],
			x, y,
		)
	case animTypeGibs:
		cw.SetStyle(tcell.ColorDarkRed, tcell.ColorBlack)
		cw.PutChar('&', x, y)
		for i := x - 1; i <= x+1; i++ {
			for j := y - 1; j <= y+1; j++ {
				if rnd.Rand(12) <= int(a.framesRemaining) {
					cw.PutChar(
						[]rune{';', '"', '=', '~', '*'}[rnd.Rand(5)],
						i, j,
					)
				}
			}
		}
	case animTypeHitscanProjectile:
		cw.SetStyle(tcell.ColorYellow, tcell.ColorBlack)
		line := r.dung.getDigitalLineOfSightBetween(a.realX, a.realY, a.realX2, a.realY2)
		if line == nil || len(line) == 0 {
			return
		}
		for i := 0; i < len(line)-int(a.framesRemaining); i++ {
			sx, sy := r.view.RealCoordsToScreenCoords(line[i].X, line[i].Y)
			cw.PutChar('*', sx, sy)
		}
	case animTypePawnIsActing:
		pawn := d.getPawnAt(a.realX, a.realY)
		if pawn != nil {
			r.renderPawn(d, pawn, true)
		}
	default:
		panic("No animation logic")
	}
}
