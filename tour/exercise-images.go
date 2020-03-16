// Remember the picture generator you wrote earlier?
// (https://tour.golang.org/moretypes/18)
//
// Let's write another one, but this time it will return an
// implementation of image.Image instead of a slice of data.
//
// Define your own Image type, implement the necessary methods,
// and call pic.ShowImage.
//
// - Bounds should return a image.Rectangle, like
// image.Rect(0, 0, w, h).
// - ColorModel should return color.RGBAModel.
// - At should return a color; the value v in the last picture
// generator corresponds to color.RGBA{v, v, 255, 255} in this one.
package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	w int
	h int
}

// ColorModel returns the Image's color model.
func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

// Bounds returns the domain for which At can return non-zero color.
// The bounds do not necessarily contain the point (0, 0).
func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.h, i.w)
}

// At returns the color of the pixel at (x, y).
// At(Bounds().Min.X, Bounds().Min.Y) returns the upper-left pixel of the grid.
// At(Bounds().Max.X-1, Bounds().Max.Y-1) returns the lower-right one.
func (i Image) At(x, y int) color.Color {
	var c uint8 = uint8((x + y) / 2)
	return color.RGBA{c, c, 255, 255}
}

func main() {
	m := Image{250, 500}
	pic.ShowImage(m)
}
