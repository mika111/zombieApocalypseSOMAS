package apocalypseServer

import (
	"math/rand/v2"
	extendedAgents "zombieApocalypeSOMAS/agent"
	"zombieApocalypeSOMAS/physicsEngine"

	"github.com/MattSScott/basePlatformSOMAS/v2/pkg/agent"
)

func (serv *ApocalypseServer) InjectAgents(numHumans, numZombies int) {
	domain := serv.MapSpawnableArea(extendedAgents.ZomboSapien)
	for i := 0; i < numZombies; i++ {
		zombie := serv.SpawnNewZombie(10.0, serv.GenerateRandomPosition(domain))
		serv.AddAgent(zombie)
	}
	domain = serv.MapSpawnableArea(extendedAgents.HomoSapien)
	for i := 0; i < numHumans; i++ {
		human := serv.SpawnNewHuman(10.0, serv.GenerateRandomPosition(domain))
		serv.AddAgent(human)
	}
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

func (serv *ApocalypseServer) SpawnNewHuman(mass int, initialPosition physicsEngine.Vector2D) *extendedAgents.Human {
	entity := &extendedAgents.ApocalypseEntity{
		BaseAgent:     agent.CreateBaseAgent(serv),
		PhysicalState: physicsEngine.CreateInitialPhysicalState(&initialPosition, mass),
	}
	human := &extendedAgents.Human{ApocalypseEntity: entity}
	return human
}

func (server *ApocalypseServer) GenerateRandomPosition(domain []physicsEngine.Vector2D) physicsEngine.Vector2D {
	position := rand.IntN(len(domain))
	vec2 := physicsEngine.Vector2D{X: domain[position].X, Y: domain[position].Y}
	return vec2
}

func (server *ApocalypseServer) MapSpawnableArea(speciesType extendedAgents.Species) []physicsEngine.Vector2D {
	validTiles := 0
	var locationBlacklist map[physicsEngine.Vector2D]struct{}
	if speciesType == extendedAgents.ZomboSapien {
		locationBlacklist = server.EntityLocationMap(extendedAgents.HomoSapien)
	} else {
		locationBlacklist = server.EntityLocationMap(extendedAgents.ZomboSapien)
	}
	for x := 0; x < server.MapSize.X; x++ {
		for y := 0; y < server.MapSize.Y; y++ {
			if server.checkLocationIsValid(x, y, locationBlacklist) && server.Maze[x][y] == 0 {
				validTiles++
			}
		}
	}
	validTilesArray := make([]physicsEngine.Vector2D, validTiles)
	i := 0
	for x := 0; x < server.MapSize.X; x++ {
		for y := 0; y < server.MapSize.Y; y++ {
			if server.checkLocationIsValid(x, y, locationBlacklist) && server.Maze[x][y] == 0 {
				validTilesArray[i] = physicsEngine.MakeVec2D(x, y)
				i++
			}
		}
	}
	return validTilesArray
}

func (server *ApocalypseServer) checkLocationIsValid(x, y int, locationBlacklist map[physicsEngine.Vector2D]struct{}) bool {
	neighbours := server.GetNeighbours(x, y)
	for _, neighbour := range neighbours {
		_, exists := locationBlacklist[neighbour]
		if exists {
			return false
		}
	}
	return true
}

func (server *ApocalypseServer) GetNeighbours(x, y int) []physicsEngine.Vector2D {
	neighbours := make([]physicsEngine.Vector2D, 0)
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			if i < 0 || i >= server.MapSize.X || j < 0 || j >= server.MapSize.Y {
				continue
			}
			if server.Maze[i][j] == 1 || server.Maze[i][j] == 2 {
				continue
			}
			neighbours = append(neighbours, physicsEngine.MakeVec2D(i, j))
		}
	}
	return neighbours
}
