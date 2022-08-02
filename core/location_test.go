package simulator

import "testing"

func TestMoveInDirection(t *testing.T) {
	directions := []Direction{
		NORTH,
		EAST,
		SOUTH,
		WEST,
	}
	expected := []Location{
		{x: 0, y: 1},
		{x: 1, y: 0},
		{x: 0, y: -1},
		{x: -1, y: 0},
	}

	for i, dir := range directions {
		actual := Location{}.MoveInDirection(dir)
		if actual != expected[i] {
			t.Fatalf(`Wrong location. Expected %v got %v`, expected[i], actual)
		}
	}
}
