package extendedAgents

import "fmt"

type IHuman interface {
	IApocalypseEntity
}

type Human struct {
	*ApocalypseEntity
}

func (h *Human) PrintPhysicalState() {
	fmt.Printf("Human %v. Position = ", h.GetID())
	h.PhysicalState.PrintPhysicalState()
}

func (h *Human) GetSpecies() Species {
	return HomoSapien
}
