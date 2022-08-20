package simulator

import (
	"fmt"
	"simulator/core/action"
	"simulator/core/agent"
	"simulator/core/logger"
	"simulator/core/world"
	"time"
)

type Simulation struct {
	world   world.IWorld
	options SimulationOptions
	actions map[*agent.Agent][]action.Action
	ticks   uint64
	output  chan SimulationEvent
}

func NewSimulation(world world.IWorld, options SimulationOptions) *Simulation {
	return &Simulation{
		world:   world,
		options: options,
		actions: make(map[*agent.Agent][]action.Action),
		output:  make(chan SimulationEvent),
	}
}

func (s *Simulation) GetEvents() <-chan SimulationEvent {
	return s.output
}

func (s *Simulation) GetTicks() uint64 {
	return s.ticks
}

func (s *Simulation) SetActions(agent *agent.Agent, actions []action.Action) {
	s.actions[agent] = append(s.actions[agent], actions...)
}

func (s *Simulation) Run(quit chan bool) <-chan SimulationEvent {
	if s.options.tickDuration == 0 {
		go func() {
			defer close(s.output)
			for {
				finished, err := s.internalRun()
				s.output <- SimulationEvent{CurrentTime: time.Now(), Err: err}
				if finished {
					return
				}
			}
		}()
	} else {
		ticker := time.NewTicker(s.options.tickDuration)

		go func() {
			defer ticker.Stop()
			defer close(s.output)

			for {
				select {
				case <-quit:
					fmt.Println("Stopping simulation")
					return
				case t := <-ticker.C:
					_, err := s.internalRun()
					s.output <- SimulationEvent{CurrentTime: t, Err: err}

					// if finished {
					// 	return
					// }
				}
			}
		}()
	}
	return s.output
}

func (s *Simulation) internalRun() (bool, error) {
	s.ticks++
	finished := true
	for agent, actions := range s.actions {
		if len(actions) > 0 {
			action := actions[0]
			fmt.Printf("Agent %c performing action %v\n", agent.GetRune(), action.ToString())
			result := action.Perform(agent, s.world)
			if result.Err != nil {
				logger.Error("Action failed. %v\nerr: %v\n", action.ToString(), result.Err.Error())
				fmt.Println(s.world.ToStringWithObjects())
				return true, result.Err
			}

			s.actions[agent] = actions[1:]
			finished = finished && len(s.actions[agent]) == 0
		}
	}

	return finished, nil
}
