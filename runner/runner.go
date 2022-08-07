package main

import (
	"fmt"
	"log"
	simulator "simulator/core"
	"simulator/core/direction"
	"simulator/core/location"
	maps "simulator/core/map"
	"simulator/core/objects"
	"simulator/core/utils"
	pathfinding "simulator/path_finding"
	"time"
	"unicode"
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
	opt.SetTickDuration(250 * time.Millisecond)
	sim := simulator.NewSimulation(world, []simulator.Agent{*agent}, opt)

	for {
		goal, err := getGoal(world)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Got goal %v\n", goal)
		boxLocation, err := findBox(world, agent.GetLocation(), *goal)
		if err != nil {
			fmt.Println(err)
			return
		}

		p, _, err := pathfinding.FindPath(world, agent.GetLocation(), *boxLocation, pathfinding.AStar)
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			continue
		}

		actions := utils.Mapi(location.PathToDirections(p), func(_ int, dir direction.Direction) simulator.Action {
			return simulator.NewMoveAction(dir)
		})

		sim.SetActions(agent, actions)
		quit := make(chan bool)
		ticker := sim.Run(quit)

		for t := range ticker {
			fmt.Printf("\n\nWorld at %s\n", t)
			fmt.Print(world.ToStringWithObjects())
		}
		fmt.Print("\nCompleted\n")
	}
}

func getGoal(w simulator.IWorld) (*objects.Goal, error) {
	goals := w.GetObjects(objects.GOAL)
	if len(goals) > 0 {
		return goals[0].(*objects.Goal), nil
	}

	return nil, fmt.Errorf("no goals found")
}

func findBox(
	world simulator.IWorld,
	start location.Location,
	goal objects.Goal) (*location.Location, error) {
	return pathfinding.FindLocation(world, start, func(l location.Location) bool {
		for _, obj := range world.GetObjectsAtLocation(l) {
			switch v := obj.(type) {
			case *objects.Box:
				if unicode.ToLower(v.GetType()) == goal.GetRune() {
					return true
				}
			}
		}

		return false
	})
}
