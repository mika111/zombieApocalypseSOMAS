package mazeGenerator

import (
	"fmt"
	"math/rand/v2"
)

type Maze [][]int

type MazeGenerator struct {
	M         int
	N         int
	maze      Maze
	generator *rand.Rand
	start     []int
	dirs      [][]int
}

func (m Maze) Print() {
	for _, row := range m {
		rowOut := ""
		for _, elem := range row {
			if elem == 0 {
				rowOut += " "
			} else if elem == 1 {
				rowOut += "#"
			} else {
				rowOut += "*"
			}
			rowOut += " "
		}
		fmt.Println(rowOut)
	}
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
		M:         M,
		N:         N,
		maze:      nil,
		dirs:      [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}},
		generator: generator,
	}
	mazeGen.generateInitialMaze()
	return mazeGen
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (mg *MazeGenerator) CreateMaze(entrance_i, entrance_j, exit_i, exit_j int) Maze {
	mg.start = []int{entrance_i, entrance_j}
	if mg.isOutOfBounds(entrance_i, entrance_j) {
		panic("Entrance to maze is outside of dimensions")
	}
	if mg.isOutOfBounds(exit_i, exit_j) {
		panic("Exit to maze is outside of dimensions")
	}

	i_dist := abs(entrance_i - exit_i)
	j_dist := abs(exit_i - exit_j)
	if (i_dist|j_dist)&1 == 1 {
		panic("Difference between both starting and ending co-ordinates must be even")
	}

	mg.maze[exit_i][exit_j] = 2
	state := mg.genMaze(entrance_i, entrance_j)
	if !state {
		mg.maze.Print()
		panic("Couldn't generate solvable maze. Did you forget to add exits?")
	}
	if mg.mazeIsSolveable() {
		panic("Some portions of open space in the maze are blocked off.")
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
	state := mg.maze[i][j] == 2
	if !state {
		mg.maze[i][j] = 0
	}
	for _, d := range GetRandomTraversal(mg.generator) {
		fmt.Println(d, i, j)
		x1, y1 := i+d[0], j+d[1]
		x2, y2 := x1+d[0], y1+d[1]
		if mg.isOutOfBounds(x2, y2) || mg.maze[x2][y2] == 0 {
			continue
		}
		mg.maze[x1][y1] = 0
		state = state || mg.genMaze(x2, y2)
	}
	return state
}

func (mg *MazeGenerator) traverseMaze(i, j int, visited [][]bool) {
	if mg.isOutOfBounds(i, j) || mg.maze[i][j] != 0 || visited[i][j] {
		return
	}
	visited[i][j] = true
	for _, dir := range mg.dirs {
		di, dj := dir[0], dir[1]
		i1, j1 := i+di, j+dj
		mg.traverseMaze(i1, j1, visited)
	}
}

func (mg *MazeGenerator) mazeIsSolveable() bool {
	visited := make([][]bool, mg.M)
	for i := range visited {
		visited[i] = make([]bool, mg.N)
	}

	regions := 0

	for i := 0; i < mg.M; i++ {
		for j := 0; j < mg.N; j++ {
			if visited[i][j] {
				continue
			}
			regions++
			mg.traverseMaze(i, j, visited)
		}
	}
	return regions == 1
}
