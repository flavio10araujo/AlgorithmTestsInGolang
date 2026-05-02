package main

import "math"

type Direction int

const (
	Idle Direction = iota
	Up
	Down
)

func (d Direction) String() string {
	switch d {
	case Up:
		return "UP"
	case Down:
		return "DOWN"
	default:
		return "IDLE"
	}
}

type Elevator struct {
	ID        int
	Floor     int
	Direction Direction
	Requests  map[Request]struct{}
}

func NewElevator(id int) *Elevator {
	return &Elevator{
		ID:        id,
		Floor:     0,
		Direction: Idle,
		Requests:  make(map[Request]struct{}),
	}
}

func (e *Elevator) AddRequest(req Request) bool {
	if req.Floor < MinFloor || req.Floor > MaxFloor {
		return false
	}

	if req.Floor == e.Floor {
		return true
	}

	e.Requests[req] = struct{}{}
	return true
}

func (e *Elevator) Step() {
	if len(e.Requests) == 0 {
		e.Direction = Idle
		return
	}

	if e.Direction == Idle {
		nearest := e.findNearestRequest()
		if nearest.Floor > e.Floor {
			e.Direction = Up
		} else if nearest.Floor < e.Floor {
			e.Direction = Down
		} else {
			e.stopAtCurrentFloor()
			return
		}
	}

	if e.shouldStopAtCurrentFloor() {
		e.stopAtCurrentFloor()
		return
	}

	if !e.hasRequestsAhead(e.Direction) {
		if e.Direction == Up {
			e.Direction = Down
		} else if e.Direction == Down {
			e.Direction = Up
		}
	}

	if e.Direction == Up && e.Floor < MaxFloor {
		e.Floor++
	} else if e.Direction == Down && e.Floor > MinFloor {
		e.Floor--
	}
}

func (e *Elevator) shouldStopAtCurrentFloor() bool {
	destination := Request{Floor: e.Floor, Type: Destination}

	if _, ok := e.Requests[destination]; ok {
		return true
	}

	if e.Direction == Up {
		_, ok := e.Requests[Request{Floor: e.Floor, Type: PickupUp}]
		return ok
	}

	if e.Direction == Down {
		_, ok := e.Requests[Request{Floor: e.Floor, Type: PickupDown}]
		return ok
	}

	return false
}

func (e *Elevator) stopAtCurrentFloor() {
	delete(e.Requests, Request{Floor: e.Floor, Type: Destination})
	delete(e.Requests, Request{Floor: e.Floor, Type: PickupUp})
	delete(e.Requests, Request{Floor: e.Floor, Type: PickupDown})

	if len(e.Requests) == 0 {
		e.Direction = Idle
	}
}

func (e *Elevator) findNearestRequest() Request {
	var nearest Request
	minDistance := math.MaxInt

	for req := range e.Requests {
		distance := abs(req.Floor - e.Floor)
		if distance < minDistance {
			minDistance = distance
			nearest = req
		}
	}

	return nearest
}

func (e *Elevator) hasRequestsAhead(dir Direction) bool {
	for req := range e.Requests {
		if dir == Up && req.Floor > e.Floor {
			return true
		}

		if dir == Down && req.Floor < e.Floor {
			return true
		}
	}

	return false
}
