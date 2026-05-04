package main

import (
	"fmt"
	"time"
)

func main() {
	aggregator := NewEventAggregator()
	now := time.Now().UnixMilli()

	aggregator.AddEvent("userA", now, now-4*60*1000)  // 4 minutes ago
	aggregator.AddEvent("userA", now, now-2*60*1000)  // 2 minutes ago
	aggregator.AddEvent("userB", now, now-1*60*1000)  // 1 minute ago
	aggregator.AddEvent("userA", now, now)            // now
	aggregator.AddEvent("userB", now, now-10*60*1000) // 10 minutes ago

	fmt.Println("userA event count:", aggregator.GetEventCount("userA", now)) // 3
	fmt.Println("userB event count:", aggregator.GetEventCount("userB", now)) // 1
}
