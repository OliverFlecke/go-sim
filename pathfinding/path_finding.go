package pathfinding

import (
	"fmt"
	"simulator/core/location"
	"simulator/core/objects"
	"simulator/core/utils"
	"simulator/core/world"

	mapset "github.com/deckarep/golang-set/v2"
	prque "github.com/ethereum/go-ethereum/common/prque"
)

type SearchCell struct {
	location location.Location
	depth    int64
	previous *SearchCell
}

type SearchStats struct {
	visited int
}

/*
Find a location that satisfies the given predicate.
Uses BFS to find the nearest location.
*/
func FindClosestObject(
	world world.IWorld,
	start location.Location,
	predicate func(location.Location) objects.WorldObject) (objects.WorldObject, error) {

	visited := mapset.NewSet[location.Location]()
	visited.Add(start)
	queue := prque.New(nil)
	queue.Push(SearchCell{location: start}, 0)

	for !queue.Empty() {
		current, _ := queue.Pop()
		cell := current.(SearchCell)

		obj := predicate(cell.location)
		if obj != nil {
			return obj, nil
		}

		for _, neighbor := range world.GetNeighbors(cell.location) {
			if !visited.Contains(neighbor) {
				visited.Add(neighbor)
				queue.Push(
					SearchCell{
						location: neighbor,
						depth:    cell.depth + 1,
						previous: &cell,
					},
					-(cell.depth + 1))
			}
		}
	}

	return nil, fmt.Errorf("no location found satisfying predicate")
}

func FindPath(
	world world.IWorld,
	start location.Location,
	goal location.Location,
	heuristic heuristic,
	filter filter) ([]location.Location, SearchStats, error) {

	visited := mapset.NewSet[location.Location]()
	visited.Add(start)
	queue := prque.New(nil)
	queue.Push(SearchCell{location: start}, 0)

	var result []location.Location
	var err error

	for {
		if queue.Empty() {
			// fmt.Printf("States visited %v\n", visited)
			err = fmt.Errorf("no path found between from %v to %v", start, goal)
			break
		}

		current, _ := queue.Pop()
		cell := current.(SearchCell)

		if cell.location == goal {
			result = cell.getLocations()
			break
		}

		neighbors := utils.Filteri(world.GetNeighbors(cell.location),
			func(_ int, l location.Location) bool {
				if filter == nil {
					return true
				}

				return filter(l, world)
			})

		for _, neighbor := range neighbors {
			if !visited.Contains(neighbor) {
				visited.Add(neighbor)
				queue.Push(SearchCell{
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

func (item *SearchCell) getLocations() []location.Location {
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
