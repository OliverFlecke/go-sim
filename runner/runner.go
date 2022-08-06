package main

import (
	"fmt"
	"log"
	simulator "simulator/core"
	"simulator/core/location"
	maps "simulator/core/map"
	"simulator/core/objects"
	pathfinding "simulator/path_finding"
	"time"
)

func main() {
	fmt.Println("Starting simulation...")
	mapName := "maps/02.map"
	world, err := maps.ParseWorldFromFile(mapName)
	if err != nil {
		log.Fatal(err)
	}

	agent := world.GetObjects(objects.AGENT)[0].(*simulator.Agent)
	fmt.Print(world.ToStringWithObjects())
	fmt.Println()

	opt := simulator.SimulationOptions{}
	opt.SetTickDuration(300 * time.Millisecond)
	sim := simulator.NewSimulation(world, []simulator.Agent{*agent}, opt)

	for {
		goal, err := getGoal(world)
		if err != nil {
			fmt.Println(err)
			continue
		}

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
			fmt.Print(world.ToStringWithObjects())
		}
		fmt.Print("\nCompleted\n")
	}
}

func getGoal(w simulator.IWorld) (location.Location, error) {
	goals := w.GetObjects(objects.GOAL)
	if len(goals) > 0 {

		return goals[0].GetLocation(), nil
	}

	fmt.Print("Enter goal: ")
	var x, y int
	_, err := fmt.Scanf("%d,%d", &x, &y)
	if err != nil {
		return location.Location{}, err
	}

	return location.New(x, y), nil
}
