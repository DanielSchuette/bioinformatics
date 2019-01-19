package plot

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"io"
	"math"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
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

// EncodePNG encodes a `Canvas' and everything that is
// drawn on it as a .png using an `io.Writer'.
func (c *Canvas) EncodePNG(w io.Writer) error {
	if err := png.Encode(w, c); err != nil {
		return err
	}
	return nil
}

// SaveToFile saves a `Canvas' to a file at `path'. The
// file does not need to exist (if it does, it will be
// overwritten).
func (c *Canvas) SaveToFile(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return fmt.Errorf("primitives: unable to create %s: %v", name, err)
	}
	defer f.Close()
	if err := c.EncodePNG(f); err != nil {
		return fmt.Errorf("primitives: unable to encode canvas to png: %v", err)
	}
	return nil
}

// Line draws a straight line on a `Canvas'. The line
// will be drawn from (`x0', `y0') to (`x1', `y1') with
// a certain thickness in pixels and a certain color.
func (c *Canvas) Line(x0, y0, x1, y1, thick int, col *color.RGBA) error {
	if (x0 > x1) || (y0 > y1) {
		return errors.New("primitives: x0,y0 must be smaller than x1,y1")
	}

	// edge case of a purely vertical line
	// the case of a horizontal line with
	// slope 0 is covered by the general
	// form of the algorithm
	if x0 == x1 {
		for y := y0; y <= y1; y++ {
			for tx := -thick; tx <= thick; tx++ {
				for ty := -thick; ty <= thick; ty++ {
					c.Set(x0+tx, y+ty, col)
				}
			}
		}
		return nil
	}

	// calculate actual linear equation
	slope := (float64(y1) - float64(y0)) /
		(float64(x1) - float64(x0)) /* calculate slope of line */
	intercept := float64(y0) -
		(slope * float64(x0)) /* calculate intercept of line */
	for x := x0; x <= x1; x++ {
		y := int(math.Round((float64(x) * slope) + intercept))
		for tx := -thick; tx <= thick; tx++ {
			for ty := -thick; ty <= thick; ty++ {
				c.Set(x+tx, y+ty, col)
			}
		}
	}
	return nil
}

// Rectangle creates a rectangle with a certain outline
// color between points (`x0', `y0') and (`x1', `y1') on
// a `Canvas'. Thickness `thick' can be specified as well.
func (c *Canvas) Rectangle(x0, y0, x1, y1, thick int, out *color.RGBA) {
	// draw horizontal and vertical lines
	// according to `thickness'
	var t, x, y int
	for t = 0; t < thick; t++ {
		// TODO: replace with `Line' implementation
		// horizontal lines
		for x = x0; x <= x1; x++ {
			c.Set(x, y0+t, out)
			c.Set(x, y1-t, out)
		}

		// vertical lines
		for y = y0; y <= y1; y++ {
			c.Set(x0+t, y, out)
			c.Set(x1-t, y, out)
		}
	}
}

// AddLabel adds a `label' at a certain `x' and `y' position
// of a `Canvas'. Currently, only a fixed-size font is
// implemented (`basicfont.Face7x13').
func (c *Canvas) AddLabel(x, y int, label string, col *color.RGBA) {
	// TODO: add a variable-size font
	point := fixed.Point26_6{fixed.Int26_6(x * 64), fixed.Int26_6(y * 64)}
	face := basicfont.Face7x13 /* fixed-size font */
	d := &font.Drawer{
		Dst:  c,
		Src:  image.NewUniform(col),
		Face: face,
		Dot:  point,
	}
	d.DrawString(label)
}
