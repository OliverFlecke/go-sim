package main

/* This file contains method for running the simulation
 * in your console and controlling agents with the
 * arrow keys
 */

import (
	"fmt"
	"log"
	"os"
	simulator "simulator/core"
	"simulator/core/action"
	"simulator/core/agent"
	"simulator/core/direction"
	maps "simulator/core/map"
	"simulator/core/objects"
	"simulator/core/world"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide path to map file")
		return
	}
	mapName := os.Args[1]
	var err error
	w, err := maps.ParseWorldFromFile(mapName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Starting simulation...")
	a := w.GetObjects(objects.AGENT)[0].(*agent.Agent)
	fmt.Print(w.ToStringWithObjects())
	fmt.Println()

	keyboardListener(w, a)
}

func keyboardListener(w world.IWorld, a *agent.Agent) {
	opt := simulator.SimulationOptions{}
	sim := simulator.NewSimulation(w, opt)

	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		if key.Code == keys.CtrlC {
			return true, nil
		}

		dir, found := keyToDirection(key)
		if found {
			clearScreen()
			sim.SetActions(a, []action.Action{action.NewMove(dir)})
			for range sim.Run(nil) {
				fmt.Print(w.ToStringWithObjects())
			}
		}

		return false, nil
	})
}

func clearScreen() {
	fmt.Println("\033[2J")
}

func keyToDirection(key keys.Key) (direction.Direction, bool) {
	switch key.Code {
	case keys.Right:
		return direction.EAST, true
	case keys.Left:
		return direction.WEST, true
	case keys.Up:
		return direction.SOUTH, true
	case keys.Down:
		return direction.NORTH, true

	default:
		return 0, false
	}
}
