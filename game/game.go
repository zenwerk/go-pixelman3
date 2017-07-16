package game

import (
	"github.com/hajimehoshi/ebiten"

	"github.com/zenwerk/go-pixelman3/camera"
	"github.com/zenwerk/go-pixelman3/field"
	"github.com/zenwerk/go-pixelman3/sprite"
)

type Game struct {
	Field  *field.Field
	Camera *camera.Camera
	Player *sprite.Player
}

func (g *Game) Init() {
	g.Field, g.Player = field.NewField(field.Field_data_1)
	g.Camera = &camera.Camera{}
}

func (g *Game) MainLoop(screen *ebiten.Image) error {
	g.Player.Move(g.Field.Sprites)
	g.Player.Action()
	g.Player.Balls.Move(g.Player.ViewPort)

	if ebiten.IsRunningSlowly() {
		return nil
	}

	g.Player.DrawImage(screen, sprite.Position{})
	for _, ball := range g.Player.Balls {
		ball.DrawImage(screen, g.Player.ViewPort)
	}
	g.Field.DrawImage(screen, g.Player.ViewPort)

	return nil
}
