package simulator

type Agent struct {
	id       int
	name     string
	location Location
}

var amountOfAgents int = 0

func NewAgent(name string) *Agent {
	defer func() {
		amountOfAgents++
	}()

	return &Agent{
		id:       amountOfAgents,
		name:     name,
		location: Location{},
	}
}

func (agent *Agent) Move(dir Direction) {
	agent.location = agent.location.MoveInDirection(dir)
}

func (agent *Agent) IsValidMove(dir Direction) {

}
