package main

// In this version, we lock by user and like that we do not lock the whole map when allowing a request.
// We still need to lock the map when creating a new bucket, but that should be less frequent (only on the first request of each user).

import "sync"

type userBucket struct {
	mu       sync.Mutex
	requests []int64
	lastSeen int64 // opcional, ajuda no cleanup
}

type RateLimiterV2 struct {
	buckets map[string]*userBucket
	mu      sync.RWMutex // protege so o map de buckets
}

func NewRateLimiterV2() *RateLimiterV2 {
	return &RateLimiterV2{
		buckets: make(map[string]*userBucket),
	}
}

func (r *RateLimiterV2) getOrCreateBucket(userID string) *userBucket {
	// caminho rapido: leitura
	r.mu.RLock()
	b := r.buckets[userID]
	r.mu.RUnlock()

	if b != nil {
		return b
	}

	// criacao (double-check)
	r.mu.Lock()
	defer r.mu.Unlock()

	if b = r.buckets[userID]; b == nil {
		b = &userBucket{}
		r.buckets[userID] = b
	}

	return b
}

func (r *RateLimiterV2) AllowRequest(userID string, now int64) bool {
	b := r.getOrCreateBucket(userID)

	b.mu.Lock()
	defer b.mu.Unlock()

	queue := b.requests
	for len(queue) > 0 && queue[0] <= now-windowMs {
		queue = queue[1:]
	}

	if len(queue) >= limit {
		b.requests = queue
		b.lastSeen = now
		return false
	}

	b.requests = append(queue, now)
	b.lastSeen = now

	return true
}
