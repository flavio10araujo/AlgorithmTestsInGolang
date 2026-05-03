package main

import "fmt"

type Robot struct {
	position  Position
	direction Direction
	grid      Grid
}

func NewRobot(grid Grid, position Position, direction Direction) *Robot {
	return &Robot{
		grid:      grid,
		position:  position,
		direction: direction,
	}
}

func (r *Robot) Status() string {
	return fmt.Sprintf(
		"Position: (%d, %d), Direction: %s",
		r.position.X,
		r.position.Y,
		r.direction,
	)
}

func (r *Robot) TurnLeft() {
	r.direction = r.direction.TurnLeft()
}

func (r *Robot) TurnRight() {
	r.direction = r.direction.TurnRight()
}

func (r *Robot) MoveForward() {
	r.move(1)
}

func (r *Robot) MoveBackward() {
	r.move(-1)
}

func (r *Robot) move(stepSign int) {
	dx, dy := r.direction.Delta()
	nextPosition := r.position.MoveBy(dx*stepSign, dy*stepSign)

	if r.grid.isInside(nextPosition) {
		r.position = nextPosition
	}
}
