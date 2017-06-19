package game

import (
	"github.com/hajimehoshi/ebiten"

	"github.com/zenwerk/go-pixelman3/field"
)

type Game struct {
	Field *field.Field
}

func (g *Game) Init() {
	g.Field = field.NewField(field.Field_data_1)
}

func (g *Game) MainLoop(screen *ebiten.Image) error {
	g.Field.Player.Move(g.Field.Sprites)
	g.Field.Player.Action()
	g.Field.Player.PlayerBalls.Move(g.Field.Player.ViewPort)

	if ebiten.IsRunningSlowly() {
		return nil
	}

	g.Field.Player.DrawImage(screen)
	for _, ball := range g.Field.Player.PlayerBalls {
		ball.DrawImage(screen, g.Field.Player.ViewPort)
	}
	g.Field.DrawImage(screen, g.Field.Player.ViewPort)

	return nil
}
