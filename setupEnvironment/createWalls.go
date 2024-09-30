package setupEnvironment

import (
	"zombieApocalypeSOMAS/apocalypseServer"
	"zombieApocalypeSOMAS/physicsEngine"
)

func CreateMaze(server *apocalypseServer.ApocalypseServer) {
	var thickness float32 = 5
	mapSize := server.MapSize
	xScale := mapSize.X / 16
	yScale := mapSize.Y / 16
	generateCoordinates := func(x1, y1, Xbeta1, YBeta1, x2, y2, Xbeta2, YBeta2 float32) apocalypseServer.PointPair {
		//Created Map using a discrete coordinate system from 1-16. This function will
		//automatically scale values according to mapSize
		scaledCoord1 := physicsEngine.Vector2D{}
		scaledCoord2 := physicsEngine.Vector2D{}
		scaledCoord1.X = xScale*x1 + Xbeta1
		scaledCoord1.Y = yScale*y1 + YBeta1
		scaledCoord2.X = xScale*x2 + Xbeta2
		scaledCoord2.Y = yScale*y2 + YBeta2
		return apocalypseServer.PointPair{PointA: scaledCoord1, PointB: scaledCoord2}
	}

	wallCoords := make([]apocalypseServer.PointPair, 48)

	//create array of walls
	wallCoords[0] = generateCoordinates(0, 2, 0, 0, 1, 2, 0, 0)
	wallCoords[1] = generateCoordinates(1, 1, 0, 0, 3, 1, 0, 0)
	wallCoords[2] = generateCoordinates(3, 1, 0, thickness, 3, 0, 0, 0)
	wallCoords[3] = generateCoordinates(2, 1, 0, 0, 2, 2, 0, thickness)
	wallCoords[4] = generateCoordinates(2, 2, -thickness, 0, 3, 2, thickness, 0)
	wallCoords[5] = generateCoordinates(1, 3, 0, 0, 2, 3, -thickness, 0)
	wallCoords[6] = generateCoordinates(3, 3, thickness, 0, 4, 3, 1*thickness, 0)
	wallCoords[7] = generateCoordinates(4, 3, 0, thickness, 4, 2, 0, -thickness)
	wallCoords[8] = generateCoordinates(4, 2, -thickness, 0, 5, 2, 0, 0)
	wallCoords[9] = generateCoordinates(5, 0, 0, 0, 5, 1, 0, 0)
	wallCoords[10] = generateCoordinates(4, 1, -thickness, 0, 6, 1, thickness, 0)
	wallCoords[11] = generateCoordinates(6, 2, 0, 0, 6, 4, 0, 0)
	wallCoords[12] = generateCoordinates(5, 3, 0, 0, 5, 5, 0, thickness)
	wallCoords[13] = generateCoordinates(2, 4, -thickness, 0, 4, 4, 0, 0)
	wallCoords[14] = generateCoordinates(4, 4, 0, -thickness, 4, 5, 0, thickness)
	wallCoords[15] = generateCoordinates(2, 4, 0, 0, 2, 6, 0, thickness)
	wallCoords[16] = generateCoordinates(3, 5, 0, 0, 3, 9, 0, 0)
	wallCoords[17] = generateCoordinates(2, 9, -thickness, 0, 4, 9, thickness, 0)
	wallCoords[18] = generateCoordinates(3, 8, 0, 0, 4, 8, thickness, 0)
	wallCoords[19] = generateCoordinates(1, 9, 0, thickness, 1, 7, 0, -thickness)
	wallCoords[20] = generateCoordinates(1, 7, -thickness, 0, 2, 7, thickness, 0)
	wallCoords[21] = generateCoordinates(2, 7, 0, -thickness, 2, 8, 0, 0)
	wallCoords[22] = generateCoordinates(2, 10, -thickness, 0, 4, 10, thickness, 0)
	wallCoords[23] = generateCoordinates(1, 10, 0, -thickness, 1, 11, 0, thickness)
	wallCoords[24] = generateCoordinates(0, 11, 0, 0, 1, 11, thickness, 0)
	wallCoords[25] = generateCoordinates(1, 4, 0, -thickness, 1, 6, 0, thickness)
	wallCoords[26] = generateCoordinates(0, 6, 0, 0, 1, 6, thickness, 0)
	wallCoords[27] = generateCoordinates(4, 6, 0, -thickness, 4, 7, 0, thickness)
	wallCoords[28] = generateCoordinates(4, 6, -thickness, 0, 7, 6, 0, 0)
	wallCoords[29] = generateCoordinates(6, 5, 0, 0, 6, 8, 0, 0)
	wallCoords[30] = generateCoordinates(5, 7, -thickness, 0, 6, 7, 0, 0)
	wallCoords[31] = generateCoordinates(5, 7, 0, -thickness, 5, 10, 0, thickness)
	wallCoords[32] = generateCoordinates(5, 5, -thickness, 0, 8, 5, 0, 0)
	wallCoords[33] = generateCoordinates(0, 13, 0, 0, 2, 13, thickness, 0)
	wallCoords[34] = generateCoordinates(3, 12, 0, -thickness, 3, 14, 0, thickness)
	wallCoords[35] = generateCoordinates(3, 14, -thickness, 0, 4, 14, 0, 0)
	wallCoords[36] = generateCoordinates(4, 12, 0, -thickness, 4, 15, 0, thickness)
	wallCoords[37] = generateCoordinates(1, 14, 0, 0, 1, 15, 0, thickness)
	wallCoords[38] = generateCoordinates(1, 15, -thickness, 0, 3, 15, 0, 0)
	wallCoords[39] = generateCoordinates(1, 12, 0, 0, 2, 12, thickness, 0)
	wallCoords[40] = generateCoordinates(2, 11, 0, -thickness, 2, 12, 0, thickness)
	wallCoords[41] = generateCoordinates(2, 11, -thickness, 0, 5, 11, 0, 0)
	wallCoords[42] = generateCoordinates(4, 12, -thickness, 0, 6, 12, thickness, 0)
	wallCoords[43] = generateCoordinates(6, 9, 0, -thickness, 6, 12, 0, thickness)
	wallCoords[44] = generateCoordinates(6, 10, 0, 0, 7, 10, thickness, 0)
	wallCoords[45] = generateCoordinates(7, 7, 0, -thickness, 7, 10, 0, thickness)
	wallCoords[46] = generateCoordinates(7, 11, 0, 0, 7, 14, 0, thickness)
	wallCoords[47] = generateCoordinates(6, 14, 0, 0, 7, 14, thickness, 0)
	// wallCoords[0] = generateCoordinates(0, 0, 0, 0, 0, 0, 0, 0)
	// wallCoords[0] = generateCoordinates(0, 0, 0, 0, 0, 0, 0, 0)
	// wallCoords[0] = generateCoordinates(0, 0, 0, 0, 0, 0, 0, 0)
	// wallCoords[0] = generateCoordinates(0, 0, 0, 0, 0, 0, 0, 0)
	// wallCoords[0] = generateCoordinates(0, 0, 0, 0, 0, 0, 0, 0)
	// wallCoords[0] = generateCoordinates(0, 0, 0, 0, 0, 0, 0, 0)
	// wallCoords[0] = generateCoordinates(0, 0, 0, 0, 0, 0, 0, 0)
	// wallCoords[0] = generateCoordinates(0, 0, 0, 0, 0, 0, 0, 0)
	// wallCoords[0] = generateCoordinates(0, 0, 0, 0, 0, 0, 0, 0)
	// wallCoords[0] = generateCoordinates(0, 0, 0, 0, 0, 0, 0, 0)
	// wallCoords[0] = generateCoordinates(0, 0, 0, 0, 0, 0, 0, 0)
	// wallCoords[0] = generateCoordinates(0, 0, 0, 0, 0, 0, 0, 0)
	// wallCoords[0] = generateCoordinates(0, 0, 0, 0, 0, 0, 0, 0)
	// wallCoords[0] = generateCoordinates(0, 0, 0, 0, 0, 0, 0, 0)
	// wallCoords[0] = generateCoordinates(0, 0, 0, 0, 0, 0, 0, 0)
	// wallCoords[0] = generateCoordinates(0, 0, 0, 0, 0, 0, 0, 0)
	// wallCoords[0] = generateCoordinates(0, 0, 0, 0, 0, 0, 0, 0)
	// wallCoords[0] = generateCoordinates(0, 0, 0, 0, 0, 0, 0, 0)
	// wallCoords[0] = generateCoordinates(0, 0, 0, 0, 0, 0, 0, 0)
	// wallCoords[0] = generateCoordinates(0, 0, 0, 0, 0, 0, 0, 0)
	// wallCoords[0] = generateCoordinates(0, 0, 0, 0, 0, 0, 0, 0)
	// wallCoords[0] = generateCoordinates(0, 0, 0, 0, 0, 0, 0, 0)
	// wallCoords[0] = generateCoordinates(0, 0, 0, 0, 0, 0, 0, 0)
	// wallCoords[0] = generateCoordinates(0, 0, 0, 0, 0, 0, 0, 0)
	// wallCoords[0] = generateCoordinates(0, 0, 0, 0, 0, 0, 0, 0)
	// wallCoords[0] = generateCoordinates(0, 0, 0, 0, 0, 0, 0, 0)
	// wallCoords[0] = generateCoordinates(0, 0, 0, 0, 0, 0, 0, 0)
	// wallCoords[0] = generateCoordinates(0, 0, 0, 0, 0, 0, 0, 0)
	// wallCoords[0] = generateCoordinates(0, 0, 0, 0, 0, 0, 0, 0)
	// wallCoords[0] = generateCoordinates(0, 0, 0, 0, 0, 0, 0, 0)
	// wallCoords[0] = generateCoordinates(0, 0, 0, 0, 0, 0, 0, 0)
	// wallCoords[0] = generateCoordinates(0, 0, 0, 0, 0, 0, 0, 0)
	// wallCoords[0] = generateCoordinates(0, 0, 0, 0, 0, 0, 0, 0)
	// wallCoords[0] = generateCoordinates(0, 0, 0, 0, 0, 0, 0, 0)
	// wallCoords[0] = generateCoordinates(0, 0, 0, 0, 0, 0, 0, 0)
	// wallCoords[0] = generateCoordinates(0, 0, 0, 0, 0, 0, 0, 0)
	// wallCoords[0] = generateCoordinates(0, 0, 0, 0, 0, 0, 0, 0)
	// wallCoords[0] = generateCoordinates(0, 0, 0, 0, 0, 0, 0, 0)
	//pass array of walls to wall constructor
	server.CreateWalls(thickness, wallCoords)
}
