package simulator

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

func (agent *Agent) Move(direction Direction) {
	switch direction {
	case NORTH:
		agent.location.y += 1
	case SOUTH:
		agent.location.y -= 1
	case EAST:
		agent.location.x += 1
	case WEST:
		agent.location.x -= 1
	}
}
