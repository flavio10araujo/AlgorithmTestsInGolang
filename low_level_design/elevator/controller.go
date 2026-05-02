package main

import (
	"errors"
	"math"
)

const (
	MinFloor      = 0
	MaxFloor      = 9
	ElevatorCount = 3
)

type Controller struct {
	Elevators []*Elevator
}

func NewElevatorController() *Controller {
	elevators := make([]*Elevator, 0, ElevatorCount)

	for i := 0; i < ElevatorCount; i++ {
		elevators = append(elevators, NewElevator(i))
	}

	return &Controller{
		Elevators: elevators,
	}
}

func (c *Controller) RequestElevator(floor int, direction Direction) bool {
	if floor < MinFloor || floor > MaxFloor {
		return false
	}

	if direction != Up && direction != Down {
		return false
	}

	reqType := PickupUp
	if direction == Down {
		reqType = PickupDown
	}

	req := Request{
		Floor: floor,
		Type:  reqType,
	}

	best := c.selectBestElevator(req)
	if best == nil {
		return false
	}

	return best.AddRequest(req)
}

func (c *Controller) SelectDestination(elevatorID int, floor int) bool {
	if floor < MinFloor || floor > MaxFloor {
		return false
	}

	if elevatorID < 0 || elevatorID >= len(c.Elevators) {
		return false
	}

	elevator := c.Elevators[elevatorID]

	if floor == elevator.Floor {
		return true
	}

	return elevator.AddRequest(Request{
		Floor: floor,
		Type:  Destination,
	})
}

func (c *Controller) Step() {
	for _, elevator := range c.Elevators {
		elevator.Step()
	}
}

func (c *Controller) selectBestElevator(req Request) *Elevator {
	if elevator := c.findMovingToward(req); elevator != nil {
		return elevator
	}

	if elevator := c.findNearestIdle(req.Floor); elevator != nil {
		return elevator
	}

	return c.findNearest(req.Floor)
}

func (c *Controller) findMovingToward(req Request) *Elevator {
	targetDirection, err := directionFromPickup(req.Type)
	if err != nil {
		return nil
	}

	var best *Elevator
	minDistance := math.MaxInt

	for _, elevator := range c.Elevators {
		if elevator.Direction != targetDirection {
			continue
		}

		if targetDirection == Up && elevator.Floor > req.Floor {
			continue
		}

		if targetDirection == Down && elevator.Floor < req.Floor {
			continue
		}

		distance := abs(elevator.Floor - req.Floor)
		if distance < minDistance {
			minDistance = distance
			best = elevator
		}
	}

	return best
}

func (c *Controller) findNearestIdle(floor int) *Elevator {
	var best *Elevator
	minDistance := math.MaxInt

	for _, elevator := range c.Elevators {
		if elevator.Direction != Idle {
			continue
		}

		distance := abs(elevator.Floor - floor)
		if distance < minDistance {
			minDistance = distance
			best = elevator
		}
	}

	return best
}

func (c *Controller) findNearest(floor int) *Elevator {
	var best *Elevator
	minDistance := math.MaxInt

	for _, elevator := range c.Elevators {
		distance := abs(elevator.Floor - floor)
		if distance < minDistance {
			minDistance = distance
			best = elevator
		}
	}

	return best
}

func directionFromPickup(reqType RequestType) (Direction, error) {
	switch reqType {
	case PickupUp:
		return Up, nil
	case PickupDown:
		return Down, nil
	default:
		return Idle, errors.New("request type is not a pickup request")
	}
}

func abs(value int) int {
	if value < 0 {
		return -value
	}

	return value
}
