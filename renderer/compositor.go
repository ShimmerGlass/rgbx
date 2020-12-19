package renderer

import (
	"sort"
	"sync"
	"time"

	"log"

	"github.com/shimmerglass/rgbx/rgbx"
)

type channel struct {
	priority int
	endAt    time.Time
	frame    *rgbx.Frame
}

type Compositor struct {
	lock sync.Mutex

	channels []*channel
	inner    Renderer
}

func NewCompositor(inner Renderer) *Compositor {
	return &Compositor{
		inner: inner,
	}
}

func (c *Compositor) Render(f *rgbx.Frame) error {
	c.lock.Lock()
	defer c.lock.Unlock()

	frameDuration := time.Duration(f.DurationMs) * time.Millisecond

	var frameEnd time.Time
	if f.DurationMs > 0 {
		frameEnd = time.Now().Add(frameDuration)
	}

	for _, channel := range c.channels {
		if channel.priority == int(f.Priority) {
			channel.frame = f
			channel.endAt = frameEnd
			goto Updated
		}
	}

	c.channels = append(c.channels, &channel{
		priority: int(f.Priority),
		endAt:    frameEnd,
		frame:    f,
	})

	sort.Slice(c.channels, func(i, j int) bool {
		return c.channels[i].priority > c.channels[j].priority
	})

Updated:
	if frameDuration > 0 {
		time.AfterFunc(frameDuration, func() {
			c.lock.Lock()
			defer c.lock.Unlock()

			err := c.update()
			if err != nil {
				log.Println(err)
			}
		})
	}

	return c.update()
}

func (c *Compositor) update() error {
	now := time.Now()
	filtered := c.channels[:0]
	for _, channel := range c.channels {
		if channel.endAt.IsZero() || !channel.endAt.Before(now) {
			filtered = append(filtered, channel)
		}
	}
	c.channels = filtered

	if len(c.channels) == 0 {
		return nil
	}

	return c.inner.Render(c.channels[0].frame)
}
