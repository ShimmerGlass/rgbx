package main

import (
	"flag"
	"log"
	"net"

	"github.com/shimmerglass/rgbx/renderer"
	"github.com/shimmerglass/rgbx/renderer/razer"
	"github.com/shimmerglass/rgbx/rgbx"
	"github.com/shimmerglass/rgbx/server"
	"google.golang.org/grpc"
)

func main() {
	listen := flag.String("listen", "127.0.0.1:1342", "Listen addr")
	flag.Parse()

	renderer := renderer.NewCompositor(
		renderer.NewRenderers(
			razer.NewRazerKeyboard(),
		),
	)

	err := renderer.Render(&rgbx.Frame{
		Priority: 0,
		Effect: &rgbx.Frame_Static{
			Static: &rgbx.EffectStatic{
				Color: &rgbx.Color{R: 0xff, G: 0, B: 0},
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", *listen)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	rgbx.RegisterRGBizerServer(grpcServer, server.New(renderer))
	grpcServer.Serve(lis)
}
