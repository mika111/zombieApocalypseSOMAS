package apocalypseServer

import (
	"encoding/json"
	"math/rand/v2"
	"os"
	"time"
	extendedAgents "zombieApocalypeSOMAS/agent"
	"zombieApocalypeSOMAS/mazeGenerator"
	"zombieApocalypeSOMAS/physicsEngine"

	"github.com/MattSScott/basePlatformSOMAS/v2/pkg/agent"
	"github.com/MattSScott/basePlatformSOMAS/v2/pkg/server"
)

type IApocalypseServer interface {
	server.IServer[extendedAgents.IApocalypseEntity]
	GetNumZombies() int
	GetNumHumans() int
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

func (serv *ApocalypseServer) InjectAgents(numHumans, numZombies int) {
	for i := 0; i < numZombies; i++ {
		zombie := serv.SpawnNewZombie(10.0, serv.GenerateRandomPosition())
		serv.AddAgent(zombie)
	}
	for i := 0; i < numHumans; i++ {
		human := serv.SpawnNewHuman(10.0, serv.GenerateRandomPosition())
		serv.AddAgent(human)
	}
}

func (serv *ApocalypseServer) GenerateMaze(entrance_i, entrance_j, exit_i, exit_j int) {
	mazeGen := mazeGenerator.CreateMazeGenerator(serv.MapSize.X, serv.MapSize.Y, serv.RandNumGenerator)
	serv.Maze = mazeGen.CreateMaze(entrance_i, entrance_j, exit_i, exit_j)
}

func (serv *ApocalypseServer) SpawnNewHuman(mass int, initialPosition physicsEngine.Vector2D) *extendedAgents.Human {
	entity := &extendedAgents.ApocalypseEntity{
		BaseAgent:     agent.CreateBaseAgent(serv),
		PhysicalState: physicsEngine.CreateInitialPhysicalState(&initialPosition, mass),
	}
	human := &extendedAgents.Human{ApocalypseEntity: entity}
	return human
}

func (serv *ApocalypseServer) SpawnNewZombie(mass int, initialPosition physicsEngine.Vector2D) *extendedAgents.Zombie {
	entity := &extendedAgents.ApocalypseEntity{
		BaseAgent:     agent.CreateBaseAgent(serv),
		PhysicalState: physicsEngine.CreateInitialPhysicalState(&initialPosition, mass),
	}

	zombie := &extendedAgents.Zombie{
		ApocalypseEntity: entity,
		Strength:         10,
		RandNumGenerator: serv.RandNumGenerator,
	}
	return zombie
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

func (server *ApocalypseServer) GenerateRandomPosition() physicsEngine.Vector2D {
	vec2 := physicsEngine.Vector2D{X: 0, Y: 0}
	vec2.Y = rand.IntN(server.MapSize.Y)
	vec2.X = rand.IntN(server.MapSize.X)
	return vec2
}
