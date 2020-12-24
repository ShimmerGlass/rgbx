package main

import (
	"strconv"

	"github.com/lucasb-eyer/go-colorful"
	"github.com/shimmerglass/rgbx/rgbx"
	"github.com/urfave/cli/v2"
)

var SetProgressCmd = &cli.Command{
	Name:  "progress",
	Usage: "Display a progress bar",
	Flags: []cli.Flag{
		&cli.IntSliceFlag{
			Name:  "row",
			Value: cli.NewIntSlice(0),
			Usage: "Rows",
		},
	},
	Action: func(c *cli.Context) error {
		color, err := colorful.Hex(c.Args().First())
		if err != nil {
			return err
		}

		valueStr := c.Args().Get(1)
		div := 1.0
		if valueStr[len(valueStr)-1] == '%' {
			div = 100
			valueStr = valueStr[:len(valueStr)-1]
		}
		value, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			return err
		}
		value /= div

		rows := []int32{}
		for _, v := range c.IntSlice("row") {
			rows = append(rows, int32(v))
		}

		f := buildSetRequest(c)
		f.Effect = &rgbx.SetRequest_Progress{
			Progress: &rgbx.EffectProgress{
				Color: &rgbx.Color{
					R: int32(color.R * 0xFF),
					G: int32(color.G * 0xFF),
					B: int32(color.B * 0xFF),
				},
				Value: value,
				Rows:  rows,
			},
		}
		return sendEffect(f)
	},
}
