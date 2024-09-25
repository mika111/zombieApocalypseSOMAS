package apocalypseServer

import (
	"time"
	extendedAgents "zombieApocalypeSOMAS/agent"
	"zombieApocalypeSOMAS/physicsEngine"

	//"github.com/kyroy/kdtree"
	"github.com/MattSScott/basePlatformSOMAS/v2/pkg/server"
)

type IApocalypseServer interface {
	server.IServer[extendedAgents.IApocalypseEntity]
	GetNumZombies() int
	GetNumHumans() int
}

type ApocalypseServer struct {
	*server.BaseServer[extendedAgents.IApocalypseEntity]
	//numOf<entity> used to preallocate space for return array for Get<entity>Locations
	numZombies int
	numHumans  int
}

func CreateApocalypseServer(numberZombies, numberHumans, iterations, turns int, maxDuration time.Duration, maxThreads int) *ApocalypseServer {
	server := &ApocalypseServer{
		BaseServer: server.CreateServer[extendedAgents.IApocalypseEntity](iterations, turns, maxDuration, maxThreads),
		numZombies: 0,
		numHumans:  0,
	}
	for i := 0; i < numberZombies; i++ {
		server.AddAgent(extendedAgents.SpawnNewZombie(10.0, physicsEngine.Vector2D{X: 0, Y: 0}, server))
		server.numZombies++
	}
	for i := 0; i < numberHumans; i++ {
		server.AddAgent(extendedAgents.SpawnNewHuman(10.0, physicsEngine.Vector2D{X: 0, Y: 0}, server))
		server.numHumans++
	}
	return server
}

func (server *ApocalypseServer) GetNumHumans() int {
	return server.numHumans
}

func (server *ApocalypseServer) GetNumZombies() int {
	return server.numZombies
}

func (server *ApocalypseServer) GetEntityLocations(speciesType extendedAgents.Species) []physicsEngine.Vector2D {
	var numEntities int
	if speciesType == extendedAgents.HomoSapien {
		numEntities = server.numHumans
	} else {
		numEntities = server.numZombies
	}
	entityLocations := make([]physicsEngine.Vector2D, numEntities)
	i := 0
	for _, ag := range server.GetAgentMap() {
		if ag.GetSpecies() == speciesType {
			entityLocations[i] = *ag.GetPhysicalState().Position
			i++
		}
	}
	return entityLocations
}
