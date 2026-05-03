package main

import "sync/atomic"

// AtomicCounter is thread-safe and uses the sync/atomic package.
// Better solution than using mutex but still not optimal under extreme contention.
type AtomicCounter struct {
	count int64
}

func NewAtomicCounter() *AtomicCounter {
	return &AtomicCounter{
		count: 0,
	}
}

func (c *AtomicCounter) Inc() {
	atomic.AddInt64(&c.count, 1)
}

func (c *AtomicCounter) Get() int64 {
	return atomic.LoadInt64(&c.count)
}
