package pathfinding

type priorityQueue []*node

func (p priorityQueue) Len() int {
	return len(p)
}

func (p priorityQueue) Less(i, j int) bool {
	return p[i].FCost < p[j].FCost
}

func (p priorityQueue) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *priorityQueue) Push(nodeToAdd any) {
	newNode := nodeToAdd.(node)
	*p = append(*p, &newNode)
}

func (p *priorityQueue) Pop() any {
	n := len(*p) - 1
	item := (*p)[n]
	(*p)[n] = nil
	*p = (*p)[0:n]
	return item
}
