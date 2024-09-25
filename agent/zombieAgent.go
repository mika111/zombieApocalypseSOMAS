package extendedAgents

import (
	"fmt"
	"zombieApocalypeSOMAS/physicsEngine"

	"github.com/MattSScott/basePlatformSOMAS/v2/pkg/agent"
)

type IApocalypseEntity interface {
	agent.IAgent[IApocalypseEntity]
	physicsEngine.IPhysicsObject
	PrintPhysicalState()
}

type SpeciesSpecificInfo interface {
	GetSpecies() string
}

type IZombie interface {
	IApocalypseEntity
}

type IHuman interface {
	IApocalypseEntity
}

type ApocalypseEntity struct {
	*agent.BaseAgent[IApocalypseEntity]
	*physicsEngine.PhysicalState
	SpeciesSpecificInfo
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
	entity.SpeciesSpecificInfo = human
	return human
}

func SpawnNewZombie(mass float32, initialPosition physicsEngine.Vector2D, serv agent.IExposedServerFunctions[IApocalypseEntity]) *Zombie {
	entity := &ApocalypseEntity{
		BaseAgent:     agent.CreateBaseAgent(serv),
		PhysicalState: physicsEngine.CreateInitialPhysicalState(&initialPosition, mass),
	}

	zombie := &Zombie{ApocalypseEntity: entity}
	entity.SpeciesSpecificInfo = zombie
	return zombie
}

func (entity *ApocalypseEntity) PrintPhysicalState() {
	state := entity.PhysicalState
	fmt.Printf("%v %v. Position = ", entity.GetSpecies(), entity.GetID())
	state.Position.Print()
	fmt.Printf(". Velocity = ")
	state.Velocity.Print()
	fmt.Printf(". Mass = %v\n", state.Mass)
}

func (human *Human) GetSpecies() string {
	return "human"
}

func (zombie *Zombie) GetSpecies() string {
	return "zombie"
}
