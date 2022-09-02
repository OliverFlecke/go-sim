package main

import (
	"fmt"
	"simulation/solver/strategy"
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

	settings := SolverSettings{
		SendActionsToServer: true,
	}
	totalActions, computationTime, err := solveSimulation(sim, settings)

	if err != nil {
		fmt.Printf("Failure %s", err.Error())
	}

	if sim.GetWorld().IsSolved() {
		logger.Info("Problem solved.\n")
	} else {
		logger.Error("Problem incorrectly solved\n")
	}
	logger.Verbose("Total actions:               %d\n", totalActions)
	logger.Verbose("Total computation time:      %v\n", computationTime)
	logger.Verbose("Simulation time:             %d\n", sim.GetTicks())
}

type SolverSettings struct {
	SendActionsToServer bool
}

func solveSimulation(
	sim *simulator.Simulation,
	settings SolverSettings) (uint32, time.Duration, error) {
	totalActions := 0
	var computationTime time.Duration = 0
	a := sim.GetWorld().GetObjects(objects.AGENT)[0].(*agent.Agent)

	err := runSolverLoop(a, &computationTime, &totalActions, sim, settings)

	return uint32(totalActions), computationTime, err
}

func runSolverLoop(
	a *agent.Agent,
	computationTime *time.Duration,
	totalActions *int,
	sim *simulator.Simulation,
	settings SolverSettings) error {
	w := sim.GetWorld()

	for {
		goal := getGoal(w, a.GetLocation())
		if goal == nil {
			break
		}

		actions, t, err := solveGoal(goal, w, a)
		if err != nil {
			return err
		}
		if actions == nil {
			// fmt.Printf("Unable to solve problem!")
			break
		}
		*computationTime += t
		*totalActions += len(actions)

		sim.SetActions(a, actions)

		if settings.SendActionsToServer {
			sendActions(sim, a, actions)
		}
		events := sim.Run()

		for e := range events {
			if e.Err != nil {
				return e.Err
			}

			if len(sim.GetActions(a)) == 0 {
				break
			}
		}
	}

	return nil
}

func solveGoal(goal *objects.Goal, w world.IWorld, a *agent.Agent) ([]action.Action, time.Duration, error) {
	// fmt.Printf("\nSolving goal %v\n", goal)
	startTime := time.Now()
	box, err := strategy.FindNearestBox(w, goal)
	if err != nil {
		return nil, 0, err
	}
	if box == nil {
		return nil, 0, fmt.Errorf("no box found")
	}

	p, _, err := pathfinding.FindPath(
		w,
		a.GetLocation(),
		box.GetLocation(),
		pathfinding.AStar,
		nil)
	if err != nil {
		return nil, 0, err
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
		return nil, 0, fmt.Errorf("unable to find path from box to goal. error: %s", err.Error())
	}
	actions = append(actions, utils.Mapi(location.PathToDirections(p),
		func(_ int, dir direction.Direction) action.Action {
			return action.NewMoveWithBox(dir, box)
		})...)

	return actions, time.Since(startTime), nil
}

func getGoal(w world.IWorld, start location.Location) *objects.Goal {
	if len(w.GetUnsolvedGoals()) == 0 {
		return nil
	}

	return strategy.GoalByDependencies(w, start)
	// return closestGoal(w, start)
}
