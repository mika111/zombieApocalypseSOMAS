package apocalypseServer

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	extendedAgents "zombieApocalypeSOMAS/agent"
	"zombieApocalypeSOMAS/mazeGenerator"
	"zombieApocalypeSOMAS/physicsEngine"
)

type gameState struct {
	RoundNum        int
	MapSize         physicsEngine.Vector2D
	ZombiePositions []physicsEngine.Vector2D
	HumanPositions  []physicsEngine.Vector2D
	Maze            mazeGenerator.Maze
	BorderSize      int
}

func (a *ApocalypseServer) ConnectToFrontEnd(addr string) {
	listen, _ := net.Listen("tcp", addr)
	a.Connection, _ = listen.Accept()
	fmt.Println("connection accepted")
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
	fmt.Println("Sending json")
	server.Connection.Write(gameStateJSON)
}
