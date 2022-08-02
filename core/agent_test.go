package simulator

import (
	"fmt"
	"testing"
)

func TestMove(t *testing.T) {
	agent := NewAgent("test agent")
	directions := []Direction{
		NORTH,
		EAST,
		SOUTH,
		WEST,
	}
	locations := []Location{
		{x: 0, y: 1},
		{x: 1, y: 1},
		{x: 1, y: 0},
		{x: 0, y: 0},
	}

	for i, dir := range directions {
		agent.Move(dir)
		if agent.location != locations[i] {
			fmt.Println(`Location is`, agent.location)
			t.Fatal(`Wrong location for agent`)
		}
	}
}
