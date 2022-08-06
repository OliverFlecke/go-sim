package simulator

import (
	dir "simulator/core/direction"
	"simulator/core/location"
	"strings"
)

type IWorld interface {
	GetLocation(location.Location) GridType
	GetNeighbors(location.Location) []location.Location
	ToStringWithAgents([]Agent) string
}

type Grid map[location.Location]GridType

type World struct {
	grid Grid
}

func NewWorld(grid Grid) *World {
	return &World{
		grid: grid,
	}
}

func NewGridWorld(size int) *World {
	grid := make(Grid)

	for y := 0; y <= size+1; y++ {
		grid[location.NewLocation(0, y)] = WALL
		grid[location.NewLocation(size+1, y)] = WALL

	}

	for x := 1; x <= size; x++ {
		grid[location.NewLocation(x, 0)] = WALL
		for y := 1; y <= size; y++ {
			grid[location.NewLocation(x, y)] = EMPTY
		}
		grid[location.NewLocation(x, size+1)] = WALL
	}

	return &World{
		grid: grid,
	}
}

// Getter and Setters
func (w *World) GetMap() Grid {
	return w.grid
}

// Methods

func (world *World) GetLocation(loc location.Location) GridType {
	result, found := world.grid[loc]
	if found {
		return result
	} else {
		return WALL
	}
}

func (world *World) GetNeighbors(loc location.Location) []location.Location {
	neighbors := make([]location.Location, 0)
	directions := []dir.Direction{
		dir.NORTH,
		dir.EAST,
		dir.SOUTH,
		dir.WEST,
	}

	for _, dir := range directions {
		newLocation := loc.MoveInDirection(dir)
		if world.GetLocation(newLocation) == EMPTY {
			neighbors = append(neighbors, newLocation)
		}
	}

	return neighbors
}

// Stringify

func (w *World) ToString() string {
	return w.toStringHelper(func(l location.Location) rune {
		return ToRune(w.GetLocation(l))
	})
}

func (w *World) ToStringWithAgents(agents []Agent) string {
	lookup := make(map[location.Location]Agent)
	for _, agent := range agents {
		lookup[agent.location] = agent
	}

	return w.toStringHelper(func(l location.Location) rune {
		agent, found := lookup[l]
		if found {
			return agent.callsign
		} else {
			return ToRune(w.GetLocation(l))
		}
	})
}

func (w *World) ToStringWithPath(path []location.Location) string {
	lookup := make(map[location.Location]rune)
	for _, location := range path {
		lookup[location] = 'x'
	}

	return w.toStringHelper(func(l location.Location) rune {
		r, found := lookup[l]
		if found {
			return r
		} else {
			return ToRune(w.GetLocation(l))
		}
	})
}

func (w *World) toStringHelper(toRune func(location.Location) rune) string {
	var str strings.Builder
	corner := w.lowerRightCorner()

	for y := 0; y <= corner.Y; y++ {
		for x := 0; x <= corner.X; x++ {
			str.WriteRune(toRune(location.NewLocation(x, y)))
		}
		str.WriteString("\n")
	}

	return str.String()[:str.Len()-1]
}

// Private methods

func (w *World) lowerRightCorner() location.Location {
	result := location.Location{}
	for key := range w.grid {
		if result.X < key.X {
			result.X = key.X
		}
		if result.Y < key.Y {
			result.Y = key.Y
		}
	}

	return result
}
