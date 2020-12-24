package device

import "github.com/shimmerglass/rgbx/rgbx"

type Device interface {
	ID() string
	Type() rgbx.DeviceType
	Width() int
	Height() int
	Render([][]rgbx.Color) error
}
