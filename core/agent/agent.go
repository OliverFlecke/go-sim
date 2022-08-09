package agent

import (
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

// IMPL: WorldObject interface
func (a Agent) GetLocation() location.Location {
	return a.location
}

func (a *Agent) SetLocation(l location.Location) {
	a.location = l
}

func (a *Agent) GetRune() rune {
	return a.callsign
}
