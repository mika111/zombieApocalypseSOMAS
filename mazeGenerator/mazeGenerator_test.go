package mazeGenerator_test

import (
	"math/rand/v2"
	"testing"
	"zombieApocalypeSOMAS/mazeGenerator"
)

var seed uint64 = 43

var randGen *rand.Rand = rand.New(rand.NewPCG(seed, seed))

func TestMazePrint(t *testing.T) {
	D := 19
	mazeGen := mazeGenerator.CreateMazeGenerator(D, D, D-1, D-1, randGen)
	maze := mazeGen.CreateMaze(0, 0)
	maze.Print()
	t.Fail()
}
