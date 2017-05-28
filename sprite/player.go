package sprite

import "github.com/hajimehoshi/ebiten"

// isOverlap は x1-x2 の範囲の整数が x3-x4 の範囲と重なるかを判定する
func isOverlap(x1, x2, x3, x4 int) bool {
	if x1 <= x4 && x2 >= x3 {
		return true
	}
	return false
}

type Player struct {
	BaseSprite
}

func NewPlayer(images []*ebiten.Image) *Player {
	player := new(Player)
	player.Images = images
	player.ImageNum = len(images)
	return player
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
		dy = -1
		p.count++
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		dy = 1
		p.count++
	}

	for _, object := range objects {
		dx, dy = p.IsCollide(dx, dy, object)
	}

	p.Position.X += dx
	p.Position.Y += dy
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
		}
	}

	return dx, dy
}
