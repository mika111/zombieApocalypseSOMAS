package agents

import (
	"github.com/MattSScott/basePlatformSOMAS/v2/pkg/agent"
)

type IZombie interface {
	agent.IAgent[IZombie]
}

type Zombie struct {
	*agent.BaseAgent[IZombie]
	XPosition float32
	YPosition float32
}
