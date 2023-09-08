package util

type Viewport struct {
	viewW, viewH                     int
	viewCenterRealX, viewCenterRealY int
}

func (v *Viewport) SetViewportSize(w, h int) {
	v.viewW, v.viewH = w, h
}

func (v *Viewport) SetViewportRealCenter(w, h int) {
	v.viewCenterRealX, v.viewCenterRealY = w, h
}

func (v *Viewport) GetViewportSize() (int, int) {
	return v.viewW, v.viewH
}

func (v *Viewport) RealCoordsToScreenCoords(x, y int) (int, int) {
	return x - v.viewCenterRealX + v.viewW/2, y - v.viewCenterRealY + v.viewH/2
}

func (v *Viewport) AreRealCoordsInViewport(x, y int) bool {
	sx, sy := v.RealCoordsToScreenCoords(x, y)
	return 0 <= sx && sx < v.viewW && 0 <= sy && sy < v.viewH
}

func (v *Viewport) ViewportCoordsToRealCoords(x, y int) (int, int) {
	return x + v.viewCenterRealX - v.viewW/2, y + v.viewCenterRealY - v.viewH/2
}
