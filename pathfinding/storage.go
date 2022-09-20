package pathfinding

import (
	"fmt"
	"simulator/core/location"
	"simulator/core/world"

	mapset "github.com/deckarep/golang-set/v2"
	prque "github.com/ethereum/go-ethereum/common/prque"
)

func FindStorageLocation(
	w world.IWorld,
	start location.Location) (*location.Location, error) {
	visited := mapset.NewSet[location.Location]()
	visited.Add(start)

	queue := prque.New(nil)
	queue.Push(SearchCell{location: start}, 0)

	for _, amount := range []int{1, 2} {
		for !queue.Empty() {
			current, _ := queue.Pop()
			cell := current.(SearchCell)

			if len(w.GetObjectsAtLocation(cell.location)) == 0 &&
				len(w.GetNeighbors(cell.location)) >= amount {
				return &cell.location, nil
			}

			for _, neighbor := range w.GetNeighbors(cell.location) {
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
	}

	return nil, fmt.Errorf("unable to find storage location")
}
