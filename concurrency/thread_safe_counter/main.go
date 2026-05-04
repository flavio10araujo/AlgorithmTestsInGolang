package main

import (
	"fmt"
	"sync"
)

const goroutines = 500
const increments = 100

type Counter interface {
	Inc()
	Get() int64
}

// A more idiomatic version using channels.
// In terms of performance, this is not the best solution, but it is a good example of how to use channels to achieve thread safety.
func channelCounter() int {
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

	return <-done
}

func main() {
	var counter Counter = NewAtomicCounter()

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

	fmt.Println(counter.Get())

	// Using channels:
	fmt.Println(channelCounter())
}
