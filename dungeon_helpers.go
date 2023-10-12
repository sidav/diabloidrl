package main

import "diabloidrl/lib/calculations/primitives"

func (d *dungeon) getDigitalLineOfSightBetween(fx, fy, tx, ty int) []primitives.Point {
	line := primitives.GetSuitableiDigitalLine(fx, fy, tx, ty, func(x, y int) bool {
		return !d.isTileOpaque(x, y) || (x == fx && y == fy) || (x == tx && y == ty)
	})
	if len(line) >= 2 {
		return line[1:]
	}
	return nil
}

func (d *dungeon) getStepForPawnToPawn(mover, targetPawn *pawn) (int, int) {
	targetX, targetY := targetPawn.x, targetPawn.y
	targetW, targetH := 1, 1
	if mover.getSize() > 1 {
		targetX -= mover.getSize() - 1
		targetY -= mover.getSize() / 2
		targetW = mover.getSize()
		targetH = mover.getSize()
	}
	cell := d.pathfinder.FindPath(
		func(x, y int) int {
			if d.isTilePassableForPawn(x, y, mover) {
				if d.getTileAt(x, y).code == tileDoor {
					return 30
				}
				return 10
			}
			return -1
		},
		mover.x, mover.y, targetX, targetY, targetW, targetH)
	if cell == nil {
		return 0, 0
	}
	return cell.GetNextStepVector()
}

// func (d *dungeon) getStepForPathFromTo(fx, fy, tx, ty int) (int, int) {
// 	cell := d.pathfinder.FindPath(
// 		func(x, y int) int {
// 			if d.isTilePassableAndEmpty(x, y) {
// 				return 10
// 			}
// 			if d.getTileAt(x, y).code == tileDoor {
// 				return 20
// 			}
// 			return -1
// 		},
// 		fx, fy, tx, ty)
// 	if cell == nil {
// 		return 0, 0
// 	}
// 	return cell.GetNextStepVector()
// }
