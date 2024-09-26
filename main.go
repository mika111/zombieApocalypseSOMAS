package main

import (
	"fmt"
	"time"
	extendedAgents "zombieApocalypeSOMAS/agent"
	apocalypseServer "zombieApocalypeSOMAS/environments"
	"zombieApocalypeSOMAS/physicsEngine"
)

func main() {
	ApocalypseServer := apocalypseServer.CreateApocalypseServer(1, 1, 1, 1, time.Millisecond, 100)
	fmt.Printf("Number of Zombies: %v. Number of Humans: %v\n", ApocalypseServer.GetNumEntity(extendedAgents.ZomboSapien), ApocalypseServer.GetNumEntity(extendedAgents.HomoSapien))
	for i := 0; i < 5; i++ {
		for _, ag := range ApocalypseServer.GetAgentMap() {
			fmt.Println("Initial physical state")
			ag.PrintPhysicalState()
			ag.UpdatePhysicalState(physicsEngine.Vector2D{X: 100, Y: 100})
			fmt.Println("New physical state")
			ag.PrintPhysicalState()
		}
	}
	zombLocations := ApocalypseServer.GetEntityLocations(extendedAgents.ZomboSapien)
	humanLocations := ApocalypseServer.GetEntityLocations(extendedAgents.HomoSapien)
	fmt.Printf("Human Locations: %v. Zombie Locations: %v\n", humanLocations, zombLocations)
}
