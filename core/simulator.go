package simulator

import (
	"fmt"
	"time"
)

type Simulation struct {
	world   *IWorld
	options SimulationOptions
	agents  []Agent
	actions map[*Agent][]Direction // TODO: Maybe introduce an action interface for this
}

func NewSimulation(world *IWorld, agents []Agent, options SimulationOptions) *Simulation {
	return &Simulation{
		world:   world,
		agents:  agents,
		options: options,
		actions: make(map[*Agent][]Direction),
	}
}

func (s *Simulation) Run(quit chan bool) <-chan time.Time {
	output := make(chan time.Time)
	if s.options.tickDuration == 0 {
		defer close(output)
		return output
	}

	ticker := time.NewTicker(s.options.tickDuration)

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

				finished := true
				for agent, actions := range s.actions {
					if len(actions) > 0 {
						dir := actions[0]
						agent.MoveInWorld(s.world, dir)
						s.actions[agent] = actions[1:]
						finished = finished && len(s.actions[agent]) == 0
					}
				}

				output <- t
				if finished {
					return
				}
			}
		}
	}()

	return output
}

func (s *Simulation) SetActions(agent *Agent, directions []Direction) {
	s.actions[agent] = append(s.actions[agent], directions...)
}
