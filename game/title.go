package game

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/freetype/truetype"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/text"
	"golang.org/x/image/font"

	"github.com/zenwerk/go-pixelman3/camera"
	"github.com/zenwerk/go-pixelman3/utils"
)

const (
	start int = iota
	exit
)

var (
	mplusNormalFont font.Face
)

func init() {
	f, err := ebitenutil.OpenFile("_resources/fonts/mplus-1p-regular.ttf")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}

	tt, err := truetype.Parse(b)
	if err != nil {
		log.Fatal(err)
	}

	const dpi = 72
	mplusNormalFont = truetype.NewFace(tt, &truetype.Options{
		Size:    24,
		DPI:     dpi,
		Hinting: font.HintingFull,
	})
}

type Title struct {
	current int
}

func NewTitle() *Title {
	return &Title{
		current: start,
	}
}

func (t *Title) Update(game *Game) {
	if ebiten.IsKeyPressed(ebiten.KeyUp) && t.current == exit {
		t.current = start
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) && t.current == start {
		t.current = exit
	}
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		switch t.current {
		case start:
			game.CurrentScene = stage1
		case exit:
			os.Exit(0)
		}
	}
}

func (t *Title) Draw(screen *ebiten.Image, _ *camera.Camera) {
	startColor := utils.White
	exitColor := utils.White
	if t.current == start {
		startColor = utils.Red
	} else if t.current == exit {
		exitColor = utils.Red
	}

	text.Draw(screen, "start", mplusNormalFont, 100, 100, startColor)
	text.Draw(screen, "exit", mplusNormalFont, 100, 150, exitColor)
}
