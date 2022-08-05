package pathfinding

import (
	sim "simulator/core"
)

type Item[T any] struct {
	previous *Item[T]
	data     T
	cost     int
	depth    int
}

func (item *Item[T]) GetDataList() []T {
	result := make([]T, 0)

	for item != nil {
		result = append(result, item.data)
		item = item.previous
	}

	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}

	return result
}

type PriorityQueue []*Item[sim.Location]

func (h PriorityQueue) Len() int {
	return len(h)
}
func (h PriorityQueue) Less(i, j int) bool {
	return h[i].cost < h[j].cost
}
func (h PriorityQueue) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *PriorityQueue) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(*Item[sim.Location]))
}

func (h *PriorityQueue) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *PriorityQueue) Peek() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	return x
}
