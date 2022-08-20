package agent

import (
	"simulator/core/location"
)

type Agent struct {
	Location location.Location `json:"location"`
	Callsign rune              `json:"callsign"`
	name     string
}

// Simple constructor
func NewAgent(name string, callsign rune) *Agent {
	return &Agent{
		Callsign: callsign,
		name:     name,
		Location: location.Location{},
	}
}

func NewAgentWithStartLocation(name string, callsign rune, start location.Location) *Agent {
	return &Agent{
		Callsign: callsign,
		name:     name,
		Location: start,
	}
}

// IMPL: WorldObject interface
func (a Agent) GetLocation() location.Location {
	return a.Location
}

func (a *Agent) SetLocation(l location.Location) {
	a.Location = l
}

func (a *Agent) GetRune() rune {
	return a.Callsign
}
