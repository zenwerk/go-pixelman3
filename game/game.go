package game

import (
	"github.com/hajimehoshi/ebiten"

	"github.com/zenwerk/go-pixelman3/camera"
	"github.com/zenwerk/go-pixelman3/field"
	"github.com/zenwerk/go-pixelman3/sprite"
)

type Game struct {
	ScreenWidth  int
	ScreenHeight int
	Field        *field.Field
	Camera       *camera.Camera
	ViewPort     *sprite.Position // 後で消す
	Player       *sprite.Player
}

func (g *Game) Init() {
	g.Field, g.Player = field.NewField(field.Field_data_1)
	g.ViewPort = &sprite.Position{}
	g.Camera = &camera.Camera{
		Width:  g.ScreenWidth,
		Height: g.ScreenHeight,
	}
}

func (g *Game) MainLoop(screen *ebiten.Image) error {
	g.Player.Move(g.Field.Sprites, g.Camera)
	g.Player.Action(g.Camera)
	g.Player.Balls.Move(g.Camera)

	if ebiten.IsRunningSlowly() {
		return nil
	}

	g.Player.DrawImage(screen, nil)
	for _, ball := range g.Player.Balls {
		ball.DrawImage(screen, g.Camera)
	}
	g.Field.DrawImage(screen, g.Camera)

	return nil
}
