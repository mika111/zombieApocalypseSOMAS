package mazeGenerator

import (
	"fmt"
	"math/rand/v2"
)

type Maze [][]int

type MazeGenerator struct {
	M             int
	N             int
	maze          Maze
	dirs          [][]int
	generator     *rand.Rand
	start         []int
	order         [][][]int
	openSpace     int
	exitableSpace int
}
type SpaceType int

const (
	open     SpaceType = 0
	exitable SpaceType = 3
)

func (m Maze) Print() {
	for _, row := range m {
		for _, elem := range row {
			fmt.Printf("%3d ", elem) // Adjust the width for alignment (e.g., %3d for 3 spaces)
		}
		fmt.Println() // Print a new line at the end of each row
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

func (mg *MazeGenerator) countSpace(spaceType SpaceType) {
	counter := 0
	for i := range mg.maze {
		for j := range mg.maze[i] {
			if SpaceType(mg.maze[i][j]) == spaceType {
				counter++
			}
		}
	}
	if spaceType == open {
		mg.openSpace = counter
	}
	if spaceType == exitable {
		mg.exitableSpace = counter
	}
}

func CreateMazeGenerator(M, N int, generator *rand.Rand) *MazeGenerator {
	if M&N&1 == 0 {
		panic("Both maze dimensions must be odd")
	}
	mazeGen := &MazeGenerator{
		M:         M,
		N:         N,
		maze:      nil,
		dirs:      nil,
		generator: generator,
		order: [][][]int{
			{{-1, 0}, {0, -1}, {1, 0}, {0, 1}},
			{{-1, 0}, {0, -1}, {0, 1}, {1, 0}},
			{{-1, 0}, {1, 0}, {0, -1}, {0, 1}},
			{{-1, 0}, {1, 0}, {0, 1}, {0, -1}},
			{{-1, 0}, {0, 1}, {0, -1}, {1, 0}},
			{{-1, 0}, {0, 1}, {1, 0}, {0, -1}},
			{{0, -1}, {-1, 0}, {1, 0}, {0, 1}},
			{{0, -1}, {-1, 0}, {0, 1}, {1, 0}},
			{{0, -1}, {1, 0}, {-1, 0}, {0, 1}},
			{{0, -1}, {1, 0}, {0, 1}, {-1, 0}},
			{{0, -1}, {0, 1}, {-1, 0}, {1, 0}},
			{{0, -1}, {0, 1}, {1, 0}, {-1, 0}},
			{{1, 0}, {-1, 0}, {0, -1}, {0, 1}},
			{{1, 0}, {-1, 0}, {0, 1}, {0, -1}},
			{{1, 0}, {0, -1}, {-1, 0}, {0, 1}},
			{{1, 0}, {0, -1}, {0, 1}, {-1, 0}},
			{{1, 0}, {0, 1}, {-1, 0}, {0, -1}},
			{{1, 0}, {0, 1}, {0, -1}, {-1, 0}},
			{{0, 1}, {-1, 0}, {0, -1}, {1, 0}},
			{{0, 1}, {-1, 0}, {1, 0}, {0, -1}},
			{{0, 1}, {0, -1}, {-1, 0}, {1, 0}},
			{{0, 1}, {0, -1}, {1, 0}, {-1, 0}},
			{{0, 1}, {1, 0}, {-1, 0}, {0, -1}},
			{{0, 1}, {1, 0}, {0, -1}, {-1, 0}},
		},
	}
	mazeGen.generateInitialMaze()
	return mazeGen
}

func (mg *MazeGenerator) CreateMaze(entrance_i, entrance_j, exit_i, exit_j int) Maze {
	mg.start = []int{entrance_i, entrance_j}
	if mg.isOutOfBounds(entrance_i, entrance_j) {
		panic("Entrance to maze is outside of dimensions")
	}
	if mg.isOutOfBounds(exit_i, exit_j) {
		panic("Exit to maze is outside of dimensions")
	}
	mg.maze[exit_i][exit_j] = 2
	state := mg.genMaze(entrance_i, entrance_j)
	fmt.Println(state)
	mg.checkSolvability()
	// if !state {
	// 	panic("couldnt generate solvable maze. Did you forget to add exits?")
	// }
	if mg.exitableSpace != mg.openSpace {
		panic(fmt.Sprintf("some portions of open space in the maze are blocked off. Open Space:%v, Exitable Space: %v", mg.openSpace, mg.exitableSpace))
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
	mg.maze[i][j] = 0
	randint := mg.generator.IntN(23)
	mg.dirs = mg.order[randint]
	for _, d := range mg.dirs {
		x1, y1 := i+d[0], j+d[1]
		x2, y2 := x1+d[0], y1+d[1]
		if mg.isOutOfBounds(x2, y2) || mg.maze[x2][y2] == 0 {
			continue
		}
		mg.maze[x1][y1] = 0
		state = (state || mg.genMaze(x2, y2))
	}
	return state
}

func (mg *MazeGenerator) getNeighbours(i, j int) [][]int {
	potentialNeighbours := [][]int{{i + 1, j},
		{i - 1, j},
		{i, j - 1},
		{i, j + 1}}
	neighbours := make([][]int, 0)
	for _, neighbour := range potentialNeighbours {
		if !mg.isOutOfBounds(neighbour[0], neighbour[1]) && mg.maze[neighbour[0]][neighbour[1]] == 0 {
			neighbours = append(neighbours, neighbour)
		}
	}
	return neighbours
}

func (mg *MazeGenerator) checkSolvability() {
	mg.countSpace(open)
	stack := [][]int{mg.start}

stackNotEmpty:
	for {
		lastElemDir := len(stack) - 1
		if lastElemDir == -1 {
			break stackNotEmpty
		}
		node := stack[lastElemDir]
		stack = stack[:lastElemDir]
		i, j := node[0], node[1]
		mg.maze[i][j] = 3
		newNeighbours := mg.getNeighbours(i, j)
		stack = append(stack, newNeighbours...)
	}
	mg.countSpace(exitable)
}
