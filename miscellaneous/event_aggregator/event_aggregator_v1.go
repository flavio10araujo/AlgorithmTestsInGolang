package main

import (
	"container/heap"
	"sync"
)

type EventAggregatorV1 struct {
	mu            sync.RWMutex
	eventsPerUser map[string]*MinHeap
}

func NewEventAggregatorV1() *EventAggregatorV1 {
	return &EventAggregatorV1{
		eventsPerUser: make(map[string]*MinHeap),
	}
}

func (e *EventAggregatorV1) GetEventCount(userID string, now int64) int {
	e.mu.Lock()
	defer e.mu.Unlock()

	h := e.eventsPerUser[userID]
	if h == nil {
		return 0
	}

	e.cleanupOldEvents(h, now)

	if h.Len() == 0 {
		delete(e.eventsPerUser, userID)
		return 0
	}

	return h.Len()
}

func (e *EventAggregatorV1) AddEvent(userID string, now int64, timestamp int64) {
	if timestamp < now-windowMs {
		return // event is too old
	}

	e.mu.Lock()
	defer e.mu.Unlock()

	h := e.eventsPerUser[userID]
	if h == nil {
		newHeap := &MinHeap{}
		heap.Init(newHeap)
		e.eventsPerUser[userID] = newHeap
		h = newHeap
	}

	heap.Push(h, timestamp)
	e.cleanupOldEvents(h, now)
}

func (e *EventAggregatorV1) cleanupOldEvents(h *MinHeap, now int64) {
	for h.Len() > 0 && (h.Peek() < (now - windowMs)) {
		heap.Pop(h)
	}
}
