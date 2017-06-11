package utils

import (
	"image"
	"strings"
)

func CreateImageFromString(charString string, img *image.RGBA) {
	width := img.Rect.Size().X
	for indexY, line := range strings.Split(charString, "\n") {
		for indexX, str := range line {
			pos := 4*indexY*width + 4*indexX
			if string(str) == "+" {
				img.Pix[pos] = 0xff   // R
				img.Pix[pos+1] = 0xff // G
				img.Pix[pos+2] = 0xff // B
				img.Pix[pos+3] = 0xff // A
			} else {
				img.Pix[pos] = 0
				img.Pix[pos+1] = 0
				img.Pix[pos+2] = 0
				img.Pix[pos+3] = 0
			}
		}
	}
}
