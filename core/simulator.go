package simulator

import (
	"fmt"
	"simulator/core/agent"
	"simulator/core/world"
	"time"
)

type Simulation struct {
	world   world.IWorld
	options SimulationOptions
	actions map[*agent.Agent][]Action
}

func NewSimulation(world world.IWorld, options SimulationOptions) *Simulation {
	return &Simulation{
		world:   world,
		options: options,
		actions: make(map[*agent.Agent][]Action),
	}
}

func (s *Simulation) SetActions(agent *agent.Agent, actions []Action) {
	s.actions[agent] = append(s.actions[agent], actions...)
}

func (s *Simulation) Run(quit chan bool) <-chan time.Time {
	output := make(chan time.Time)
	if s.options.tickDuration == 0 {
		go func() {
			defer close(output)
			for {
				finished := s.internalRun()
				output <- time.Now()
				if finished {
					return
				}
			}
		}()
	} else {
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
					finished := s.internalRun()
					output <- t

					if finished {
						return
					}
				}
			}
		}()
	}
	return output

}

func (s *Simulation) internalRun() bool {
	finished := true
	for agent, actions := range s.actions {
		if len(actions) > 0 {
			action := actions[0]
			// fmt.Printf("\nAgent %c performing action %v\n", agent.callsign, reflect.TypeOf(action))
			action.Perform(agent, &s.world)
			s.actions[agent] = actions[1:]
			finished = finished && len(s.actions[agent]) == 0
		}
	}

	return finished
}
