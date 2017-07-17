package sprite

import (
	"image"

	"github.com/hajimehoshi/ebiten"

	"github.com/zenwerk/go-pixelman3/utils"
)

const (
	block_img = `++++++++++++++++
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
	blockWidth  = 16
	blockHeight = 16
)

var (
	blockImg *ebiten.Image
)

type Block struct {
	BaseSprite
}

func init() {
	tmpImage := image.NewRGBA(image.Rect(0, 0, blockWidth, blockHeight))
	utils.CreateImageFromString(block_img, tmpImage, utils.Brown)
	blockImg, _ = ebiten.NewImage(blockWidth, blockHeight, ebiten.FilterNearest)
	blockImg.ReplacePixels(tmpImage.Pix)
}

func NewBlock() *Block {
	block := new(Block)
	block.Images = []*ebiten.Image{
		blockImg,
	}
	block.ImageNum = len(block.Images)
	return block
}
