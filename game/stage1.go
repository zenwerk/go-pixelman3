package game

import (
	"github.com/hajimehoshi/ebiten"

	"github.com/zenwerk/go-pixelman3/camera"
	"github.com/zenwerk/go-pixelman3/field"
	"github.com/zenwerk/go-pixelman3/sprite"
)

type Stage1 struct {
	Field  *field.Field
	Player *sprite.Player
}

func NewStage1() *Stage1 {
	st1 := &Stage1{}
	st1.Field, st1.Player = field.NewField(field.Field_data_1)
	return st1

}

func (s *Stage1) Update(game *Game) {
	game.Camera.MaxWidth = s.Field.Width
	game.Camera.MaxHeight = s.Field.Height
	game.Camera.Move(s.Player.Position.X, s.Player.Position.Y)

	s.Player.Move(s.Field.Sprites)
	s.Player.Action()
	s.Player.Balls.Move(game.Camera)
}

func (s *Stage1) Draw(screen *ebiten.Image, camera *camera.Camera) {
	s.Player.DrawImage(screen, camera)
	for _, ball := range s.Player.Balls {
		ball.DrawImage(screen, camera)
	}
	s.Field.DrawImage(screen, camera)
}
