package sprite

import (
	"image"

	"github.com/hajimehoshi/ebiten"

	"github.com/zenwerk/go-pixelman3/utils"
)

const (
	spikes_img = `---+---+----+---
---+---+----+---
---+---+----+---
--++--+++--+++--
--++--+++--+++--
--++--+++--+++--
--+++-++++-+++--
--+++-++++-+++--
--+++-++++-++++-
-++++-++++-++++-
-++++-++++-++++-
-++++++++++++++-
-++++++++++++++-
+++++++++++++++-
++++++++++++++++
++++++++++++++++`
	spikeWidth  = 16
	spikeHeight = 16
)

var (
	spikesImg *ebiten.Image
)

type Spikes struct {
	BaseSprite
}

func init() {
	tmpImage := image.NewRGBA(image.Rect(0, 0, spikeWidth, spikeHeight))
	utils.CreateImageFromString(spikes_img, tmpImage, utils.Red)
	spikesImg, _ = ebiten.NewImage(spikeWidth, spikeHeight, ebiten.FilterNearest)
	spikesImg.ReplacePixels(tmpImage.Pix)
}

func NewSpikes() *Spikes {
	spikes := new(Spikes)
	spikes.Images = []*ebiten.Image{
		spikesImg,
	}
	spikes.ImageNum = len(spikes.Images)
	return spikes
}
