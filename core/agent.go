package simulator

type Agent struct {
	callsign rune
	name     string
	location Location
}

func NewAgent(name string, callsign rune) *Agent {
	return &Agent{
		callsign: callsign,
		name:     name,
		location: Location{},
	}
}

func NewAgentWithStartLocation(name string, callsign rune, start Location) *Agent {
	return &Agent{
		callsign: callsign,
		name:     name,
		location: start,
	}
}

func (a *Agent) GetLocation() Location {
	return a.location
}

func (agent *Agent) move(dir Direction) {
	agent.location = agent.location.MoveInDirection(dir)
}

func (agent *Agent) MoveInWorld(world *World, dir Direction) bool {
	if agent.IsValidMove(world, dir) {
		agent.move(dir)
		return true
	} else {
		return false
	}
}

func (agent *Agent) IsValidMove(world *World, dir Direction) bool {
	newLocation := agent.location.MoveInDirection(dir)
	return world.GetLocation(newLocation) == EMPTY
}
