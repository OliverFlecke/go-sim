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

func (agent *Agent) Move(dir Direction) {
	agent.location = agent.location.MoveInDirection(dir)
}

func (agent *Agent) IsValidMove(dir Direction) {

}
