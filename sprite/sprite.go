package sprite

import (
	"github.com/hajimehoshi/ebiten"
	log "github.com/sirupsen/logrus"
)

type Sprite interface {
	GetCoordinates() (int, int, int, int)
	DrawImage(*ebiten.Image, *Position)
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

func (s *BaseSprite) DrawImage(screen *ebiten.Image, viewPort *Position) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(s.Position.X+viewPort.X), float64(s.Position.Y+viewPort.Y))
	screen.DrawImage(s.currentImage(), op)
}

func (s *BaseSprite) GetCoordinates() (int, int, int, int) {
	w, h := s.currentImage().Size()
	return s.Position.X, s.Position.Y, w, h
}

func (s *BaseSprite) Collision(object Sprite, dx, dy *int, cm *CollideMap) {
	log.Info("overwrite this method.")
}
