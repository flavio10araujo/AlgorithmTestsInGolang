package main

import (
	"fmt"
	"time"
)

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
