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

func (agent *Agent) Move(dir Direction) {
	agent.location = agent.location.MoveInDirection(dir)
}

func (agent *Agent) IsValidMove(dir Direction) {

}
