package simulator

import (
	"testing"
)

func TestMove(t *testing.T) {
	agent := newAgent("test agent")
	directions := []Direction{NORTH, EAST, SOUTH, WEST}
	locations := []Location{{x: 0, y: 1}, {x: 1, y: 1}, {x: 1, y: 0}, {x: 0, y: 0}}

	for i, dir := range directions {
		agent.move(dir)
		if agent.location != locations[i] {
			t.Fatal(`Wrong location for agent`)
		}
	}
}
