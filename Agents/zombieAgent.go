package agents

import (
	"github.com/MattSScott/basePlatformSOMAS/v2/pkg/agent"
)

type IZombieAgent interface {
	agent.IAgent[IZombieAgent]
}

type Zombie struct {
	XPosition float32
	YPosition float32
	*agent.BaseAgent[IZombieAgent]
}
