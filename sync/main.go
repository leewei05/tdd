package main

import "sync"

type Counter struct {
	mu    sync.Mutex
	count int
}

func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func (c *Counter) Value() int {
	return c.count
}
