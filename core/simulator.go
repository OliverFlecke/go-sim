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
	Id      string
	world   world.IWorld
	options SimulationOptions
	actions map[*agent.Agent][]action.Action
	ticks   uint64
	output  chan SimulationEvent
	pause   chan bool
	state   SimulationStatus
}

func NewSimulation(world world.IWorld, options SimulationOptions) *Simulation {
	return &Simulation{
		world:   world,
		options: options,
		actions: make(map[*agent.Agent][]action.Action),
		output:  make(chan SimulationEvent),
		pause:   make(chan bool),
		state:   NONE,
	}
}

// Getters

func (s *Simulation) GetWorld() world.IWorld {
	return s.world
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

func (s *Simulation) GetActions(a *agent.Agent) []action.Action {
	return s.actions[a]
}

func (s *Simulation) Pause() {
	s.state = PAUSED
	s.pause <- true
}

func (s *Simulation) Run() <-chan SimulationEvent {
	if s.state == RUNNING {
		return s.output
	}

	s.state = RUNNING

	if s.options.TickDuration == 0 {
		go func() {
			defer close(s.output)
			for {
				noMoreActions, err := s.internalRun()
				s.updateOutput(time.Now(), err)

				if s.state != COMPLETED && noMoreActions {
					s.state = PAUSED
					return
				}
			}
		}()
	} else {
		ticker := time.NewTicker(s.options.TickDuration)

		go func() {
			defer ticker.Stop()

			for {
				select {
				case <-s.pause:
					fmt.Println("Stopping simulation")
					return
				case t := <-ticker.C:
					_, err := s.internalRun()
					s.updateOutput(t, err)
				}
			}
		}()
	}
	return s.output
}

func (s *Simulation) updateOutput(t time.Time, err error) {
	s.output <- SimulationEvent{CurrentTime: t, Err: err, Status: RUNNING}

	if s.world.IsSolved() {
		s.state = COMPLETED
		s.output <- SimulationEvent{CurrentTime: t, Err: err, Status: COMPLETED}
	}
}

func (s *Simulation) internalRun() (bool, error) {
	s.ticks++
	noMoreActions := true
	for agent, actions := range s.actions {
		if len(actions) > 0 {
			action := actions[0]
			// fmt.Printf("Agent %c performing action %v. Remaining: %d\n", agent.GetRune(), action.ToString(), len(actions)-1)
			result := action.Perform(agent, s.world)
			s.actions[agent] = actions[1:]
			if result.Err != nil {
				logger.Error("Action failed. %v\nerr: %v\n", action.ToString(), result.Err.Error())
				fmt.Println(s.world.ToStringWithObjects())
				return true, result.Err
			}

			noMoreActions = noMoreActions && len(s.actions[agent]) == 0
		}
	}

	return noMoreActions, nil
}
