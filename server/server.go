package server

import (
	"context"
	"log"

	"github.com/shimmerglass/rgbx/renderer"
	"github.com/shimmerglass/rgbx/rgbx"
)

type Server struct {
	rgbx.UnimplementedRGBizerServer
	renderer renderer.Renderer
}

func New(renderer renderer.Renderer) *Server {
	return &Server{renderer: renderer}
}

func (s *Server) SetFrame(ctx context.Context, in *rgbx.Frame) (*rgbx.SetFrameResponse, error) {
	err := s.renderer.Render(in)
	if err != nil {
		log.Println(err)
	}

	return &rgbx.SetFrameResponse{
		Success: err == nil,
	}, nil
}
