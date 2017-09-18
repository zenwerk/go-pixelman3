package sprite

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten"
	log "github.com/sirupsen/logrus"

	"github.com/zenwerk/go-pixelman3/camera"
	"github.com/zenwerk/go-pixelman3/utils"
)

const (
	charWidth  = 16
	charHeight = 16

	player_anim0 = `---+++++++++++--
--++++--+++--++-
--++++--+++--++-
--++++--+++--++-
---+++++++++++--
----++++++++----
--++++++++++++--
-+++-++++++-+++-
-++--++++++--++-
-++--++++++--++-
-----++++++-----
-----++--++-----
-----++--++-----
-----++--++-----
----+++--+++----
---++++--++++---`

	player_anim1 = `---+++++++++++--
--++++--+++--++-
--++++--+++--++-
--++++--+++--++-
---+++++++++++--
----++++++++----
---++++++++++---
--++-++++++-++--
--++-++++++--++-
--++-++++++---++
-++--++++++-----
-----++--++-----
----++----++----
---++------++---
--++--------++--
--+++--------+++`

	player_anim2 = `---+++++++++++--
--++++--+++--++-
--++++--+++--++-
--++++--+++--++-
---+++++++++++--
----++++++++----
---++++++++++---
--++-++++++-++--
--++-++++++-++--
--++-++++++--++-
---++++++++---++
----++----++----
---++------++---
-+++--------++--
-+----------++--
------------+++-`
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

// state はゲーム全体に関連するプレイヤーの状態を保存する
type state struct {
	ArrivedAtNextPoint bool // 次のステージへ移動するか
	//RemainingLives     int  // 残機
	//Point              int  // 取得ポイント
}

type Player struct {
	BaseSprite
	jumping    bool    // 現在ジャンプ中か
	jumpSpeed  float64 // 現在のジャンプ力
	fallSpeed  float64 // 落下速度
	Balls      Balls
	Speed      float64 // 現在のスピード
	AccelSpeed float64 // 加速度
	MaxSpeed   float64 // 最大速度
	State      *state  // 現在の状態
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
	player.AccelSpeed = 0.25
	player.MaxSpeed = 3.0
	player.State = &state{}
	return player
}

func (p *Player) jump() {
	if !p.jumping {
		p.jumping = true
		p.jumpSpeed = -6
	}
}

func (p *Player) Move(objects []Sprite, maxX, maxY int) {
	// dx, dy はユーザーの移動方向を保存する
	var dx, dy int
	if p.IsKeyPressed(ebiten.KeyLeft) {
		if p.Speed > -p.MaxSpeed {
			p.Speed -= p.AccelSpeed
		}
		p.count++
	} else if p.IsKeyPressed(ebiten.KeyRight) {
		if p.Speed < p.MaxSpeed {
			p.Speed += p.AccelSpeed
		}
		p.count++
	} else {
		if p.Speed > 0 {
			p.Speed -= p.AccelSpeed
		} else if p.Speed < 0 {
			p.Speed += p.AccelSpeed
		}
	}
	if p.IsKeyPressedOneTime(ebiten.KeyUp) {
		p.jump()
		p.count++
	}
	dx = round(p.Speed)

	// 落下速度の計算
	if p.jumpSpeed < 5 {
		p.jumpSpeed += p.fallSpeed
	}
	dy = round(p.jumpSpeed)

	// 画面端に達していたら移動できない
	if p.Position.X+dx < 0 || p.Position.X+p.Width()+dx > maxX {
		dx = 0
	}
	if p.Position.Y+dy < 0 || p.Position.Y+p.Height()+dy > maxY {
		dy = 0
	}

	if dx != 0 {
		p.moveX(dx, objects)
	}
	if dy != 0 {
		p.moveY(dy, objects)
	}
}

func (p *Player) moveX(dx int, sprites []Sprite) {
	p.Position.X += dx
	// 衝突判定
	for _, s := range sprites {
		if p.Intersect(s) {
			p.Collision(s, dx, 0)
		}
	}
}

func (p *Player) moveY(dy int, sprites []Sprite) {
	p.Position.Y += dy
	// 衝突判定
	for _, s := range sprites {
		if p.Intersect(s) {
			p.Collision(s, 0, dy)
		}
	}
}

func (p *Player) Action() {
	if p.IsKeyPressedOneTime(ebiten.KeySpace) {
		pos := Position{
			X: (p.Position.X) + 8,
			Y: (p.Position.Y) + 4,
		}
		ball := NewBall(pos)
		p.Balls = append(p.Balls, ball)
	}
}

func (p *Player) DrawImage(screen *ebiten.Image, camera *camera.Camera) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(p.Position.X+camera.X), float64(p.Position.Y+camera.Y))
	screen.DrawImage(p.currentImage(), op)
}

func (p *Player) Collision(object Sprite, dx, dy int) {
	switch v := object.(type) {
	case *Block:
		p.collideBlock(v, dx, dy)
	case *Coin:
		p.collideCoin(v)
	case *NextPoint:
		p.collideNextPoint(v)
	default:
		log.Warn("unknown type")
	}
}

func (p *Player) collideBlock(b *Block, dx, dy int) {
	// 右に移動して衝突
	if dx > 0 {
		// プレイヤーの右端の座標がブロックの左端座標になるようにする
		p.Position.X = b.Position.X - p.Width()
	}
	// 左に移動して衝突
	if dx < 0 {
		// プレイヤーの左端の座標がブロックの右端座標になるようにする
		p.Position.X = b.Position.X + p.Width()
	}
	// 下に移動して衝突
	if dy > 0 {
		// プレイヤーの下端の座標がブロックの上端座標になるようにする
		p.Position.Y = b.Position.Y - p.Height()
		p.jumping = false
		p.jumpSpeed = 0
	}
	// 上に移動して衝突
	if dy < 0 {
		// プレイヤーの上端の座標がブロックの下端座標になるようにする
		p.Position.Y = b.Position.Y + p.Height()
	}
}

func (p *Player) collideCoin(c *Coin) {
	c.Alive = false
}

func (p *Player) collideNextPoint(np *NextPoint) {
	// 到達したら次のステージに進む
	if p.Position.X+p.Width() >= np.Position.X+np.Width() {
		p.State.ArrivedAtNextPoint = true
	}
}
