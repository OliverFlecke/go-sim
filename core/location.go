package simulator

import (
	"math"
	"simulator/core/direction"
)

type Location struct {
	x int
	y int
}

func NewLocation(x int, y int) Location {
	return Location{x: x, y: y}
}

func (loc Location) MoveInDirection(dir direction.Direction) Location {
	result := loc
	switch dir {
	case direction.NORTH:
		result.y += 1
	case direction.SOUTH:
		result.y -= 1
	case direction.EAST:
		result.x += 1
	case direction.WEST:
		result.x -= 1
	}

	return result
}

func (loc Location) ManhattanDistance(other Location) int {
	return Abs(loc.x-other.x) + Abs(loc.y-other.y)
}

func (loc Location) DirectDistance(other Location) float64 {
	return math.Sqrt(math.Pow(float64(loc.x), 2) + math.Pow(float64(loc.y), 2))
}

func Subtract(a, b Location) Location {
	return NewLocation(a.x-b.x, a.y-b.y)
}

func PathToDirections(locations []Location) []direction.Direction {
	result := make([]direction.Direction, 0)
	if len(locations) == 0 {
		return result
	}

	lookup := make(map[Location]direction.Direction)
	lookup[Location{x: 0, y: 1}] = direction.NORTH
	lookup[Location{x: 0, y: -1}] = direction.SOUTH
	lookup[Location{x: 1, y: 0}] = direction.EAST
	lookup[Location{x: -1, y: 0}] = direction.WEST

	for i := 1; i < len(locations); i++ {
		result = append(result, lookup[Subtract(locations[i], locations[i-1])])
	}

	return result
}
