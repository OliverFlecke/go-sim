package main

/* This file contains method for running the simulation
 * in your console and controlling agents with the
 * arrow keys
 */

import (
	"fmt"
	simulator "simulator/core"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

func keyboardListener(world *simulator.IWorld, agent *simulator.Agent) {
	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		if key.Code == keys.CtrlC {
			return true, nil
		}

		dir, found := keyToDirection(key)
		if found {
			clearScreen()
			if !agent.MoveInWorld(world, dir) {
				fmt.Println("Invalid move")
			}
			fmt.Print((*world).ToStringWithAgents([]simulator.Agent{*agent}))
		}

		return false, nil
	})
}

func clearScreen() {
	fmt.Println("\033[2J")
}

func keyToDirection(key keys.Key) (simulator.Direction, bool) {
	switch key.Code {
	case keys.Right:
		return simulator.EAST, true
	case keys.Left:
		return simulator.WEST, true
	case keys.Up:
		return simulator.SOUTH, true
	case keys.Down:
		return simulator.NORTH, true

	default:
		return 0, false
	}
}
