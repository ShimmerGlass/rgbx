package render

import (
	"log"
	"sort"
	"sync"
	"time"

	"github.com/shimmerglass/rgbx/device"
	"github.com/shimmerglass/rgbx/render/effect"
	"github.com/shimmerglass/rgbx/rgbx"
)

type channel struct {
	priority int
	endAt    time.Time
	duration time.Duration
	effect   effect.Effect
}

type Compositor struct {
	lock sync.Mutex

	channels []*channel
	timers   map[int]*time.Timer
	device   device.Device
	clock    *Clock
	stop     chan struct{}

	frame [][]rgbx.Color
}

func NewCompositor(device device.Device) *Compositor {
	frame := make([][]rgbx.Color, device.Height())
	for i := range frame {
		frame[i] = make([]rgbx.Color, device.Width())
	}

	clock := NewClock()
	go clock.Run()

	return &Compositor{
		device: device,
		timers: map[int]*time.Timer{},
		clock:  clock,
		frame:  frame,
		stop:   make(chan struct{}),
	}
}

func (c *Compositor) Run() {
	for {
		select {
		case elapsed := <-c.clock.C:
			err := c.render(elapsed)
			if err != nil {
				log.Println(err)
			}
		case <-c.stop:
			c.clock.Stop()
			return
		}
	}
}

func (c *Compositor) Stop() {
	close(c.stop)
}

func (c *Compositor) Set(f *rgbx.SetRequest) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	effect, err := effect.New(f.Effect)
	if err != nil {
		return err
	}
	duration := time.Duration(f.DurationMs) * time.Millisecond

	for _, channel := range c.channels {
		if channel.priority == int(f.Priority) {
			c.clock.Remove(channel.effect.Frequency(), false)
			channel.effect = effect
			channel.endAt = time.Now().Add(duration)
			channel.duration = duration
			goto Add
		}
	}

	c.channels = append(c.channels, &channel{
		priority: int(f.Priority),
		effect:   effect,
		endAt:    time.Now().Add(duration),
		duration: duration,
	})

Add:
	sort.Slice(c.channels, func(i, j int) bool {
		return c.channels[i].priority < c.channels[j].priority
	})

	c.clock.Add(effect.Frequency(), duration == 0)
	if duration > 0 {
		c.clock.Add(duration, true)
	}

	return nil
}

func (c *Compositor) Remove(f *rgbx.RemoveRequest) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	for i, channel := range c.channels {
		if channel.priority == int(f.Priority) {
			c.channels = append(c.channels[:i], c.channels[i+1:]...)
			c.clock.Remove(channel.duration, true)
		}
	}

	return nil
}

func (c *Compositor) render(elapsed time.Duration) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	now := time.Now()
	c.zeroFrame()

	cleaned := c.channels[:0]
	for _, channel := range c.channels {
		if channel.endAt.Before(now) && channel.duration > 0 {
			c.clock.Remove(channel.duration, false)
			continue
		}
		cleaned = append(cleaned, channel)
		continue
	}
	c.channels = cleaned

	for _, channel := range c.channels {
		channel.effect.Render(elapsed, c.frame)
	}

	return c.device.Render(c.frame)
}

func (c *Compositor) zeroFrame() {
	for y := range c.frame {
		for x := range c.frame[y] {
			c.frame[y][x] = rgbx.Color{}
		}
	}
}
