package pticker

import "time"

func NewTicker(d time.Duration) *Ticker {
	return &Ticker{
		Duration: d,
		C:        make(chan time.Time),
		done:     make(chan struct{}),
	}
}

func (t *Ticker) Start() {
	go func() {
		for {
			x := time.Now()
			y := x.Round(t.Duration)
			if x.After(y) {
				y = y.Add(t.Duration)
			}
			select {
			case m := <-time.After(y.Sub(time.Now())):
				t.C <- m
			case <-t.done:
				close(t.C)
				return
			}
		}
	}()
}

func (t *Ticker) Stop() {
	t.done <- struct{}{}
}

type Ticker struct {
	Duration time.Duration
	C        chan time.Time

	done chan struct{}
}
