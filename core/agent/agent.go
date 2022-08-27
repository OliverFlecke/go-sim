package agent

import (
	"simulator/core/location"
)

type Agent struct {
	id       uint32
	Location location.Location `json:"location"`
	Callsign rune              `json:"callsign"`
}

var agentCount uint32 = 0

func NewAgentWithStartLocation(callsign rune, start location.Location) *Agent {
	id := agentCount
	agentCount += 1

	return NewAgent(id, callsign, start)
}
func NewAgent(id uint32, callsign rune, start location.Location) *Agent {
	return &Agent{
		id:       id,
		Callsign: callsign,
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
