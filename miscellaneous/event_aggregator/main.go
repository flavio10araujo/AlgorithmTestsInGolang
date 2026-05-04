package main

import (
	"fmt"
	"sync"
	"time"
)

const windowMs = 5 * 60 * 1000

func main() {
	aggregator := NewEventAggregatorV2()
	now := time.Now().UnixMilli()
	users := []string{"userA", "userB", "userC"}
	workers := 12           // 12 goroutines
	eventsPerWorker := 1000 // 1000 events
	staleEvery := 10        // each 10 valid events, insert an old event

	expectedPerUser := make(map[string]int, len(users))
	staleAttempts := 0
	for workerID := 0; workerID < workers; workerID++ {
		for eventIndex := 0; eventIndex < eventsPerWorker; eventIndex++ {
			userID := users[(workerID+eventIndex)%len(users)]
			expectedPerUser[userID]++
			if eventIndex%staleEvery == 0 {
				staleAttempts++
			}
		}
	}

	var wg sync.WaitGroup
	start := make(chan struct{})

	for workerID := 0; workerID < workers; workerID++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			<-start

			for eventIndex := 0; eventIndex < eventsPerWorker; eventIndex++ {
				userID := users[(workerID+eventIndex)%len(users)]
				timestamp := now - int64(eventIndex%5)*60*1000
				aggregator.AddEvent(userID, now, timestamp)

				if eventIndex%staleEvery == 0 {
					aggregator.AddEvent(userID, now, now-10*60*1000)
				}
			}
		}(workerID)
	}

	close(start) // each goroutine will start only when close is called
	wg.Wait()

	fmt.Printf("workers: %d | events/worker: %d | valid events expected: %d | stale events ignored: %d\n",
		workers,
		eventsPerWorker,
		workers*eventsPerWorker,
		staleAttempts,
	)

	for _, userID := range users {
		actual := aggregator.GetEventCount(userID, now)
		expected := expectedPerUser[userID]
		fmt.Printf("%s -> expected: %d | actual: %d\n", userID, expected, actual)
	}
}
