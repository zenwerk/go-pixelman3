package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"

	"github.com/zenwerk/go-pixelman3/game"
)

func main() {
	g := game.NewGame(320, 240)
	g.Init()

	if err := ebiten.Run(g.MainLoop, g.ScreenWidth, g.ScreenHeight, 2, "go-pixelman3"); err != nil {
		log.Fatal(err)
	}
}
