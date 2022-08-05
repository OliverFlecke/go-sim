package simulator

import (
	"strings"
)

type GridType = byte

const (
	EMPTY GridType = iota
	WALL
)

func ToRune(g GridType) rune {
	switch g {
	case WALL:
		return '#'
	default:
		return ' '
	}
}

type World struct {
	grid map[Location]GridType
}

func NewWorld(grid map[Location]GridType) *World {
	return &World{
		grid: grid,
	}
}

func NewGridWorld(size int) *World {
	grid := make(map[Location]GridType)

	for y := 0; y <= size+1; y++ {
		grid[NewLocation(0, y)] = WALL
		grid[NewLocation(size+1, y)] = WALL

	}

	for x := 1; x <= size; x++ {
		grid[NewLocation(x, 0)] = WALL
		for y := 1; y <= size; y++ {
			grid[Location{x: x, y: y}] = EMPTY
		}
		grid[NewLocation(x, size+1)] = WALL
	}

	return &World{
		grid: grid,
	}
}

// Getter and Setters
func (w *World) GetMap() map[Location]GridType {
	return w.grid
}

// Methods

func (world *World) GetLocation(loc Location) GridType {
	result, found := world.grid[loc]
	if found {
		return result
	} else {
		return WALL
	}
}

func (w *World) ToString() string {
	return w.toStringHelper(func(l Location) rune {
		return ToRune(w.GetLocation(l))
	})
}

func (w *World) ToStringWithAgents(agents []Agent) string {
	lookup := make(map[Location]Agent)
	for _, agent := range agents {
		lookup[agent.location] = agent
	}

	return w.toStringHelper(func(l Location) rune {
		agent, found := lookup[l]
		if found {
			return agent.callsign
		} else {
			return ToRune(w.GetLocation(l))
		}
	})
}

func (w *World) ToStringWithPath(path []Location) string {
	lookup := make(map[Location]rune)
	for _, location := range path {
		lookup[location] = 'x'
	}

	return w.toStringHelper(func(l Location) rune {
		r, found := lookup[l]
		if found {
			return r
		} else {
			return ToRune(w.GetLocation(l))
		}
	})
}

func (w *World) toStringHelper(toRune func(Location) rune) string {
	var str strings.Builder
	corner := w.lowerRightCorner()

	for y := 0; y <= corner.y; y++ {
		for x := 0; x <= corner.x; x++ {
			str.WriteRune(toRune(Location{x: x, y: y}))
		}
		str.WriteString("\n")
	}

	return str.String()[:str.Len()-1]
}

func (w *World) lowerRightCorner() Location {
	result := Location{}
	for key := range w.grid {
		if result.x < key.x {
			result.x = key.x
		}
		if result.y < key.y {
			result.y = key.y
		}
	}

	return result
}
