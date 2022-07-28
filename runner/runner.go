package main

import (
	"fmt"
	simulator "simulator/core"
)

func main() {
	fmt.Println("Running simulation")
	agent := simulator.NewAgent("Agent A")
	fmt.Println(agent)
}
