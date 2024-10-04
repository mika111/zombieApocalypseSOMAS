package main

import (
	"time"
	"zombieApocalypeSOMAS/apocalypseServer"
)

func main() {
	serv := apocalypseServer.CreateApocalypseServer(1, 1, time.Millisecond, 100, 49, 49, 1257)
	serv.ConnectToFrontEnd("localhost:8080")
	serv.GenerateMaze(0, 0, 48, 48)
	serv.ExportState("state.json")
	time.Sleep(5 * time.Second)
	serv.InjectAgents(1000, 100)
	serv.ExportState("state.json")
}
