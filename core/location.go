package simulator

import "math"

type Location struct {
	x int
	y int
}

func NewLocation(x int, y int) Location {
	return Location{x: x, y: y}
}

func (loc Location) MoveInDirection(dir Direction) Location {
	result := loc
	switch dir {
	case NORTH:
		result.y += 1
	case SOUTH:
		result.y -= 1
	case EAST:
		result.x += 1
	case WEST:
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
