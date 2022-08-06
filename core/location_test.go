package simulator

import (
	"simulator/core/direction"
	"testing"
)

func TestMoveInDirection(t *testing.T) {
	directions := []direction.Direction{
		direction.NORTH,
		direction.EAST,
		direction.SOUTH,
		direction.WEST,
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

type TestSubtraction struct {
	a        Location
	b        Location
	expected Location
}

func TestSubstract(t *testing.T) {
	data := []TestSubtraction{
		{
			a:        NewLocation(0, 0),
			b:        NewLocation(1, 1),
			expected: NewLocation(-1, -1),
		},
	}

	for _, value := range data {
		result := Subtract(value.a, value.b)
		if result != value.expected {
			t.Fatalf("Wrong value. Expected %d, got %d", value.expected, result)
		}
	}
}

func TestPathToDirections(t *testing.T) {
	path := []Location{
		NewLocation(0, 0),
		NewLocation(0, 1),
		NewLocation(0, 2),
		NewLocation(1, 2),
		NewLocation(2, 2),
		NewLocation(1, 2),
		NewLocation(1, 1),
	}
	expected := []direction.Direction{
		direction.NORTH,
		direction.NORTH,
		direction.EAST,
		direction.EAST,
		direction.WEST,
		direction.SOUTH,
	}

	actual := PathToDirections(path)

	AssertEqualSlices(t, actual, expected)
}
