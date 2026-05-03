package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Command rune

const (
	CommandRight    Command = 'R'
	CommandLeft     Command = 'L'
	CommandForward  Command = 'F'
	CommandBackward Command = 'B'
)

func main() {
	grid := NewGrid(10, 10)
	position := Position{X: 0, Y: 0}
	direction := North
	robot := NewRobot(grid, position, direction)
	controller := NewController(robot)

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

			if err := controller.ExecuteCommand(command); err != nil {
				fmt.Println(err)
				continue
			}
		}

		fmt.Println(controller.Status())
	}
}
