package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	switch d {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	default:
		return "Unknown"
	}
}

func (d Direction) TurnLeft() Direction {
	return Direction((int(d) + 3) % 4)
}

func (d Direction) TurnRight() Direction {
	return Direction((int(d) + 1) % 4)
}

func (d Direction) Delta() (dx, dy int) {
	switch d {
	case North:
		return 0, 1
	case East:
		return 1, 0
	case South:
		return 0, -1
	case West:
		return -1, 0
	default:
		return 0, 0
	}
}

type Command rune

const (
	CommandRight    Command = 'R'
	CommandLeft     Command = 'L'
	CommandForward  Command = 'F'
	CommandBackward Command = 'B'
)

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

func main() {
	grid := NewGrid(10, 10)
	position := Position{X: 0, Y: 0}
	direction := North
	robot := NewRobot(grid, position, direction)
	game := NewGame(robot)

	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())

		if strings.EqualFold(input, "exit") {
			fmt.Println("Exiting...")
			return
		}

		input = strings.ToUpper(input)

		for _, rawCommand := range input {
			command := Command(rawCommand)

			if err := game.ExecuteCommand(command); err != nil {
				fmt.Println(err)
				continue
			}
		}

		fmt.Println(game.Status())
	}
}

type Game struct {
	robot *Robot
}

func NewGame(robot *Robot) *Game {
	return &Game{robot: robot}
}

func (g *Game) Status() string {
	return g.robot.Status()
}

func (g *Game) ExecuteCommand(command Command) error {
	switch command {
	case CommandLeft:
		g.robot.TurnLeft()
	case CommandRight:
		g.robot.TurnRight()
	case CommandForward:
		g.robot.MoveForward()
	case CommandBackward:
		g.robot.MoveBackward()
	default:
		return fmt.Errorf("invalid command: %c", command)
	}
	return nil
}
