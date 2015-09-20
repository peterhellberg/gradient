package gradient

import (
	"image"
	"image/color"
)

// Vertical is a vertical gradient
type Vertical struct {
	width  int
	height int
	stops  []Stop
}

// NewVertical creates a new vertical gradient
func NewVertical(width, height int, stops []Stop) *Vertical {
	return &Vertical{
		width:  width,
		height: height,
		stops:  stops,
	}
}

// ColorModel used by the gradient
func (p *Vertical) ColorModel() color.Model {
	return color.NRGBAModel
}

// Bounds is the gradient bounds
func (p *Vertical) Bounds() image.Rectangle {
	return image.Rect(0, 0, p.width, p.height)
}

// At returns the color of the pixel at (x, y)
func (p *Vertical) At(x, y int) color.Color {
	if len(p.stops) == 0 {
		return color.Black
	}

	if y == 0 {
		return p.stops[0].Color
	}

	if y >= p.height {
		return p.stops[len(p.stops)+1].Color
	}

	if x >= p.width {
		return color.Transparent
	}

	return getColor(float64(y)/float64(p.height), p.stops)
}

// Horizontal is a horizontal gradient
type Horizontal struct {
	width  int
	height int
	stops  []Stop
}

// NewHorizontal creates a new horizontal gradient
func NewHorizontal(width, height int, stops []Stop) *Horizontal {
	return &Horizontal{
		width:  width,
		height: height,
		stops:  stops,
	}
}

// ColorModel used by the gradient
func (p *Horizontal) ColorModel() color.Model {
	return color.NRGBAModel
}

// Bounds is the gradient bounds
func (p *Horizontal) Bounds() image.Rectangle {
	return image.Rect(0, 0, p.width, p.height)
}

// At returns the color of the pixel at (x, y)
func (p *Horizontal) At(x, y int) color.Color {
	if x == 0 {
		return p.stops[0].Color
	}

	if x >= p.width {
		return p.stops[len(p.stops)+1].Color
	}

	if y >= p.height {
		return color.Transparent
	}

	return getColor(float64(x)/float64(p.width), p.stops)
}

// Stop contains a gradient position and color
type Stop struct {
	Position float64
	Color    color.Color
}

func getColor(pos float64, stops []Stop) color.Color {
	if pos <= 0.0 || len(stops) == 1 {
		return stops[0].Color
	}

	last := stops[len(stops)-1]

	if pos >= last.Position {
		return last.Color
	}

	for i, stop := range stops[1:] {
		if pos < stop.Position {
			pos = (pos - stops[i].Position) / (stop.Position - stops[i].Position)

			return colorLerp(stops[i].Color, stop.Color, pos)
		}
	}

	return last.Color
}

func colorLerp(c0, c1 color.Color, t float64) color.Color {
	r0, g0, b0, a0 := c0.RGBA()
	r1, g1, b1, a1 := c1.RGBA()

	return color.NRGBA{
		lerp(r0, r1, t),
		lerp(g0, g1, t),
		lerp(b0, b1, t),
		lerp(a0, a1, t),
	}
}

func lerp(a, b uint32, t float64) uint8 {
	return uint8(int32(float64(a)*(1.0-t)+float64(b)*t) >> 8)
}
