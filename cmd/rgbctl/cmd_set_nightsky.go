package main

import (
	"github.com/lucasb-eyer/go-colorful"
	"github.com/shimmerglass/rgbx/rgbx"
	"github.com/urfave/cli/v2"
)

var SetNightskyCmd = &cli.Command{
	Name:  "nightsky",
	Usage: "A sparkling nightsky",
	Flags: []cli.Flag{},
	Action: func(c *cli.Context) error {
		color1, err := colorful.Hex(c.Args().Get(0))
		if err != nil {
			return err
		}
		color2, err := colorful.Hex(c.Args().Get(1))
		if err != nil {
			return err
		}

		f := buildSetRequest(c)
		f.Effect = &rgbx.SetRequest_Nightsky{
			Nightsky: &rgbx.EffectNightSky{
				Color1: &rgbx.Color{
					R: int32(color1.R * 0xFF),
					G: int32(color1.G * 0xFF),
					B: int32(color1.B * 0xFF),
				},
				Color2: &rgbx.Color{
					R: int32(color2.R * 0xFF),
					G: int32(color2.G * 0xFF),
					B: int32(color2.B * 0xFF),
				},
			},
		}
		return sendEffect(f)
	},
}
