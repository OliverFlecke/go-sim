package world

import (
	"simulator/core/location"
	"simulator/core/objects"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGridWorld(t *testing.T) {
	world := NewGridWorld(10)
	expectedSize := 144

	if len(world.grid) != expectedSize {
		t.Fatalf(`World does not have the right size. Expected %d but got %d`, expectedSize, len(world.grid))
	}
}

func TestToString(t *testing.T) {
	world := NewGridWorld(3)
	expected := "#####\n#   #\n#   #\n#   #\n#####"

	actual := world.GetStaticMapAsString()
	if actual != expected {
		t.Fatalf("World does not look the way it should! Actual '%s'", actual)
	}
}

func TestToStringWithAgents(t *testing.T) {
	world := NewGridWorld(3)
	// agents := []agent.Agent{
	// 	*agent.NewAgentWithStartLocation("Test agent", '0', location.New(2, 2)),
	// }
	expected := "#####\n#   #\n# 0 #\n#   #\n#####"

	actual := world.ToStringWithAgents()

	if actual != expected {
		t.Fatalf(`World with agents does not look as expected '%s'. Actual: '%s'`, expected, actual)
	}
}

func TestWorldImplementsIWord(t *testing.T) {
	var w IWorld = (*World)(NewGridWorld(1)) // Verify that *T implements I.
	if w.GetLocation(location.New(1, 1)) != EMPTY {
		t.Fatal("Incorrect location returned from interface")
	}
}

func TestIsSolved(t *testing.T) {
	objs := make(objects.ObjectMap)
	objs[objects.GOAL] = append(objs[objects.GOAL], objects.NewGoal(location.New(1, 1), 'a'))
	objs[objects.BOX] = append(objs[objects.BOX], objects.NewBox(location.New(1, 1), 'a'))

	w := NewWorld(NewGrid(3), objs)

	assert.True(t, w.IsSolved(), "All goals have a valid box at the same position, and should therefore be solved")
}

func TestIsSolvedEmptyWorld(t *testing.T) {
	assert.True(t, NewGridWorld(3).IsSolved(), "No goals, therefore problem is always solved")
}

func TestIsSolvedFailing(t *testing.T) {
	objs := make(objects.ObjectMap)
	objs[objects.GOAL] = append(objs[objects.GOAL], objects.NewGoal(location.New(1, 1), 'a'))
	objs[objects.BOX] = append(objs[objects.BOX], objects.NewBox(location.New(2, 2), 'a'))

	w := NewWorld(NewGrid(3), objs)

	assert.False(t, w.IsSolved(), "Box is not at the same position as the goal, and problem is therefore not solved")
}

func TestIsSolvedFailingNoBox(t *testing.T) {
	objs := make(objects.ObjectMap)
	objs[objects.GOAL] = append(objs[objects.GOAL], objects.NewGoal(location.New(1, 1), 'a'))

	w := NewWorld(NewGrid(3), objs)

	assert.False(t, w.IsSolved(), "No box to solve problem with")
}
