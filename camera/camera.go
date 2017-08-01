package camera

type Camera struct {
	X int // 現在のカメラのX座標
	Y int // 現在のカメラのY座標

	Width     int // カメラの幅
	Height    int // カメラの高さ
	MaxWidth  int // カメラの最大幅
	MaxHeight int // カメラの最大幅高さ

	//XLeftLimit  int // 左方向移動の画面上の限界
	//XRightLimit int // 右方向移動の画面上の限界
	//YUpperLimit int // 上方向移動の画面上の限界
	//YLowerLimit int // 下方向移動の画面上の限界
}

func (c *Camera) SimpleMove(x, y int) {
	c.X = (c.Width / 2) - x
	c.Y = (c.Height / 2) - y
}

func (c *Camera) Move(x, y int) {
	maxXOffset := -(c.MaxWidth - c.Width)
	maxYOffset := -(c.MaxHeight - c.Height)

	restX := (c.Width / 2) - x
	restY := (c.Height / 2) - y

	if restX > 0 {
		c.X = 0
	} else if restX < maxXOffset {
		c.X = maxXOffset
	} else {
		c.X = restX
	}

	if restY > 0 {
		c.Y = 0
	} else if restY < maxYOffset {
		c.Y = maxYOffset
	} else {
		c.Y = restY
	}
}
