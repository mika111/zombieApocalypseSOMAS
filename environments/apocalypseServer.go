package apocalypseServer

import (
	"time"
	extendedAgents "zombieApocalypeSOMAS/agent"

	"github.com/MattSScott/basePlatformSOMAS/v2/pkg/server"
)

type IApocalypseServer interface {
	server.IServer[extendedAgents.IZombie]
}

type ApocalypseServer struct {
	*server.BaseServer[extendedAgents.IZombie]
}

func CreateApocalypseServer(numAgents, iterations, turns int, maxDuration time.Duration, maxThreads int) *ApocalypseServer {
	server := &ApocalypseServer{
		BaseServer: server.CreateServer[extendedAgents.IZombie](iterations, turns, maxDuration, maxThreads),
	}
	for i := 0; i < numAgents; i++ {
		server.AddAgent(extendedAgents.SpawnNewZombie(server))
	}
	return server
}
