package game

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/zenwerk/go-pixelman3/camera"
)

// Scene は現在のゲームの状態を表す
type Scene interface {
	Update(game *Game)
	Draw(image *ebiten.Image, camera *camera.Camera)
}
