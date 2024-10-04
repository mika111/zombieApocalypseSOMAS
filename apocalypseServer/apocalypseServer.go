package apocalypseServer

import (
	"encoding/json"
	"math/rand/v2"
	"os"
	"time"
	extendedAgents "zombieApocalypeSOMAS/agent"
	"zombieApocalypeSOMAS/mazeGenerator"
	"zombieApocalypeSOMAS/physicsEngine"

	"github.com/MattSScott/basePlatformSOMAS/v2/pkg/server"
)

type IApocalypseServer interface {
	server.IServer[extendedAgents.IApocalypseEntity]
	GetNumZombies() int
	GetNumHumans() int
	MapSpawnableArea() int
}

type ApocalypseServer struct {
	*server.BaseServer[extendedAgents.IApocalypseEntity]
	RandNumGenerator *rand.Rand
	MapSize          physicsEngine.Vector2D
	Maze             mazeGenerator.Maze
}

type gameState struct {
	RoundNum        int
	MapSize         physicsEngine.Vector2D
	ZombiePositions []physicsEngine.Vector2D
	HumanPositions  []physicsEngine.Vector2D
	Maze            mazeGenerator.Maze
	BorderSize      int
}

type ApocalypeSeed int

func (seed ApocalypeSeed) Uint64() uint64 {
	return uint64(seed)
}

func CreateApocalypseServer(iterations, turns int, maxDuration time.Duration, maxThreads int, width, height int, mazeSeed uint64) *ApocalypseServer {
	return &ApocalypseServer{
		BaseServer:       server.CreateServer[extendedAgents.IApocalypseEntity](iterations, turns, maxDuration, maxThreads),
		RandNumGenerator: rand.New(rand.NewPCG(mazeSeed, mazeSeed)),
		MapSize:          physicsEngine.MakeVec2D(width, height),
		Maze:             nil,
	}
}

func (serv *ApocalypseServer) GenerateMaze(entrance_i, entrance_j, exit_i, exit_j int) {
	mazeGen := mazeGenerator.CreateMazeGenerator(serv.MapSize.X, serv.MapSize.Y, serv.RandNumGenerator)
	serv.Maze = mazeGen.CreateMaze(entrance_i, entrance_j, exit_i, exit_j)
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

func (server *ApocalypseServer) EntityLocationMap(entity extendedAgents.Species) map[physicsEngine.Vector2D]struct{} {
	entityLocations := make(map[physicsEngine.Vector2D]struct{})
	for _, ag := range server.GetAgentMap() {
		if ag.GetSpecies() == entity {
			entityLocations[*ag.GetPhysicalState().Position] = struct{}{}
		}
	}
	return entityLocations
}

func (server *ApocalypseServer) ExportState(filePath string) {
	state := gameState{
		RoundNum:        2,
		MapSize:         server.MapSize,
		ZombiePositions: server.GetEntityLocations(extendedAgents.ZomboSapien),
		HumanPositions:  server.GetEntityLocations(extendedAgents.HomoSapien),
		Maze:            server.Maze,
		BorderSize:      10,
	}

	gameStateJSON, _ := json.Marshal(state)
	file, _ := os.Create(filePath)
	defer file.Close()
	file.Write(gameStateJSON)
}
