package sprite

import (
	"image"

	log "github.com/sirupsen/logrus"
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

func (b *Block) Collision(object Sprite, dx, dy *int, cm *CollideMap) {
	switch v := object.(type) {
	case *Player:
		b.collidePlayer(v, dx, dy, cm)
	default:
		log.Warn("unknown type")
	}
}

func (b *Block) collidePlayer(p *Player, dx, dy *int, cm *CollideMap) {
	if cm.Left || cm.Right {
		*dx = 0
	}
	if cm.Top {
		*dy = 0
	}
	if cm.Bottom {
		*dy = 0
		// ジャンプ中フラグをオフにする
		p.jumping = false
		p.jumpSpeed = 0
	}
}
