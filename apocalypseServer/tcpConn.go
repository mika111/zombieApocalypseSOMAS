package apocalypseServer

import (
	"encoding/json"
	"fmt"
	"net"
	extendedAgents "zombieApocalypeSOMAS/agent"
	"zombieApocalypeSOMAS/mazeGenerator"
	"zombieApocalypeSOMAS/physicsEngine"
)

type staticGameState struct {
	//this is only sent once; it can never change
	Maze       mazeGenerator.Maze
}

type gameState struct {
	RoundNum        int
	ZombiePositions []physicsEngine.Vector2D
	HumanPositions  []physicsEngine.Vector2D
}

func (a *ApocalypseServer) ConnectToFrontEnd(addr string) {
	listen, _ := net.Listen("tcp", addr)
	a.Connection, _ = listen.Accept()
	fmt.Println("connection accepted")
}

func (server *ApocalypseServer) ExportState() {
	state := gameState{
		RoundNum:        2,
		ZombiePositions: server.GetEntityLocations(extendedAgents.ZomboSapien),
		HumanPositions:  server.GetEntityLocations(extendedAgents.HomoSapien),
	}

	gameStateJSON, _ := json.Marshal(state)
	server.Connection.Write(gameStateJSON)
}

func (server *ApocalypseServer) ExportInitialState() {
	//export all the data for setting up the renderer
	state := staticGameState{
		Maze: server.Maze,
	}
	gameStateJSON, _ := json.Marshal(state)
	//fmt.Println("Sending json")
	server.Connection.Write(gameStateJSON)
}
