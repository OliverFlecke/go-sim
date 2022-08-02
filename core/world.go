package simulator

type GridType = byte

const (
	NONE GridType = iota
	EMPTY
	WALL
)

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
		return NONE
	}
}
