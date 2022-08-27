package agent

import (
	"simulator/core/location"
	"simulator/core/utils"
	"testing"
)

// func TestMove(t *testing.T) {
// 	agent := NewAgent("test agent", 0)
// 	directions := []direction.Direction{
// 		direction.NORTH,
// 		direction.EAST,
// 		direction.SOUTH,
// 		direction.WEST,
// 	}
// 	locations := []location.Location{
// 		location.New(0, 1),
// 		location.New(1, 1),
// 		location.New(1, 0),
// 		location.New(0, 0),
// 	}

// 	for i, dir := range directions {
// 		agent.move(dir)
// 		if agent.location != locations[i] {
// 			fmt.Println(`Location is`, agent.location)
// 			t.Fatal(`Wrong location for agent`)
// 		}
// 	}
// }

func TestNewAgent(t *testing.T) {
	a := NewAgent(0, 'A', location.New(0, 0))
	b := NewAgent(1, 'B', location.New(0, 0))
	utils.AssertEqual(t, a.Callsign, 'A')
	utils.AssertEqual(t, b.Callsign, 'B')
}

// func TestIsValidMove(t *testing.T) {
// 	a := NewAgent("Test agent", 0)
// 	w := NewGridWorld(3)

// 	if a.IsValidMove(w, direction.WEST) {
// 		t.Fatal("WEST is not a valid move in this situation")
// 	}
// }
