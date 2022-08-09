package action

import (
	"simulator/core/agent"
	"simulator/core/world"
)

type Action interface {
	Perform(*agent.Agent, world.IWorld)
}
