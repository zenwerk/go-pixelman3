package game

import (
	"github.com/hajimehoshi/ebiten"

	"github.com/zenwerk/go-pixelman3/camera"
)

type StateKey string

const (
	title  StateKey = "title"
	stage1 StateKey = "stage1"
)

type Game struct {
	ScreenWidth  int
	ScreenHeight int
	Camera       *camera.Camera
	CurrentState StateKey
	States       map[StateKey]State
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

	g.States = map[StateKey]State{
		title:  NewTitle(),
		stage1: NewStage1(),
	}
	g.CurrentState = title
}

func (g *Game) MainLoop(screen *ebiten.Image) error {
	g.States[g.CurrentState].Update(g)

	if ebiten.IsRunningSlowly() {
		return nil
	}

	g.States[g.CurrentState].Draw(screen, g.Camera)

	return nil
}
