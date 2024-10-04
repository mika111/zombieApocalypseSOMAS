package main

import (
	"fmt"
	"time"
	extendedAgents "zombieApocalypeSOMAS/agent"
	"zombieApocalypeSOMAS/apocalypseServer"
	pathfinding "zombieApocalypeSOMAS/pathFinding"
	"zombieApocalypeSOMAS/physicsEngine"
)

func main() {
	serv := apocalypseServer.CreateApocalypseServer(1, 1, time.Millisecond, 100, 49, 49, 1257)
	serv.ConnectToFrontEnd("localhost:8080")
	serv.GenerateMaze(0, 0, 48, 48)
	serv.ExportInitialState()
	zombie, human := spawnAgentsDemo(serv)
	agentPos := human.PhysicalState.Position
	zombiePos := zombie.PhysicalState.Position
	fmt.Println(zombiePos, agentPos)
findAgent:
	for {
		fmt.Println(zombiePos, agentPos)

		if *zombiePos == *agentPos {
			fmt.Println("breaking")
			break findAgent
		}
		solnPath := pathfinding.FindPath(zombiePos.X, zombiePos.Y, agentPos.X, agentPos.Y, serv.Maze)
		// if len(solnPath) == 0 {
		// 	fmt.Print("breaking")
		// 	break findAgent
		// }
		vec2 := physicsEngine.MakeVec2D(solnPath[len(solnPath)-1][0], solnPath[len(solnPath)-1][1])
		zombie.PhysicalState.Position = &vec2
		zombiePos = zombie.PhysicalState.Position
		serv.ExportState()
		//time.Sleep(time.Millisecond)
	}
	// for i := 0; i < 100; i++ {

	// 	serv.ExportState()
	// }
}

func spawnAgentsDemo(serv *apocalypseServer.ApocalypseServer) (extendedAgents.Zombie, extendedAgents.Human) {
	domain := serv.MapSpawnableArea(extendedAgents.ZomboSapien)
	zombie := serv.SpawnNewZombie(10.0, serv.GenerateRandomPosition(domain))
	serv.AddAgent(zombie)
	domain = serv.MapSpawnableArea(extendedAgents.HomoSapien)
	human := serv.SpawnNewHuman(10.0, serv.GenerateRandomPosition(domain))
	serv.AddAgent(human)
	return *zombie, *human
}
