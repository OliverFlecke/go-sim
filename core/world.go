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
		return EMPTY
	}
}

func (w *World) ToString() string {
	var str strings.Builder
	corner := w.lowerRightCorner()

	for y := 0; y <= corner.y; y++ {
		for x := 0; x <= corner.x; x++ {
			str.WriteRune(ToRune(w.GetLocation(Location{x: x, y: y})))
		}
		str.WriteRune('\n')
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
