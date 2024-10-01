package main

import (
	"time"
	"zombieApocalypeSOMAS/apocalypseServer"
	"zombieApocalypeSOMAS/physicsEngine"
)

func main() {
	ApocalypseServer := apocalypseServer.CreateApocalypseServer(10, 10, 1, 1, time.Millisecond, 100, 100, 100, 8570573485792)

	for i := 0; i < 5; i++ {
		for _, ag := range ApocalypseServer.GetAgentMap() {

			ag.UpdatePhysicalState(physicsEngine.Vector2D{X: 100, Y: 100})

		}
	}
	pointA := physicsEngine.MakeVec2D(ApocalypseServer.MapSize.X-10, ApocalypseServer.MapSize.Y-10)
	//pointB := physicsEngine.MakeVec2D(10, 0)
	ApocalypseServer.AddExit(pointA, pointA)
	ApocalypseServer.CreateMaze(34769876598)
	ApocalypseServer.ExportState("state.json")
}
