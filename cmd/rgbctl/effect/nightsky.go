package effect

import (
	"context"
	"image/color"
	"log"
	"math/rand"
	"time"

	"github.com/shimmerglass/rgbx/rgbx"
)

const (
	width  = 21
	height = 6
)

type NightSky struct {
	client rgbx.RGBizerClient
	state  []int
	color1 color.Color
	color2 color.Color
}

func NewNightSky(client rgbx.RGBizerClient, color1 color.Color, color2 color.Color) *NightSky {
	return &NightSky{
		client: client,
		color1: color1,
		color2: color2,
		state:  make([]int, height*width),
	}
}

func (e *NightSky) Run() error {
	for range time.Tick(50 * time.Millisecond) {
		for i := range e.state {
			if e.state[i] > 0 {
				e.state[i]--
			}
		}

		// if rand.Intn(5) == 0 {
		e.state[rand.Intn(len(e.state))] = 10
		// }

		err := e.render()
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}

func (e *NightSky) render() error {
	effect := &rgbx.EffectMatrix{
		Rows: make([]*rgbx.Row, height),
	}

	for i := 0; i < height; i++ {
		effect.Rows[i] = &rgbx.Row{}
		for _, s := range e.state[i*width : (i+1)*width] {
			c := grad(e.color1, e.color2, float64(s)/10)
			effect.Rows[i].Data = append(effect.Rows[i].Data, &rgbx.Color{
				R: int32(c.R),
				G: int32(c.G),
				B: int32(c.B),
			})
		}
	}

	_, err := e.client.SetFrame(context.Background(), &rgbx.Frame{
		Priority:   10,
		DurationMs: 100,
		DeviceType: rgbx.Frame_KEYBOARD,
		Effect: &rgbx.Frame_Matrix{
			Matrix: effect,
		},
	})
	return err
}

func grad(color1, color2 color.Color, f float64) color.RGBA {
	c1 := color.RGBAModel.Convert(color1).(color.RGBA)
	c2 := color.RGBAModel.Convert(color2).(color.RGBA)

	return color.RGBA{
		R: uint8(float64(c2.R)*f + float64(c1.R)*(1-f)),
		G: uint8(float64(c2.G)*f + float64(c1.G)*(1-f)),
		B: uint8(float64(c2.B)*f + float64(c1.B)*(1-f)),
	}
}
