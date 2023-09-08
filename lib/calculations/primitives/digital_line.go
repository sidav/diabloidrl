package primitives

func GetAllDigitalLines(fromx, fromy, tox, toy int) [][]Point {
	// uses "digital lines" algorithm  (modification of Bresenham's)

	if fromx == tox && fromy == toy {
		return [][]Point{{Point{
			X: fromx,
			Y: fromy,
		}}}
	}

	lines := make([][]Point, 0)
	deltax := abs(tox - fromx)
	deltay := abs(toy - fromy)
	xmod := 1
	ymod := 1
	if tox < fromx {
		xmod = -1
	}
	if toy < fromy {
		ymod = -1
	}
	if deltax >= deltay {
		startEpsMod := Gcd(deltax, deltay) // needed to reduce the number of repeating lines
		for startEps := 0; startEps < deltax; startEps += startEpsMod {
			eps := startEps
			line := make([]Point, 0)
			y := fromy
			for x := fromx; x != tox+xmod; x += xmod {
				line = append(line, Point{x, y})
				eps += deltay
				if eps >= deltax {
					y += ymod
					eps -= deltax
				}
			}
			lines = append(lines, line)
		}
	} else {
		startEpsMod := Gcd(deltay, deltax)
		for startEps := 0; startEps < deltay; startEps += startEpsMod {
			eps := startEps
			x := fromx
			line := make([]Point, 0)
			for y := fromy; y != toy+ymod; y += ymod {
				line = append(line, Point{x, y})
				eps += deltax
				if eps >= deltay {
					x += xmod
					eps -= deltay
				}
			}
			lines = append(lines, line)
		}
	}
	return lines
}

func GetSuitableiDigitalLine(fx, fy, tx, ty int, isSuitable func(int, int) bool) []Point {
	lines := GetAllDigitalLines(fx, fy, tx, ty)
selectSuitableLine:
	for i := range lines {
		// that index is needed to start the search from the middle (middle lines look more natural)
		line := lines[(i+len(lines)/2)%len(lines)]
		for _, p := range line {
			if !isSuitable(p.X, p.Y) {
				continue selectSuitableLine
			}
		}
		return line
	}
	return nil
}

func Gcd(a, b int) int {
	if a == 0 {
		return b
	}
	if b == 0 {
		return a
	}
	for a != b {
		if a > b {
			a -= b
		} else {
			b -= a
		}
	}
	return a
}
