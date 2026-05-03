package main

import "fmt"

type Controller struct {
	robot *Robot
}

func NewController(robot *Robot) *Controller {
	return &Controller{robot: robot}
}

func (g *Controller) Status() string {
	return g.robot.Status()
}

func (g *Controller) ExecuteCommand(command Command) error {
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
