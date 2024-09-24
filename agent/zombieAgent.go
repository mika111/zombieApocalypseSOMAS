package extendedAgents

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

func SpawnNewZombie(serv agent.IExposedServerFunctions[IZombie]) *Zombie {
	return &Zombie{
		BaseAgent: agent.CreateBaseAgent(serv),
	}
}


