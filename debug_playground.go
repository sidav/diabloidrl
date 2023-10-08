package main

import (
	roomgrowinggenerator "diabloidrl/lib/dungeon_generators/room_growing_generator"
	"fmt"
	"math"
	"strings"

	"github.com/gdamore/tcell/v2"
	"github.com/kr/pretty"
)

func calculateLevels() {
	const calcLevels = 15
	curr := 0
	increment := 5
	var arr []int
	for lvl := 1; lvl <= calcLevels; lvl++ {
		curr += increment
		fmt.Printf("For level %d: %d (+%d)\n", lvl, curr, increment)
		arr = append(arr, curr)
		increment += int(math.Round(math.Pow(float64(lvl), 1.5))) + 4 // 2 * lvl / 3
		// increment = int(math.Round(math.Pow(float64(increment), 1.025)) + 5)
	}
	fmt.Printf("{ ")
	for i := 0; i < len(arr)-1; i++ {
		fmt.Printf("%d, ", arr[i])
	}
	fmt.Printf("%d }\n", arr[len(arr)-1])

	fmt.Printf("ALTERNATE: \n  ")
	for i := 0; i < calcLevels; i++ {
		expReq := int(math.Round(math.Pow(float64(i)/0.25, 2))) + 5*(i+1)
		expReq /= 5
		expReq *= 5
		fmt.Printf(
			//"%d, ", int(math.Round(math.Pow(1.5, float64(i)/2)*float64(5*(i+1)))),
			"%d, ", expReq,
		)
	}
	fmt.Printf("\n")
}

func gaussTest() {
	tries := 10000
	min := 0
	max := 8
	for weight := 0.0; weight < 5; weight += 0.5 {
		fmt.Printf("\n Weight factor %.1f: ", weight)
		res := make([]int, max-min+1)
		for i := 0; i < tries; i++ {
			ind := gauss(min, max, max/2, weight)
			res[ind-min]++
		}
		fmt.Printf("\n")
		putGraphFromArray(res, tries)
		pretty.Println(res)
	}
	weight := 1.5
	for leanTo := min; leanTo <= max; leanTo++ {
		fmt.Printf("\nLeaning to %d: ", leanTo)
		res := make([]int, max-min+1)
		for i := 0; i < tries; i++ {
			ind := gauss(min, max, leanTo, weight)
			res[ind-min]++
		}
		fmt.Printf("\n")
		putGraphFromArray(res, tries)
		pretty.Println(res)
	}
}

func ipow(base, exp int) int {
	result := 1
	for exp != 0 {
		if (exp & 1) == 1 {
			result *= base
		}
		exp >>= 1
		base *= base
	}
	return result
}

func gauss(min, max, leanTo int, leanWeight float64) int {
	// maxWeight := (max - min + 1) * leanWeight
	maxDist := (max - min + 1)
	maxDist = max - leanTo
	if leanTo-min > maxDist {
		maxDist = leanTo - min
	}
	weightFunc := func(x int) int {
		distance := leanTo - x
		if distance < 0 {
			distance = -distance
		}
		return int(math.Pow(float64(maxDist-distance+1), leanWeight)) // (maxDist - distance + 1) + leanWeight*(maxDist-distance)
		// return (maxWeight*(maxDist-distance) + maxDist/2) / maxDist
	}
	// fmt.Println("WEIGHTS:")
	// for i := 0; i < max-min+1; i++ {
	// 	fmt.Printf("%d ", weightFunc(i))
	// }
	// fmt.Println("")
	return rnd.SelectRandomIndexFromWeighted(max-min+1, weightFunc) + min
}

func putGraphFromArray(arr []int, sum int) {
	const lenMax = 50
	biggest := 0
	for _, val := range arr {
		if val > biggest {
			biggest = val
		}
	}
	biggest = sum
	for i, val := range arr {
		length := (lenMax*val + biggest/2) / biggest
		fmt.Printf(fmt.Sprintf("%d: %s\n", i, strings.Repeat("#", length)))
	}
}

func testGen() {
	gen := roomgrowinggenerator.Generator{
		MinRoomSide: 3,
	}
	roomgrowinggenerator.SetDebugCw(cw)
	gen.Init()
	key := ""
	for key != "ESCAPE" {
		gen.Generate(80, 40, rnd)
		cw.ClearScreen()
		for x := range gen.Tiles {
			for y := range gen.Tiles[x] {
				rune := '?'
				switch gen.Tiles[x][y].Code {
				case roomgrowinggenerator.TILE_UNFILLED:
					cw.SetStyle(tcell.ColorWhite, tcell.ColorBlack)
					rune = '.'
				case roomgrowinggenerator.TILE_FLOOR:
					cw.SetStyle(tcell.ColorWhite, tcell.ColorBlack)
					rune = ' '
				case roomgrowinggenerator.TILE_DOOR:
					cw.SetStyle(tcell.ColorWhite, tcell.ColorBlack)
					rune = '+'
				case roomgrowinggenerator.TILE_WALL:
					cw.SetStyle(tcell.ColorBlack, tcell.ColorDarkRed)
					rune = '#'
				}
				if !gen.Tiles[x][y].Connected && gen.Tiles[x][y].Code != roomgrowinggenerator.TILE_UNFILLED {
					cw.InverseStyle()
				}
				cw.PutChar(rune, x, y)
			}
		}
		cw.FlushScreen()
		key = cw.ReadKey()
	}
}
