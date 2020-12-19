package renderer

import "github.com/shimmerglass/rgbx/rgbx"

type Renderer interface {
	Render(*rgbx.Frame) error
}

type Renderers struct {
	inner []Renderer
}

func NewRenderers(inner ...Renderer) *Renderers {
	return &Renderers{
		inner: inner,
	}
}

func (r *Renderers) Render(f *rgbx.Frame) error {
	for _, i := range r.inner {
		err := i.Render(f)
		if err != nil {
			return err
		}
	}

	return nil
}
