package main

import (
	simulator "simulator/core"
	"simulator/core/agent"
	"simulator/core/level"
	"strings"
	"time"

	"github.com/google/uuid"
)

const PATH_TO_LEVELS = "../maps"

func getAgent(id uint32, sim *simulator.Simulation) *agent.Agent {
	for _, a := range sim.GetWorld().GetAgents() {
		if a.GetId() == id {
			return a
		}
	}

	return nil
}

func startSimulation(levelName string) string {
	id := generateId()
	w, _ := level.ParseWorldFromFile(PATH_TO_LEVELS, levelName)

	opt := simulator.SimulationOptions{}
	opt.SetTickDuration(50 * time.Millisecond)

	simulations[id] = simulator.NewSimulation(w, opt)
	simulations[id].Id = id

	return id
}

func generateId() string {
	return strings.ReplaceAll(uuid.NewString(), `-`, ``)
}
