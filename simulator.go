package simulator

func Run() string {
	return "Hello from simulator"
}

func move(agent *Agent, direction Direction) {
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

func newAgent(name string) *Agent {
	return &Agent{
		name:     name,
		location: Location{},
	}
}

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
