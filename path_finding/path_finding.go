package pathfinding

import (
	"container/heap"
	"errors"
	"fmt"
	sim "simulator/core"

	mapset "github.com/deckarep/golang-set/v2"
)

func FindPath(
	world *sim.World,
	start sim.Location,
	end sim.Location) ([]sim.Location, error) {

	visited := mapset.NewSet[sim.Location]()
	visited.Add(start)

	frontier := make(PriorityQueue, 0)
	heap.Push(&frontier, &Item[sim.Location]{
		data: start,
		cost: 0,
	})

	for {
		if frontier.Len() == 0 {
			break
		}

		current := frontier.Pop().(*Item[sim.Location])

		fmt.Printf("Exploring %v. Cost: %d\n",
			current.data, current.cost)

		if current.data == end {
			return current.GetDataList(), nil
		}

		for _, neighbor := range neighbors(world, current.data) {
			if !visited.Contains(neighbor) {
				visited.Add(neighbor)
				heap.Push(&frontier, &Item[sim.Location]{
					previous: current,
					data:     neighbor,
					cost:     -(current.depth + 1),
					depth:    current.depth + 1,
				})
			}
		}
	}

	return nil, errors.New("no path found")
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
