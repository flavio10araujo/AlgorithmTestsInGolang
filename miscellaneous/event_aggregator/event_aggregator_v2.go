package main

import (
	"container/heap"
	"sync"
)

// userBucket isolates synchronization and data per user.
type userBucket struct {
	mu     sync.Mutex
	events MinHeap
}

type EventAggregatorV2 struct {
	usersMu sync.RWMutex
	users   map[string]*userBucket
}

func NewEventAggregatorV2() *EventAggregatorV2 {
	return &EventAggregatorV2{
		users: make(map[string]*userBucket),
	}
}

func (e *EventAggregatorV2) GetEventCount(userID string, now int64) int {
	bucket := e.getBucket(userID)
	if bucket == nil {
		return 0
	}

	bucket.mu.Lock()
	e.cleanupOldEvents(bucket, now)
	size := bucket.events.Len()
	bucket.mu.Unlock()

	if size == 0 {
		e.deleteBucketIfEmpty(userID, bucket)
	}

	return size
}

func (e *EventAggregatorV2) AddEvent(userID string, now int64, timestamp int64) {
	if timestamp < now-windowMs {
		return // event is too old
	}

	bucket := e.getOrCreateBucket(userID)

	bucket.mu.Lock()
	heap.Push(&bucket.events, timestamp)
	e.cleanupOldEvents(bucket, now)
	bucket.mu.Unlock()
}

func (e *EventAggregatorV2) cleanupOldEvents(bucket *userBucket, now int64) {
	for bucket.events.Len() > 0 && bucket.events.Peek() < now-windowMs {
		heap.Pop(&bucket.events)
	}
}

func (e *EventAggregatorV2) getBucket(userID string) *userBucket {
	e.usersMu.RLock()
	bucket := e.users[userID]
	e.usersMu.RUnlock()
	return bucket
}

func (e *EventAggregatorV2) getOrCreateBucket(userID string) *userBucket {
	if bucket := e.getBucket(userID); bucket != nil {
		return bucket
	}

	e.usersMu.Lock()
	defer e.usersMu.Unlock()

	if bucket := e.users[userID]; bucket != nil {
		return bucket
	}

	bucket := &userBucket{}
	heap.Init(&bucket.events)
	e.users[userID] = bucket
	return bucket
}

func (e *EventAggregatorV2) deleteBucketIfEmpty(userID string, expected *userBucket) {
	e.usersMu.Lock()
	defer e.usersMu.Unlock()

	current := e.users[userID]
	if current == nil || current != expected {
		return
	}

	current.mu.Lock()
	empty := current.events.Len() == 0
	current.mu.Unlock()

	if empty {
		delete(e.users, userID)
	}
}
