package render

import "github.com/shimmerglass/rgbx/rgbx"

type Renderer interface {
	Set(*rgbx.SetRequest) error
	Remove(*rgbx.RemoveRequest) error
}
