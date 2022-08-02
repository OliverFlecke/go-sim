package main

import (
	"fmt"
	simulator "simulator/core"
)

func main() {
	fmt.Println("Running simulation")
	agent := simulator.NewAgent("Agent A")
	moves := []simulator.Direction{simulator.NORTH, simulator.EAST, simulator.EAST, simulator.EAST}

	for _, d := range moves {
		agent.Move(d)
		fmt.Println(agent)
	}
}
