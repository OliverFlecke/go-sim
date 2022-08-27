package main

import (
	simulator "simulator/core"
	"simulator/core/agent"
	"simulator/core/level"
	"strings"
	"time"

	"github.com/google/uuid"
)

func getAgent(id uint32, sim *simulator.Simulation) *agent.Agent {
	for _, a := range sim.GetWorld().GetAgents() {
		if a.GetId() == id {
			return a
		}
	}

	return nil
}

func startSimulation(levelName string) string {
	id := strings.ReplaceAll(uuid.NewString(), `-`, ``)

	levelWithPath := "../maps/" + levelName
	w, _ := level.ParseWorldFromFile(levelWithPath)

	opt := simulator.SimulationOptions{}
	opt.SetTickDuration(50 * time.Millisecond)

	simulations[id] = simulator.NewSimulation(w, opt)
	simulations[id].Id = id

	return id
}
