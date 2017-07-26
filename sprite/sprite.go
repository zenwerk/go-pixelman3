package sprite

import (
	"github.com/hajimehoshi/ebiten"
	log "github.com/sirupsen/logrus"

	"github.com/zenwerk/go-pixelman3/camera"
)

type Sprite interface {
	GetCoordinates() (int, int, int, int)
	DrawImage(*ebiten.Image, *camera.Camera)
	Collision(Sprite, *int, *int, *CollideMap)
}

type Position struct {
	X int
	Y int
}

// CollideMap は Sprite が上下左右で衝突しているかを表す
type CollideMap struct {
	Left   bool
	Right  bool
	Top    bool
	Bottom bool
}

func (cm *CollideMap) HasCollision() bool {
	return cm.Left || cm.Right || cm.Top || cm.Bottom
}

type BaseSprite struct {
	Images     []*ebiten.Image     // アニメーションさせる画像の配列
	ImageNum   int                 // 総イメージ数
	CurrentNum int                 // 現在何枚目の画像が表示されているか
	Position   Position            // 現在表示されている位置(左上)
	count      int                 // フレーム数のカウンター
	keyPressed map[ebiten.Key]bool // 現在なんのキーが押されているか
}

func NewSprite(images []*ebiten.Image) *BaseSprite {
	return &BaseSprite{
		Images:     images,
		ImageNum:   len(images),
		keyPressed: make(map[ebiten.Key]bool),
	}
}

// currentImage は現在表示する画像を選択して返す
func (s *BaseSprite) currentImage() *ebiten.Image {
	// 毎フレーム画像を更新するとアニメーションが早すぎるため
	// 5フレーム毎に画像を更新する
	if s.count > 5 {
		s.count = 0
		s.CurrentNum++
		s.CurrentNum %= s.ImageNum
	}
	return s.Images[s.CurrentNum]
}

// IsKeyPressed は指定したキーが現在押下されているか判断し返す
func (s *BaseSprite) IsKeyPressed(key ebiten.Key) bool {
	pressed := ebiten.IsKeyPressed(key)
	if pressed {
		s.keyPressed[key] = true
	}
	return pressed
}

// IsKeyReleased は指定されたキーが現在押下されていないか判断し返す
func (s *BaseSprite) IsKeyReleased(key ebiten.Key) bool {
	// キーは押されていると記録されていたが, 実際にはキーは押されていない -> キーがリリースされた
	if s.keyPressed[key] && !ebiten.IsKeyPressed(key) {
		s.keyPressed[key] = false
		return true
	}
	return false
}

// IsKeyPressedOneTime はキーが押下された初回のみ true を返す
func (s *BaseSprite) IsKeyPressedOneTime(key ebiten.Key) bool {
	pressed := s.keyPressed[key]
	if pressed && !s.IsKeyReleased(key) {
		return false
	}
	return s.IsKeyPressed(key)
}

func (s *BaseSprite) DrawImage(screen *ebiten.Image, camera *camera.Camera) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(s.Position.X+camera.X), float64(s.Position.Y+camera.Y))
	screen.DrawImage(s.currentImage(), op)
}

func (s *BaseSprite) GetCoordinates() (int, int, int, int) {
	w, h := s.currentImage().Size()
	return s.Position.X, s.Position.Y, w, h
}

// isOverlap は x1-x2 の範囲の整数が x3-x4 の範囲と重なるかを判定する
func isOverlap(x1, x2, x3, x4 int) bool {
	if x1 <= x4 && x2 >= x3 {
		return true
	}
	return false
}

func (s *BaseSprite) detectCollisions(object Sprite, dx, dy *int, camera *camera.Camera) *CollideMap {
	var cm CollideMap
	// 自身の座標
	x := s.Position.X // x座標の位置
	y := s.Position.Y // y座標の位置
	img := s.currentImage()
	w, h := img.Size() // 自身の幅と高さ

	// 対象のオブジェクトの x,y座標の位置と幅と高さを取得する
	x1, y1, w1, h1 := object.GetCoordinates()

	// 対象オブジェクトは相対座標付与して衝突判定を行う
	x1 += camera.X
	y1 += camera.Y

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

	return &cm
}

// IsCollide は自身が対象の object と衝突しているか判定する
func (s *BaseSprite) IsCollide(object Sprite, dx, dy *int, camera *camera.Camera) {
	log.Info("overwrite this method.")
}

func (s *BaseSprite) Collision(object Sprite, dx, dy *int, cm *CollideMap) {
	log.Info("overwrite this method.")
}
