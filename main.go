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
	human := serv.SpawnNewHuman(10.0, physicsEngine.MakeVec2D(0, 0))
	serv.AddAgent(human)
	agentPos := human.PhysicalState.Position
	targetPos := physicsEngine.MakeVec2D(48, 48)
	fmt.Println(targetPos, agentPos)
findAgent:
	for {
		if targetPos == *agentPos {
			break findAgent
		}
		solnPath := pathfinding.FindPath(agentPos.X, agentPos.Y, targetPos.X, targetPos.Y, serv.Maze)
		vec2 := physicsEngine.MakeVec2D(solnPath[len(solnPath)-1][0], solnPath[len(solnPath)-1][1])
		human.PhysicalState.Position = &vec2
		agentPos = human.PhysicalState.Position
		serv.ExportState()
		//time.Sleep(time.Millisecond)
	}
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
