package main

import (
	"time"
	"zombieApocalypeSOMAS/apocalypseServer"
	"zombieApocalypeSOMAS/physicsEngine"
)

func main() {
	serv := apocalypseServer.CreateApocalypseServer(1, 1, time.Millisecond, 100, 101, 101, 54)
	serv.GenerateMaze(0, 0, serv.MapSize.X-1, serv.MapSize.Y-1)
	serv.InjectAgents(10, 10)

	for i := 0; i < 5; i++ {
		for _, ag := range serv.GetAgentMap() {
			ag.UpdatePhysicalState(physicsEngine.Vector2D{X: 100, Y: 100})
		}
	}

	serv.ExportState("state.json")
}
