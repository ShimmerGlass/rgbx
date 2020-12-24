package server

import (
	"context"

	"github.com/shimmerglass/rgbx/render"
	"github.com/shimmerglass/rgbx/rgbx"
)

type Server struct {
	rgbx.UnimplementedRGBizerServer
	renderer render.Renderer
}

func New(renderer render.Renderer) *Server {
	return &Server{renderer: renderer}
}

func (s *Server) Set(ctx context.Context, in *rgbx.SetRequest) (*rgbx.SuccessResponse, error) {
	err := s.renderer.Set(in)
	if err != nil {
		return nil, err
	}

	return &rgbx.SuccessResponse{
		Success: true,
	}, nil
}

func (s *Server) Remove(ctx context.Context, in *rgbx.RemoveRequest) (*rgbx.SuccessResponse, error) {
	err := s.renderer.Remove(in)
	if err != nil {
		return nil, err
	}

	return &rgbx.SuccessResponse{
		Success: true,
	}, nil
}
