package simulator

import (
	"simulator/core/direction"
	"simulator/core/location"
)

type Agent struct {
	callsign rune
	name     string
	location location.Location
}

// Simple constructor
func NewAgent(name string, callsign rune) *Agent {
	return &Agent{
		callsign: callsign,
		name:     name,
		location: location.Location{},
	}
}

func NewAgentWithStartLocation(name string, callsign rune, start location.Location) *Agent {
	return &Agent{
		callsign: callsign,
		name:     name,
		location: start,
	}
}

// Implements WorldObject interface
func (a *Agent) GetLocation() location.Location {
	return a.location
}

// Movement functions
func (agent *Agent) move(dir direction.Direction) {
	agent.location = agent.location.MoveInDirection(dir)
}

func (agent *Agent) MoveInWorld(world IWorld, dir direction.Direction) bool {
	if agent.IsValidMove(world, dir) {
		agent.move(dir)
		return true
	} else {
		return false
	}
}

func (agent *Agent) IsValidMove(world IWorld, dir direction.Direction) bool {
	newLocation := agent.location.MoveInDirection(dir)
	return world.GetLocation(newLocation) == EMPTY
}
