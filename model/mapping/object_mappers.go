package mapping

import (
	"simulator/core/agent"
	"simulator/core/location"
	"simulator/core/objects"
	"simulator/model/dto"
)

func LocationToDto(l location.Location) *dto.Location {
	return &dto.Location{
		X: uint32(l.X),
		Y: uint32(l.Y),
	}
}

func AgentToDto(a *agent.Agent) *dto.Agent {
	return &dto.Agent{
		Id:       a.GetId(),
		Location: LocationToDto(a.GetLocation()),
		Callsign: uint32(a.GetRune()) - uint32('0'),
	}
}

func BoxToDto(b *objects.Box) *dto.Box {
	return &dto.Box{
		Id:       b.GetId(),
		Location: LocationToDto(b.GetLocation()),
		Type:     uint32(b.GetType()) - uint32('a'),
	}
}

func GoalToDto(g *objects.Goal) *dto.Goal {
	return &dto.Goal{
		Location: LocationToDto(g.GetLocation()),
		Type:     uint32(g.GetType()) - uint32('a'),
	}
}
