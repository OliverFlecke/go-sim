package simulator

import (
	dir "simulator/core/direction"
	"simulator/core/location"
	"simulator/core/objects"
	"simulator/core/utils"
	"strings"
)

type IWorld interface {
	GetLocation(location.Location) GridType
	GetNeighbors(location.Location) []location.Location
	GetObjects(objects.WorldObjectKey) []objects.WorldObject
	GetObjectsAtLocation(location.Location) []objects.WorldObject
	MoveObject(o objects.WorldObject, newLoc location.Location)

	ToStringWithObjects() string
}

type WorldObjectMap map[location.Location][]objects.WorldObject
type Grid map[location.Location]GridType

type World struct {
	grid      Grid
	objects   objects.ObjectMap
	objectMap WorldObjectMap
}

func NewWorld(grid Grid, xs objects.ObjectMap) *World {
	return &World{
		grid:      grid,
		objects:   xs,
		objectMap: objectsToMap(xs),
	}
}

func NewGridWorld(size int) *World {
	return &World{
		grid: NewGrid(size),
	}
}

func NewGrid(size int) Grid {
	grid := make(Grid)

	for y := 0; y <= size+1; y++ {
		grid[location.New(0, y)] = WALL
		grid[location.New(size+1, y)] = WALL
	}

	for x := 1; x <= size; x++ {
		grid[location.New(x, 0)] = WALL
		for y := 1; y <= size; y++ {
			grid[location.New(x, y)] = EMPTY
		}
		grid[location.New(x, size+1)] = WALL
	}

	return grid
}

// Getter and Setters
func (w *World) GetMap() Grid {
	return w.grid
}

// IMPL: IWorld interface

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

func (w *World) GetObjects(key objects.WorldObjectKey) []objects.WorldObject {
	return w.objects[key]
}

func (w *World) MoveObject(o objects.WorldObject, newLoc location.Location) {
	for i, v := range w.objectMap[o.GetLocation()] {
		if v == o {
			utils.Remove(w.objectMap[o.GetLocation()], i)
		}
	}

	w.objectMap[newLoc] = append(w.objectMap[newLoc], o)
	o.SetLocation(newLoc)
}

func (w *World) GetObjectsAtLocation(loc location.Location) []objects.WorldObject {
	return w.objectMap[loc]
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

func (w *World) ToStringWithObjects() string {
	lookup := make(map[location.Location]objects.WorldObject)
	for _, objs := range w.objects {
		for _, obj := range objs {
			if lookup[obj.GetLocation()] == nil {
				lookup[obj.GetLocation()] = obj
			}
		}
	}

	return w.toStringHelper(func(l location.Location) rune {
		obj, found := lookup[l]
		if found {
			return obj.GetRune()
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
			str.WriteRune(toRune(location.New(x, y)))
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

func objectsToMap(values objects.ObjectMap) WorldObjectMap {
	result := make(WorldObjectMap)
	for _, vs := range values {
		for _, o := range vs {
			result[o.GetLocation()] = append(result[o.GetLocation()], o)
		}
	}

	return result
}
