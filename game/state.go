package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/zenwerk/go-pixelman3/camera"
)

// State は現在のゲームの状態を表す
type State interface {
	Update(game *Game)
	Draw(image *ebiten.Image, camera *camera.Camera)
}
