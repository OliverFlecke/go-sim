package main

import (
	"fmt"
	"log"
	simulator "simulator/core"
	maps "simulator/core/map"
	pathfinding "simulator/path_finding"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

func main() {
	fmt.Println("Starting simulation...")
	mapName := "maps/01.map"
	world, err := maps.GetStringFromFile(mapName)
	if err != nil {
		log.Fatal(err)
	}

	agent := simulator.NewAgentWithStartLocation("Agent 0", '0', simulator.NewLocation(1, 1))
	fmt.Print(world.ToStringWithAgents([]simulator.Agent{*agent}))
	fmt.Println()

	sim := simulator.NewSimulation(
		world,
		[]simulator.Agent{*agent},
		simulator.SimulationOptions{})

	for {
		fmt.Print("Enter goal: ")
		var x, y int
		_, err := fmt.Scanf("%d,%d", &x, &y)
		if err != nil {
			fmt.Print(err)
			continue
		}

		goal := simulator.NewLocation(x, y)
		fmt.Printf("Got goal %v\n", goal)

		p, _, err := pathfinding.FindPath(world, agent.GetLocation(), goal, pathfinding.AStar)
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			continue
		}

		sim.SetActions(agent, simulator.PathToDirections(p))
		quit := make(chan bool)
		ticker := sim.Run(quit)

		for t := range ticker {
			fmt.Printf("\n\nWorld at %s\n", t)
			fmt.Print(world.ToStringWithAgents([]simulator.Agent{*agent}))
		}
		fmt.Print("\nCompleted\n")
	}
}

func keyboardListener(world *simulator.World, agent *simulator.Agent) {
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
			fmt.Print(world.ToStringWithAgents([]simulator.Agent{*agent}))
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
