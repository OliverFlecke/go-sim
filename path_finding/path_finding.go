package pathfinding

import (
	"errors"
	sim "simulator/core"

	mapset "github.com/deckarep/golang-set/v2"
	prque "github.com/ethereum/go-ethereum/common/prque"
)

type Cell struct {
	location sim.Location
	previous *Cell
}

type SearchStats struct {
	visited int
}

func FindPath(
	world *sim.World,
	start sim.Location,
	end sim.Location) ([]sim.Location, SearchStats, error) {

	visited := mapset.NewSet[sim.Location]()
	visited.Add(start)
	queue := prque.New(nil)
	queue.Push(Cell{location: start}, 0)

	var result []sim.Location
	var err error

	for {
		if queue.Empty() {
			err = errors.New("no path found")
			break
		}

		current, cost := queue.Pop()
		cell := current.(Cell)

		if cell.location == end {
			result = cell.getDataList()
			break
		}

		for _, neighbor := range neighbors(world, cell.location) {
			if !visited.Contains(neighbor) {
				visited.Add(neighbor)
				queue.Push(Cell{
					location: neighbor,
					previous: &cell,
				},
					cost-1)
			}
		}
	}

	return result,
		SearchStats{
			visited: visited.Cardinality(),
		}, err
}

func neighbors(world *sim.World, location sim.Location) []sim.Location {
	neighbors := make([]sim.Location, 0)
	directions := []sim.Direction{
		sim.NORTH,
		sim.EAST,
		sim.SOUTH,
		sim.WEST,
	}

	for _, dir := range directions {
		newLocation := location.MoveInDirection(dir)
		if world.GetLocation(newLocation) == sim.EMPTY {
			neighbors = append(neighbors, newLocation)
		}
	}

	return neighbors
}

func (item *Cell) getDataList() []sim.Location {
	result := make([]sim.Location, 0)

	for item != nil {
		result = append(result, item.location)
		item = item.previous
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}
