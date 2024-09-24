package agent

import "github.com/MattSScott/basePlatformSOMAS/v2/pkg/agent"

type IZombieAgent interface {
	agent.IAgent[IZombieAgent]
}
