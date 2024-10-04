package pathfinding

import (
	"container/heap"
)

func (s *solver) tracePath(startX, startY, endX, endY int) {
	//first find length of path and preallocate arr
	pathLen := 0
	frontierNode := s.nodeList[endX][endY]
	startNode := s.nodeList[startX][startY]
	for frontierNode != startNode {
		frontierNode = s.nodeList[frontierNode.parentXCoord][frontierNode.parentYCoord]
		pathLen++
	}
	s.solutionPath = make([][]int, pathLen)
	i := 0
	frontierNode = s.nodeList[endX][endY]
	coords := []int{endX, endY}
	for frontierNode != startNode {
		s.solutionPath[i] = []int{coords[0], coords[1]}
		coords[0] = frontierNode.parentXCoord
		coords[1] = frontierNode.parentYCoord
		frontierNode = s.nodeList[frontierNode.parentXCoord][frontierNode.parentYCoord]
		i++
	}
}

func (s *solver) search(startX, startY, goalX, goalY int) bool {
	s.nodeList[startX][startY].UpdateNode(0, 0, 0, startX, startY)
	heap.Push(&s.openList, s.nodeList[startX][startY])
	for s.openList.Len() > 0 {
		frontierNode := heap.Pop(&s.openList).(*priorityQueueNode).node
		s.visited[frontierNode.X][frontierNode.Y] = true
		neighbourList := s.generateNeighbours(frontierNode.X, frontierNode.Y)
		for _, neighbourCoords := range neighbourList {
			neiX := neighbourCoords[0]
			neiY := neighbourCoords[1]
			if neiX == goalX && neiY == goalY {
				s.nodeList[goalX][goalY].parentXCoord = frontierNode.X
				s.nodeList[goalX][goalY].parentYCoord = frontierNode.Y
				s.tracePath(startX, startY, goalX, goalY)
				return true
			}
			newG := s.nodeList[frontierNode.X][frontierNode.Y].GCost + 1
			newH := s.hCost(neiX, neiY, goalX, goalY, chebyshev)
			newF := newG + newH
			if newF < s.nodeList[neiX][neiY].FCost {
				s.nodeList[neiX][neiY].UpdateNode(newF, newG, newH, frontierNode.X, frontierNode.Y)
				heap.Push(&s.openList, s.nodeList[neiX][neiY])
			}
		}
	}
	return false
}

func FindPath(startX, startY, goalX, goalY int, maze [][]int) [][]int {
	solver := newSolver(maze)
	solver.search(startX, startY, goalX, goalY)
	return solver.solutionPath
}
