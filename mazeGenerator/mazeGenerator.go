package mazeGenerator

import (
	"fmt"
	"math/rand/v2"
	"zombieApocalypeSOMAS/physicsEngine"
)

type Maze [][]int

func (m Maze) Print() {
	for _, row := range m {
		rowStr := ""
		for _, elem := range row {
			if elem == 0 {
				rowStr += " "
			} else if elem == 1 {
				rowStr += "#"
			} else {
				rowStr += "*"
			}
			rowStr += " "
		}
		fmt.Println(rowStr)
	}
}

type MazeGenerator struct {
	M         int
	N         int
	maze      Maze
	dirs      []physicsEngine.Vector2D
	generator *rand.Rand
}

func (mg *MazeGenerator) generateInitialMaze(i, j int) {
	if mg.isOutOfBounds(i, j) {
		panic("Exit to maze is outside of dimensions")
	}
	array := make([][]int, mg.M)
	for i := range array {
		array[i] = make([]int, mg.N)
		for j := range array[i] {
			array[i][j] = 1
		}
	}
	array[i][j] = 2
	mg.maze = array
}

func CreateMazeGenerator(M, N, exit_i, exit_j int, generator *rand.Rand) *MazeGenerator {
	if M&N&1 == 0 {
		panic("Both maze dimensions must be odd")
	}
	mazeGen := &MazeGenerator{
		M:    M,
		N:    N,
		maze: nil,
		dirs: []physicsEngine.Vector2D{
			{X: -1, Y: 0},
			{X: 0, Y: -1},
			{X: 1, Y: 0},
			{X: 0, Y: 1},
		},
		generator: generator,
	}
	mazeGen.generateInitialMaze(exit_i, exit_j)
	return mazeGen
}

func (mg *MazeGenerator) CreateMaze(entrance_i, entrance_j int) Maze {
	if mg.isOutOfBounds(entrance_i, entrance_j) {
		panic("Entrance to maze is outside of dimensions")
	}
	state := mg.genMaze(entrance_i, entrance_j)
	if !state {
		panic("couldnt generate solvable maze. Did you forget to add exits?")
	}
	return mg.maze
}

func (mg *MazeGenerator) isOutOfBounds(i, j int) bool {
	if i < 0 || i >= mg.M {
		return true
	}
	if j < 0 || j >= mg.N {
		return true
	}
	return false
}

func (mg *MazeGenerator) genMaze(i, j int) bool {
	if mg.maze[i][j] == 2 {
		return true
	}
	state := false
	mg.generator.Shuffle(4, func(i, j int) {
		mg.dirs[i], mg.dirs[j] = mg.dirs[j], mg.dirs[i]
	})
	// fmt.Println(mg.dirs)
	mg.maze[i][j] = 0
	for _, d := range mg.dirs {
		x1, y1 := i+d.X, j+d.Y
		x2, y2 := x1+d.X, y1+d.Y
		if mg.isOutOfBounds(x2, y2) || mg.maze[x2][y2] == 0 {
			continue
		}
		mg.maze[x1][y1] = 0
		state = state || mg.genMaze(x2, y2)
	}
	return state
}

// func (mg *MazeGenerator) GetMaze() Maze {
// 	mg.maze.Print()
// 	return mg.maze
// }
