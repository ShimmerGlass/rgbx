package effect

import (
	"time"

	"github.com/shimmerglass/rgbx/rgbx"
)

var _ Effect = (*Static)(nil)

type Static struct {
	color rgbx.Color
}

func NewStatic(color rgbx.Color) *Static {
	return &Static{color: color}
}

func (e *Static) Frequency() time.Duration {
	return 30 * time.Second
}

func (e *Static) Render(elapsed time.Duration, frame [][]rgbx.Color) {
	for y := range frame {
		for x := range frame[y] {
			frame[y][x] = e.color
		}
	}
}
