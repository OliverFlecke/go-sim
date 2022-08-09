package location

import (
	"math"
	"simulator/core/direction"
	"simulator/core/utils"
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
		{X: 0, Y: 1},
		{X: 1, Y: 0},
		{X: 0, Y: -1},
		{X: -1, Y: 0},
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
			a:        New(0, 0),
			b:        New(2, 2),
			expected: 4,
		},
		{
			a:        New(10, 10),
			b:        New(20, 20),
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

type TestDirectDistanceData struct {
	a        Location
	b        Location
	expected float64
}

func TestDirectDistance(t *testing.T) {
	data := []TestDirectDistanceData{
		{
			a:        New(0, 0),
			b:        New(2, 2),
			expected: math.Sqrt(8),
		},
		{
			a:        New(10, 10),
			b:        New(20, 20),
			expected: math.Sqrt(200),
		},
	}

	for _, value := range data {
		dist := value.a.DirectDistance(value.b)
		if dist != value.expected {
			t.Fatalf("Wrong distance. Expected %f, got %f",
				value.expected, dist)
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
			a:        New(0, 0),
			b:        New(1, 1),
			expected: New(-1, -1),
		},
	}

	for _, value := range data {
		result := Subtract(value.a, value.b)
		if result != value.expected {
			t.Fatalf("Wrong value. Expected %d, got %d", value.expected, result)
		}
	}
}

func TestEmptyPathToDirections(t *testing.T) {
	utils.AssertEqualSlices(t,
		PathToDirections(make([]Location, 0)),
		make([]direction.Direction, 0))
}

func TestPathToDirections(t *testing.T) {
	path := []Location{
		New(0, 0),
		New(0, 1),
		New(0, 2),
		New(1, 2),
		New(2, 2),
		New(1, 2),
		New(1, 1),
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

	utils.AssertEqualSlices(t, actual, expected)
}
