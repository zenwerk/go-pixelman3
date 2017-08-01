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
	Player       *sprite.Player
}

func (g *Game) Init() {
	g.Field, g.Player = field.NewField(field.Field_data_1)
	g.Camera = &camera.Camera{
		Width:     g.ScreenWidth,
		Height:    g.ScreenHeight,
		MaxWidth:  g.Field.Width,
		MaxHeight: g.Field.Height,
	}
}

func (g *Game) MainLoop(screen *ebiten.Image) error {
	g.Camera.Move(g.Player.Position.X, g.Player.Position.Y)
	g.Player.Move(g.Field.Sprites)
	g.Player.Action()
	g.Player.Balls.Move(g.Camera)

	if ebiten.IsRunningSlowly() {
		return nil
	}

	g.Player.DrawImage(screen, g.Camera)
	for _, ball := range g.Player.Balls {
		ball.DrawImage(screen, g.Camera)
	}
	g.Field.DrawImage(screen, g.Camera)

	return nil
}
