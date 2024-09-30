package apocalypseServer

import (
	"encoding/json"
	"math/rand/v2"
	"os"
	"time"
	extendedAgents "zombieApocalypeSOMAS/agent"
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
	Walls            []Wall
	Exits            []Exit
}

type Exit struct {
	//Exit is a line segment in 2D space defined by two points
	PointA physicsEngine.Vector2D
	PointB physicsEngine.Vector2D
}

type Wall struct {
	TopLeftCorner     physicsEngine.Vector2D
	BottomRightCorner physicsEngine.Vector2D
}

type PointPair struct {
	PointA physicsEngine.Vector2D
	PointB physicsEngine.Vector2D
}

func CreatePointPair(x, y physicsEngine.Vector2D) PointPair {
	return PointPair{PointA: x, PointB: y}
}

type gameState struct {
	//This json will be fed to a renderer to render a visualisation of the current game state
	RoundNum        int
	MapSize         physicsEngine.Vector2D
	Walls           []Wall
	Exits           []Exit
	ZombiePositions []physicsEngine.Vector2D
	HumanPositions  []physicsEngine.Vector2D
	BorderSize      float32
}

func CreateApocalypseServer(numZombies, numHumans, iterations, turns int, maxDuration time.Duration, maxThreads int, width, height float32) *ApocalypseServer {
	server := &ApocalypseServer{
		BaseServer: server.CreateServer[extendedAgents.IApocalypseEntity](iterations, turns, maxDuration, maxThreads),
		Walls:      make([]Wall, 0),
		Exits:      make([]Exit, 0),
		MapSize:    physicsEngine.MakeVec2D(width, height),
	}
	for i := 0; i < numZombies; i++ {
		zombie := server.SpawnNewZombie(10.0, server.GenerateRandomPosition())
		server.AddAgent(zombie)
	}
	for i := 0; i < numHumans; i++ {
		human := server.SpawnNewHuman(10.0, server.GenerateRandomPosition())
		server.AddAgent(human)
	}
	return server
}

func (serv *ApocalypseServer) SpawnNewHuman(mass float32, initialPosition physicsEngine.Vector2D) *extendedAgents.Human {
	entity := &extendedAgents.ApocalypseEntity{
		BaseAgent:     agent.CreateBaseAgent(serv),
		PhysicalState: physicsEngine.CreateInitialPhysicalState(&initialPosition, mass),
	}
	human := &extendedAgents.Human{ApocalypseEntity: entity}
	return human
}

func (serv *ApocalypseServer) SpawnNewZombie(mass float32, initialPosition physicsEngine.Vector2D) *extendedAgents.Zombie {
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

func (server *ApocalypseServer) AddWall(topLeft, bottomRight physicsEngine.Vector2D) {
	wall := Wall{TopLeftCorner: topLeft,
		BottomRightCorner: bottomRight}
	server.Walls = append(server.Walls, wall)
}

func (server *ApocalypseServer) AddExit(point1, point2 physicsEngine.Vector2D) {
	exit := Exit{PointA: point1,
		PointB: point2}
	server.Exits = append(server.Exits, exit)
}

func (server *ApocalypseServer) ExportState(filePath string) {

	state := gameState{
		RoundNum:        2,
		MapSize:         server.MapSize,
		Walls:           server.Walls,
		Exits:           server.Exits,
		ZombiePositions: server.GetEntityLocations(extendedAgents.ZomboSapien),
		HumanPositions:  server.GetEntityLocations(extendedAgents.HomoSapien),
		BorderSize:      5,
	}

	gameStateJSON, _ := json.Marshal(state)
	file, _ := os.Create(filePath)
	defer file.Close()
	_, _ = file.Write(gameStateJSON)
}

func (server *ApocalypseServer) GenerateRandomPosition() physicsEngine.Vector2D {
	//Generate a normally distributed position
	vec2 := physicsEngine.Vector2D{X: 0, Y: 0}
	vec2.Y = server.MapSize.Y * (float32(rand.NormFloat64()) + 1) / 2
	vec2.X = server.MapSize.X * (float32(rand.NormFloat64()) + 1) / 2
	return vec2
}

func (server *ApocalypseServer) CreateWalls(thickness float32, wallCoords []PointPair) {
	for _, wallCoord := range wallCoords {
		pointA := wallCoord.PointA
		pointB := wallCoord.PointB
		var wall Wall
		if pointA.X == pointB.X {
			if pointA.Y > pointB.Y {
				wall.TopLeftCorner = pointA
				wall.TopLeftCorner.X -= thickness
				wall.BottomRightCorner = pointB
				wall.BottomRightCorner.X += thickness
			} else {
				wall.TopLeftCorner = pointB
				wall.TopLeftCorner.X -= thickness
				wall.BottomRightCorner = pointA
				wall.BottomRightCorner.X += thickness
			}
		} else {
			if pointA.X > pointB.X {
				wall.TopLeftCorner = pointB
				wall.TopLeftCorner.Y += thickness
				wall.BottomRightCorner = pointA
				wall.BottomRightCorner.Y -= thickness
			} else {
				wall.TopLeftCorner = pointA
				wall.TopLeftCorner.Y += thickness
				wall.BottomRightCorner = pointB
				wall.BottomRightCorner.Y -= thickness
			}
		}
		server.Walls = append(server.Walls, wall)
	}

}
