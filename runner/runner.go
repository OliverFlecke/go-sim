package main

import (
	"fmt"
	"log"
	simulator "simulator/core"
	"simulator/core/location"
	maps "simulator/core/map"
	pathfinding "simulator/path_finding"
	"time"
)

func main() {
	fmt.Println("Starting simulation...")
	mapName := "maps/01.map"
	world, err := maps.ParseWorldFromFile(mapName)
	if err != nil {
		log.Fatal(err)
	}

	agent := simulator.NewAgentWithStartLocation("Agent 0", '0', location.New(1, 1))
	fmt.Print(world.ToStringWithAgents([]simulator.Agent{*agent}))
	fmt.Println()

	opt := simulator.SimulationOptions{}
	opt.SetTickDuration(300 * time.Millisecond)
	sim := simulator.NewSimulation(world, []simulator.Agent{*agent}, opt)

	for {
		fmt.Print("Enter goal: ")
		var x, y int
		_, err := fmt.Scanf("%d,%d", &x, &y)
		if err != nil {
			fmt.Print(err)
			continue
		}

		goal := location.New(x, y)
		fmt.Printf("Got goal %v\n", goal)

		p, _, err := pathfinding.FindPath(world, agent.GetLocation(), goal, pathfinding.AStar)
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			continue
		}

		sim.SetActions(agent, location.PathToDirections(p))
		quit := make(chan bool)
		ticker := sim.Run(quit)

		for t := range ticker {
			fmt.Printf("\n\nWorld at %s\n", t)
			fmt.Print(world.ToStringWithAgents([]simulator.Agent{*agent}))
		}
		fmt.Print("\nCompleted\n")
	}
}
