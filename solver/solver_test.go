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

const LEVEL_DIRECTORY = "../level"

func TestAllSALevels(t *testing.T) {
	log.Printf("Attempting to solve all SA levels")

	testLevels(t, LEVEL_DIRECTORY, func(s string) bool {
		return !strings.Contains(s, "unsolvable") && !strings.Contains(s, "MA")
	})
}

func TestAllMALevels(t *testing.T) {
	log.Printf("Attempting to solve all SA levels")

	testLevels(t, LEVEL_DIRECTORY, func(s string) bool {
		return strings.Contains(s, "MA")
	})
}

func testLevels(t *testing.T, level_directory string, filter func(string) bool) {
	level.GetMaps(level_directory, func(levelName string) {
		if !filter(levelName) {
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
		totalActions, duration, _ := solveSimulation(sim, settings)

		stats := fmt.Sprintf("Actions: %5d, duration: %15s", totalActions, duration)
		if sim.GetWorld().IsSolved() {
			logger.Info("Solved           \t%-30s %s\n", name, stats)
		} else {
			logger.Error("Failed to solve: \t%-30s %s\n", name, stats)
			t.Fail()
		}
	})
}
