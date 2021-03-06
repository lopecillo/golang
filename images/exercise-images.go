package main

import (
	"image"
	"image/color"
	"golang.org/x/tour/pic"
)

type Image struct {
	Width, Height int
}

func (m Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (m Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, m.Width, m.Height)
}

func (m Image) At(x, y int) color.Color {
	c := uint8(x ^ y)
	return color.RGBA{c, c, 255, 255}
}

func main() {
	m := Image{512, 256}
	pic.ShowImage(m)
}