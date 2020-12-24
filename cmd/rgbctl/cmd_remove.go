package main

import (
	"context"
	"log"
	"strconv"

	"github.com/shimmerglass/rgbx/rgbx"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

var RemoveCmd = &cli.Command{
	Name:  "remove",
	Usage: "Remove an effect",
	Action: func(c *cli.Context) error {
		pri, err := strconv.Atoi(c.Args().First())
		if err != nil {
			return err
		}

		return sendRemove(&rgbx.RemoveRequest{Priority: uint32(pri)})
	},
}

func sendRemove(f *rgbx.RemoveRequest) error {
	conn, err := grpc.Dial("127.0.0.1:1342", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := rgbx.NewRGBizerClient(conn)

	_, err = client.Remove(context.Background(), f)
	return err
}
