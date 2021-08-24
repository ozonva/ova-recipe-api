package ticker

import "time"

type Ticker interface {
	Stop()
	Chanel() <-chan time.Time
}

func New(d time.Duration) Ticker {
	return &ticker{t: time.NewTicker(d)}
}

type ticker struct {
	t *time.Ticker
}

func (t *ticker) Stop() {
	t.t.Stop()
}

func (t *ticker) Chanel() <-chan time.Time {
	return t.t.C
}
