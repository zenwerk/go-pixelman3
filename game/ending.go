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
	restart int = iota
	end
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

type Ending struct {
	key     SceneKey
	current int
}

func NewEnding() *Ending {
	return &Ending{
		key:     ending,
		current: start,
	}
}

func (e *Ending) Update(game *Game) {
	if ebiten.IsKeyPressed(ebiten.KeyUp) && e.current == end {
		e.current = restart
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) && e.current == restart {
		e.current = end
	}
	if ebiten.IsKeyPressed(ebiten.KeyEnter) {
		switch e.current {
		case restart:
			game.Init()
		case end:
			os.Exit(0)
		}
	}
}

func (e *Ending) Draw(screen *ebiten.Image, _ *camera.Camera) {
	restartColor := utils.White
	endColor := utils.White
	if e.current == restart {
		restartColor = utils.Red
	} else if e.current == end {
		endColor = utils.Red
	}

	text.Draw(screen, "Clear! Congratulations!", mplusNormalFont, 20, 50, utils.Yellow)
	text.Draw(screen, "restart", mplusNormalFont, 100, 100, restartColor)
	text.Draw(screen, "end", mplusNormalFont, 100, 150, endColor)
}

func (e *Ending) SceneKey() SceneKey {
	return e.key
}
