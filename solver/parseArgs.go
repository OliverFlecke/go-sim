package main

import (
	"fmt"
	"log"
	"os"
	simulator "simulator/core"
	"simulator/core/level"
	"time"
)

func parseArgs() *simulator.Simulation {
	switch len(os.Args) {
	case 2:
		return parseServerArgs()
	case 3, 4:
		return parseLocal()
	default:
		log.Fatalf("Invalid number of args %d", len(os.Args))
		return nil // UNREACHABLE
	}
}

func parseServerArgs() *simulator.Simulation {
	simId := os.Args[1]
	opt := simulator.SimulationOptions{
		TickDuration: 10 * time.Millisecond,
	}

	w := downloadLevel(simId)
	sim := simulator.NewSimulation(w, opt)
	sim.Id = simId

	return sim
}

func parseLocal() *simulator.Simulation {
	mapName := os.Args[1]
	var speed = defaultSpeed
	var err error
	if len(os.Args) > 2 {
		speed, err = time.ParseDuration(os.Args[2])
		if err != nil {
			log.Fatalf("Time must be an integer. Error: %s", err)
		}
	}

	w, err := level.ParseWorldFromFile("level", mapName)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print(w.ToStringWithObjects())
	fmt.Println()

	opt := simulator.SimulationOptions{}
	opt.SetTickDuration(speed)
	sim := simulator.NewSimulation(w, opt)
	if len(os.Args) > 3 {
		sim.Id = os.Args[3]
	}

	return sim
}
