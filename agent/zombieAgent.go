package agents

import (
	"zombieApocalypeSOMAS/physicsEngine"

	"github.com/MattSScott/basePlatformSOMAS/v2/pkg/agent"
)

type IZombie interface {
	agent.IAgent[IZombie]
	physicsEngine.IPhysicsObject
}

type Zombie struct {
	*agent.BaseAgent[IZombie]
	*physicsEngine.PhysicalState
}
