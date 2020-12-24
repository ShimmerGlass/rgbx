package main

import (
	"github.com/lucasb-eyer/go-colorful"
	"github.com/shimmerglass/rgbx/rgbx"
	"github.com/urfave/cli/v2"
)

var SetK2000Cmd = &cli.Command{
	Name:  "k2000",
	Usage: "Vroom",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:  "row",
			Value: 0,
			Usage: "Row",
		},
	},
	Action: func(c *cli.Context) error {
		color, err := colorful.Hex(c.Args().First())
		if err != nil {
			return err
		}

		f := buildSetRequest(c)
		f.Effect = &rgbx.SetRequest_K2000{
			K2000: &rgbx.EffectK2000{
				Color: &rgbx.Color{
					R: int32(color.R * 0xFF),
					G: int32(color.G * 0xFF),
					B: int32(color.B * 0xFF),
				},
				Row: int32(c.Int("row")),
			},
		}
		return sendEffect(f)
	},
}
