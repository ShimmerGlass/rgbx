package effect

import (
	"math"
	"time"

	"github.com/lucasb-eyer/go-colorful"
	"github.com/shimmerglass/rgbx/rgbx"
)

var _ Effect = (*Static)(nil)

type Wave struct {
	offset float64
}

func NewWave() *Wave {
	return &Wave{}
}

func (e *Wave) Frequency() time.Duration {
	return 50 * time.Millisecond
}

func (e *Wave) Render(elapsed time.Duration, frame [][]rgbx.Color) {
	e.offset += elapsed.Seconds() * 100
	ks := 360 / float64(len(frame[0]))

	for y := range frame {
		for x := range frame[y] {
			c := colorful.Hsv(math.Mod(e.offset+float64(x)*ks, 360), 1, 1)
			frame[y][x] = rgbx.Color{
				R: int32(c.R * 255),
				G: int32(c.G * 255),
				B: int32(c.B * 255),
			}
		}
	}
}
