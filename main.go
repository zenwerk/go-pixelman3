package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"

	"github.com/zenwerk/go-pixelman3/game"
)

func main() {
	game := game.Game{
		ScreenWidth:  320,
		ScreenHeight: 240,
	}
	game.Init()

	if err := ebiten.Run(game.MainLoop, game.ScreenWidth, game.ScreenHeight, 2, "go-pixelman3"); err != nil {
		log.Fatal(err)
	}
}
