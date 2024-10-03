package mazeGenerator

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand/v2"
	"os"
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
	DirArray  [][]physicsEngine.Vector2D
	i         int
}

type ShuffleOrderings struct {
	DirArray [][]physicsEngine.Vector2D
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
	jsonObject := openJSON("mazeData.json")
	mazeData := jsonObject.MazeData
	mazeArr := make([][]physicsEngine.Vector2D, len(mazeData))
	for i := range mazeData {
		subArray := make([]physicsEngine.Vector2D, 4)
		for j, mazePoint := range mazeData[i] {
			subArray[j] = physicsEngine.Vector2D{X: mazePoint[0], Y: mazePoint[1]}
		}
		mazeArr[i] = subArray
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
		DirArray:  mazeArr,
		i:         0,
	}
	mazeGen.generateInitialMaze()
	return mazeGen
}

func (mg *MazeGenerator) CreateMaze(entrance_i, entrance_j, exit_i, exit_j int) Maze {
	fmt.Println("length of instructions", len(mg.DirArray))
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
	mg.writeToJSON("mazeObject.json")
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
	if mg.i >= len(mg.DirArray) {
		return true
	}
	if mg.maze[i][j] == 2 {
		return true
	}
	state := false
	mg.maze[i][j] = 0
	mg.dirs = mg.DirArray[mg.i]
	mg.i += 1
	for _, d := range mg.dirs {
		x1, y1 := i+d.X, j+d.Y
		x2, y2 := x1+d.X, y1+d.Y
		if mg.isOutOfBounds(x2, y2) || mg.maze[x2][y2] == 0 {
			continue
		}
		mg.maze[x1][y1] = 0
		state = (mg.genMaze(x2, y2))
	}
	return state
}

func (mg *MazeGenerator) writeToJSON(filePath string) {
	shuffles := ShuffleOrderings{
		DirArray: mg.DirArray,
	}
	gameStateJSON, _ := json.Marshal(shuffles)
	file, _ := os.Create(filePath)
	defer file.Close()
	file.Write(gameStateJSON)
}

type mazeData struct {
	MazeData [][][]int
}

func openJSON(filename string) mazeData {
	file, error := os.Open(filename)
	if error != nil {
		log.Fatalf("Failed to open file: %v", error)
	}
	defer file.Close()

	byteval, error := io.ReadAll(file)
	if error != nil {
		log.Fatalf("Failed to read file: %v", error)
	}
	var result mazeData
	error = json.Unmarshal(byteval, &result)
	if error != nil {
		log.Fatalf("Failed to unmarshal JSON: %v", error)
	}
	fmt.Println("Got JSON data")
	return result
}
