package main

import (
	"context"
	"log"
	"time"

	"github.com/shimmerglass/rgbx/rgbx"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

var Set = &cli.Command{
	Name:  "set",
	Usage: "Set an effect",
	Flags: []cli.Flag{
		&cli.IntFlag{
			Name:    "priority",
			Value:   0,
			Usage:   "Effect priority",
			Aliases: []string{"p"},
		},
		&cli.DurationFlag{
			Name:    "duration",
			Value:   0,
			Usage:   "Effect duration",
			Aliases: []string{"d"},
		},
	},
	Subcommands: []*cli.Command{
		SetStaticCmd,
		SetNightskyCmd,
		SetWaveCmd,
		SetK2000Cmd,
		SetProgressCmd,
		SetVUMeterCmd,
	},
}

func buildSetRequest(c *cli.Context) *rgbx.SetRequest {
	return &rgbx.SetRequest{
		Priority:   uint32(c.Int("priority")),
		DurationMs: int32(c.Duration("duration") / time.Millisecond),
		DeviceType: rgbx.DeviceType_KEYBOARD,
	}
}

func sendEffect(f *rgbx.SetRequest) error {
	conn, err := grpc.Dial("127.0.0.1:1342", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := rgbx.NewRGBizerClient(conn)

	_, err = client.Set(context.Background(), f)
	return err
}
