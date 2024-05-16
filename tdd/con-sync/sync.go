package consync

import "sync"

// It is important to note that the Counter type has a Mutex field, as
// opposed to embedding the Mutex type. While it might look nicer to
// invoke 'counter.Lock()', you're making the Mutex lock part of the
// public Counter API, which is a big NO NO.
type Counter struct {
	mu    sync.Mutex
	value int
}

func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

func NewCounter() *Counter {
	return &Counter{}
}
