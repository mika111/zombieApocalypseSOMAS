package main

import (
	"fmt"
	"time"
	"zombieApocalypeSOMAS/apocalypseServer"
	pathfinding "zombieApocalypeSOMAS/pathFinding"
)

func main() {
	serv := apocalypseServer.CreateApocalypseServer(1, 1, time.Millisecond, 100, 49, 49, 1257)
	serv.GenerateMaze(0, 0, 18, 18)
	//serv.Maze.Print()
	//serv.InjectAgents(10, 10)
	solnPath := pathfinding.FindPath(0, 0, 18, 18, serv.Maze)
	fmt.Println(solnPath)
	for _, coord := range solnPath {
		serv.Maze[coord[0]][coord[1]] = 4
	}

	serv.ExportState("state.json")
}
