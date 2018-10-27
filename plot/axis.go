package plot

import (
	"fmt"
	"image/color"
)

const (
	xOrigin = 0.2 /* x axis begins at 0.8*Canvas.Width */
	xEnd    = 0.8 /* x axis ends at 0.2*Canvas.Width */
	yOrigin = 0.2 /* y axis begins at 0.2*Canvas.Height */
	yEnd    = 0.8 /* y axis end at 0.8*Canvas.Height */

)

// AddAxis adds an axis to a `Canvas' in
// either horizontal or vertical position.
func (c *Canvas) AddAxis(ticks []int, thick int, position string) error {
	xs := int(float64(c.Width) * xOrigin)  /* x start */
	xe := int(float64(c.Width) * xEnd)     /* x end */
	ys := int(float64(c.Height) * yOrigin) /* y start */
	ye := int(float64(c.Height) * yEnd)    /* y end */
	if position == "horizontal" {
		if err := c.Line(xs, ye, xe, ye, thick,
			&color.RGBA{0, 0, 0, 255}); err != nil {
			return fmt.Errorf("axis: %v", err)
		}
		return nil
	}
	if position == "vertical" {
		if err := c.Line(xs, ys, xs, ye, thick,
			&color.RGBA{0, 0, 0, 255}); err != nil {
			return fmt.Errorf("axis: %v", err)
		}
		return nil
	}
	return fmt.Errorf("axis: don't know position %s", position)
}
