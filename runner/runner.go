package main

import (
	"fmt"
	simulator "simulator/core"
	dir "simulator/core/direction"
)

func main() {
	fmt.Println("Running simulation")
	agent := simulator.NewAgent("Agent A")
	moves := []dir.Direction{dir.NORTH, dir.EAST, dir.EAST, dir.EAST}

	for _, d := range moves {
		agent.Move(d)
		fmt.Println(agent)
	}
}
