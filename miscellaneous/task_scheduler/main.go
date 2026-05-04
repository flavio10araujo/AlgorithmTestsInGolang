package main

import (
	"fmt"
	"sync"
	"time"
)

type Scheduler struct {
	wg sync.WaitGroup
}

func (s *Scheduler) Schedule(task func(), delay time.Duration) {
	s.wg.Add(1)

	time.AfterFunc(delay, func() {
		defer s.wg.Done()
		task()
	})
}

func (s *Scheduler) ShutdownAndAwait(timeout time.Duration) {
	done := make(chan struct{})

	go func() {
		s.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		// all tasks finished
	case <-time.After(timeout):
		// timeout reached (não dá pra "matar" goroutines facilmente)
		fmt.Println("Timeout reached")
	}
}

func main() {
	scheduler := &Scheduler{}

	scheduler.Schedule(
		func() { fmt.Println("Task A") },
		2*time.Second,
	)

	scheduler.Schedule(
		func() { fmt.Println("Task B") },
		500*time.Millisecond,
	)

	scheduler.ShutdownAndAwait(3 * time.Second)
}
