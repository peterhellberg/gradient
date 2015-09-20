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

	for i, tt := range []struct {
		x int
		y int
		c color.Color
	}{
		{100, 100, color.NRGBA{106, 0, 149, 255}},
	} {
		c := v.At(tt.x, tt.y)

		r0, g0, b0, a0 := c.RGBA()
		r1, g1, b1, a1 := tt.c.RGBA()

		if r0 != r1 || g0 != g1 || b0 != b1 || a0 != a1 {
			t.Fatalf("[%d] v.At(%d, %d) = %+v, want %+v", tt.x, tt.y, i, c, tt.c)
		}
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

	for i, tt := range []struct {
		x int
		y int
		c color.Color
	}{
		{0, 0, color.NRGBA{0, 255, 255, 0}},
		{100, 100, color.NRGBA{0, 177, 0, 177}},
		{239, 319, color.NRGBA{0, 255, 0, 255}},
	} {
		c := h.At(tt.x, tt.y)

		r0, g0, b0, a0 := c.RGBA()
		r1, g1, b1, a1 := tt.c.RGBA()

		if r0 != r1 || g0 != g1 || b0 != b1 || a0 != a1 {
			t.Fatalf("[%d] h.At(%d, %d) = %+v, want %+v", i, tt.x, tt.y, c, tt.c)
		}
	}
}