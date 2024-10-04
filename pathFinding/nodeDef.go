package pathfinding

import "math"

type node struct {
	X            int
	Y            int
	parentXCoord int
	parentYCoord int
	FCost        int
	HCost        int
	GCost        int
}

func NewNode(x, y int) node {
	return node{
		X:            x,
		Y:            y,
		parentXCoord: -1,
		parentYCoord: -1,
		FCost:        math.MaxInt,
		HCost:        math.MaxInt,
		GCost:        math.MaxInt,
	}
}

func (n *node) UpdateNode(f, g, h, parentX, parentY int) {
	n.parentXCoord = parentX
	n.parentYCoord = parentY
	n.FCost = f
	n.HCost = h
	n.GCost = g
}
