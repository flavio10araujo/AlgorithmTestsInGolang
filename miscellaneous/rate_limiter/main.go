package main

import (
	"fmt"
	"time"
)

const (
	limit    = 100
	windowMs = 60_000
)

func main() {
	rateLimiter := NewRateLimiterV2()
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
