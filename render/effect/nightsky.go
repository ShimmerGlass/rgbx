package effect

import (
	"image/color"
	"math/rand"
	"time"

	"github.com/shimmerglass/rgbx/rgbx"
)

var _ Effect = (*NightSky)(nil)

type NightSky struct {
	state  []int
	color1 rgbx.Color
	color2 rgbx.Color
}

func NewNightSky(color1, color2 rgbx.Color) *NightSky {
	return &NightSky{
		color1: color1,
		color2: color2,
	}
}

func (e *NightSky) Frequency() time.Duration {
	return 50 * time.Millisecond
}

func (e *NightSky) Render(elapsed time.Duration, f [][]rgbx.Color) {
	if e.state == nil {
		e.state = make([]int, len(f)*len(f[0]))
	}

	for i, v := range e.state {
		if v > 0 {
			e.state[i]--
		}
	}

	a := int(elapsed / (50 * time.Millisecond))
	for i := 0; i < a; i++ {
		e.state[rand.Intn(len(e.state))] = 10
	}

	for y := range f {
		for x := range f[y] {
			s := e.state[y*len(f[0])+x]
			c := grad(e.color1, e.color2, float64(s)/10)
			f[y][x] = rgbx.Color{
				R: int32(c.R),
				G: int32(c.G),
				B: int32(c.B),
			}
		}
	}
}

func (e *NightSky) init(w, h int) {
	e.state = make([]int, w*h)
}

func grad(c1, c2 rgbx.Color, f float64) color.RGBA {
	return color.RGBA{
		R: uint8(float64(c2.R)*f + float64(c1.R)*(1-f)),
		G: uint8(float64(c2.G)*f + float64(c1.G)*(1-f)),
		B: uint8(float64(c2.B)*f + float64(c1.B)*(1-f)),
	}
}
