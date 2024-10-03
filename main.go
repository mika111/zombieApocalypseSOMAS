package main

import (
	"time"
	"zombieApocalypeSOMAS/apocalypseServer"
	"zombieApocalypeSOMAS/physicsEngine"
)

func main() {
	serv := apocalypseServer.CreateApocalypseServer(1, 1, time.Millisecond, 100, 15, 15, 1234)
	serv.GenerateMaze(1, 1, 8, 8)
	serv.InjectAgents(10, 10)

	for i := 0; i < 5; i++ {
		for _, ag := range serv.GetAgentMap() {
			ag.UpdatePhysicalState(physicsEngine.Vector2D{X: 100, Y: 100})
		}
	}

	serv.ExportState("state.json")
}
