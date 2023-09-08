package main

import "diabloidrl/lib/calculations"

func (d *dungeon) addItemAt(i *item, x, y int) {
	sx, sy := x, y
	if d.isAnyItemAt(x, y) || !d.isTilePassable(x, y) {
		sx, sy = calculations.SpiralSearchForClosestConditionFrom(
			func(xx, yy int) bool {
				return !d.isAnyItemAt(xx, yy) && d.isTilePassableAndEmpty(xx, yy)
			}, x, y, 3, rnd.Rand(4),
		)
	}
	i.x, i.y = sx, sy
	d.items = append(d.items, i)
}

func (d *dungeon) isAnyItemAt(x, y int) bool {
	for _, i := range d.items {
		if i.x == x && i.y == y {
			return true
		}
	}
	return false
}

func (d *dungeon) getItemsAt(x, y int) (items []*item) {
	for i := len(d.items) - 1; i >= 0; i-- {
		if d.items[i].x == x && d.items[i].y == y {
			items = append(items, d.items[i])
		}
	}
	return
}

func (d *dungeon) removeItem(itm *item) {
	for i := len(d.items) - 1; i >= 0; i-- {
		if d.items[i] == itm {
			d.items = append(d.items[:i], d.items[i+1:]...)
			return
		}
	}
}
