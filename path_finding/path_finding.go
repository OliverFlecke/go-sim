package pathfinding

import (
	"fmt"
	sim "simulator/core"
	"simulator/core/location"

	mapset "github.com/deckarep/golang-set/v2"
	prque "github.com/ethereum/go-ethereum/common/prque"
)

type Cell struct {
	location location.Location
	depth    int64
	previous *Cell
}

type SearchStats struct {
	visited int
}

func FindPath(
	world sim.IWorld,
	start location.Location,
	goal location.Location,
	heuristic heuristic) ([]location.Location, SearchStats, error) {

	visited := mapset.NewSet[location.Location]()
	visited.Add(start)
	queue := prque.New(nil)
	queue.Push(Cell{location: start}, 0)

	var result []location.Location
	var err error

	for {
		if queue.Empty() {
			err = fmt.Errorf("no path found between from %v to %v", start, goal)
			break
		}

		current, _ := queue.Pop()
		cell := current.(Cell)

		if cell.location == goal {
			result = cell.getLocations()
			break
		}

		for _, neighbor := range world.GetNeighbors(cell.location) {
			if !visited.Contains(neighbor) {
				visited.Add(neighbor)
				queue.Push(Cell{
					location: neighbor,
					depth:    cell.depth + 1,
					previous: &cell,
				},
					// Heuristic value is inverted to get priority queue to act as a min heap
					-heuristic(cell.depth, cell.location, goal, world))
			}
		}
	}

	return result,
		SearchStats{
			visited: visited.Cardinality(),
		}, err
}

func (item *Cell) getLocations() []location.Location {
	result := make([]location.Location, 0)

	for item != nil {
		result = append(result, item.location)
		item = item.previous
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}
