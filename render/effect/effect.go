package effect

import (
	"fmt"
	"time"

	"github.com/shimmerglass/rgbx/rgbx"
)

type Effect interface {
	Frequency() time.Duration
	Render(elapsed time.Duration, frame [][]rgbx.Color)
}

func New(config interface{}) (Effect, error) {
	switch e := config.(type) {
	case *rgbx.SetRequest_Static:
		return NewStatic(*e.Static.Color), nil
	case *rgbx.SetRequest_Wave:
		return NewWave(), nil
	case *rgbx.SetRequest_Nightsky:
		return NewNightSky(*e.Nightsky.Color1, *e.Nightsky.Color2), nil
	case *rgbx.SetRequest_K2000:
		return NewK2000(*e.K2000.Color, int(e.K2000.Row)), nil
	case *rgbx.SetRequest_VUMeter:
		return NewVUMeter(*e.VUMeter.Color1, *e.VUMeter.Color2), nil
	case *rgbx.SetRequest_Progress:
		rows := []int{}
		for _, v := range e.Progress.Rows {
			rows = append(rows, int(v))
		}
		return NewProgress(*e.Progress.Color, e.Progress.Value, rows), nil
	}

	return nil, fmt.Errorf("effect not found")
}
