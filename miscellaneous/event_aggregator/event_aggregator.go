package main

import "container/heap"

const windowMs = 5 * 60 * 1000

type MinHeap []int64

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(int64))
}

func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func (h MinHeap) Peek() int64 {
	return h[0]
}

type EventAggregator struct {
	eventsPerUser map[string]*MinHeap
}

func NewEventAggregator() *EventAggregator {
	return &EventAggregator{
		eventsPerUser: make(map[string]*MinHeap),
	}
}

func (e *EventAggregator) GetEventCount(userID string, now int64) int {
	heap := e.eventsPerUser[userID]
	if heap == nil {
		return 0
	}

	e.cleanupOldEvents(heap, now)
	size := heap.Len()

	if size == 0 {
		delete(e.eventsPerUser, userID)
	}

	return size
}

func (e *EventAggregator) AddEvent(userID string, now int64, timestamp int64) {
	if timestamp < now-windowMs {
		return // event is too old
	}

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

func (e *EventAggregator) cleanupOldEvents(h *MinHeap, now int64) {
	for h.Len() > 0 && (h.Peek() < (now - windowMs)) {
		heap.Pop(h)
	}
}
