package main

import "sync"

// MutexCounter is thread-safe because it uses a mutex to control access to the count variable
type MutexCounter struct {
	count int64
	mu    sync.Mutex
}

func NewMutexCounter() *MutexCounter {
	return &MutexCounter{
		count: 0,
	}
}

func (c *MutexCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *MutexCounter) Get() int64 {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}
