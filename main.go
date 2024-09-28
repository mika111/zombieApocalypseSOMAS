package main

import (
	"time"
	apocalypseServer "zombieApocalypeSOMAS/environments"
	"zombieApocalypeSOMAS/physicsEngine"
)

func CreateEnvironment(server *apocalypseServer.ApocalypseServer) {
	server.AddWall(physicsEngine.MakeVec2D(10.2, 40.2), physicsEngine.MakeVec2D(15.2, 10.2))
}

func CreateExits(server *apocalypseServer.ApocalypseServer) {
	server.AddExit(physicsEngine.MakeVec2D(0, 0), physicsEngine.MakeVec2D(0, 100))
	server.AddExit(physicsEngine.MakeVec2D(100, 0), physicsEngine.MakeVec2D(110, 0))
}
func main() {
	ApocalypseServer := apocalypseServer.CreateApocalypseServer(100, 100, 1, 1, time.Millisecond, 100, 700, 500)
	//fmt.Printf("Number of Zombies: %v. Number of Humans: %v\n", ApocalypseServer.GetNumEntity(extendedAgents.ZomboSapien), ApocalypseServer.GetNumEntity(extendedAgents.HomoSapien))
	CreateExits(ApocalypseServer)
	CreateEnvironment(ApocalypseServer)
	for i := 0; i < 5; i++ {
		for _, ag := range ApocalypseServer.GetAgentMap() {
			//fmt.Println("Initial physical state")
			//ag.PrintPhysicalState()
			ag.UpdatePhysicalState(physicsEngine.Vector2D{X: 100, Y: 100})
			//fmt.Println("New physical state")
			//ag.PrintPhysicalState()
		}
	}
	// zombLocations := ApocalypseServer.GetEntityLocations(extendedAgents.ZomboSapien)
	// humanLocations := ApocalypseServer.GetEntityLocations(extendedAgents.HomoSapien)
	// fmt.Printf("Human Locations: %v. Zombie Locations: %v\n", humanLocations, zombLocations)
	ApocalypseServer.ExportState("state.json")
}
