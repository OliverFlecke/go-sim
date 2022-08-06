package simulator

import (
	"fmt"
	"testing"
)

func TestMove(t *testing.T) {
	agent := NewAgent("test agent", 0)
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
		agent.move(dir)
		if agent.location != locations[i] {
			fmt.Println(`Location is`, agent.location)
			t.Fatal(`Wrong location for agent`)
		}
	}
}

func TestNewAgent(t *testing.T) {
	a := NewAgent("Agent A", 'A')
	b := NewAgent("Agent B", 'B')
	AssertEqual(t, a.callsign, 'A')
	AssertEqual(t, b.callsign, 'B')
}

func TestIsValidMove(t *testing.T) {
	a := NewAgent("Test agent", 0)
	w := (IWorld)(NewGridWorld(3))

	if a.IsValidMove(&w, WEST) {
		t.Fatal("WEST is not a valid move in this situation")
	}
}
