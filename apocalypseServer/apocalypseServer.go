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
	Maze             [][]int
}

type Exit struct {
	//Exit is a line segment in 2D space defined by two points
	PointA physicsEngine.Vector2D
	PointB physicsEngine.Vector2D
}

type gameState struct {
	//This json will be fed to a renderer to render a visualisation of the current game state
	RoundNum        int
	MapSize         physicsEngine.Vector2D
	ZombiePositions []physicsEngine.Vector2D
	HumanPositions  []physicsEngine.Vector2D
	Maze            [][]int
	BorderSize      int
}

func initiliaseMaze(rows, cols int) [][]int {
	maze := make([][]int, rows)
	for x := range maze {
		maze[x] = make([]int, cols)
		for y := range maze[x] {
			maze[x][y] = 1
		}
	}
	return maze
}

func CreateApocalypseServer(numZombies, numHumans, iterations, turns int, maxDuration time.Duration, maxThreads int, width, height, mazeSeed int) *ApocalypseServer {

	server := &ApocalypseServer{
		BaseServer: server.CreateServer[extendedAgents.IApocalypseEntity](iterations, turns, maxDuration, maxThreads),
		MapSize:    physicsEngine.MakeVec2D(width, height),
		Maze:       initiliaseMaze(width, height),
	}
	// for i := 0; i < numZombies; i++ {
	// 	zombie := server.SpawnNewZombie(10.0, server.GenerateRandomPosition())
	// 	server.AddAgent(zombie)
	// }
	// for i := 0; i < numHumans; i++ {
	// 	human := server.SpawnNewHuman(10.0, server.GenerateRandomPosition())
	// 	server.AddAgent(human)
	// }
	return server
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

func (server *ApocalypseServer) AddExit(point1, point2 physicsEngine.Vector2D) {
	exit := Exit{PointA: point1,
		PointB: point2}
	xMin, xMax := exit.PointA.X, exit.PointB.X

	if xMin > xMax {
		xMin, xMax = exit.PointB.X, exit.PointA.X
	}
	yMin, yMax := exit.PointA.Y, exit.PointB.Y
	if yMin > yMax {
		yMin, yMax = yMax, yMin
	}

	for x := xMin; x <= xMax; x++ {

		for y := yMin; y <= yMax; y++ {

			server.Maze[y][x] = 2
		}
	}
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
	_, _ = file.Write(gameStateJSON)
}

func (server *ApocalypseServer) GenerateRandomPosition() physicsEngine.Vector2D {
	//Generate a normally distributed position
	vec2 := physicsEngine.Vector2D{X: 0, Y: 0}
	vec2.Y = rand.IntN(server.MapSize.Y)
	vec2.X = rand.IntN(server.MapSize.X)
	return vec2
}

func (server *ApocalypseServer) CreateMaze(seed uint64) {
	//this function is seeded so that "random" mazes can be regenerated for multiple experiments
	randomSource := rand.NewPCG(seed, seed+26)
	generator := rand.New(randomSource)
	unitVectors := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
	state := genMaze(server.Maze, unitVectors, server.MapSize.X, server.MapSize.Y, 0, 0, generator)
	if !state {
		panic("couldnt generate solvable maze. Did you forget to add exits?")
	}
}

func outOfBoundsCheck(xMax, yMax, x, y int) bool {
	//fmt.Println("checking", x, y)
	if x < 0 || x >= xMax {
		return true
	} else if y < 0 || y >= yMax {
		return true
	} else {
		return false
	}
}

func genMaze(maze, unitVectors [][]int, xMax, yMax, x, y int, generator *rand.Rand) bool {
	state := false
	if maze[x][y] == 2 {
		return true
	}
	generator.Shuffle(4, func(i, j int) {
		unitVectors[i], unitVectors[j] = unitVectors[j], unitVectors[i]
	})
	maze[x][y] = 0
	for _, coords := range unitVectors {
		dx := coords[0]
		dy := coords[1]
		x1, y1 := x+dx, y+dy
		x2, y2 := x1+dx, y1+dy
		if outOfBoundsCheck(xMax, yMax, x2, y2) {
			continue
		}
		if maze[x2][y2] == 0 {
			continue
		}
		maze[x1][y1] = 0
		if genMaze(maze, unitVectors, xMax, yMax, x2, y2, generator) {
			state = true
		}
	}
	return state
}
