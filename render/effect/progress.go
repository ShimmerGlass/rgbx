package effect

import (
	"math"
	"time"

	"github.com/shimmerglass/rgbx/rgbx"
)

var _ Effect = (*Progress)(nil)

type Progress struct {
	color rgbx.Color
	rows  []int
	value float64
}

func NewProgress(color rgbx.Color, value float64, rows []int) *Progress {
	if value > 1 {
		value = 1
	}
	if value < 0 {
		value = 0
	}
	return &Progress{
		color: color,
		value: value,
		rows:  rows,
	}
}

func (e *Progress) Frequency() time.Duration {
	return 30 * time.Second
}

func (e *Progress) Render(elapsed time.Duration, frame [][]rgbx.Color) {
	v := int(math.Round(float64(len(frame[0])) * e.value))
	for _, y := range e.rows {
		for x := range frame[y] {
			c := rgbx.Color{}
			if x <= v {
				c = e.color
			}

			frame[y][x] = c
		}
	}
}
