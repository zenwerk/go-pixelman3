package sprite

import (
	"image"

	"github.com/hajimehoshi/ebiten"

	"github.com/zenwerk/go-pixelman3/camera"
	"github.com/zenwerk/go-pixelman3/utils"
)

const (
	coin_img = `-----++++++-----
---++++++++++---
--++++++++++++--
-++++++++++++++-
-++++++++++++++-
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
-++++++++++++++-
-++++++++++++++-
--++++++++++++--
---++++++++++---
-----++++++-----`
	coinWidth  = 16
	coinHeight = 16
)

var (
	coinImg *ebiten.Image
)

type Coin struct {
	BaseSprite
	Alive bool
}

func init() {
	tmpImage := image.NewRGBA(image.Rect(0, 0, coinWidth, coinHeight))
	utils.CreateImageFromString(coin_img, tmpImage, utils.Yellow)
	coinImg, _ = ebiten.NewImage(coinWidth, coinHeight, ebiten.FilterNearest)
	coinImg.ReplacePixels(tmpImage.Pix)
}

func NewCoin() *Coin {
	coin := new(Coin)
	coin.Images = []*ebiten.Image{
		coinImg,
	}
	coin.ImageNum = len(coin.Images)
	coin.Alive = true
	return coin
}

func (c *Coin) DrawImage(screen *ebiten.Image, camera *camera.Camera) {
	if c.Alive {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(float64(c.Position.X+camera.X), float64(c.Position.Y+camera.Y))
		screen.DrawImage(c.currentImage(), op)
	}
}
