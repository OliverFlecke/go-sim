package simulator

import (
	"fmt"
	"time"
)

type Simulation struct {
	world   *World
	options SimulationOptions
	agents  []Agent
	actions map[*Agent][]Direction // TODO: Maybe introduce an action interface for this
}

type SimulationOptions struct {
	waitForAction bool
	tickDuration  time.Duration
}

func NewSimulation(world *World, agents []Agent, options SimulationOptions) *Simulation {
	return &Simulation{
		world:   world,
		agents:  agents,
		options: options,
		actions: make(map[*Agent][]Direction),
	}
}

func (s *Simulation) Run(quit chan bool) <-chan time.Time {
	ticker := time.NewTicker(1 * time.Second)
	output := make(chan time.Time)

	go func() {
		defer ticker.Stop()
		defer close(output)

		for {
			select {
			case <-quit:
				fmt.Println("Stopping simulation")
				return
			case t := <-ticker.C:
				// fmt.Printf("Tick %v\n", t)

				for agent, actions := range s.actions {
					if len(actions) > 0 {
						dir := actions[0]
						agent.MoveInWorld(s.world, dir)
						s.actions[agent] = actions[1:]
					}
				}

				output <- t
			}
		}
	}()

	return output
}

func (s *Simulation) SetActions(agent *Agent, directions []Direction) {
	s.actions[agent] = append(s.actions[agent], directions...)
}
