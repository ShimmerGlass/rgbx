package device

type Driver interface {
	DeviceAdded() chan Device
	DeviceRemoved() chan string
}

type DriverPool struct {
	added   chan Device
	removed chan string
}

func NewDriverPool(inner ...Driver) *DriverPool {
	p := &DriverPool{
		added:   make(chan Device),
		removed: make(chan string),
	}

	for _, d := range inner {
		go func(d Driver) {
			for device := range d.DeviceAdded() {
				p.added <- device
			}
		}(d)
		go func(d Driver) {
			for id := range d.DeviceRemoved() {
				p.removed <- id
			}
		}(d)
	}

	return p
}

func (p *DriverPool) DeviceAdded() chan Device {
	return p.added
}

func (p *DriverPool) DeviceRemoved() chan string {
	return p.removed
}
