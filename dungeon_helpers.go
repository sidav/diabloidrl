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
	cell := d.pathfinder.FindPath(
		func(x, y int) int {
			if d.isTilePassableForPawn(x, y, mover) {
				return 10
			}
			if d.getTileAt(x, y).code == tileDoor {
				return 20
			}
			return -1
		},
		mover.x, mover.y, targetPawn.x, targetPawn.y)
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
