package main

import (
	"github.com/lucasb-eyer/go-colorful"
	"github.com/shimmerglass/rgbx/rgbx"
	"github.com/urfave/cli/v2"
)

var SetStaticCmd = &cli.Command{
	Name:  "static",
	Usage: "Set a static color",
	Flags: []cli.Flag{},
	Action: func(c *cli.Context) error {
		color, err := colorful.Hex(c.Args().First())
		if err != nil {
			return err
		}

		f := buildSetRequest(c)
		f.Effect = &rgbx.SetRequest_Static{
			Static: &rgbx.EffectStatic{
				Color: &rgbx.Color{
					R: int32(color.R * 0xFF),
					G: int32(color.G * 0xFF),
					B: int32(color.B * 0xFF),
				},
			},
		}
		return sendEffect(f)
	},
}
