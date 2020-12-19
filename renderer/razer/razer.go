package razer

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/shimmerglass/rgbx/rgbx"
)

type RazerKeyboard struct {
	device string
}

func NewRazerKeyboard() *RazerKeyboard {
	r := &RazerKeyboard{}
	err := r.updateDevice()
	if err != nil {
		log.Fatal(err)
	}
	return r
}

func (r *RazerKeyboard) Render(f *rgbx.Frame) error {
	switch e := f.Effect.(type) {
	case *rgbx.Frame_Static:
		return r.effectStatic(e.Static)
	case *rgbx.Frame_Matrix:
		return r.effectMatrix(e.Matrix)
	}

	return nil
}

func (r *RazerKeyboard) effectStatic(e *rgbx.EffectStatic) error {
	return ioutil.WriteFile(
		filepath.Join(r.device, "matrix_effect_static"),
		[]byte{byte(e.Color.R), byte(e.Color.G), byte(e.Color.B)},
		0,
	)
}

func (r *RazerKeyboard) effectMatrix(e *rgbx.EffectMatrix) error {
	buf := &bytes.Buffer{}
	for i, row := range e.Rows {
		buf.Write([]byte{byte(i), 0x0, byte(len(row.Data) - 1)})

		for _, c := range row.Data {
			buf.Write([]byte{byte(c.R), byte(c.G), byte(c.B)})
		}
	}

	err := ioutil.WriteFile(
		filepath.Join(r.device, "matrix_effect_custom"),
		[]byte{1},
		0,
	)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(
		filepath.Join(r.device, "matrix_custom_frame"),
		buf.Bytes(),
		0,
	)
}

func (r *RazerKeyboard) updateDevice() error {
	matches, err := filepath.Glob("/sys/bus/hid/drivers/razerkbd/*/device_type")
	if err != nil {
		return err
	}

	if len(matches) == 0 {
		return fmt.Errorf("no devices found")
	}

	r.device = filepath.Dir(matches[0])
	return nil
}
