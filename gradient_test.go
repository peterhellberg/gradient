package gradient

import (
	"image/color"
	"testing"
)

func TestNewVertical(t *testing.T) {
	x := 320
	y := 240
	v := NewVertical(x, y, []Stop{
		{0.0, color.NRGBA{0, 0, 255, 255}},
		{1.0, color.NRGBA{255, 0, 0, 255}},
	})

	if got, want := v.Bounds().Dx(), x; got != want {
		t.Fatalf(`v.Bounds().Dx() = %d, want %d`, got, want)
	}

	if got, want := v.Bounds().Dy(), y; got != want {
		t.Fatalf(`v.Bounds().Dy() = %d, want %d`, got, want)
	}

	if got, want := v.ColorModel(), color.NRGBAModel; got != want {
		t.Fatalf(`v.ColorModel() = %v, want %v`, got, want)
	}

	for i, tt := range []struct {
		x int
		y int
		c color.Color
	}{
		{0, 0, color.NRGBA{0, 0, 255, 255}},
		{100, 100, color.NRGBA{106, 0, 149, 255}},
		{200, 200, color.NRGBA{213, 0, 42, 255}},
		{319, 239, color.NRGBA{254, 0, 1, 255}},
		{900, 900, color.Transparent},
	} {
		c := v.At(tt.x, tt.y)

		r0, g0, b0, a0 := c.RGBA()
		r1, g1, b1, a1 := tt.c.RGBA()

		if r0 != r1 || g0 != g1 || b0 != b1 || a0 != a1 {
			t.Fatalf("[%d] v.At(%d, %d) = %+v, want %+v", i, tt.x, tt.y, c, tt.c)
		}
	}
}

func TestNewVertical_noStops(t *testing.T) {
	v := NewVertical(200, 400, []Stop{}).At(100, 150)
	c := color.NRGBA{255, 0, 255, 255}

	r0, g0, b0, a0 := v.RGBA()
	r1, g1, b1, a1 := c.RGBA()

	if r0 != r1 || g0 != g1 || b0 != b1 || a0 != a1 {
		t.Fatalf("v.At(%d, %d) = %+v, want %+v", 200, 400, v, c)
	}
}

func TestNewVertical_singleStop(t *testing.T) {
	c := color.NRGBA{0, 255, 255, 0}
	v := NewVertical(200, 400, []Stop{{0.0, c}}).At(100, 150)

	r0, g0, b0, a0 := v.RGBA()
	r1, g1, b1, a1 := c.RGBA()

	if r0 != r1 || g0 != g1 || b0 != b1 || a0 != a1 {
		t.Fatalf("v.At(%d, %d) = %+v, want %+v", 200, 400, v, c)
	}
}

func TestNewHorizontal(t *testing.T) {
	x := 240
	y := 320
	h := NewHorizontal(x, y, []Stop{
		{0.0, color.NRGBA{0, 255, 255, 0}},
		{0.6, color.NRGBA{0, 255, 0, 255}},
		{1.0, color.NRGBA{0, 255, 0, 255}},
	})

	if got, want := h.Bounds().Dx(), x; got != want {
		t.Fatalf(`h.Bounds().Dx() = %d, want %d`, got, want)
	}

	if got, want := h.Bounds().Dy(), y; got != want {
		t.Fatalf(`h.Bounds().Dy() = %d, want %d`, got, want)
	}

	if got, want := h.ColorModel(), color.NRGBAModel; got != want {
		t.Fatalf(`h.ColorModel() = %v, want %v`, got, want)
	}

	for i, tt := range []struct {
		x int
		y int
		c color.Color
	}{
		{0, 0, color.NRGBA{0, 255, 255, 0}},
		{100, 100, color.NRGBA{0, 177, 0, 177}},
		{239, 319, color.NRGBA{0, 255, 0, 255}},
		{900, 900, color.Transparent},
	} {
		c := h.At(tt.x, tt.y)

		r0, g0, b0, a0 := c.RGBA()
		r1, g1, b1, a1 := tt.c.RGBA()

		if r0 != r1 || g0 != g1 || b0 != b1 || a0 != a1 {
			t.Fatalf("[%d] h.At(%d, %d) = %+v, want %+v", i, tt.x, tt.y, c, tt.c)
		}
	}
}

func TestNewHorizontal_noStops(t *testing.T) {
	h := NewHorizontal(200, 400, []Stop{}).At(100, 150)
	c := color.NRGBA{255, 0, 255, 255}

	r0, g0, b0, a0 := h.RGBA()
	r1, g1, b1, a1 := c.RGBA()

	if r0 != r1 || g0 != g1 || b0 != b1 || a0 != a1 {
		t.Fatalf("h.At(%d, %d) = %+v, want %+v", 200, 400, h, c)
	}
}

func TestNewHorizontal_singleStop(t *testing.T) {
	c := color.NRGBA{0, 255, 255, 0}
	h := NewHorizontal(200, 400, []Stop{{0.0, c}}).At(100, 150)

	r0, g0, b0, a0 := h.RGBA()
	r1, g1, b1, a1 := c.RGBA()

	if r0 != r1 || g0 != g1 || b0 != b1 || a0 != a1 {
		t.Fatalf("h.At(%d, %d) = %+v, want %+v", 200, 400, h, c)
	}
}
