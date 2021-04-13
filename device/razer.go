package device

import (
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/shimmerglass/rgbx/rgbx"
	log "github.com/sirupsen/logrus"
)

type RazerKeyboardDriver struct {
	added   chan Device
	removed chan string
	devices map[string]struct{}
}

func NewRazerKeyboardDriver() *RazerKeyboardDriver {
	r := &RazerKeyboardDriver{
		devices: map[string]struct{}{},
		added:   make(chan Device),
		removed: make(chan string),
	}
	go func() {
		r.updateDevices()
		for range time.Tick(10 * time.Second) {
			r.updateDevices()
		}
	}()

	return r
}

func (r *RazerKeyboardDriver) DeviceAdded() chan Device {
	return r.added
}

func (r *RazerKeyboardDriver) DeviceRemoved() chan string {
	return r.removed
}

func (r *RazerKeyboardDriver) updateDevices() error {
	matches, err := filepath.Glob("/sys/bus/hid/drivers/razerkbd/*/device_type")
	if err != nil {
		return err
	}

	found := map[string]bool{}
	for _, devicePath := range matches {
		devicePath = path.Dir(devicePath)
		found[devicePath] = true

		if _, ok := r.devices[devicePath]; ok {
			continue
		}

		d, err := NewRazerKeyboard(devicePath)
		if err != nil {
			return err
		}
		r.devices[devicePath] = struct{}{}
		r.added <- d
	}

	for id := range r.devices {
		if !found[id] {
			delete(r.devices, id)
			r.removed <- id
		}
	}

	return nil
}

type RazerKeyboard struct {
	path    string
	f       *os.File
	limiter *time.Ticker
}

func NewRazerKeyboard(path string) (*RazerKeyboard, error) {
	err := ioutil.WriteFile(filepath.Join(path, "matrix_effect_custom"), []byte{1}, 0)
	if err != nil {
		return nil, err
	}

	f, err := os.OpenFile(filepath.Join(path, "matrix_custom_frame"), os.O_WRONLY, 0)
	if err != nil {
		return nil, err
	}

	return &RazerKeyboard{
		path:    path,
		limiter: time.NewTicker(50 * time.Millisecond),
		f:       f,
	}, nil
}

func (k *RazerKeyboard) ID() string {
	return k.path
}

func (k *RazerKeyboard) Type() rgbx.DeviceType {
	return rgbx.DeviceType_KEYBOARD
}

func (k *RazerKeyboard) Width() int {
	return 21
}

func (k *RazerKeyboard) Height() int {
	return 6
}

func (r *RazerKeyboard) Render(frame [][]rgbx.Color) error {
	log.Debugf("razer: %s: render", r.path)

	// <-r.limiter.C

	buf := &bytes.Buffer{}
	for i, row := range frame {
		buf.Write([]byte{byte(i), 0x0, byte(len(row) - 1)})

		for _, c := range row {
			buf.Write([]byte{byte(c.R), byte(c.G), byte(c.B)})
		}
	}
	_, err := r.f.Write(buf.Bytes())
	if err != nil {
		return err
	}

	return nil
}
