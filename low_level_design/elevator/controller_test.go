package main

import "testing"

func TestNewElevatorController(t *testing.T) {
	controller := NewElevatorController()

	if len(controller.Elevators) != ElevatorCount {
		t.Fatalf("got %d elevators, want %d", len(controller.Elevators), ElevatorCount)
	}

	for i, elevator := range controller.Elevators {
		if elevator == nil {
			t.Fatalf("elevator %d is nil", i)
		}

		if elevator.ID != i {
			t.Errorf("elevator %d has id %d, want %d", i, elevator.ID, i)
		}

		if elevator.Floor != MinFloor {
			t.Errorf("elevator %d floor = %d, want %d", i, elevator.Floor, MinFloor)
		}

		if elevator.Direction != Idle {
			t.Errorf("elevator %d direction = %v, want %v", i, elevator.Direction, Idle)
		}

		if elevator.Requests == nil {
			t.Errorf("elevator %d requests map should be initialized", i)
		}
	}
}

func TestRequestElevatorRejectsInvalidInput(t *testing.T) {
	controller := NewElevatorController()

	tests := []struct {
		name      string
		floor     int
		direction Direction
	}{
		{name: "floor below min", floor: MinFloor - 1, direction: Up},
		{name: "floor above max", floor: MaxFloor + 1, direction: Down},
		{name: "idle direction", floor: 3, direction: Idle},
		{name: "unknown direction", floor: 3, direction: Direction(99)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := controller.RequestElevator(tt.floor, tt.direction); got {
				t.Fatalf("RequestElevator(%d, %v) = true, want false", tt.floor, tt.direction)
			}
		})
	}
}

func TestRequestElevatorAssignsPickupToMovingTowardElevator(t *testing.T) {
	controller := &Controller{
		Elevators: []*Elevator{
			{ID: 0, Floor: 1, Direction: Up, Requests: make(map[Request]struct{})},
			{ID: 1, Floor: 0, Direction: Idle, Requests: make(map[Request]struct{})},
			{ID: 2, Floor: 6, Direction: Up, Requests: make(map[Request]struct{})},
		},
	}

	if ok := controller.RequestElevator(3, Up); !ok {
		t.Fatal("RequestElevator returned false, want true")
	}

	want := Request{Floor: 3, Type: PickupUp}
	if _, ok := controller.Elevators[0].Requests[want]; !ok {
		t.Fatalf("expected request %+v on elevator 0", want)
	}

	if len(controller.Elevators[1].Requests) != 0 {
		t.Fatalf("expected elevator 1 to receive no requests, got %d", len(controller.Elevators[1].Requests))
	}

	if len(controller.Elevators[2].Requests) != 0 {
		t.Fatalf("expected elevator 2 to receive no requests, got %d", len(controller.Elevators[2].Requests))
	}
}

func TestSelectBestElevatorPriority(t *testing.T) {
	t.Run("prefers moving toward elevator over idle", func(t *testing.T) {
		controller := &Controller{
			Elevators: []*Elevator{
				{ID: 0, Floor: 2, Direction: Up, Requests: make(map[Request]struct{})},
				{ID: 1, Floor: 1, Direction: Idle, Requests: make(map[Request]struct{})},
			},
		}

		best := controller.selectBestElevator(Request{Floor: 4, Type: PickupUp})
		if best != controller.Elevators[0] {
			t.Fatalf("got elevator %d, want elevator %d", best.ID, controller.Elevators[0].ID)
		}
	})

	t.Run("falls back to nearest idle when no elevator is moving toward request", func(t *testing.T) {
		controller := &Controller{
			Elevators: []*Elevator{
				{ID: 0, Floor: 7, Direction: Down, Requests: make(map[Request]struct{})},
				{ID: 1, Floor: 5, Direction: Idle, Requests: make(map[Request]struct{})},
				{ID: 2, Floor: 8, Direction: Up, Requests: make(map[Request]struct{})},
			},
		}

		best := controller.selectBestElevator(Request{Floor: 4, Type: PickupUp})
		if best != controller.Elevators[1] {
			t.Fatalf("got elevator %d, want elevator %d", best.ID, controller.Elevators[1].ID)
		}
	})

	t.Run("falls back to nearest elevator when no idle elevator exists", func(t *testing.T) {
		controller := &Controller{
			Elevators: []*Elevator{
				{ID: 0, Floor: 6, Direction: Down, Requests: make(map[Request]struct{})},
				{ID: 1, Floor: 8, Direction: Up, Requests: make(map[Request]struct{})},
			},
		}

		best := controller.selectBestElevator(Request{Floor: 4, Type: PickupUp})
		if best != controller.Elevators[0] {
			t.Fatalf("got elevator %d, want elevator %d", best.ID, controller.Elevators[0].ID)
		}
	})
}

func TestSelectDestination(t *testing.T) {
	controller := NewElevatorController()

	if ok := controller.SelectDestination(-1, 3); ok {
		t.Fatal("SelectDestination should reject negative elevator id")
	}

	if ok := controller.SelectDestination(0, MaxFloor+1); ok {
		t.Fatal("SelectDestination should reject floor above max")
	}

	controller.Elevators[1].Floor = 4
	if ok := controller.SelectDestination(1, 4); !ok {
		t.Fatal("SelectDestination should succeed when elevator is already at floor")
	}

	if len(controller.Elevators[1].Requests) != 0 {
		t.Fatalf("same-floor destination should not add a request, got %d", len(controller.Elevators[1].Requests))
	}

	if ok := controller.SelectDestination(1, 7); !ok {
		t.Fatal("SelectDestination should add destination request")
	}

	want := Request{Floor: 7, Type: Destination}
	if _, ok := controller.Elevators[1].Requests[want]; !ok {
		t.Fatalf("expected destination request %+v to be added", want)
	}
}
