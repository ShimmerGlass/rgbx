package main

import (
	"image/color"
	"log"

	"github.com/shimmerglass/rgbx/cmd/rgbctl/effect"
	"github.com/shimmerglass/rgbx/rgbx"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:1342", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := rgbx.NewRGBizerClient(conn)

	e := effect.NewNightSky(client,
		color.RGBA{R: 0xFF, G: 0x25, B: 0x5B},
		color.RGBA{R: 0x00, G: 0xE9, B: 0xFF},
	)
	e.Run()
}
