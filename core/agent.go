package simulator

import (
	direction "simulator/core/direction"
)

type Agent struct {
	name     string
	location Location
}

func NewAgent(name string) *Agent {
	return &Agent{
		name:     name,
		location: Location{},
	}
}

func (agent *Agent) Move(dir direction.Direction) {
	switch dir {
	case direction.NORTH:
		agent.location.y += 1
	case direction.SOUTH:
		agent.location.y -= 1
	case direction.EAST:
		agent.location.x += 1
	case direction.WEST:
		agent.location.x -= 1
	}
}
