package tyschedule

import (
	"errors"
	"time"
)

type Func func(currentRetryCount int, maxRetryCount int) (ok bool)

type Monitor struct {
	done          chan struct{}
	maxRetryCount int
	interval      int
	fn            Func
	running       bool
}

func New() *Monitor {
	return &Monitor{
		done: make(chan struct{}, 1),
	}
}

func (mon *Monitor) HandleFunc(fn Func) {
	if fn == nil {
		fn = func(currentRetryCount int, maxRetryCount int) (ok bool) {
			return true
		}
	}
	mon.fn = fn
}

func (mon *Monitor) Start(interval int, maxRetryCount int) error {
	if interval <= 0 {
		return errors.New("interval error, interval must more 0")
	}

	if maxRetryCount <= 0 {
		return errors.New("maxRetryCount error, maxRetryCount must more 0")
	}

	mon.interval = interval
	mon.maxRetryCount = maxRetryCount

	if !mon.running {
		mon.running = true
		go mon.run()

	}

	return nil
}

func (mon *Monitor) run() {
	ticker := time.NewTicker(time.Second * time.Duration(mon.interval))
	defer ticker.Stop()

	var currentRetryCount int
	currentRetryCount = 1

loop:
	for {
		select {
		case <-ticker.C:
			if currentRetryCount >= mon.maxRetryCount {
				currentRetryCount = 1
			}
			ok := mon.fn(currentRetryCount, mon.maxRetryCount)
			if ok {
				currentRetryCount = 1
			} else {
				currentRetryCount++
			}
		case <-mon.done:
			close(mon.done)
			break loop
		}
	}
}

func (mon *Monitor) Stop() {
	if mon == nil {
		return
	}
	mon.done <- struct{}{}
}
