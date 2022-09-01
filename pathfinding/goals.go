package pathfinding

import (
	"fmt"
	"simulator/core/location"
	"simulator/core/objects"
	"simulator/core/world"

	mapset "github.com/deckarep/golang-set/v2"
	prque "github.com/ethereum/go-ethereum/common/prque"
)

type GoalSearchCell struct {
	location location.Location
	depth    int64
	tree     *GoalTree
}

type GoalTree struct {
	children []*GoalTree
	item     *objects.Goal
}

func NewGoalTree(g *objects.Goal) *GoalTree {
	return &GoalTree{
		children: make([]*GoalTree, 0),
		item:     g,
	}
}

func GoalDependencies(
	w world.IWorld,
	start location.Location) GoalTree {

	visited := mapset.NewSet[location.Location]()
	visited.Add(start)

	root := GoalTree{}
	queue := prque.New(nil)
	queue.Push(GoalSearchCell{location: start, tree: &root}, 0)

	for !queue.Empty() {
		current, _ := queue.Pop()
		cell := current.(GoalSearchCell)

		goal := getGoalAtLocation(w, cell.location)

		tree := cell.tree
		if goal != nil && !w.IsGoalSolved(goal) {
			tree = NewGoalTree(goal)
			cell.tree.children = append(cell.tree.children, tree)
		}

		for _, neighbor := range w.GetNeighbors(cell.location) {
			if !visited.Contains(neighbor) {
				visited.Add(neighbor)
				queue.Push(
					GoalSearchCell{
						location: neighbor,
						depth:    cell.depth + 1,
						tree:     tree,
					},
					-(cell.depth + 1))
			}
		}
	}

	return root
}

func getGoalAtLocation(w world.IWorld, l location.Location) *objects.Goal {
	for _, o := range w.GetObjectsAtLocation(l) {
		switch g := o.(type) {
		case *objects.Goal:
			if !w.IsGoalSolved(g) {
				return g
			}
		}
	}

	return nil
}

func printTree(tree *GoalTree, depth uint) {
	fmt.Printf("%*sGoal %v children: %d\n", depth*4, "",
		tree.item, len(tree.children))
	for _, child := range tree.children {
		printTree(child, depth+1)
	}
}

func GoalTreeToList(root *GoalTree) []*objects.Goal {
	output := make([]*objects.Goal, 0)

	var helper func(*GoalTree)
	helper = func(node *GoalTree) {
		for _, child := range node.children {
			helper(child)
		}
		output = append(output, node.item)
	}

	helper(root)

	return output
}
