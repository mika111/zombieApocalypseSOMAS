package extendedAgents

import (
	"fmt"
	"zombieApocalypeSOMAS/physicsEngine"

	"github.com/MattSScott/basePlatformSOMAS/v2/pkg/agent"
)

type IApocalypseEntity interface {
	agent.IAgent[IApocalypseEntity]
	physicsEngine.IPhysicsObject
	GetSpecies() Species
	PrintPhysicalState()
}

type Species int

const (
	HomoSapien = iota
	ZomboSapien
)

type IZombie interface {
	IApocalypseEntity
}

type IHuman interface {
	IApocalypseEntity
}

type ApocalypseEntity struct {
	*agent.BaseAgent[IApocalypseEntity]
	*physicsEngine.PhysicalState
}

type Zombie struct {
	*ApocalypseEntity
}

type Human struct {
	*ApocalypseEntity
}

func SpawnNewHuman(mass float32, initialPosition physicsEngine.Vector2D, serv agent.IExposedServerFunctions[IApocalypseEntity]) *Human {
	entity := &ApocalypseEntity{
		BaseAgent:     agent.CreateBaseAgent(serv),
		PhysicalState: physicsEngine.CreateInitialPhysicalState(&initialPosition, mass),
	}

	human := &Human{ApocalypseEntity: entity}
	return human
}

func SpawnNewZombie(mass float32, initialPosition physicsEngine.Vector2D, serv agent.IExposedServerFunctions[IApocalypseEntity]) *Zombie {
	entity := &ApocalypseEntity{
		BaseAgent:     agent.CreateBaseAgent(serv),
		PhysicalState: physicsEngine.CreateInitialPhysicalState(&initialPosition, mass),
	}

	zombie := &Zombie{ApocalypseEntity: entity}
	return zombie
}

func (h *Human) PrintPhysicalState() {
	fmt.Printf("Human %v. Position = ", h.GetID())
	h.PhysicalState.PrintPhysicalState()
}

func (z *Zombie) PrintPhysicalState() {
	fmt.Printf("Zombie %v. Position = ", z.GetID())
	z.PhysicalState.PrintPhysicalState()
}

func (h *Human) GetSpecies() Species {
	return HomoSapien
}

func (z *Zombie) GetSpecies() Species {
	return ZomboSapien
}
