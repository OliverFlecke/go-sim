package main

// /* This file contains method for running the simulation
//  * in your console and controlling agents with the
//  * arrow keys
//  */

// import (
// 	"fmt"
// 	sim "simulator/core"
// 	"simulator/core/direction"

// 	"atomicgo.dev/keyboard"
// 	"atomicgo.dev/keyboard/keys"
// )

// func keyboardListener(world sim.IWorld, agent *sim.Agent) {
// 	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
// 		if key.Code == keys.CtrlC {
// 			return true, nil
// 		}

// 		dir, found := keyToDirection(key)
// 		if found {
// 			clearScreen()
// 			if !agent.MoveInWorld(world, dir) {
// 				fmt.Println("Invalid move")
// 			}
// 			fmt.Print(world.ToStringWithAgents([]sim.Agent{*agent}))
// 		}

// 		return false, nil
// 	})
// }

// func clearScreen() {
// 	fmt.Println("\033[2J")
// }

// func keyToDirection(key keys.Key) (direction.Direction, bool) {
// 	switch key.Code {
// 	case keys.Right:
// 		return direction.EAST, true
// 	case keys.Left:
// 		return direction.WEST, true
// 	case keys.Up:
// 		return direction.SOUTH, true
// 	case keys.Down:
// 		return direction.NORTH, true

// 	default:
// 		return 0, false
// 	}
// }
