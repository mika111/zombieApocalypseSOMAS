package main

import (
	"fmt"
	"time"
	apocalypseServer "zombieApocalypeSOMAS/environments"
	"zombieApocalypeSOMAS/physicsEngine"
)

func main() {
	ApocalypseServer := apocalypseServer.CreateApocalypseServer(1, 10, 1, 1, time.Millisecond, 100)
	fmt.Printf("Number of Zombies: %v. Number of Humans: %v\n", ApocalypseServer.GetNumZombies(), ApocalypseServer.GetNumHumans())
	for _, ag := range ApocalypseServer.GetAgentMap() {
		fmt.Println("Initial physical state")
		ag.PrintPhysicalState()
		ag.UpdatePhysicalState(physicsEngine.Vector2D{X: 100, Y: 100})
		fmt.Println("New physical state")
		ag.PrintPhysicalState()
	}
	zombLocations := ApocalypseServer.GetZombieLocations()
	humanLocations := ApocalypseServer.GetHumanLocations()
	fmt.Printf("Human Locations: %v. Zombie Locations: %v\n", humanLocations, zombLocations)
}
