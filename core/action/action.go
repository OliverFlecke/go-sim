package action

import (
	"simulator/core/agent"
	"simulator/core/world"
)

type Action interface {
	Perform(*agent.Agent, world.IWorld) ActionResult
}

type ActionResult struct {
	Err error
}

func success() ActionResult {
	return ActionResult{}
}

func failure(err error) ActionResult {
	return ActionResult{
		Err: err,
	}
}
