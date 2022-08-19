package simulator

import (
	"fmt"
	"reflect"
	"simulator/core/action"
	"simulator/core/agent"
	"simulator/core/world"
	"time"
)

type Simulation struct {
	world   world.IWorld
	options SimulationOptions
	actions map[*agent.Agent][]action.Action
	ticks   uint64
}

func NewSimulation(world world.IWorld, options SimulationOptions) *Simulation {
	return &Simulation{
		world:   world,
		options: options,
		actions: make(map[*agent.Agent][]action.Action),
	}
}

func (s *Simulation) GetTicks() uint64 {
	return s.ticks
}

func (s *Simulation) SetActions(agent *agent.Agent, actions []action.Action) {
	s.actions[agent] = append(s.actions[agent], actions...)
}

func (s *Simulation) Run(quit chan bool) <-chan SimulationEvent {
	output := make(chan SimulationEvent)
	if s.options.tickDuration == 0 {
		go func() {
			defer close(output)
			for {
				finished, err := s.internalRun()
				output <- SimulationEvent{CurrentTime: time.Now(), Err: err}
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
					finished, err := s.internalRun()
					output <- SimulationEvent{CurrentTime: t, Err: err}

					if finished {
						return
					}
				}
			}
		}()
	}
	return output

}

func (s *Simulation) internalRun() (bool, error) {
	s.ticks++
	finished := true
	for agent, actions := range s.actions {
		if len(actions) > 0 {
			action := actions[0]
			// fmt.Printf("\nAgent %c performing action %v\n", agent.callsign, reflect.TypeOf(action))
			result := action.Perform(agent, s.world)
			if result.Err != nil {
				fmt.Printf("Action failed. %v err: %v\n", reflect.TypeOf(action), result.Err.Error())
				fmt.Println(s.world.ToStringWithObjects())
				return true, result.Err
			}

			s.actions[agent] = actions[1:]
			finished = finished && len(s.actions[agent]) == 0
		}
	}

	return finished, nil
}
