package game

import (
	"github.com/hajimehoshi/ebiten"

	"github.com/zenwerk/go-pixelman3/camera"
	"github.com/zenwerk/go-pixelman3/field"
)

type SceneKey int

const (
	title SceneKey = iota
	stage1
	stage2
	ending
)

type Game struct {
	ScreenWidth  int
	ScreenHeight int
	Camera       *camera.Camera
	CurrentScene SceneKey
	Scenes       map[SceneKey]Scene
}

func NewGame(width, heigh int) *Game {
	return &Game{
		ScreenWidth:  width,
		ScreenHeight: heigh,
	}
}

func (g *Game) Init() {
	g.Camera = &camera.Camera{
		Width:  g.ScreenWidth,
		Height: g.ScreenHeight,
	}

	g.Scenes = map[SceneKey]Scene{
		title:  NewTitle(),
		stage1: NewStage(field.Level_data_1, stage2),
		stage2: NewStage(field.Level_data_2, ending),
	}
	g.CurrentScene = title
}

func (g *Game) MainLoop(screen *ebiten.Image) error {
	g.Scenes[g.CurrentScene].Update(g)

	if ebiten.IsRunningSlowly() {
		return nil
	}

	g.Scenes[g.CurrentScene].Draw(screen, g.Camera)

	return nil
}
