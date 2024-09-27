package extendedAgents

import (
	"zombieApocalypeSOMAS/physicsEngine"

	"github.com/MattSScott/basePlatformSOMAS/v2/pkg/agent"
)

type Species int

const (
	HomoSapien = iota
	ZomboSapien
)

type IApocalypseEntity interface {
	agent.IAgent[IApocalypseEntity]
	physicsEngine.IPhysicsObject
	GetSpecies() Species
}

type ApocalypseEntity struct {
	*agent.BaseAgent[IApocalypseEntity]
	*physicsEngine.PhysicalState
}
