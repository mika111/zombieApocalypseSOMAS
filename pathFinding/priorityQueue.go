package pathfinding

import "container/heap"

type priorityQueueNode struct {
	node  node
	index int
}

type priorityQueue []*priorityQueueNode

func (p priorityQueue) Len() int {
	return len(p)
}

func (p priorityQueue) Less(i, j int) bool {
	return p[i].node.FCost < p[j].node.FCost
}

func (p priorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].index = i
	p[j].index = j
}

func (p *priorityQueue) Push(nodeToAdd any) {
	newNode := nodeToAdd.(node)
	pNode := &priorityQueueNode{node: newNode, index: len(*p)}
	*p = append(*p, pNode)
}

func (p *priorityQueue) Pop() any {
	n := len(*p) - 1
	item := (*p)[n]
	(*p)[n] = nil
	*p = (*p)[0:n]
	return item
}

func (p *priorityQueue) update(node *priorityQueueNode, newF,newG,newH int) {
	node.node.FCost = newF
	node.node.GCost = newG
	node.node.HCost = newH
	heap.Fix(p, node.index)
}
