package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"

	"github.com/zenwerk/go-pixelman3/game"
)

func main() {
	game := game.Game{}
	game.Init()

	if err := ebiten.Run(game.MainLoop, 320, 240, 2, "go-pixelman3"); err != nil {
		log.Fatal(err)
	}
}
