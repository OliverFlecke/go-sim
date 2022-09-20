package main

import (
	"fmt"
	"log"
	"path/filepath"
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
	log.Printf("Attempting to solve all MA levels")

	testLevels(t, LEVEL_DIRECTORY, func(s string) bool {
		return strings.Contains(s, "MA")
	})
}

func TestSASolving(t *testing.T) {
	lvl := filepath.Join(LEVEL_DIRECTORY, "02.map")
	sim, stats, _ := executeSimulation(t, lvl)

	assert.True(t, sim.GetWorld().IsSolved(), "world is not solved")
	assert.Equal(t, uint64(4), stats.TotalActions, "it should take 2 actions to get to the box, and another two to get the box to its goal")
	assert.Equal(t, uint64(4), stats.TotalTicks, "it should take 4 actions and therefore 4 ticks to solve the problem")
}

func TestMASolving(t *testing.T) {
	lvl := filepath.Join(LEVEL_DIRECTORY, "MA01.map")
	sim, stats, _ := executeSimulation(t, lvl)

	assert.True(t, sim.GetWorld().IsSolved(), "world is not solved")
	assert.Equal(t, uint64(2), stats.TotalTicks, "total number of ticks should only be two, as both problems can be solved in parallel in two steps")
	assert.Equal(t, uint64(4), stats.TotalActions, "total number of actions required is 2 + 2 = 4")
}

func testLevels(t *testing.T, level_directory string, filter func(string) bool) {
	level.GetMaps(level_directory, func(lvlNameWithPath string) {
		if !filter(lvlNameWithPath) {
			return
		}

		name := strings.TrimPrefix(lvlNameWithPath, level_directory+"/")

		sim, stats, _ := executeSimulation(t, lvlNameWithPath)

		statsStr := fmt.Sprintf("Actions: %5d, duration: %15s", stats.TotalActions, stats.ComputationDuration)
		if sim.GetWorld().IsSolved() {
			logger.Info("Solved           \t%-30s %s\n", name, statsStr)
		} else {
			logger.Error("Failed to solve: \t%-30s %s\n", name, statsStr)
			t.Fail()
		}
	})
}

func executeSimulation(t *testing.T, levelName string) (*simulator.Simulation, SimulationStatistics, error) {
	w, err := level.ParseWorldFromFile("", levelName)
	assert.NoError(t, err)

	opt := simulator.SimulationOptions{}
	opt.SetTickDuration(10 * time.Microsecond)
	sim := simulator.NewSimulation(w, opt)

	settings := SolverSettings{
		SendActionsToServer: false,
	}

	stats, err := solveSimulation(sim, settings)
	return sim, stats, err
}
