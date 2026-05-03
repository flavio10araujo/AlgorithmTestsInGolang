package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	limit    = 100
	windowMs = 60_000
)

type RateLimiter struct {
	requests map[string][]int64
	mu       sync.Mutex
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{
		requests: make(map[string][]int64),
	}
}

func (r *RateLimiter) AllowRequest(userID string, now int64) bool {
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

func main() {
	rateLimiter := NewRateLimiter()
	now := time.Now().UnixMilli()
	userIds := []string{"userA", "userB"}

	// Simulate 100 requests:
	for i := 0; i < 100; i++ {
		allowed := rateLimiter.AllowRequest(userIds[0], now+int64(i*500))
		fmt.Printf("User %s Request %d: %s\n", userIds[0], i+1, status(allowed))

		// every 10 requests, userB makes a request
		if i%10 == 0 {
			allowed = rateLimiter.AllowRequest(userIds[1], now+int64(i*500))
			fmt.Printf("User %s Request %d: %s\n", userIds[1], i+1, status(allowed))
		}
	}

	// UserB still has quota left:
	allowed := rateLimiter.AllowRequest(userIds[1], now+int64(100*500))
	fmt.Printf("User %s New Request: %s\n", userIds[1], status(allowed))

	// UserA's 101st request should be blocked:
	allowed = rateLimiter.AllowRequest(userIds[0], now+int64(100*500))
	fmt.Printf("User %s Request 101: %s\n", userIds[0], status(allowed))

	// UserB still has quota left:
	allowed = rateLimiter.AllowRequest(userIds[1], now+int64(100*500))
	fmt.Printf("User %s New Request: %s\n", userIds[1], status(allowed))
}

func status(allowed bool) string {
	if allowed {
		return "Allowed"
	}
	return "Blocked"
}
