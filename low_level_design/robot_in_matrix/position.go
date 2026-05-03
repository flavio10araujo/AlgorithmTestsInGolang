package main

type Position struct {
	X int
	Y int
}

func (p Position) MoveBy(dx, dy int) Position {
	return Position{
		X: p.X + dx,
		Y: p.Y + dy,
	}
}
