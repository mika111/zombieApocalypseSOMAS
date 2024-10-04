package main

import (
	"time"
	extendedAgents "zombieApocalypeSOMAS/agent"
	"zombieApocalypeSOMAS/apocalypseServer"
)

func main() {
	serv := apocalypseServer.CreateApocalypseServer(1, 1, time.Millisecond, 100, 49, 49, 1257)
	serv.ConnectToFrontEnd("localhost:8080")
	serv.GenerateMaze(0, 0, 48, 48)
	serv.ExportInitialState()
	spawnAgentsDemo(serv)
	//serv.ExportState()
	serv.InjectAgents(100, 100)
	for i := 0; i < 100; i++ {
		time.Sleep(1 * time.Millisecond)
		serv.ExportState()
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
