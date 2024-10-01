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

// func (z *Zombie) GenerateRandomForce() physicsEngine.Vector2D {
// 	Xcomponent := z.Strength * float32(z.RandNumGenerator.NormFloat64())
// 	YComponent := z.Strength * float32(z.RandNumGenerator.NormFloat64())
// 	return physicsEngine.Vector2D{X: Xcomponent,
// 		Y: YComponent}
// }
