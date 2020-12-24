package render

import (
	"sync"

	"github.com/shimmerglass/rgbx/device"
	"github.com/shimmerglass/rgbx/rgbx"
	log "github.com/sirupsen/logrus"
)

type deviceMatcher interface {
	GetDeviceType() rgbx.DeviceType
}

type deviceCompositor struct {
	device     device.Device
	compositor *Compositor
}

type Compositors struct {
	lock        sync.Mutex
	compositors map[string]*deviceCompositor
}

func NewCompositors(driver device.Driver) *Compositors {
	c := &Compositors{compositors: map[string]*deviceCompositor{}}

	go func() {
		for device := range driver.DeviceAdded() {
			if _, ok := c.compositors[device.ID()]; ok {
				continue
			}
			log.Infof("device %s added", device.ID())

			compositor := NewCompositor(device)
			go compositor.Run()
			c.lock.Lock()
			c.compositors[device.ID()] = &deviceCompositor{
				compositor: compositor,
				device:     device,
			}
			c.lock.Unlock()
		}
	}()

	go func() {
		for id := range driver.DeviceRemoved() {
			devcomp, ok := c.compositors[id]
			if !ok {
				continue
			}
			log.Infof("device %s removed", id)

			devcomp.compositor.Stop()
			c.lock.Lock()
			delete(c.compositors, id)
			c.lock.Unlock()
		}
	}()

	return c
}

func (c *Compositors) Set(f *rgbx.SetRequest) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	for _, devcomp := range c.compositors {
		if !c.match(f, devcomp.device) {
			continue
		}

		err := devcomp.compositor.Set(f)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Compositors) Remove(f *rgbx.RemoveRequest) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	for _, devcomp := range c.compositors {
		if !c.match(f, devcomp.device) {
			continue
		}

		err := devcomp.compositor.Remove(f)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *Compositors) match(f deviceMatcher, d device.Device) bool {
	return f.GetDeviceType() == d.Type()
}
