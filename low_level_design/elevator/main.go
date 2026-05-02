package main

import (
	"fmt"
)

func main() {
	controller := NewElevatorController()

	controller.RequestElevator(3, Up)
	controller.RequestElevator(7, Down)

	controller.SelectDestination(0, 8)
	controller.SelectDestination(1, 1)

	for i := 0; i < 10; i++ {
		fmt.Printf("Step %d\n", i)

		for _, elevator := range controller.Elevators {
			fmt.Printf(
				"Elevator %d | Floor: %d | Direction: %s | Requests: %d\n",
				elevator.ID,
				elevator.Floor,
				elevator.Direction,
				len(elevator.Requests),
			)
		}

		fmt.Println()
		controller.Step()
	}
}
