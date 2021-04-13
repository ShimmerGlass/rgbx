package effect

// #cgo pkg-config: libpulse
// #include "./pulse.h"
import "C"

import (
	"math"
	"sync"
	"time"

	"github.com/lucasb-eyer/go-colorful"
	"github.com/shimmerglass/rgbx/rgbx"
)

var lock sync.Mutex
var levels [][]float64
var _ Effect = (*VUMeter)(nil)

func init() {
	go C.run_pulse()
}

type VUMeter struct {
	color1 colorful.Color
	color2 colorful.Color
}

func NewVUMeter(color1, color2 rgbx.Color) *VUMeter {
	return &VUMeter{
		color1: colorful.Color{R: float64(color1.R) / 255, G: float64(color1.G) / 255, B: float64(color1.B) / 255},
		color2: colorful.Color{R: float64(color2.R) / 255, G: float64(color2.G) / 255, B: float64(color2.B) / 255},
	}
}

func (e *VUMeter) Frequency() time.Duration {
	return 50 * time.Millisecond
}

func (e *VUMeter) Render(elapsed time.Duration, frame [][]rgbx.Color) {
	lock.Lock()
	defer lock.Unlock()

	avg := []float64{}
	for _, l := range levels {
		v := 0.0
		for _, ll := range l {
			v += ll
		}
		v /= float64(len(l))
		avg = append(avg, v)
	}

	val0, val1 := 0.0, 0.0
	switch len(levels) {
	case 0:
		return
	case 1:
		val0 = avg[0]
		val1 = avg[0]
	default:
		val0 = avg[0]
		val1 = avg[1]
	}

	val0 = math.Log10(val0*9 + 1)
	val1 = math.Log10(val1*9 + 1)

	for y := range frame {
		rowLen := len(frame[y])
		rowLenH := float64(rowLen) / 2
		th0 := int(math.Round(rowLenH * val0))
		th1 := int(math.Round(rowLenH * val1))
		for x := range frame[y] {
			var color colorful.Color
			switch {
			case x < th0:
				t := float64(x) / ((rowLenH - 1) / 2)
				if t > 1 {
					t = 1
				}
				color = e.color1.BlendHsv(e.color2, t)
			case x == th0:
				color = e.color2
			case rowLen-x < th1:
				t := float64(rowLen-1-x) / ((rowLenH - 1) / 2)
				if t > 1 {
					t = 1
				}
				color = e.color1.BlendHsv(e.color2, t)
			case rowLen-x == th1:
				color = e.color2
			default:
				frame[y][x] = rgbx.Color{}
			}

			r, g, b := color.RGB255()
			frame[y][x] = rgbx.Color{
				R: int32(r),
				G: int32(g),
				B: int32(b),
			}
		}
	}

	for i := range levels {
		m := len(levels[i]) / 4
		levels[i] = levels[i][m*3:]
	}
}

//export setLevels
func setLevels(channel C.int, val C.float) {
	lock.Lock()
	defer lock.Unlock()

	l := math.Abs(float64(val))
	c := int(channel)

	for len(levels)-1 < c {
		levels = append(levels, []float64{})
	}

	levels[c] = append(levels[c], l)
}
