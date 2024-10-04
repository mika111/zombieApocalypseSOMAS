package pathfinding

type solver struct {
	maze         [][]int
	width        int
	height       int
	dirs         [][]int
	visited      [][]bool
	openList     priorityQueue
	nodeList     [][]node
	solutionPath [][]int
}

func newSolver(maze [][]int) solver {
	solver := solver{
		maze:     maze,
		height:   len(maze),
		width:    len(maze[0]),
		dirs:     [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}},
		openList: make(priorityQueue, 0),
	}
	solver.generateVisitedAndNodeArray()
	return solver
}

func (s *solver) generateVisitedAndNodeArray() {
	s.visited = make([][]bool, s.height)
	s.nodeList = make([][]node, s.height)
	for x := 0; x < s.width; x++ {
		s.visited[x] = make([]bool, s.width)
		s.nodeList[x] = make([]node, s.width)
		for y := 0; y < s.height; y++ {
			s.visited[x][y] = false
			s.nodeList[x][y] = NewNode(x, y)
		}
	}
}

func (s *solver) isOutOfBounds(x, y int) bool {
	if x < 0 || x >= s.width {
		return true
	}
	if y < 0 || y >= s.height {
		return true
	}
	return s.maze[x][y] == 1
}

func (s *solver) generateNeighbours(x, y int) [][]int {
	neighbours := make([][]int, 0)
	for _, dir := range s.dirs {
		xNew, yNew := x+dir[0], y+dir[1]
		if s.isOutOfBounds(xNew, yNew) || s.visited[xNew][yNew] {
			continue
		} else {
			neighbours = append(neighbours, []int{xNew, yNew})
		}
	}
	return neighbours
}

func (s *solver) hCost(xStart, yStart, xTarget, yTarget int, heuristicFunction func(int, int) int) int {
	xDistance := xStart - xTarget
	yDistance := yStart - yTarget
	return heuristicFunction(xDistance, yDistance)
}
