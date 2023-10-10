package main

import (
	// bspdung "diabloidrl/lib/dungeon_generators/bsp_dung"
	roomgrowinggenerator "diabloidrl/lib/dungeon_generators/room_growing_generator"
	"diabloidrl/lib/game_log"
	"diabloidrl/lib/random"
	"diabloidrl/lib/random/pcgrandom"
	"diabloidrl/lib/tcell_console_wrapper"
	"diabloidrl/static"
	"time"

	"github.com/gdamore/tcell/v2"
)

const (
	LOG_SIZE = 5
)

var (
	cw       *tcell_console_wrapper.ConsoleWrapper
	log      game_log.GameLog
	stopGame bool
	rnd      random.PRNG
	renderer rendererStruct
	player   *pawn
)

func main() {
	// defer gaussTest()
	// defer calculateLevels()

	static.PerformAffixSanityCheck()

	rnd = pcgrandom.New(int(time.Now().UnixNano()))
	static.SetRandom(rnd)
	// rnd = pcgrandom.New(1)

	cw = &tcell_console_wrapper.ConsoleWrapper{}
	cw.Init()
	defer cw.Close()
	testGen()
	log.Init(LOG_SIZE)

	cw.SetStyle(tcell.ColorRed, tcell.ColorBlack)
	gen := roomgrowinggenerator.Generator{
		MinRoomSide: 3,
		MaxRoomSide: 25,
	} // bspdung.Generator{Cw: cw}
	generatedMap := gen.Generate(80, 50, rnd)
	cw.ReadKey()

	dung := &dungeon{}
	renderer.attachToDungeonStruct(dung)
	pc := &playerController{}
	player = &pawn{
		hitpoints: 10,
		playerStats: &playerStruct{
			experience: 0,
		},
		inv: &inventory{},
	}
	player.playerStats.setDefaultStats()
	dung.init(generatedMap)
	log.AppendMessage("Init complete")
	// log.AppendMessage("Dice test: " + random.NewDice(2, 6, 1).GetDescriptionString())
	game(dung, pc)
}
