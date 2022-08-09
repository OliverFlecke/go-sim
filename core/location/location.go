package location

import (
	"math"
	"simulator/core/direction"
	"simulator/core/utils"
)

type Location struct {
	X int
	Y int
}

// CONSTRUCTOR
func New(x int, y int) Location {
	return Location{X: x, Y: y}
}

func (loc Location) MoveInDirection(dir direction.Direction) Location {
	result := loc
	switch dir {
	case direction.NORTH:
		result.Y += 1
	case direction.SOUTH:
		result.Y -= 1
	case direction.EAST:
		result.X += 1
	case direction.WEST:
		result.X -= 1
	}

	return result
}

func (loc Location) ManhattanDistance(other Location) int {
	return utils.Abs(loc.X-other.X) + utils.Abs(loc.Y-other.Y)
}

func (loc Location) DirectDistance(other Location) float64 {
	return math.Sqrt(
		math.Pow(float64(other.X-loc.X), 2) +
			math.Pow(float64(other.Y-loc.Y), 2))
}

func Subtract(a, b Location) Location {
	return New(a.X-b.X, a.Y-b.Y)
}

func PathToDirections(locations []Location) []direction.Direction {
	result := make([]direction.Direction, 0)
	if len(locations) == 0 {
		return result
	}

	lookup := make(map[Location]direction.Direction)
	lookup[Location{X: 0, Y: 1}] = direction.NORTH
	lookup[Location{X: 0, Y: -1}] = direction.SOUTH
	lookup[Location{X: 1, Y: 0}] = direction.EAST
	lookup[Location{X: -1, Y: 0}] = direction.WEST

	for i := 1; i < len(locations); i++ {
		result = append(result, lookup[Subtract(locations[i], locations[i-1])])
	}

	return result
}
