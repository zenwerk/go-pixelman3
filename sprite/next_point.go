package sprite

import (
	"image"

	"github.com/hajimehoshi/ebiten"

	"github.com/zenwerk/go-pixelman3/utils"
)

const (
	nextpoint_img = `++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++
++++++++++++++++`
	nextPointWidth  = 16
	nextPointHeight = 16
)

var (
	nextPointImg *ebiten.Image
)

type NextPoint struct {
	BaseSprite
}

func init() {
	tmpImage := image.NewRGBA(image.Rect(0, 0, nextPointWidth, nextPointHeight))
	utils.CreateImageFromString(nextpoint_img, tmpImage, utils.Transparent)
	nextPointImg, _ = ebiten.NewImage(nextPointWidth, nextPointHeight, ebiten.FilterNearest)
	nextPointImg.ReplacePixels(tmpImage.Pix)
}

func NewNextPoint() *NextPoint {
	nextPoint := new(NextPoint)
	nextPoint.Images = []*ebiten.Image{
		nextPointImg,
	}
	nextPoint.ImageNum = len(nextPoint.Images)
	return nextPoint
}
