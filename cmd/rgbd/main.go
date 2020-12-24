package main

import (
	"flag"
	"net"

	log "github.com/sirupsen/logrus"

	"github.com/shimmerglass/rgbx/device"
	"github.com/shimmerglass/rgbx/render"
	"github.com/shimmerglass/rgbx/rgbx"
	"github.com/shimmerglass/rgbx/server"
	"google.golang.org/grpc"
)

func main() {
	listen := flag.String("listen", "127.0.0.1:1342", "Listen addr")
	flag.Parse()

	// log.SetLevel(log.DebugLevel)

	compositors := render.NewCompositors(device.NewDriverPool(
		device.NewRazerKeyboardDriver(),
	))

	lis, err := net.Listen("tcp", *listen)
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()
	rgbx.RegisterRGBizerServer(grpcServer, server.New(compositors))
	grpcServer.Serve(lis)
}
