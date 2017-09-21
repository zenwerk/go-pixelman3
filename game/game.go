package game

import (
	"fmt"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"

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

const (
	defaultLives = 5
)

type playerState struct {
	RemainingLives int       // 残機
	Point          int       // 取得ポイント
	StartTime      time.Time // 開始時刻
}

type Game struct {
	ScreenWidth  int
	ScreenHeight int
	Camera       *camera.Camera
	CurrentScene SceneKey
	Scenes       map[SceneKey]Scene
	PState       *playerState
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

	g.PState = &playerState{
		RemainingLives: defaultLives,
		StartTime:      time.Now(),
	}

	g.Scenes = map[SceneKey]Scene{
		title:  NewTitle(),
		stage1: NewStage(field.Level_data_1, stage2),
		stage2: NewStage(field.Level_data_2, ending),
		ending: NewEnding(),
	}
	g.CurrentScene = title
}

func (g *Game) DrawStatus(screen *ebiten.Image) {
	stage := ""
	switch g.CurrentScene {
	case stage1:
		stage = "1"
	case stage2:
		stage = "2"
	}

	now := time.Now().Truncate(time.Nanosecond)
	elapsedTime := now.Sub(g.PState.StartTime.Truncate(time.Nanosecond))

	str := fmt.Sprintf(" Score:%d        Lives:%d        State:%s      Time:%.0f",
		g.PState.Point, g.PState.RemainingLives, stage, elapsedTime.Seconds())
	ebitenutil.DebugPrint(screen, str)
}

func (g *Game) MainLoop(screen *ebiten.Image) error {
	g.Scenes[g.CurrentScene].Update(g)

	if ebiten.IsRunningSlowly() {
		return nil
	}

	g.Scenes[g.CurrentScene].Draw(screen, g.Camera)
	if g.CurrentScene != title && g.CurrentScene != ending {
		g.DrawStatus(screen)
	}

	return nil
}
