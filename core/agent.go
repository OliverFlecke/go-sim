package simulator

type Agent struct {
	callsign rune
	name     string
	location Location
}

// Simple constructor
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

// Implements WorldObject interface
func (a *Agent) GetLocation() Location {
	return a.location
}

// Movement functions
func (agent *Agent) move(dir Direction) {
	agent.location = agent.location.MoveInDirection(dir)
}

func (agent *Agent) MoveInWorld(world IWorld, dir Direction) bool {
	if agent.IsValidMove(world, dir) {
		agent.move(dir)
		return true
	} else {
		return false
	}
}

func (agent *Agent) IsValidMove(world IWorld, dir Direction) bool {
	newLocation := agent.location.MoveInDirection(dir)
	return world.GetLocation(newLocation) == EMPTY
}
