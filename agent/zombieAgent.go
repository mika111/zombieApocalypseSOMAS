package extendedAgents

import (
	"fmt"
	"math/rand/v2"
)

type IZombie interface {
	IApocalypseEntity
}

type Zombie struct {
	*ApocalypseEntity
	Strength         float32
	RandNumGenerator *rand.Rand
}

func (z *Zombie) PrintPhysicalState() {
	fmt.Printf("Zombie %v. Position = ", z.GetID())
	z.PhysicalState.PrintPhysicalState()
}

func (z *Zombie) GetSpecies() Species {
	return ZomboSapien
}
