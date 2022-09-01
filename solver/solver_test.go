package main

import (
	"fmt"
	"log"
	simulator "simulator/core"
	"simulator/core/level"
	"simulator/core/logger"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAllLevels(t *testing.T) {
	log.Printf("Attempting to solve all levels")
	const level_directory = "../maps"

	level.GetMaps(level_directory, func(levelName string) {
		if strings.Contains(levelName, "unsolveable") {
			return
		}

		name := strings.TrimPrefix(levelName, level_directory+"/")
		// t.Logf("Solving level: %s")

		w, err := level.ParseWorldFromFile("", levelName)
		assert.NoError(t, err)

		opt := simulator.SimulationOptions{}
		opt.SetTickDuration(10 * time.Microsecond)
		sim := simulator.NewSimulation(w, opt)

		settings := SolverSettings{
			SendActionsToServer: false,
		}

		// Run solver
		totalActions, duration := solveSimulation(sim, settings)

		stats := fmt.Sprintf("Actions: %5d, duration: %15s", totalActions, duration)
		if sim.GetWorld().IsSolved() {
			logger.Info("Solved           \t%-30s %s\n", name, stats)
		} else {
			logger.Error("Failed to solve: \t%-30s %s\n", name, stats)
			t.Fail()
		}
	})
}
