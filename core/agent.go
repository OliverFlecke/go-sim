package simulator

type Agent struct {
	id       int
	name     string
	location Location
}

func NewAgent(name string, id int) *Agent {
	return &Agent{
		id:       id,
		name:     name,
		location: Location{},
	}
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
