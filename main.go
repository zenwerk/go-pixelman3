package main

import (
	"image"
	"log"

	"github.com/hajimehoshi/ebiten"

	"github.com/zenwerk/go-pixelman3/sprite"
	"github.com/zenwerk/go-pixelman3/utils"
)

var player_anim0 = `------+++++-----
----+++++++++---
---+++++++++++--
--+++++++++++++-
--++++--+++--++-
-+++++--+++--++-
+++++++++++++++-
++++++++++++++++
++++++++++++++++
++++++-+++++-+++
+++++++-----++++
+-++++++++++++++
--++++++++++++-+
---++++++++++---
---++-----++----
--+++++--+++++--`

var player_anim1 = `------+++++-----
----+++++++++---
---+++++++++++--
--+++++++++++++-
--++++--+++--++-
-+++++--+++--++-
+++++++++++++++-
++++++++++++++++
++++++++++++++++
++++++-+++++-+++
+++++++-----++++
+-++++++++++++++
--++++++++++++-+
---++++++++++---
--+++++---++----
---------+++++--`

var player_anim2 = `------+++++-----
----+++++++++---
---+++++++++++--
--+++++++++++++-
--++++--+++--++-
-+++++--+++--++-
+++++++++++++++-
++++++++++++++++
++++++++++++++++
++++++-+++++-+++
+++++++-----++++
+-++++++++++++++
--++++++++++++-+
---++++++++++---
---++----+++++--
--+++++---------`

var block_img = `++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++`

var (
	charWidth   = 16
	charHeight  = 16
	tmpImage    *image.RGBA
	playerAnim0 *ebiten.Image
	playerAnim1 *ebiten.Image
	playerAnim2 *ebiten.Image
	blockImg    *ebiten.Image
)

type Game struct {
	Player *sprite.Player
	Blocks []*sprite.Block
}

func (g *Game) Init() {
	tmpImage = image.NewRGBA(image.Rect(0, 0, charWidth, charHeight))

	utils.CreateImageFromString(player_anim0, tmpImage)
	playerAnim0, _ = ebiten.NewImage(charWidth, charHeight, ebiten.FilterNearest)
	playerAnim0.ReplacePixels(tmpImage.Pix)

	utils.CreateImageFromString(player_anim1, tmpImage)
	playerAnim1, _ = ebiten.NewImage(charWidth, charHeight, ebiten.FilterNearest)
	playerAnim1.ReplacePixels(tmpImage.Pix)

	utils.CreateImageFromString(player_anim2, tmpImage)
	playerAnim2, _ = ebiten.NewImage(charWidth, charHeight, ebiten.FilterNearest)
	playerAnim2.ReplacePixels(tmpImage.Pix)

	utils.CreateImageFromString(block_img, tmpImage)
	blockImg, _ = ebiten.NewImage(charWidth, charHeight, ebiten.FilterNearest)
	blockImg.ReplacePixels(tmpImage.Pix)

	// プレイヤー
	images := []*ebiten.Image{
		playerAnim0,
		playerAnim1,
		playerAnim2,
	}
	g.Player = sprite.NewPlayer(images)
	g.Player.Position.X = 160
	g.Player.Position.Y = 50

	// ブロック
	// 床
	for x := 0; x < 640; x += 17 {
		block := sprite.NewBlock([]*ebiten.Image{blockImg})
		block.Position.X = x
		block.Position.Y = 204
		g.Blocks = append(g.Blocks, block)
	}

	// 左の壁
	for y := 0; y < 200; y += 17 {
		block := sprite.NewBlock([]*ebiten.Image{blockImg})
		block.Position.X = 0
		block.Position.Y = y
		g.Blocks = append(g.Blocks, block)
	}

	// 右の壁
	for y := 0; y < 200; y += 17 {
		block := sprite.NewBlock([]*ebiten.Image{blockImg})
		block.Position.X = 629
		block.Position.Y = y
		g.Blocks = append(g.Blocks, block)
	}

	// 第2床
	for x := 8 * 17; x < 17*13; x += 17 {
		block := sprite.NewBlock([]*ebiten.Image{blockImg})
		block.Position.X = x
		block.Position.Y = 115
		g.Blocks = append(g.Blocks, block)
	}

	block1 := sprite.NewBlock([]*ebiten.Image{blockImg})
	block1.Position.X = 60
	block1.Position.Y = 165
	g.Blocks = append(g.Blocks, block1)

	block2 := sprite.NewBlock([]*ebiten.Image{blockImg})
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

	if ebiten.IsRunningSlowly() {
		return nil
	}

	g.Player.DrawImage(screen)
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
