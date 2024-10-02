package mazeGenerator

import (
	"fmt"
	"math/rand/v2"
	"zombieApocalypeSOMAS/physicsEngine"
)

type Maze [][]int

func (m Maze) Print() {
	for _, row := range m {
		for _, elem := range row {
			fmt.Printf("%3d ", elem) // Adjust the width for alignment (e.g., %3d for 3 spaces)
		}
		fmt.Println() // Print a new line at the end of each row
	}
}

type MazeGenerator struct {
	M         int
	N         int
	maze      Maze
	dirs      []physicsEngine.Vector2D
	generator *rand.Rand
}

func (mg *MazeGenerator) generateInitialMaze() {

	array := make([][]int, mg.M)
	for i := range array {
		array[i] = make([]int, mg.N)
		for j := range array[i] {
			array[i][j] = 1
		}
	}

	mg.maze = array
}

func CreateMazeGenerator(M, N int, generator *rand.Rand) *MazeGenerator {
	if M&N&1 == 0 {
		panic("Both maze dimensions must be odd")
	}
	mazeGen := &MazeGenerator{
		M:    M,
		N:    N,
		maze: nil,
		dirs: []physicsEngine.Vector2D{
			{X: 0, Y: 1},
			{X: 1, Y: 0},
			{X: 0, Y: -1},
			{X: -1, Y: 0},
		},
		generator: generator,
	}
	mazeGen.generateInitialMaze()
	return mazeGen
}

func (mg *MazeGenerator) CreateMaze(entrance_i, entrance_j, exit_i, exit_j int) Maze {
	if mg.isOutOfBounds(entrance_i, entrance_j) {
		panic("Entrance to maze is outside of dimensions")
	}
	if mg.isOutOfBounds(exit_i, exit_j) {
		panic("Exit to maze is outside of dimensions")
	}
	mg.maze[exit_i][exit_j] = 2
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
	// mg.generator.Shuffle(4, func(i, j int) {
	// 	mg.dirs[i], mg.dirs[j] = mg.dirs[j], mg.dirs[i]
	// })
	rand.Shuffle(4, func(i, j int) {
		mg.dirs[i], mg.dirs[j] = mg.dirs[j], mg.dirs[i]
	})
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


