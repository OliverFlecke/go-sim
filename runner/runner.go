package main

import (
	"fmt"
	"log"
	"os"
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

const defaultSpeed time.Duration = 250 * time.Millisecond

func main() {
	fmt.Println("Starting simulation...")
	mapName := os.Args[1]
	var speed = defaultSpeed
	var err error
	if len(os.Args) > 2 {
		speed, err = time.ParseDuration(os.Args[2])
		if err != nil {
			fmt.Printf("Time must be an integer. Error: %s", err)
			return
		}
	}
	world, err := maps.ParseWorldFromFile(mapName)
	if err != nil {
		log.Fatal(err)
	}

	agent := world.GetObjects(objects.AGENT)[0].(*simulator.Agent)
	fmt.Print(world.ToStringWithObjects())
	fmt.Println()

	opt := simulator.SimulationOptions{}
	opt.SetTickDuration(speed)
	sim := simulator.NewSimulation(world, []simulator.Agent{*agent}, opt)

	goalId := 0

	for {
		goal := getGoal(world, goalId)
		if goal == nil {
			fmt.Printf("\nNo more goals to solve. Stopping simulator")
			return
		}

		fmt.Printf("\nSolving goal %v\n", goal)
		box, err := findBox(world, agent.GetLocation(), *goal)
		if err != nil {
			fmt.Println(err)
			return
		}
		if box == nil {
			fmt.Println("No box found")
			return
		}

		p, _, err := pathfinding.FindPath(
			world,
			agent.GetLocation(),
			box.GetLocation(),
			pathfinding.AStar,
			nil)
		if err != nil {
			fmt.Printf("Error: %s\n", err.Error())
			continue
		}

		actions := utils.Mapi(location.PathToDirections(p), func(_ int, dir direction.Direction) simulator.Action {
			return simulator.NewActionMove(dir)
		})
		p, _, err = pathfinding.FindPath(
			world,
			box.GetLocation(),
			goal.GetLocation(),
			pathfinding.AStar,
			func(l location.Location, w simulator.IWorld) bool {
				for _, o := range w.GetObjectsAtLocation(l) {
					switch o.(type) {
					case *objects.Box:
						return false
					}
				}

				return true
			})
		if err != nil {
			fmt.Printf("Unable to find path from box to goal. Error: %s\n", err.Error())
			return
		}
		actions = append(actions, utils.Mapi(location.PathToDirections(p),
			func(_ int, dir direction.Direction) simulator.Action {
				return simulator.NewActionMoveWithBox(dir, box)
			})...)

		sim.SetActions(agent, actions)
		quit := make(chan bool)
		ticker := sim.Run(quit)

		for t := range ticker {
			fmt.Printf("\n\nWorld at %s\n", t)
			fmt.Print(world.ToStringWithObjects())
		}
		goalId += 1
	}
}

func getGoal(w simulator.IWorld, i int) *objects.Goal {
	goals := w.GetObjects(objects.GOAL)
	if len(goals) > i {
		return goals[i].(*objects.Goal)
	}

	return nil
}

func findBox(
	world simulator.IWorld,
	start location.Location,
	goal objects.Goal) (*objects.Box, error) {
	obj, err := pathfinding.FindLocation(world, start, func(l location.Location) objects.WorldObject {
		var box *objects.Box = nil
		var otherGoal objects.WorldObject

		for _, obj := range world.GetObjectsAtLocation(l) {
			switch v := obj.(type) {
			case *objects.Box:
				if unicode.ToLower(v.GetType()) == goal.GetRune() {
					box = v
				}
			case *objects.Goal:
				otherGoal = v
			}
		}

		if box == nil || (otherGoal != nil && unicode.ToLower(box.GetType()) == goal.GetRune()) {
			return nil
		}

		return box
	})

	if err != nil {
		return nil, err
	}

	return obj.(*objects.Box), nil
}
