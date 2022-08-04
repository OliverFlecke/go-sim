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

func NewGridWorld(size int) *World {
	grid := make(map[Location]GridType)

	for x := 0; x < size; x++ {
		for y := 0; y < size; y++ {
			grid[Location{x: x, y: y}] = EMPTY
		}
	}

	return &World{
		grid: grid,
	}
}

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
			return rune('0' + agent.id%10)
		} else {
			return ToRune(w.GetLocation(l))
		}
	})
}

func (w *World) toStringHelper(toRune func(Location) rune) string {
	var str strings.Builder
	corner := w.lowerRightCorner()

	writeFullLineWall(&str, corner.x+3)
	str.WriteRune('\n')
	for y := 0; y <= corner.y; y++ {
		str.WriteRune('#')
		for x := 0; x <= corner.x; x++ {
			str.WriteRune(toRune(Location{x: x, y: y}))
		}
		str.WriteString("#\n")
	}
	writeFullLineWall(&str, corner.x+3)

	return str.String()
}

func writeFullLineWall(str *strings.Builder, size int) {
	for x := 0; x < size; x++ {
		str.WriteRune('#')
	}
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
