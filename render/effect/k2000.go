package effect

import (
	"time"

	"github.com/shimmerglass/rgbx/rgbx"
)

var _ Effect = (*K2000)(nil)

type K2000 struct {
	row   int
	color rgbx.Color
	i     float64
	rtl   bool
}

func NewK2000(color rgbx.Color, row int) *K2000 {
	return &K2000{color: color}
}

func (e *K2000) Frequency() time.Duration {
	return 250 * time.Millisecond
}

func (e *K2000) Render(elapsed time.Duration, frame [][]rgbx.Color) {
	for x := range frame[e.row] {
		c := rgbx.Color{}
		if int(e.i) == x {
			c = e.color
		}
		frame[e.row][x] = c
	}

	inc := float64(elapsed) / float64((250 * time.Millisecond))

	if e.rtl {
		if e.i <= 0 {
			e.i = 1
			e.rtl = false
		} else {
			e.i -= inc
		}
	} else {
		if e.i >= float64(len(frame[e.row]))-1 {
			e.i = float64(len(frame[e.row])) - 2
			e.rtl = true
		} else {
			e.i += inc
		}
	}
}
