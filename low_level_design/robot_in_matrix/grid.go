package main

type Grid struct {
	Width, Height int
}

func NewGrid(width, height int) Grid {
	return Grid{
		Width:  width,
		Height: height,
	}
}

func (g Grid) isInside(position Position) bool {
	return position.X >= 0 &&
		position.X < g.Width &&
		position.Y >= 0 &&
		position.Y < g.Height
}
