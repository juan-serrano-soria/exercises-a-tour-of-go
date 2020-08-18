package main

import (
	"image"
	"image/color"
	
	"golang.org/x/tour/pic"
)

// Image implements a rectangular grid of colors
// with a width and a height.
type Image struct {
	w, h int
}

// ColorModel returns the color model used in Image.
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds returns the domain for which At returns a non-zero value,
// where our pic is populated in.
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.w, i.h)
}

// At returns the color of the pixel at coords x, y.
func (i Image) At(x, y int) color.Color {
	return color.RGBA{uint8(x^y), uint8(x^y), 255, 255}
}

func main() {
	m := Image{250, 250}
	pic.ShowImage(m)
}
