package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	simulator "simulator/core"
	"simulator/core/action"
	"simulator/core/agent"
	"simulator/core/direction"
	"simulator/core/level"
	"simulator/core/location"
	"simulator/core/logger"
	"simulator/core/objects"
	"simulator/core/utils"
	"simulator/core/world"
	"simulator/model/mapping"
	"simulator/pathfinding"
	"time"

	"google.golang.org/protobuf/encoding/protojson"
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
	w, err := level.ParseWorldFromFile(mapName)
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
	var computationTime time.Duration = 0

	// fmt.Printf("\n\nWorld at %s\n", t)
	// fmt.Print(w.ToStringWithObjects())
	runSolverLoop(w, goalId, a, &computationTime, &totalActions, sim)

	if w.IsSolved() {
		logger.Info("Problem solved.\n")
	} else {
		logger.Error("Problem incorrectly solved\n")
	}
	logger.Verbose("Total actions:               %d\n", totalActions)
	logger.Verbose("Total computation time:      %v\n", computationTime)
	logger.Verbose("Simulation time:             %d\n", sim.GetTicks())
}

func runSolverLoop(
	w world.IWorld,
	goalId int,
	a *agent.Agent,
	computationTime *time.Duration,
	totalActions *int,
	sim *simulator.Simulation) {
	for {
		goal := getGoal(w, goalId)
		if goal == nil {
			break
		}

		actions, t := solveGoal(goal, w, a)
		if actions == nil {
			fmt.Printf("Unable to solve problem!")
			break
		}
		*computationTime += t
		*totalActions += len(actions)

		sim.SetActions(a, actions)
		sendActions(a, actions)
		quit := make(chan bool)
		events := sim.Run(quit)

		for e := range events {
			if e.Err != nil {
				return
			}
			// logger.Verbose("%s\n", w.ToStringWithObjects())

			if len(sim.GetActions(a)) == 0 {
				break
			}
		}
		goalId += 1
	}
}

func sendActions(a *agent.Agent, acts []action.Action) {
	httpposturl := "http://localhost:8080/agent/0"

	dtos := mapping.ToDtos(acts)
	data, err := protojson.Marshal(dtos)
	if err != nil {
		logger.Error("%v\n", err.Error())
		return
	}

	request, err := http.NewRequest("POST", httpposturl, bytes.NewBuffer(data))
	if err != nil {
		logger.Error("%v\n", err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	client := &http.Client{}
	res, err := client.Do(request)
	if err != nil {
		logger.Error("%v\n", err.Error())
		return
	}

	logger.Info("Status: %d\n", res.StatusCode)
}

func solveGoal(goal *objects.Goal, w world.IWorld, a *agent.Agent) ([]action.Action, time.Duration) {
	// fmt.Printf("\nSolving goal %v\n", goal)
	startTime := time.Now()
	box, err := findBox(w, a.GetLocation(), goal)
	if err != nil {
		fmt.Println(err)
		return nil, 0
	}
	if box == nil {
		fmt.Println("No box found")
		return nil, 0
	}

	p, _, err := pathfinding.FindPath(
		w,
		a.GetLocation(),
		box.GetLocation(),
		pathfinding.AStar,
		nil)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
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
		fmt.Printf("Unable to find path from box to goal. Error: %s\n", err.Error())
		return nil, 0
	}
	actions = append(actions, utils.Mapi(location.PathToDirections(p),
		func(_ int, dir direction.Direction) action.Action {
			return action.NewMoveWithBox(dir, box)
		})...)

	return actions, time.Since(startTime)
}

func getGoal(w world.IWorld, i int) *objects.Goal {
	goals := w.GetUnsolvedGoals()
	if len(goals) > 0 {
		return &goals[0]
	}

	return nil
}

func findBox(
	w world.IWorld,
	start location.Location,
	goal *objects.Goal) (*objects.Box, error) {
	obj, err := pathfinding.FindLocation(w, start, func(l location.Location) objects.WorldObject {
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
