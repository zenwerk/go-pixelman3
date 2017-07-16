package sprite

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten"

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

// isOverlap は x1-x2 の範囲の整数が x3-x4 の範囲と重なるかを判定する
func isOverlap(x1, x2, x3, x4 int) bool {
	if x1 <= x4 && x2 >= x3 {
		return true
	}
	return false
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
		p.IsCollide(&dx, &dy, object, viewPort)
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

// TODO: 衝突判定は true/false を返す関数にする
// TODO: 衝突後になにが起きるかは別関数(各構造体に持つのが良さそう)にする
// IsCollide はプレイヤーが対象の object と衝突しているか判定する
func (p *Player) IsCollide(dx, dy *int, object Sprite, viewPort *Position) {
	var cm CollideMap
	// プレイヤーの座標
	x := p.Position.X // x座標の位置
	y := p.Position.Y // y座標の位置
	img := p.currentImage()
	w, h := img.Size() // プレイヤーの幅と高さ

	// 対象のオブジェクトの x,y座標の位置と幅と高さを取得する
	x1, y1, w1, h1 := object.GetCoordinates()

	// 対象オブジェクトは相対座標付与して衝突判定を行う
	x1 += viewPort.X
	y1 += viewPort.Y

	overlappedX := isOverlap(x, x+w, x1, x1+w1) // x軸で重なっているか
	overlappedY := isOverlap(y, y+h, y1, y1+h1) // y軸で重なっているか

	if overlappedY {
		if *dx < 0 && x+*dx <= x1+w1 && x+w+*dx >= x1 {
			// 左方向の移動の衝突判定
			cm.Left = true
		} else if *dx > 0 && x+w+*dx >= x1 && x+*dx <= x1+w1 {
			// 右方向の移動の衝突判定
			cm.Right = true
		}
	}
	if overlappedX {
		if *dy < 0 && y+*dy <= y1+w1 && y+h+*dy >= y1 {
			// 上方向の移動の衝突判定
			cm.Top = true
		} else if *dy > 0 && y+h+*dy >= y1 && y+*dy <= y1+h1 {
			// 下方向の移動の衝突判定
			cm.Bottom = true
		}
	}

	if cm.HasCollision() {
		object.Collision(p, dx, dy, &cm)
	}

	return
}

func (p *Player) DrawImage(screen *ebiten.Image, _ *Position) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.Position.X), float64(p.Position.Y))
	screen.DrawImage(p.currentImage(), op)
}
