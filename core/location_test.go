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

type TestData struct {
	a        Location
	b        Location
	expected int
}

func TestManhattanDistance(t *testing.T) {
	data := []TestData{
		{
			a:        NewLocation(0, 0),
			b:        NewLocation(2, 2),
			expected: 4,
		},
		{
			a:        NewLocation(10, 10),
			b:        NewLocation(20, 20),
			expected: 20,
		},
	}

	for _, value := range data {
		dist := value.a.ManhattanDistance(value.b)
		if dist != value.expected {
			t.Fatalf("Wrong distance. Expected %d, got %d", value.expected, dist)
		}
	}
}
