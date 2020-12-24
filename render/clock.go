package render

import (
	"math/big"
	"time"

	"github.com/sirupsen/logrus"
)

type Clock struct {
	freqs  map[time.Duration]int
	freq   time.Duration
	C      chan time.Duration
	ticker *time.Ticker
	stop   chan struct{}
}

func NewClock() *Clock {
	return &Clock{
		freqs:  map[time.Duration]int{},
		C:      make(chan time.Duration),
		ticker: time.NewTicker(10 * time.Minute),
		stop:   make(chan struct{}),
	}
}

func (c *Clock) Run() {
	for {
		select {
		case <-c.ticker.C:
			c.notify()
		case <-c.stop:
			return
		}
	}
}

func (c *Clock) Stop() {
	close(c.stop)
}

func (c *Clock) Add(freq time.Duration, notify bool) {
	c.freqs[freq]++
	c.update(notify)
}

func (c *Clock) Remove(freq time.Duration, notify bool) {
	c.freqs[freq]--
	if c.freqs[freq] == 0 {
		delete(c.freqs, freq)
	}
	c.update(notify)
}

func (c *Clock) update(notify bool) {
	last := big.NewInt(int64(10 * time.Minute))
	for f := range c.freqs {
		last = new(big.Int).GCD(nil, nil, last, big.NewInt(int64(f)))
	}

	c.freq = time.Duration(last.Int64())
	logrus.Debugf("new clock frequency: %s", c.freq)

	c.ticker.Reset(c.freq)
	if notify {
		go c.notify()
	}
}

func (c *Clock) notify() {
	c.C <- c.freq
}
