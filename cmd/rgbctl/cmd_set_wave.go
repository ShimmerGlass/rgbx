package main

import (
	"github.com/shimmerglass/rgbx/rgbx"
	"github.com/urfave/cli/v2"
)

var SetWaveCmd = &cli.Command{
	Name:  "wave",
	Usage: "A rainbow wave",
	Flags: []cli.Flag{},
	Action: func(c *cli.Context) error {
		f := buildSetRequest(c)
		f.Effect = &rgbx.SetRequest_Wave{
			Wave: &rgbx.EffectWave{},
		}
		return sendEffect(f)
	},
}
