package main

import (
	"sync"
	"sync/atomic"
)

const goroutines = 500
const increments = 100

type Counter interface {
	Inc()
	Get() int64
}

// CounterNotThreadSafe is not thread-safe
type CounterNotThreadSafe struct {
	count int64
}

func (c *CounterNotThreadSafe) Inc() {
	c.count++
}

func (c *CounterNotThreadSafe) Get() int64 {
	return c.count
}

// MutexCounter is thread-safe because it uses a mutex to control access to the count variable
type MutexCounter struct {
	count int64
	mu    sync.Mutex
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

// AtomicCounter is thread-safe and uses the sync/atomic package.
// Better solution than using mutex but still not optimal under extreme contention.
type AtomicCounter struct {
	count int64
}

func (c *AtomicCounter) Inc() {
	atomic.AddInt64(&c.count, 1)
}

func (c *AtomicCounter) Get() int64 {
	return atomic.LoadInt64(&c.count)
}

// A more idiomatic version using channels.
// In terms of performance, this is not the best solution but it is a good example of how to use channels to achieve thread safety.
func channelCounter() {
	increment := make(chan int)
	done := make(chan int)

	go func() {
		var counter int
		for v := range increment {
			counter += v
		}
		done <- counter
	}()

	var wg sync.WaitGroup

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < increments; j++ {
				increment <- 1
			}
		}()
	}

	wg.Wait()
	close(increment)

	//fmt.Println(<-done)
}

func main() {
	counter := &AtomicCounter{}

	var wg sync.WaitGroup

	worker := func() {
		defer wg.Done()
		for i := 0; i < increments; i++ {
			counter.Inc()
		}
	}

	for i := 0; i < goroutines; i++ {
		wg.Add(1)
		go worker()
	}

	wg.Wait()

	//println(counter.Get())

	// Using channels:
	channelCounter()
}
