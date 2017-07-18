package sprite

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten"
	log "github.com/sirupsen/logrus"

	"github.com/zenwerk/go-pixelman3/utils"
)

const (
	xLeftLimit  = 16 * 3         // 左方向移動の画面上の限界
	xRightLimit = 320 - (16 * 3) // 右方向移動の画面上の限界
	yUpperLimit = 16 * 2         // 上方向移動の画面上の限界
	yLowerLimit = 240 - (16 * 2) // 下方向移動の画面上の限界

	charWidth  = 16
	charHeight = 16

	player_anim0 = `------+++++-----
----+++++++++---
---+++++++++++--
--+++++++++++++-
--++++--+++--++-
-+++++--+++--++-
+++++++++++++++-
++++++++++++++++
++++++++++++++++
++++++-+++++-+++
+++++++-----++++
+-++++++++++++++
--++++++++++++-+
---++++++++++---
---++-----++----
--+++++--+++++--`

	player_anim1 = `------+++++-----
----+++++++++---
---+++++++++++--
--+++++++++++++-
--++++--+++--++-
-+++++--+++--++-
+++++++++++++++-
++++++++++++++++
++++++++++++++++
++++++-+++++-+++
+++++++-----++++
+-++++++++++++++
--++++++++++++-+
---++++++++++---
--+++++---++----
---------+++++--`

	player_anim2 = `------+++++-----
----+++++++++---
---+++++++++++--
--+++++++++++++-
--++++--+++--++-
-+++++--+++--++-
+++++++++++++++-
++++++++++++++++
++++++++++++++++
++++++-+++++-+++
+++++++-----++++
+-++++++++++++++
--++++++++++++-+
---++++++++++---
---++----+++++--
--+++++---------`
)

var (
	playerAnim0 *ebiten.Image
	playerAnim1 *ebiten.Image
	playerAnim2 *ebiten.Image
)

func init() {
	tmpImage := image.NewRGBA(image.Rect(0, 0, charWidth, charHeight))

	utils.CreateImageFromString(player_anim0, tmpImage, utils.White)
	playerAnim0, _ = ebiten.NewImage(charWidth, charHeight, ebiten.FilterNearest)
	playerAnim0.ReplacePixels(tmpImage.Pix)

	utils.CreateImageFromString(player_anim1, tmpImage, utils.White)
	playerAnim1, _ = ebiten.NewImage(charWidth, charHeight, ebiten.FilterNearest)
	playerAnim1.ReplacePixels(tmpImage.Pix)

	utils.CreateImageFromString(player_anim2, tmpImage, utils.White)
	playerAnim2, _ = ebiten.NewImage(charWidth, charHeight, ebiten.FilterNearest)
	playerAnim2.ReplacePixels(tmpImage.Pix)
}

// 四捨五入関数
func round(f float64) int {
	return int(math.Floor(f + .5))
}

type Player struct {
	BaseSprite
	jumping   bool    // 現在ジャンプ中か
	jumpSpeed float64 // 現在のジャンプ力
	fallSpeed float64 // 落下速度
	Balls     Balls
}

func NewPlayer() *Player {
	player := new(Player)
	player.Images = []*ebiten.Image{
		playerAnim0,
		playerAnim1,
		playerAnim2,
	}
	player.ImageNum = len(player.Images)
	player.jumpSpeed = 0
	player.fallSpeed = 0.4
	player.keyPressed = make(map[ebiten.Key]bool)
	return player
}

func (p *Player) jump() {
	if !p.jumping {
		p.jumping = true
		p.jumpSpeed = -6
	}
}

func (p *Player) Move(objects []Sprite, viewPort *Position) {
	// dx, dy はユーザーの移動方向を保存する
	var dx, dy int
	if p.IsKeyPressed(ebiten.KeyLeft) {
		dx = -1
		p.count++
	}
	if p.IsKeyPressed(ebiten.KeyRight) {
		dx = 1
		p.count++
	}
	if p.IsKeyPressedOneTime(ebiten.KeyUp) {
		p.jump()
		p.count++
	}

	// 落下速度の計算
	if p.jumpSpeed < 5 {
		p.jumpSpeed += p.fallSpeed
	}
	dy = round(p.jumpSpeed)

	for _, object := range objects {
		p.IsCollide(object, &dx, &dy, viewPort)
	}

	// 画面上の左右の移動限界に達しているか確認する
	if p.Position.X+dx < xLeftLimit || p.Position.X+dx > xRightLimit {
		// 移動限界に達しているなら相対座標を更新する
		// 他のオブジェクトがプレイヤーが移動する方向の逆方向に進んで欲しいので反転して(-=)代入する
		viewPort.X -= dx
	} else {
		// 移動限界に達していないなら自身の絶対座標を更新する
		p.Position.X += dx
	}

	if p.Position.Y+dy < yUpperLimit || p.Position.Y+dy > yLowerLimit {
		viewPort.Y -= dy
	} else {
		p.Position.Y += dy
	}
}

func (p *Player) Action(viewPort *Position) {
	if p.IsKeyPressedOneTime(ebiten.KeySpace) {
		pos := Position{
			X: (p.Position.X - viewPort.X) + 8,
			Y: (p.Position.Y - viewPort.Y) + 4,
		}
		ball := NewBall(pos)
		p.Balls = append(p.Balls, ball)
	}
}

func (p *Player) DrawImage(screen *ebiten.Image, _ *Position) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.Position.X), float64(p.Position.Y))
	screen.DrawImage(p.currentImage(), op)
}

// IsCollide は自身が対象の object と衝突しているか判定する
func (p *Player) IsCollide(object Sprite, dx, dy *int, viewPort *Position) {
	cm := p.detectCollisions(object, dx, dy, viewPort)

	if cm.HasCollision() {
		p.Collision(object, dx, dy, cm)
	}

	return
}

func (p *Player) Collision(object Sprite, dx, dy *int, cm *CollideMap) {
	switch v := object.(type) {
	case *Block:
		p.collideBlock(v, dx, dy, cm)
	case *Coin:
		p.collideCoin(v, dx, dy, cm)
	default:
		log.Warn("unknown type")
	}
}

func (p *Player) collideBlock(_ *Block, dx, dy *int, cm *CollideMap) {
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

func (p *Player) collideCoin(c *Coin, _, _ *int, cm *CollideMap) {
	c.Alive = false
}
