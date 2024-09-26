package apocalypseServer

import (
	"time"
	extendedAgents "zombieApocalypeSOMAS/agent"
	"zombieApocalypeSOMAS/physicsEngine"

	//"github.com/kyroy/kdtree"
	"github.com/MattSScott/basePlatformSOMAS/v2/pkg/server"
	"github.com/google/uuid"
)

type IApocalypseServer interface {
	server.IServer[extendedAgents.IApocalypseEntity]
	GetNumZombies() int
	GetNumHumans() int
}

type ApocalypseServer struct {
	*server.BaseServer[extendedAgents.IApocalypseEntity]
	zombies map[uuid.UUID]extendedAgents.IZombie
}

func CreateApocalypseServer(numberZombies, numberHumans, iterations, turns int, maxDuration time.Duration, maxThreads int) *ApocalypseServer {
	server := &ApocalypseServer{
		BaseServer: server.CreateServer[extendedAgents.IApocalypseEntity](iterations, turns, maxDuration, maxThreads),
		zombies:    make(map[uuid.UUID]extendedAgents.IZombie),
	}
	for i := 0; i < numberZombies; i++ {
		zombie := extendedAgents.SpawnNewZombie(10.0, physicsEngine.Vector2D{X: 0, Y: 0}, server)
		server.AddAgent(zombie)
	}
	for i := 0; i < numberHumans; i++ {
		human := extendedAgents.SpawnNewHuman(10.0, physicsEngine.Vector2D{X: 0, Y: 0}, server)
		server.AddAgent(human)
	}
	return server
}

func (server *ApocalypseServer) GetNumEntity(entity extendedAgents.Species) int {
	ans := 0
	for _, ag := range server.GetAgentMap() {
		if ag.GetSpecies() == entity {
			ans += 1
		}
	}
	return ans
}


func (server *ApocalypseServer) GetEntityLocations(entity extendedAgents.Species) []physicsEngine.Vector2D {
	entityLocations := make([]physicsEngine.Vector2D, 0)
	for _, ag := range server.GetAgentMap() {
		if ag.GetSpecies() == entity {
			entityLocations = append(entityLocations, *ag.GetPhysicalState().Position)
		}
	}
	return entityLocations
}
