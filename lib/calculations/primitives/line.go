package primitives

type Point struct {
	X, Y int
}

func (p *Point) GetCoords() (int, int) {
	return p.X, p.Y
}

func (p *Point) Equals(other *Point) bool {
	return p.X == other.X && p.Y == other.Y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func GetLine(fromx, fromy, tox, toy int) []Point {
	line := make([]Point, 0)
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
		y := fromy
		eps := deltax >> 1
		for x := fromx; x != tox+xmod; x += xmod {
			line = append(line, Point{x, y})
			eps += deltay
			if eps >= deltax {
				y += ymod
				eps -= deltax
			}
		}
	} else {
		x := fromx
		eps := deltay >> 1
		for y := fromy; y != toy+ymod; y += ymod {
			line = append(line, Point{x, y})
			eps += deltax
			if eps >= deltay {
				x += xmod
				eps -= deltay
			}
		}
	}
	return line
}

func GetLineOfLength(fromx, fromy, throughX, throughY, length int) []Point {
	line := make([]Point, 0)
	deltax := abs(throughX - fromx)
	deltay := abs(throughY - fromy)
	xmod := 1
	ymod := 1
	if throughX < fromx {
		xmod = -1
	}
	if throughY < fromy {
		ymod = -1
	}
	if deltax >= deltay {
		y := fromy
		eps := deltax >> 1
		for x := fromx; len(line) < length; x += xmod {
			line = append(line, Point{x, y})
			eps += deltay
			if eps >= deltax {
				y += ymod
				eps -= deltax
			}
		}
	} else {
		x := fromx
		eps := deltay >> 1
		for y := fromy; len(line) < length; y += ymod {
			line = append(line, Point{x, y})
			eps += deltax
			if eps >= deltay {
				x += xmod
				eps -= deltay
			}
		}
	}
	return line
}
