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
	"simulator/core/level"
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
	w, err := level.ParseWorldFromFile("", mapName)
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
	var box *objects.Box

	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		if key.Code == keys.CtrlC {
			return true, nil
		}

		switch key.String() {
		case "p":
			if box != nil {
				box = nil
			} else {
				for _, x := range w.GetObjectsAtLocation(a.GetLocation()) {
					switch o := x.(type) {
					case *objects.Box:
						box = o
					}
				}
			}
		default:
			dir, found := keyToDirection(key)
			if found {
				clearScreen()
				var act action.Action
				if box != nil {
					act = action.NewMoveWithBox(dir, box)
				} else {
					act = action.NewMove(dir)
				}
				sim.SetActions(a, []action.Action{act})
				for range sim.Run() {
					fmt.Print(w.ToStringWithObjects())
				}
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
