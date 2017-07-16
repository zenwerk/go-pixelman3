package field

import (
	"strings"

	"github.com/hajimehoshi/ebiten"

	"github.com/zenwerk/go-pixelman3/sprite"
)

const (
	width  = 17
	height = 17

	blockMark  = "+"
	playerMark = "P"
	coinMark   = "C"
)

type Field struct {
	Sprites []sprite.Sprite
}

func NewField(fieldData string) (*Field, *sprite.Player) {
	player := new(sprite.Player)
	field := new(Field)

	for indexY, line := range strings.Split(fieldData, "\n") {
		for indexX, str := range line {
			switch string(str) {
			case blockMark:
				block := sprite.NewBlock()
				block.Position.X = indexX * width
				block.Position.Y = indexY * height
				field.Sprites = append(field.Sprites, block)
			case playerMark:
				player = sprite.NewPlayer()
				player.Position.X = indexX * width
				player.Position.Y = indexY * height
			case coinMark:
				coin := sprite.NewCoin()
				coin.Position.X = indexX * width
				coin.Position.Y = indexY * width
				field.Sprites = append(field.Sprites, coin)
			}
		}
	}

	return field, player
}

func (f *Field) DrawImage(screen *ebiten.Image, viewport sprite.Position) {
	for _, sprite := range f.Sprites {
		sprite.DrawImage(screen, viewport)
	}
}
