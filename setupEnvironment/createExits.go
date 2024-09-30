package setupEnvironment

import (
	"zombieApocalypeSOMAS/apocalypseServer"
	"zombieApocalypeSOMAS/physicsEngine"
)

func CreateExits(server *apocalypseServer.ApocalypseServer) {
	server.AddExit(physicsEngine.MakeVec2D(0, 0), physicsEngine.MakeVec2D(0, 100))
	server.AddExit(physicsEngine.MakeVec2D(10, 0), physicsEngine.MakeVec2D(200, 0))
}
