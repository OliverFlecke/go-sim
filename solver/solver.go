package main

import (
	"fmt"
	"log"
	"os"
	simulator "simulator/core"
	"simulator/core/action"
	"simulator/core/agent"
	"simulator/core/direction"
	"simulator/core/location"
	maps "simulator/core/map"
	"simulator/core/objects"
	"simulator/core/utils"
	"simulator/core/world"
	"simulator/pathfinding"
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
	w, err := maps.ParseWorldFromFile(mapName)
	if err != nil {
		log.Fatal(err)
	}

	a := w.GetObjects(objects.AGENT)[0].(*agent.Agent)
	fmt.Print(w.ToStringWithObjects())
	fmt.Println()

	opt := simulator.SimulationOptions{}
	opt.SetTickDuration(speed)
	sim := simulator.NewSimulation(w, opt)

	goalId := 0
	totalActions := 0

	for {
		goal := getGoal(w, goalId)
		if goal == nil {
			break
		}

		actions := solveGoal(goal, w, a)
		if actions == nil {
			fmt.Printf("Unable to solve problem!")
			return
		}
		totalActions += len(actions)

		sim.SetActions(a, actions)
		quit := make(chan bool)
		ticker := sim.Run(quit)

		for range ticker {
			// fmt.Printf("\n\nWorld at %s\n", t)
			// fmt.Print(w.ToStringWithObjects())
		}
		goalId += 1
	}

	fmt.Printf("Problem solved.\n")
	fmt.Printf("Total actions:   %d\n", totalActions)
	fmt.Printf("Simulation time: %d\n", sim.GetTicks())
}

func solveGoal(goal *objects.Goal, w world.IWorld, a *agent.Agent) []action.Action {
	// fmt.Printf("\nSolving goal %v\n", goal)
	box, err := findBox(w, a.GetLocation(), *goal)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if box == nil {
		fmt.Println("No box found")
		return nil
	}

	p, _, err := pathfinding.FindPath(
		w,
		a.GetLocation(),
		box.GetLocation(),
		pathfinding.AStar,
		nil)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return nil
	}

	actions := utils.Mapi(location.PathToDirections(p), func(_ int, dir direction.Direction) action.Action {
		return action.NewMove(dir)
	})
	p, _, err = pathfinding.FindPath(
		w,
		box.GetLocation(),
		goal.GetLocation(),
		pathfinding.AStar,
		func(l location.Location, w world.IWorld) bool {
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
		return nil
	}
	actions = append(actions, utils.Mapi(location.PathToDirections(p),
		func(_ int, dir direction.Direction) action.Action {
			return action.NewMoveWithBox(dir, box)
		})...)

	return actions
}

func getGoal(w world.IWorld, i int) *objects.Goal {
	goals := w.GetObjects(objects.GOAL)
	if len(goals) > i {
		return goals[i].(*objects.Goal)
	}

	return nil
}

func findBox(
	w world.IWorld,
	start location.Location,
	goal objects.Goal) (*objects.Box, error) {
	obj, err := pathfinding.FindLocation(w, start, func(l location.Location) objects.WorldObject {
		var box *objects.Box = nil
		var otherGoal objects.WorldObject

		for _, obj := range w.GetObjectsAtLocation(l) {
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
