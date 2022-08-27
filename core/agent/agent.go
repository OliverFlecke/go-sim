package agent

import (
	"simulator/core/location"
)

type Agent struct {
	id       uint32
	Location location.Location `json:"location"`
	Callsign rune              `json:"callsign"`
	name     string
}

var agentCount uint32 = 0

// Simple constructor
func NewAgent(name string, callsign rune) *Agent {
	return NewAgentWithStartLocation(name, callsign, location.New(0, 0))
}

func NewAgentWithStartLocation(name string, callsign rune, start location.Location) *Agent {
	id := agentCount
	agentCount += 1

	return &Agent{
		id:       id,
		Callsign: callsign,
		name:     name,
		Location: start,
	}
}

func (a *Agent) GetId() uint32 {
	return a.id
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
