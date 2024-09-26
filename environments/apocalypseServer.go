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
		zombie := (extendedAgents.SpawnNewZombie(10.0, physicsEngine.Vector2D{X: 0, Y: 0}, server))
		server.zombies[zombie.GetID()] = zombie
	}
	for i := 0; i < numberHumans; i++ {
		server.AddAgent(extendedAgents.SpawnNewHuman(10.0, physicsEngine.Vector2D{X: 0, Y: 0}, server))
	}
	return server
}

func (server *ApocalypseServer) GetNumHumans() int {
	return len(server.GetAgentMap())
}

func (server *ApocalypseServer) GetNumZombies() int {
	return len(server.zombies)
}

func (server *ApocalypseServer) GetZombieLocations() []physicsEngine.Vector2D {
	entityLocations := make([]physicsEngine.Vector2D, server.GetNumZombies())
	i := 0
	for _, zom := range server.zombies {
		entityLocations[i] = *zom.GetPhysicalState().Position
		i++
	}
	return entityLocations
}

func (server *ApocalypseServer) GetHumanLocations() []physicsEngine.Vector2D {
	entityLocations := make([]physicsEngine.Vector2D, server.GetNumHumans())
	i := 0
	for _, ag := range server.GetAgentMap() {
		entityLocations[i] = *ag.GetPhysicalState().Position
		i++
	}
	return entityLocations
}
