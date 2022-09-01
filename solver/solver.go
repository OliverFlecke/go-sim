package main

import (
	"fmt"
	simulator "simulator/core"
	"simulator/core/action"
	"simulator/core/agent"
	"simulator/core/direction"
	"simulator/core/location"
	"simulator/core/logger"
	"simulator/core/objects"
	"simulator/core/utils"
	"simulator/core/world"
	"simulator/pathfinding"
	"time"
)

const defaultSpeed time.Duration = 250 * time.Millisecond

func main() {
	sim := parseArgs()

	fmt.Println("Starting simulation...")

	totalActions, computationTime := solveSimulation(sim)

	if sim.GetWorld().IsSolved() {
		logger.Info("Problem solved.\n")
	} else {
		logger.Error("Problem incorrectly solved\n")
	}
	logger.Verbose("Total actions:               %d\n", totalActions)
	logger.Verbose("Total computation time:      %v\n", computationTime)
	logger.Verbose("Simulation time:             %d\n", sim.GetTicks())
}

func solveSimulation(sim *simulator.Simulation) (uint32, time.Duration) {
	totalActions := 0
	var computationTime time.Duration = 0
	a := sim.GetWorld().GetObjects(objects.AGENT)[0].(*agent.Agent)

	runSolverLoop(a, &computationTime, &totalActions, sim)

	return uint32(totalActions), computationTime
}

func runSolverLoop(
	a *agent.Agent,
	computationTime *time.Duration,
	totalActions *int,
	sim *simulator.Simulation) {
	w := sim.GetWorld()

	for {
		goal := getGoal(w, a.GetLocation())
		if goal == nil {
			break
		}

		actions, t := solveGoal(goal, w, a)
		if actions == nil {
			// fmt.Printf("Unable to solve problem!")
			break
		}
		*computationTime += t
		*totalActions += len(actions)

		sim.SetActions(a, actions)
		// sendActions(sim, a, actions)
		events := sim.Run()

		for e := range events {
			if e.Err != nil {
				return
			}

			if len(sim.GetActions(a)) == 0 {
				break
			}
		}
	}
}

func solveGoal(goal *objects.Goal, w world.IWorld, a *agent.Agent) ([]action.Action, time.Duration) {
	// fmt.Printf("\nSolving goal %v\n", goal)
	startTime := time.Now()
	box, err := findNearestBox(w, goal)
	if err != nil {
		// fmt.Println(err)
		return nil, 0
	}
	if box == nil {
		// fmt.Println("No box found")
		return nil, 0
	}

	p, _, err := pathfinding.FindPath(
		w,
		a.GetLocation(),
		box.GetLocation(),
		pathfinding.AStar,
		nil)
	if err != nil {
		// fmt.Printf("Error: %s\n", err.Error())
		return nil, 0
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
		// fmt.Printf("Unable to find path from box to goal. Error: %s\n", err.Error())
		return nil, 0
	}
	actions = append(actions, utils.Mapi(location.PathToDirections(p),
		func(_ int, dir direction.Direction) action.Action {
			return action.NewMoveWithBox(dir, box)
		})...)

	return actions, time.Since(startTime)
}

func getGoal(w world.IWorld, start location.Location) *objects.Goal {
	if len(w.GetUnsolvedGoals()) == 0 {
		return nil
	}

	result, _ := pathfinding.FindClosestObject(w, start,
		func(l location.Location) objects.WorldObject {
			for _, x := range w.GetObjectsAtLocation(l) {
				switch g := x.(type) {
				case *objects.Goal:
					if !w.IsGoalSolved(g) {
						return g
					}
				}
			}
			return nil
		})

	return result.(*objects.Goal)
}

func findNearestBox(w world.IWorld, goal *objects.Goal) (*objects.Box, error) {
	obj, err := pathfinding.FindClosestObject(w, goal.GetLocation(),
		func(l location.Location) objects.WorldObject {
			var box *objects.Box = nil
			var otherGoal *objects.Goal

			for _, obj := range w.GetObjectsAtLocation(l) {
				switch v := obj.(type) {
				case *objects.Box:
					if v.Matches(*goal) {
						box = v
					}
				case *objects.Goal:
					otherGoal = v
				}
			}

			if box == nil || (otherGoal != nil && box.Matches(*otherGoal)) {
				return nil
			}

			return box
		})

	if err != nil {
		return nil, err
	}

	return obj.(*objects.Box), nil
}
