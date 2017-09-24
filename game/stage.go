package game

import (
	"github.com/hajimehoshi/ebiten"

	"github.com/zenwerk/go-pixelman3/camera"
	"github.com/zenwerk/go-pixelman3/field"
	"github.com/zenwerk/go-pixelman3/sprite"
)

type Stage struct {
	Field  *field.Field
	Player *sprite.Player
	key    SceneKey
	next   SceneKey
}

func NewStage(level string, key, nextScene SceneKey) *Stage {
	st := &Stage{}
	st.Field, st.Player = field.NewField(level)
	st.key = key
	st.next = nextScene
	return st
}

func (s *Stage) Update(game *Game) {
	game.Camera.MaxWidth = s.Field.Width
	game.Camera.MaxHeight = s.Field.Height
	game.Camera.Move(s.Player.Position.X, s.Player.Position.Y)

	s.Player.Move(s.Field.Sprites, s.Field.Width, s.Field.Height)
	s.Player.Action()
	s.Player.Balls.Move(game.Camera)

	// 次のステージへ進む
	if s.Player.State.ArrivedAtNextPoint {
		game.CurrentScene = s.next
	}
}

func (s *Stage) Draw(screen *ebiten.Image, camera *camera.Camera) {
	s.Player.DrawImage(screen, camera)
	for _, ball := range s.Player.Balls {
		ball.DrawImage(screen, camera)
	}
	s.Field.DrawImage(screen, camera)
}

func (s *Stage) SceneKey() SceneKey {
	return s.key
}

func (s *Stage) GetPlayer() *sprite.Player {
	return s.Player
}
