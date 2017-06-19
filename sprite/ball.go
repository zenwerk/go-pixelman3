package sprite

import (
	"image"

	"github.com/hajimehoshi/ebiten"
	"github.com/zenwerk/go-pixelman3/utils"
)

const (
	ball_img = `--+++--
-+++++-
-+++++-
--+++--`
	ballSpeed    = 4
	screenWidth  = 320
	screenHeight = 240
)

type Balls []*Ball

var (
	ballImg *ebiten.Image
)

func init() {
	tmpImage := image.NewRGBA(image.Rect(0, 0, 7, 4))
	utils.CreateImageFromString(ball_img, tmpImage)
	ballImg, _ = ebiten.NewImage(7, 4, ebiten.FilterNearest)
	ballImg.ReplacePixels(tmpImage.Pix)
}

type Ball struct {
	BaseSprite
}

func NewBall(pos Position) *Ball {
	ball := new(Ball)
	ball.Images = []*ebiten.Image{
		ballImg,
	}
	ball.ImageNum = len(ball.Images)
	ball.Position = pos
	return ball
}

func (bs *Balls) Move(viewport Position) {
	balls := *bs

	for i := 0; i < len(balls); i++ {
		b := balls[i]
		b.Position.X += ballSpeed

		// 表示領域外に出たら配列から削除する
		if b.Position.X > (screenWidth-viewport.X) || b.Position.Y > (screenHeight-viewport.Y) || b.Position.X < 0 || b.Position.Y < 0 {
			balls = append(balls[:i], balls[i+1:]...)
			i--
		}
	}
	*bs = balls
}
