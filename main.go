package main

import (
	"fmt"
	"time"
	apocalypseServer "zombieApocalypeSOMAS/environments"
	"zombieApocalypeSOMAS/physicsEngine"
)

func main() {
	ApocalypseServer := apocalypseServer.CreateApocalypseServer(10,10, 1, 1, time.Millisecond, 100)
	for _, ag := range ApocalypseServer.GetAgentMap() {
		fmt.Println("Initial physical state")
		ag.PrintPhysicalState()
		ag.UpdatePhysicalState(physicsEngine.Vector2D{X: 100, Y: 100})
		fmt.Println("New physical state")
		ag.PrintPhysicalState()
	}
}
