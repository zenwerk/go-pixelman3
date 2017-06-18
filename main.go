package main

import (
	"log"

	"github.com/hajimehoshi/ebiten"

	"github.com/zenwerk/go-pixelman3/sprite"
)

type Game struct {
	Player *sprite.Player
	Blocks []*sprite.Block
}

func (g *Game) Init() {
	// プレイヤー
	g.Player = sprite.NewPlayer()
	g.Player.Position.X = 160
	g.Player.Position.Y = 50

	// ブロック
	// 床
	for x := 0; x < 640; x += 17 {
		block := sprite.NewBlock()
		block.Position.X = x
		block.Position.Y = 204
		g.Blocks = append(g.Blocks, block)
	}

	// 左の壁
	for y := 0; y < 200; y += 17 {
		block := sprite.NewBlock()
		block.Position.X = 0
		block.Position.Y = y
		g.Blocks = append(g.Blocks, block)
	}

	// 右の壁
	for y := 0; y < 200; y += 17 {
		block := sprite.NewBlock()
		block.Position.X = 629
		block.Position.Y = y
		g.Blocks = append(g.Blocks, block)
	}

	// 第2床
	for x := 8 * 17; x < 17*13; x += 17 {
		block := sprite.NewBlock()
		block.Position.X = x
		block.Position.Y = 115
		g.Blocks = append(g.Blocks, block)
	}

	block1 := sprite.NewBlock()
	block1.Position.X = 60
	block1.Position.Y = 165
	g.Blocks = append(g.Blocks, block1)

	block2 := sprite.NewBlock()
	block2.Position.X = 95
	block2.Position.Y = 135
	g.Blocks = append(g.Blocks, block2)

}

func (g *Game) MainLoop(screen *ebiten.Image) error {
	sprites := []sprite.Sprite{}
	for _, b := range g.Blocks {
		sprites = append(sprites, b)
	}
	g.Player.Move(sprites)
	g.Player.Action()
	g.Player.PlayerBalls.Move(g.Player.ViewPort)

	if ebiten.IsRunningSlowly() {
		return nil
	}

	g.Player.DrawImage(screen)
	for _, ball := range g.Player.PlayerBalls {
		ball.DrawImage(screen, g.Player.ViewPort)
	}
	for _, block := range g.Blocks {
		block.DrawImage(screen, g.Player.ViewPort)
	}

	return nil
}

func main() {
	game := Game{}
	game.Init()

	if err := ebiten.Run(game.MainLoop, 320, 240, 2, "go-pixelman3"); err != nil {
		log.Fatal(err)
	}
}
