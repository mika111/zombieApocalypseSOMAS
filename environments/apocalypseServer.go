package apocalypseServer

import (
	"time"
	extendedAgents "zombieApocalypeSOMAS/agent"
	"zombieApocalypeSOMAS/physicsEngine"

	"github.com/MattSScott/basePlatformSOMAS/v2/pkg/server"
)

type IApocalypseServer interface {
	server.IServer[extendedAgents.IApocalypseEntity]
}

type ApocalypseServer struct {
	*server.BaseServer[extendedAgents.IApocalypseEntity]
}

func CreateApocalypseServer(numZombies, numHumans, iterations, turns int, maxDuration time.Duration, maxThreads int) *ApocalypseServer {
	server := &ApocalypseServer{
		BaseServer: server.CreateServer[extendedAgents.IApocalypseEntity](iterations, turns, maxDuration, maxThreads),
	}
	for i := 0; i < numZombies; i++ {
		server.AddAgent(extendedAgents.SpawnNewZombie(10.0, physicsEngine.Vector2D{X: 0, Y: 0}, server))
	}
	for i := 0; i < numHumans; i++ {
		server.AddAgent(extendedAgents.SpawnNewHuman(10.0, physicsEngine.Vector2D{X: 0, Y: 0}, server))
	}

	return server
}
