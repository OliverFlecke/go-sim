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

func Subtract(a, b Location) Location {
	return NewLocation(a.x-b.x, a.y-b.y)
}

func PathToDirections(locations []Location) []Direction {
	result := make([]Direction, 0)
	if len(locations) == 0 {
		return result
	}

	lookup := make(map[Location]Direction)
	lookup[Location{x: 0, y: 1}] = NORTH
	lookup[Location{x: 0, y: -1}] = SOUTH
	lookup[Location{x: 1, y: 0}] = EAST
	lookup[Location{x: -1, y: 0}] = WEST

	for i := 1; i < len(locations); i++ {
		result = append(result, lookup[Subtract(locations[i], locations[i-1])])
	}

	return result
}
