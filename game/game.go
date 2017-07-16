package game

import (
	"github.com/hajimehoshi/ebiten"

	"github.com/zenwerk/go-pixelman3/field"
	"github.com/zenwerk/go-pixelman3/sprite"
)

type Game struct {
	Field    *field.Field
	ViewPort *sprite.Position
	Player   *sprite.Player
}

func (g *Game) Init() {
	g.Field, g.Player = field.NewField(field.Field_data_1)
	g.ViewPort = &sprite.Position{}
}

func (g *Game) MainLoop(screen *ebiten.Image) error {
	g.Player.Move(g.Field.Sprites, g.ViewPort)
	g.Player.Action(g.ViewPort)
	g.Player.Balls.Move(g.ViewPort)

	if ebiten.IsRunningSlowly() {
		return nil
	}

	g.Player.DrawImage(screen, nil)
	for _, ball := range g.Player.Balls {
		ball.DrawImage(screen, g.ViewPort)
	}
	g.Field.DrawImage(screen, g.ViewPort)

	return nil
}
