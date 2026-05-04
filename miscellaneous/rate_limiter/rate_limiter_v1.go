package main

import "sync"

type RateLimiterV1 struct {
	requests map[string][]int64
	mu       sync.Mutex
}

func NewRateLimiterV1() *RateLimiterV1 {
	return &RateLimiterV1{
		requests: make(map[string][]int64),
	}
}

func (r *RateLimiterV1) AllowRequest(userID string, now int64) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	queue := r.requests[userID]

	for len(queue) > 0 && queue[0] <= now-windowMs {
		queue = queue[1:]
	}

	r.requests[userID] = queue

	if len(queue) >= limit {
		return false
	}

	r.requests[userID] = append(queue, now)
	return true
}
