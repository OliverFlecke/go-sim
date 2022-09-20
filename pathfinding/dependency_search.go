package pathfinding

import (
	"fmt"
	"simulator/core/location"
	"simulator/core/objects"
	"simulator/core/world"

	mapset "github.com/deckarep/golang-set/v2"
	prque "github.com/ethereum/go-ethereum/common/prque"
)

type searchCell struct {
	location location.Location
	depth    int64
	node     *SearchNode
}

type SearchNode struct {
	Children []*SearchNode
	Item     objects.WorldObject
}

func NewSearchNode(item objects.WorldObject) *SearchNode {
	return &SearchNode{
		Children: make([]*SearchNode, 0),
		Item:     item,
	}
}

// TODO: Naming
type dependency_mapper func(world.IWorld, location.Location, *SearchNode) (*SearchNode, bool)

func FindDepencyTree(
	w world.IWorld,
	start location.Location,
	mapper dependency_mapper) SearchNode {
	visited := mapset.NewSet[location.Location]()
	visited.Add(start)

	root := SearchNode{}
	queue := prque.New(nil)
	queue.Push(searchCell{location: start, node: &root}, 0)

	for !queue.Empty() {
		current, _ := queue.Pop()
		cell := current.(searchCell)

		node, finished := mapper(w, cell.location, cell.node)
		if finished {
			break
		}

		for _, neighbor := range w.GetNeighbors(cell.location) {
			if !visited.Contains(neighbor) {
				visited.Add(neighbor)
				queue.Push(
					searchCell{
						location: neighbor,
						depth:    cell.depth + 1,
						node:     node,
					},
					-(cell.depth + 1))
			}
		}
	}

	return root
}

func PrintTree(node *SearchNode, depth uint) {
	fmt.Printf("%*sNode %v children: %d\n", depth*4, "",
		node.Item, len(node.Children))
	for _, child := range node.Children {
		PrintTree(child, depth+1)
	}
}

func SearchNodeToList(root *SearchNode) []objects.WorldObject {
	list := make([]objects.WorldObject, 0)
	var helper func(*SearchNode)
	helper = func(node *SearchNode) {
		for _, child := range node.Children {
			helper(child)
		}
		list = append(list, node.Item)
	}

	helper(root)

	return list
}
