package plot

import (
	"image"
	"image/color"
)

// Canvas is the basis for all other drawing
// primitives. Its only properties are a width,
// a height and a background color.
// Canvas is based on `image.RGBA'.
type Canvas struct {
	img    *image.RGBA
	Width  int
	Height int
}

// At ensures that `Canvas' implements `image.Image'
func (c *Canvas) At(x, y int) color.Color {
	return c.img.At(x, y)
}

// Set ensures that `Canvas' implements `image.Image'
func (c *Canvas) Set(x, y int, col color.Color) {
	c.img.Set(x, y, col)
}

// Bounds ensures that `Canvas' implements `image.Image'
func (c *Canvas) Bounds() image.Rectangle {
	return c.img.Bounds()
}

// ColorModel ensures that `Canvas' implements `image.Image'
func (c *Canvas) ColorModel() color.Model {
	return c.img.ColorModel()
}

// NewCanvas creates a new canvas for plotting.
func NewCanvas(width, height int, bg *color.RGBA) *Canvas {
	// create a new rectangular `Canvas'
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	c := &Canvas{
		img:    img,
		Width:  width,
		Height: height,
	}

	// set background color of `Canvas'
	b := c.img.Bounds()
	for y := b.Min.Y; y < b.Max.Y; y++ {
		for x := b.Min.X; x < b.Max.X; x++ {
			c.img.Set(x, y, bg)
		}
	}

	return c
}

// Rectangle creates a rectangle with a certain
// color between points (x0, y0) and (x1, y1) on
// a `Canvas'. `outline' color and `thickness' can
// be specified as well.
func (c *Canvas) Rectangle(x0, y0, x1, y1, thickness int, bg, outline *color.RGBA) {
}
