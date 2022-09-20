package main

import (
	simulator "simulator/core"
	"simulator/core/agent"
	"strings"

	"github.com/google/uuid"
)

const PATH_TO_LEVELS = "../level"

func getAgent(id uint32, sim *simulator.Simulation) *agent.Agent {
	for _, a := range sim.GetWorld().GetAgents() {
		if a.GetId() == id {
			return a
		}
	}

	return nil
}

func generateId() string {
	return strings.ReplaceAll(uuid.NewString(), `-`, ``)
}
