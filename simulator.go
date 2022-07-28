package simulator

type Direction byte

const (
	NORTH Direction = iota
	EAST
	SOUTH
	WEST
)

type Location struct {
	x int64
	y int64
}

type Agent struct {
	name     string
	location Location
}

func newAgent(name string) *Agent {
	return &Agent{
		name:     name,
		location: Location{},
	}
}

func (agent Agent) move(direction Direction) {
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
