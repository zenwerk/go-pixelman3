package sprite

import (
	"math"

	"github.com/hajimehoshi/ebiten"
)

const (
	xLeftLimit  = 16 * 3         // 左方向移動の画面上の限界
	xRightLimit = 320 - (16 * 3) // 右方向移動の画面上の限界
	yUpperLimit = 16 * 2         // 上方向移動の画面上の限界
	yLowerLimit = 240 - (16 * 2) // 下方向移動の画面上の限界
)

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
	jumping   bool     // 現在ジャンプ中か
	jumpSpeed float64  // 現在のジャンプ力
	fallSpeed float64  // 落下速度
	ViewPort  position // スクリーン上の相対座標
}

func NewPlayer(images []*ebiten.Image) *Player {
	player := new(Player)
	player.Images = images
	player.ImageNum = len(images)
	player.jumpSpeed = 0
	player.fallSpeed = 0.4
	return player
}

func (p *Player) jump() {
	if !p.jumping {
		p.jumping = true
		p.jumpSpeed = -6
	}
}

func (p *Player) Move(objects []Sprite) {
	// dx, dy はユーザーの移動方向を保存する
	var dx, dy int
	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		dx = -1
		p.count++
	}
	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		dx = 1
		p.count++
	}
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.jump()
		p.count++
	}

	// 落下速度の計算
	if p.jumpSpeed < 5 {
		p.jumpSpeed += p.fallSpeed
	}
	dy = round(p.jumpSpeed)

	for _, object := range objects {
		dx, dy = p.IsCollide(dx, dy, object)
	}

	// 画面上の左右の移動限界に達しているか確認する
	if p.Position.X+dx < xLeftLimit || p.Position.X+dx > xRightLimit {
		// 移動限界に達しているなら相対座標を更新する
		// 他のオブジェクトがプレイヤーが移動する方向の逆方向に進んで欲しいので反転して(-=)代入する
		p.ViewPort.X -= dx
	} else {
		// 移動限界に達していないなら自身の絶対座標を更新する
		p.Position.X += dx
	}

	if p.Position.Y+dy < yUpperLimit || p.Position.Y+dy > yLowerLimit {
		p.ViewPort.Y -= dy
	} else {
		p.Position.Y += dy
	}
}

// IsCollide はプレイヤーが対象の object と衝突しているか判定する
func (p *Player) IsCollide(dx, dy int, object Sprite) (int, int) {
	// プレイヤーの座標
	x := p.Position.X // x座標の位置
	y := p.Position.Y // y座標の位置
	img := p.currentImage()
	w, h := img.Size() // プレイヤーの幅と高さ

	// 対象のオブジェクトの x,y座標の位置と幅と高さを取得する
	x1, y1, w1, h1 := object.GetCordinates()

	// 対象オブジェクトは相対座標付与して衝突判定を行う
	x1 += p.ViewPort.X
	y1 += p.ViewPort.Y

	overlappedX := isOverlap(x, x+w, x1, x1+w1) // x軸で重なっているか
	overlappedY := isOverlap(y, y+h, y1, y1+h1) // y軸で重なっているか

	if overlappedY {
		if dx < 0 && x+dx <= x1+w1 && x+w+dx >= x1 {
			// 左方向の移動の衝突判定
			// 衝突していたらx軸の移動速度を 0 にする
			dx = 0
		} else if dx > 0 && x+w+dx >= x1 && x+dx <= x1+w1 {
			// 右方向の移動の衝突判定
			// 衝突していたらx軸の移動速度を 0 にする
			dx = 0
		}
	}
	if overlappedX {
		if dy < 0 && y+dy <= y1+w1 && y+h+dy >= y1 {
			// 上方向の移動の衝突判定
			// 衝突していたらy軸の移動速度を 0 にする
			dy = 0
		} else if dy > 0 && y+h+dy >= y1 && y+dy <= y1+h1 {
			// 下方向の移動の衝突判定
			// 衝突していたらy軸の移動速度を 0 にする
			dy = 0

			// ジャンプ中フラグをオフにする
			p.jumping = false
			p.jumpSpeed = 0
		}
	}

	return dx, dy
}

func (p *Player) DrawImage(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.Position.X), float64(p.Position.Y))
	screen.DrawImage(p.currentImage(), op)
}
